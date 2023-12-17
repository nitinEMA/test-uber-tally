package metrics

type Metrics interface {
	Counter(name string) Counter
	Timer(name string) Timer
	HistogramVec(name string) HistogramVec
}

type Counter interface {
	Inc()
}

type HistogramVec interface {
	Start()
	Stop()
	Observe(value float64)
}

type Timer interface {
	Start()
	Stop()
}