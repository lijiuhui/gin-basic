package logger

type AppLog struct{}

func (a *AppLog) Write(p []byte) (n int, err error) {
	AppChannel <- string(p)
	return 0, nil
}
