package command

type DebugLog interface {
	EnableDebug() bool
	PrintLog(key string, value ...any)
}
