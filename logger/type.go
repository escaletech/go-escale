package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Level string

const DEBUG Level = "DEBUG"
const INFO Level = "INFO"
const WARN Level = "WARN"
const ERROR Level = "ERROR"
const FATAL Level = "FATAL"

type Adapter interface {
	Error(msg string)
	Warn(msg string)
	Info(msg string)
	Debug(msg string)
	Fatal(msg string)
	Access(msg string)
	Deprecated() *logrus.Logger
}

type Fields map[string]interface{}

type Logger struct {
	adapter Adapter
	ctx     context.Context
}

type TrackingId string
