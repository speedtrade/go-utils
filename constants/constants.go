package constants

// Log Levels
const (
	DebugLevel  = "debug"
	InfoLevel   = "info"
	WarnLevel   = "warn"
	ErrorLevel  = "error"
	FatalLevel  = "fatal"
	PanicLevel  = "panic"
	DPanicLevel = "dpanic"
)

const (
	EnvKey                     = "env"
	EnvDefaultValue            = "dev"
	EnvUsage                   = "runtime environment"
	PortKey                    = "port"
	PortDefaultValue           = 8080
	PortUsage                  = "application port number"
	BaseConfigPathKey          = "baseconfigpath"
	BaseConfigPathDefaultValue = "resources"
	BaseConfigPathUsage        = "path to folder that stores your configurations"
)
