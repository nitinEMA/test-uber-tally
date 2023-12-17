package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	prom "test-uber-tally/metrics/impl"
	promreporter "github.com/uber-go/tally/v4/prometheus"
)

func main() {
	m := prom.New()
	r := promreporter.NewReporter(promreporter.Options{})

	counter := m.Counter("test_counter")
	timer := m.Timer("test_timer_summary")
	histogram := m.HistogramVec("test_histogram")

	go func() {
		for {
			counter.Inc()
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			timer.Start()
			histogram.Start()

			time.Sleep(time.Duration(rand.Float64() * float64(time.Second)))

			timer.Stop()
			histogram.Stop()
		}
	}()

	http.Handle("/metrics", r.HTTPHandler())
	fmt.Printf("Serving :8080/metrics\n")
	fmt.Printf("%v\n", http.ListenAndServe(":8080", nil))
	select {}
}