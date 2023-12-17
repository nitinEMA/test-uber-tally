package prom

import (
	"time"

	"test-uber-tally/metrics"
	tally "github.com/uber-go/tally/v4"
	promreporter "github.com/uber-go/tally/v4/prometheus"
)

type prom struct {
	Scope tally.Scope
}

func New() metrics.Metrics {
	r := promreporter.NewReporter(promreporter.Options{})

	scope, _ := tally.NewRootScope(tally.ScopeOptions{
		Prefix:         "ts19",
		Tags:           map[string]string{},
		CachedReporter: r,
		Separator:      promreporter.DefaultSeparator,
	}, 1*time.Second)
	
	Scope := scope.Tagged(map[string]string{})

	return &prom{
		Scope: Scope,
	}
}

type counterTally struct {
	counter tally.Counter
}

func (s *prom) Counter(name string) metrics.Counter {
	counter := s.Scope.Counter(name)
	
	return &counterTally{
		counter: counter,
	}
}

func (c *counterTally) Inc() {
	c.counter.Inc(1)
}

type histogramTally struct {
	histogram tally.Histogram
	Stopwatch tally.Stopwatch
}

func (s *prom) HistogramVec(name string) metrics.HistogramVec {
	histogram := s.Scope.Histogram(name, tally.DefaultBuckets)	

	return &histogramTally{
		histogram: histogram,
	}
}

func (h *histogramTally) Observe(value float64){
	h.histogram.RecordValue(value)
}

func (h *histogramTally) Start(){
	h.Stopwatch = h.histogram.Start()
}

func (h *histogramTally) Stop() {
	h.Stopwatch.Stop()
}

type timerTally struct {
	timer tally.Timer
	Stopwatch tally.Stopwatch
}

func (s *prom) Timer(name string) metrics.Timer {
	timer := s.Scope.Timer(name)	

	return &timerTally{
		timer: timer,
	}
}

func (t *timerTally) Start(){
	t.Stopwatch = t.timer.Start()
}

func (t *timerTally) Stop() {
	t.Stopwatch.Stop()
}
