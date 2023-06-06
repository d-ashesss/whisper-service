package whisper

type options struct {
	Format string
}

func defaultOptions() options {
	return options{
		Format: "json",
	}
}

type Option interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{f: f}
}

func (f *funcOption) apply(o *options) {
	f.f(o)
}

func WithFormat(format string) Option {
	return newFuncOption(func(o *options) {
		o.Format = format
	})
}
