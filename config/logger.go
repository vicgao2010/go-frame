package config

type Logger struct {
	Level    string  `yaml:"level"`
	Path     string  `yaml:"app"`
	Encoding string  `yaml:"encoding"`
	Encoder  Encoder `yaml:"encoder"`
	Rotate   Rotate  `yaml:"rotate"`
}

type Encoder struct {
	MessageKey       string `yaml:"messageKey"`
	LevelKey         string `yaml:"levelKey"`
	TimeKey          string `yaml:"timeKey"`
	NameKey          string `yaml:"nameKey"`
	CallerKey        string `yaml:"callerKey"`
	FunctionKey      string `yaml:"functionKey"`
	StacktraceKey    string `yaml:"stacktraceKey"`
	LineEnding       string `yaml:"lineEnding"`
	EncodeLevel      string `yaml:"encodeLevel"`
	EncodeTime       string `yaml:"encodeTime"`
	EncodeDuration   string `yaml:"encodeDuration"`
	EncodeCaller     string `yaml:"encodeCaller"`
	EncodeName       string `yaml:"encodeName"`
	ConsoleSeparator string `yaml:"consoleSeparator"`
}

type Rotate struct {
	MaxSize    int  `yaml:"maxSize"`
	MaxBackups int  `yaml:"maxBackups"`
	MaxAge     int  `yaml:"maxAge"`
	Compress   bool `yaml:"compress"`
}
