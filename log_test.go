package aliyunlog

import (
	"github.com/sirupsen/logrus"
	"testing"
)

var logs *logrus.Logger

func init() {
	logs = NewLogger("/Users/zhouyang/Code/golang/demo", "def")
}

func TestInfo(t *testing.T) {
	logs.Info("this is info ")
}

func TestWarn(t *testing.T) {
	logs.Warn("this is warn")
}

func TestError(t *testing.T) {
	logs.Error("this is error")
}

func TestFatal(t *testing.T) {
	logs.Fatal("this is Fatal")
}

func TestPanic(t *testing.T) {
	logs.Panic("this is Panic")
}
