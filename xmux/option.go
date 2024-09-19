package xmux

type Options struct {
	AllowsDuplicate bool
}

type Option func(*Options)

func WithAllowsDuplicate(allowsDuplicate bool) Option {
	return func(o *Options) {
		o.AllowsDuplicate = allowsDuplicate
	}
}
