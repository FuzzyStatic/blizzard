package blizzard

import (
	"time"

	"github.com/avast/retry-go"
	"golang.org/x/time/rate"
)

// NewRateOpt sets the rate of which the client can do requests
// Blizzard allows 36000 per hour or up to 100 per second
//
// Example:
//
//  blizzard.NewRateOpt(1*time.Second/100, 10)
func NewRateOpt(rate time.Duration, burst int) ClientOpts {
	return &RateOpt{
		rate: rate,
		b:    burst,
	}
}

type RateOpt struct {
	rate time.Duration
	b    int
}

func (r *RateOpt) Apply(c *Client) {
	c.ratelimiter = rate.NewLimiter(
		rate.Every(r.rate),
		r.b,
	)
}

func NewRetryOpt(opts ...retry.Option) ClientOpts {
	return &RetryOpt{
		opts: opts,
	}
}

type RetryOpt struct {
	opts []retry.Option
}

func (r *RetryOpt) Apply(c *Client) {
	c.retryopts = r.opts
}
