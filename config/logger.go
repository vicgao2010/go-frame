package config

type Logger struct {
	Encoder  Encoder `yaml:"encoder,omitempty"`
	Level    string  `yaml:"level,omitempty"`
	Path      string  `yaml:"app,omitempty"`
	Encoding string  `yaml:"encoding,omitempty"`
}

type Encoder struct {
	MessageKey       string `yaml:"message-key,omitempty"`
	LevelKey         string `yaml:"level-key,omitempty"`
	TimeKey          string `yaml:"time-key,omitempty"`
	NameKey          string `yaml:"name-key,omitempty"`
	CallerKey        string `yaml:"caller-key,omitempty"`
	FunctionKey      string `yaml:"function-key,omitempty"`
	StacktraceKey    string `yaml:"stacktrace-key,omitempty"`
	LineEnding       string `yaml:"line-ending,omitempty"`
	EncodeLevel      string `yaml:"level-encoder,omitempty"`
	EncodeTime       string `yaml:"time-encoder,omitempty"`
	EncodeDuration   string `yaml:"duration-encoder,omitempty"`
	EncodeCaller     string `yaml:"caller-encoder,omitempty"`
	EncodeName       string `yaml:"name-encoder,omitempty"`
	ConsoleSeparator string `yaml:"console-separator,omitempty"`
}

type Rotate struct {
	MaxSize    int  `yaml:"max-size,omitempty"`
	MaxBackups int  `yaml:"max-backups,omitempty"`
	MaxAge     int  `yaml:"max-age,omitempty"`
	Compress   bool `yaml:"compress,omitempty"`
}
