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
	fastlySubroutineHit     = "vcl_hit"
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
		var hit bool
		hit, err = r.LookupCache()
		if err != nil {
			return errors.WithStack(err)
		}
		if hit {
			err = c.lifecycleHit(ctx, r)
		} else {
			err = c.lifecycleMiss(ctx, r)
		}
	default:
		err = fmt.Errorf("Unexpected state returned: %s in RECV", state)
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

	rh := r.CreateBackendRequest()
	c.BackendRequestHeader = NewHeader(rh)

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
		err = fmt.Errorf("Unexpected state returned: %s in MISS", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecycleHit(ctx context.Context, r T) error {
	var state vintage.State = vintage.DELIVER
	var err error

	if vclHit, ok := c.Subroutines[fastlySubroutineHit]; ok {
		state, err = vclHit(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case vintage.DELIVER, vintage.NONE:
		err = c.lifecycleDeliver(ctx, r)
	case vintage.PASS:
		err = c.lifecyclePass(ctx, r)
	case vintage.ERROR:
		err = c.lifecycleError(ctx, r)
	case vintage.RESTART:
		err = c.lifecycleRestart(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s in HIT", state)
	}

	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecyclePass(ctx context.Context, r T) error {
	var state vintage.State = vintage.PASS
	var err error

	rh := r.CreateBackendRequest()
	c.BackendRequestHeader = NewHeader(rh)

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
		err = fmt.Errorf("Unexpected state returned: %s in PASS", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecycleFetch(ctx context.Context, r T) error {
	var state vintage.State = vintage.DELIVER
	var err error
	var backendName string

	// if stored backend is a director, select backend from its type
	if c.Backend.Director != nil {
		backendName = c.Backend.Director.Backend(vintage.RequestIdentity{
			Hash:   c.RequestHash,
			Client: c.GetClientIdentity(),
		})
	} else {
		backendName = c.Backend.Name
	}

	if rh, err := r.Proxy(ctx, backendName); err != nil {
		return errors.WithStack(err)
	} else {
		c.BackendRequestHeader = NewHeader(rh)
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
		err = fmt.Errorf("Unexpected state returned: %s in FETCH", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *Runtime[T]) lifecycleError(ctx context.Context, r T) error {
	var state vintage.State = vintage.DELIVER
	var err error

	// Possibility response object is nil, then need to construct via runtime
	if c.BackendResponseHeader == nil {
		if rh, err := r.CreateObjectResponse(int(c.ObjectStatus), c.ObjectResponse); err != nil {
			return errors.WithStack(err)
		} else {
			c.BackendResponseHeader = NewHeader(rh)
		}
	}

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
		err = fmt.Errorf("Unexpected state returned: %s in ERROR", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecycleDeliver(ctx context.Context, r T) error {
	var state vintage.State = vintage.LOG
	var err error

	if err := r.SaveCache(); err != nil {
		return errors.WithStack(err)
	}

	if rh, err := r.CreateClientResponse(); err != nil {
		return errors.WithStack(err)
	} else {
		c.ResponseHeader = NewHeader(rh)
	}

	// Time to first bytes is calculated from restart has started to vcl_deliver will call.
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
		err = fmt.Errorf("Unexpected state returned: %s in DELIVER", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Runtime[T]) lifecycleLog(ctx context.Context, r T) error {
	sizes, err := r.WriteResponse()
	if err != nil {
		return errors.WithStack(err)
	}
	c.ResponseHeaderBytesWritten = sizes[0]
	c.ResponseBodyBytesWritten = sizes[1]
	c.ResponseBytesWritten = sizes[2]
	c.ResponseCompleted = true

	if vclLog, ok := c.Subroutines[fastlySubroutineLog]; ok {
		if _, err := vclLog(r); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
