package logger

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type keyContext string

const (
	keyTrackingId keyContext = "trackingId"
)

func NewLogger(adp Adapter, ctx context.Context) *Logger {
	return &Logger{adapter: adp, ctx: ctx}
}

func (l *Logger) HandleError(err error) {
	l.adapter.Error(l.prepareMessage(err.Error()))
}

func (l *Logger) Error(msg string) {
	l.adapter.Error(l.prepareMessage(msg))
}

func (l *Logger) Warn(msg string) {
	l.adapter.Warn(l.prepareMessage(msg))
}

func (l *Logger) Info(msg string) {
	l.adapter.Info(l.prepareMessage(msg))
}

func (l *Logger) Fatal(msg string) {
	l.adapter.Fatal(l.prepareMessage(msg))
}

func (l *Logger) Debug(msg string) {
	l.adapter.Debug(l.prepareMessage(msg))
}

func (l *Logger) prepareMessage(msg string) string {
	return fmt.Sprintf("%v [trackingId: %v]", msg, l.ctx.Value(keyTrackingId))
}

func (l *Logger) Access(msg string) {
	l.adapter.Access(l.prepareMessage(msg))
}

func (l *Logger) SetTrackingId(trackingId string) {
	l.ctx = context.WithValue(l.ctx, keyTrackingId, trackingId)
}

func (l *Logger) Deprecated() *logrus.Logger {
	return l.adapter.Deprecated()
}
