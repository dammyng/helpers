package interfaces

type LogWriter interface {
	LogInfo(s string)
	LogError(s string)
	LogFatal(s string)
}