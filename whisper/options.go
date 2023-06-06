package whisper

// options contains Whisper service configuraiton options.
type options struct {
	Format string
}

// defaultOptions creates options set with default values.
func defaultOptions() options {
	return options{
		Format: "json",
	}
}

// Option is a configuration option for Whisper service.
type Option interface {
	// apply applies value of an option to options container.
	apply(*options)
}

// funcOption creates an option from a function.
type funcOption struct {
	f func(*options)
}

// newFuncOption creates new funcOption.
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
