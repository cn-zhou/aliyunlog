package aliyunlog

import (
	log "github.com/sirupsen/logrus"
)

func NewLogger(p, n string) *log.Logger {
	ll := log.New()
	ll.SetLevel(log.InfoLevel)
	ll.SetOutput(&outputDisable{})
	ll.AddHook(getFileHook(p, n))
	return ll
}

//禁止输出到控制台
type outputDisable struct {
}

func (a *outputDisable) Write(p []byte) (n int, err error) {
	return 0, nil
}
