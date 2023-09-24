package vintage

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
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
func (c *Context[T]) Lifecycle(ctx context.Context, r T) error {
	var state State = PASS
	var err error

	if vclRecv, ok := c.Subroutines[fastlySubroutineRecv]; ok {
		state, err = vclRecv(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	switch state {
	case PASS:
		if err = c.lifecycleHash(ctx, r); err != nil {
			return errors.WithStack(err)
		}
		err = c.lifecyclePass(ctx, r)
	case ERROR:
		err = c.lifecycleError(ctx, r)
	case RESTART:
		err = c.lifecycleRestart(ctx, r)
	case LOOKUP, NONE:
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

func (c *Context[T]) lifecycleRestart(ctx context.Context, r T) error {
	c.Restarts++
	if c.Restarts > 3 {
		return errors.WithStack(
			fmt.Errorf("Max restart count exceeded. VCL runtime is limited only 3 times to restart."),
		)
	}
	return c.Lifecycle(ctx, r)
}

func (c *Context[T]) lifecycleHash(ctx context.Context, r T) error {
	if vclHash, ok := c.Subroutines[fastlySubroutineHash]; ok {
		if _, err := vclHash(r); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (c *Context[T]) lifecycleMiss(ctx context.Context, r T) error {
	var state State = FETCH
	var err error

	if miss, ok := c.Subroutines[fastlySubroutineMiss]; ok {
		state, err = miss(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case DELIVER_STALE:
		err = c.lifecycleDeliver(ctx, r)
	case PASS:
		err = c.lifecyclePass(ctx, r)
	case ERROR:
		err = c.lifecycleError(ctx, r)
	case FETCH, NONE:
		err = c.lifecycleFetch(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Context[T]) lifecycleHit(ctx context.Context, r T) error {
	// TODO: nothing to do because Edge runtime does have caching behavior on its runtime
	return nil
}

func (c *Context[T]) lifecyclePass(ctx context.Context, r T) error {
	var state State = FETCH
	var err error

	if vclPass, ok := c.Subroutines[fastlySubroutinePass]; ok {
		state, err = vclPass(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case PASS, NONE:
		err = c.lifecycleFetch(ctx, r)
	case ERROR:
		err = c.lifecycleError(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Context[T]) lifecycleFetch(ctx context.Context, r T) error {
	var state State = DELIVER
	var err error

	if err = r.Proxy(ctx, c.Backend); err != nil {
		return errors.WithStack(err)
	}

	c.RequestEndTime = time.Now()

	if vclFetch, ok := c.Subroutines[fastlySubroutineFetch]; ok {
		state, err = vclFetch(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case DELIVER, DELIVER_STALE, NONE:
		err = c.lifecycleDeliver(ctx, r)
	case ERROR:
		err = c.lifecycleError(ctx, r)
	case RESTART:
		err = c.lifecycleRestart(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Context[T]) lifecycleError(ctx context.Context, r T) error {
	var state State = DELIVER
	var err error

	if vclError, ok := c.Subroutines[fastlySubroutineError]; ok {
		state, err = vclError(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case DELIVER, NONE:
		err = c.lifecycleDeliver(ctx, r)
	case RESTART:
		err = c.lifecycleRestart(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Context[T]) lifecycleDeliver(ctx context.Context, r T) error {
	var state State = LOG
	var err error

	if vclDeliver, ok := c.Subroutines[fastlySubroutineDeliver]; ok {
		state, err = vclDeliver(r)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	switch state {
	case RESTART:
		err = c.lifecycleRestart(ctx, r)
	case LOG, DELIVER, NONE:
		err = c.lifecycleLog(ctx, r)
	default:
		err = fmt.Errorf("Unexpected state returned: %s", state)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *Context[T]) lifecycleLog(ctx context.Context, r T) error {
	if vclLog, ok := c.Subroutines[fastlySubroutineLog]; ok {
		if _, err := vclLog(r); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
