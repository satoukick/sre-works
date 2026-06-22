package metric

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

// var (
// 	apiCounter = metric.ing
// )

var (
	meter metric.Meter

	APICounter metric.Int64Counter
)

func newAPICounter() {
	c, err := meter.Int64Counter(
		"api_counter",
		metric.WithDescription("api request counter"),
	)
	if err != nil {
		panic("failed to create api_counter counter: " + err.Error())
	}

	APICounter = c

}

func newMeter() {
	meter = otel.Meter("sre-works/metrics")
}

func InitMetric() {
	newMeter()

	newAPICounter()
}
