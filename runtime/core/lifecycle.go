package core

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/ysugimoto/vintage"
)

const (
	fastlySubroutineRecv    = "vcl_recv"
	fastlySubroutineHash    = "vcl_hash"
	fastlySubroutineMiss    = "vcl_miss"
	fastlySubroutinePass    = "vcl_pass"
	fastlySubroutineError   = "vcl_error"
	fastlySubroutineFetch   = "vcl_fetch"
	fastlySubroutineDeliver = "vcl_deliver"
	fastlySubroutineLog     = "vcl_log"
)

// Lifecycle starts from RECV directive.
func (c *Runtime[T]) Lifecycle(ctx context.Context, r T) error {
	var state vintage.State = vintage.PASS
	var err error

	if vclRecv, ok := c.Subroutines[fastlySubroutineRecv]; ok {
		state, err = vclRecv(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	switch state {
	case vintage.PASS:
		if err = c.lifecycleHash(ctx, r); err != nil {
			return errors.WithStack(err)
		}
		err = c.lifecyclePass(ctx, r)
	case vintage.ERROR:
		err = c.lifecycleError(ctx, r)
	case vintage.RESTART:
		err = c.lifecycleRestart(ctx, r)
	case vintage.LOOKUP, vintage.NONE:
		if err = c.lifecycleHash(ctx, r); err != nil {
			return errors.WithStack(err)
		}
		// TODO: consider lookup cache
		err = c.lifecycleMiss(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}

	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecycleRestart(ctx context.Context, r T) error {
	c.Restarts++
	if c.Restarts > 3 {
		return errors.WithStack(
			fmt.Errorf("Max restart count exceeded. VCL runtime is limited only 3 times to restart."),
		)
	}
	return c.Lifecycle(ctx, r)
}

func (c *Runtime[T]) lifecycleHash(ctx context.Context, r T) error {
	if vclHash, ok := c.Subroutines[fastlySubroutineHash]; ok {
		if _, err := vclHash(r); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (c *Runtime[T]) lifecycleMiss(ctx context.Context, r T) error {
	var state vintage.State = vintage.FETCH
	var err error

	r.CreateBackendRequest()

	if miss, ok := c.Subroutines[fastlySubroutineMiss]; ok {
		state, err = miss(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case vintage.DELIVER_STALE:
		err = c.lifecycleDeliver(ctx, r)
	case vintage.PASS:
		err = c.lifecyclePass(ctx, r)
	case vintage.ERROR:
		err = c.lifecycleError(ctx, r)
	case vintage.FETCH, vintage.NONE:
		err = c.lifecycleFetch(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecycleHit(ctx context.Context, r T) error {
	// TODO: nothing to do because Edge runtime does have caching behavior on its runtime
	return nil
}

func (c *Runtime[T]) lifecyclePass(ctx context.Context, r T) error {
	var state vintage.State = vintage.FETCH
	var err error

	r.CreateBackendRequest()

	if vclPass, ok := c.Subroutines[fastlySubroutinePass]; ok {
		state, err = vclPass(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case vintage.PASS, vintage.NONE:
		err = c.lifecycleFetch(ctx, r)
	case vintage.ERROR:
		err = c.lifecycleError(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecycleFetch(ctx context.Context, r T) error {
	var state vintage.State = vintage.DELIVER
	var err error

	if err = r.Proxy(ctx, c.Backend); err != nil {
		return errors.WithStack(err)
	}

	if vclFetch, ok := c.Subroutines[fastlySubroutineFetch]; ok {
		state, err = vclFetch(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case vintage.DELIVER, vintage.DELIVER_STALE, vintage.NONE:
		err = c.lifecycleDeliver(ctx, r)
	case vintage.ERROR:
		err = c.lifecycleError(ctx, r)
	case vintage.RESTART:
		err = c.lifecycleRestart(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if err = r.CreateClientResponse(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *Runtime[T]) lifecycleError(ctx context.Context, r T) error {
	var state vintage.State = vintage.DELIVER
	var err error

	if vclError, ok := c.Subroutines[fastlySubroutineError]; ok {
		state, err = vclError(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case vintage.DELIVER, vintage.NONE:
		err = c.lifecycleDeliver(ctx, r)
	case vintage.RESTART:
		err = c.lifecycleRestart(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecycleDeliver(ctx context.Context, r T) error {
	var state vintage.State = vintage.LOG
	var err error

	// Time to first bytes is calculates from restart has started to vcl_deliver will call.
	// https://developer.fastly.com/reference/vcl/variables/client-response/time-to-first-byte/
	c.TimeToFirstByte = time.Since(c.RequestStartTime)
	c.RequestEndTime = time.Now()

	if vclDeliver, ok := c.Subroutines[fastlySubroutineDeliver]; ok {
		state, err = vclDeliver(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case vintage.RESTART:
		err = c.lifecycleRestart(ctx, r)
	case vintage.LOG, vintage.DELIVER, vintage.NONE:
		err = c.lifecycleLog(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecycleLog(ctx context.Context, r T) error {

	var err error
	c.ResponseHeaderBytesWritten, c.ResponseBodyBytesWritten, c.ResponseBytesWritten, err = r.WriteResponse()
	if err != nil {
		return errors.WithStack(err)
	}

	if vclLog, ok := c.Subroutines[fastlySubroutineLog]; ok {
		if _, err := vclLog(r); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
