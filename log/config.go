package log

var defaultLogger = []OutputConfig{
	{
		Writer: "file",
		Level:  "info",
		WriterConfig: WriterConfig{
			LogPath:    "./",
			Filename:   "default.log",
			MaxAge:     7,
			MaxSize:    10,
			MaxBackups: 10,
			Compress:   false,
		},
	},
	{
		Writer: "console",
		Level:  "info",
	},
}

// OutputConfig 输出配置
type OutputConfig struct {
	Writer       string       `yaml:"writer"`
	Level        string       `yaml:"level"`
	CallerSkip   int          `yaml:"caller_skip"`
	WriterConfig WriterConfig `yaml:"writer_config"`
}

// WriterConfig 日志写配置
type WriterConfig struct {
	LogPath    string `yaml:"log_path"`
	Filename   string `yaml:"filename"`
	MaxAge     int    `yaml:"max_age"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
}
