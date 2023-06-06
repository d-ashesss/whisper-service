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

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}

// WithFormat specifies desired format of the transcription.
func WithFormat(f string) Option {
	return newFuncOption(func(o *options) {
		o.Format = f
	})
}
