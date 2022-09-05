package model

// Configuration for logging
type LoggerOptions struct {
	// Enable console logging
	ConsoleLoggingEnabled bool
	// Level to log into file
	Level string
	// Filename is the name of the logfile, you can provide it with log path
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
	//Compress rolled files in .gz format
	Compress bool
}
