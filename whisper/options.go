package whisper

// options contains Whisper service configuraiton options.
type options struct {
	// Format is a format of the transcription.
	Format string
	// InitialPrompt is optional text to provide as a prompt for the first window.
	InitialPrompt string
	// Language specifies language spoken in the audio, otherwise it will be detected automatically.
	Language string
	// MaxLineCount defines maximum lines in a single captions segment.
	MaxLineCount uint64
	// MaxLineWidth defines maximum length of the line.
	MaxLineWidth uint64
	// Translate translates transcription to English.
	Translate bool
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

// WithFormat sets desired format of the transcription.
func WithFormat(v string) Option {
	return newFuncOption(func(o *options) {
		o.Format = v
	})
}

// WithInitialPrompt sets optional text to provide as a prompt for the first window.
func WithInitialPrompt(v string) Option {
	return newFuncOption(func(o *options) {
		o.InitialPrompt = v
	})
}

// WithLanguage sets specifies language spoken in the audio.
func WithLanguage(v string) Option {
	return newFuncOption(func(o *options) {
		o.Language = v
	})
}

// WithMaxLineCount sets maximum lines in a single captions segment.
func WithMaxLineCount(v uint64) Option {
	return newFuncOption(func(o *options) {
		o.MaxLineCount = v
	})
}

// WithWithMaxLineWidth sets maximum length of the line.
func WithWithMaxLineWidth(v uint64) Option {
	return newFuncOption(func(o *options) {
		o.MaxLineWidth = v
	})
}

// WithTranslation sets whether the transcriptions should be translated to English.
func WithTranslation() Option {
	return newFuncOption(func(o *options) {
		o.Translate = true
	})
}
