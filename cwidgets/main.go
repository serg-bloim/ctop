package cwidgets

import (
	"ctop/logging"
	"ctop/models"
)

var log = logging.Init()

type WidgetUpdater interface {
	SetMeta(string, string)
	SetMetrics(models.Metrics)
}
