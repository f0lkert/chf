package app

import (
	chf_context "github.com/f0lkert/chf/internal/context"
	"github.com/f0lkert/chf/pkg/factory"
)

type App interface {
	SetLogEnable(enable bool)
	SetLogLevel(level string)
	SetReportCaller(reportCaller bool)

	Start()
	Terminate()

	Context() *chf_context.CHFContext
	Config() *factory.Config
}
