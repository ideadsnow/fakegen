package fakegen

import "time"

var (
	DefaultRate    = 1
	DefaultMaxStep = 999
)

type options struct {
	rate    int
	maxStep int
	t       time.Time
}

type Option interface {
	apply(*options)
}

type rateOption struct {
	rate int
}

func (s rateOption) apply(opts *options) {
	opts.rate = s.rate
}

func WithRate(rate int) Option {
	if rate <= 0 {
		rate = DefaultRate
	}
	return rateOption{rate: rate}
}

type maxStepOption struct {
	maxStep int
}

func (s maxStepOption) apply(opts *options) {
	opts.maxStep = s.maxStep
}

func WithMaxStep(maxStep int) Option {
	if maxStep <= 0 {
		maxStep = DefaultMaxStep
	}
	return maxStepOption{maxStep: maxStep}
}

type timeOption struct {
	t time.Time
}

func (t timeOption) apply(opts *options) {
	opts.t = t.t
}

func WithTime(t time.Time) Option {
	return timeOption{t: t}
}
