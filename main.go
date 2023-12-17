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


















// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"test-uber-tally/metrics"

// 	tally "github.com/uber-go/tally/v4"
// 	promreporter "github.com/uber-go/tally/v4/prometheus"
// )

// type prom struct {
// 	Scope tally.Scope
// }

// func New() metrics.Metrics {
//     r := promreporter.NewReporter(promreporter.Options{})

//     scope, _ := tally.NewRootScope(tally.ScopeOptions{
//         Prefix:         "ts17",
//         Tags:           map[string]string{}, // Set the default tags for the root scope
//         CachedReporter: r,
//         Separator:      promreporter.DefaultSeparator,
//     }, 1*time.Second)
    
//     return prom{
//         Scope: scope,
//     }
// }

// type counterTally struct {
// 	counter tally.Counter
// }

// func (s prom) Counter(name string) metrics.Counter {
// 	counter := s.Scope.Counter(name)
	
// 	return counterTally{
// 		counter: counter,
// 	}
// }

// func (c counterTally) Inc() {
// 	c.counter.Inc(1)
// }

// func main() {
// 	m := New()

// 	// scope, closer := tally.NewRootScope(tally.ScopeOptions{
// 	// 	Prefix:         "test_run_05",
// 	// 	Tags:           map[string]string{},
// 	// 	CachedReporter: r,
// 	// 	Separator:      promreporter.DefaultSeparator,
// 	// }, 1*time.Second)
// 	// defer closer.Close()

// 	// counter := scope.Tagged(map[string]string{}).Counter("test_counter")

// 	counter := m.Counter("test_counter")

// 	go func() {
// 		for {
// 			counter.Inc()
// 			time.Sleep(time.Second)
// 		}
// 	}()

// 	r := promreporter.NewReporter(promreporter.Options{})
// 	http.Handle("/metrics", r.HTTPHandler())
// 	fmt.Printf("Serving :8080/metrics\n")
// 	fmt.Printf("%v\n", http.ListenAndServe(":8080", nil))
// 	select {}
// }










// package main

// import (
//     "fmt"
//     "net/http"
//     "time"

//     tally "github.com/uber-go/tally/v4"
//     promreporter "github.com/uber-go/tally/v4/prometheus"
// )


// //////////// METRICS INTERFACE //////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// type Metrics interface {
// 	Counter(name string) Counter
// }

// type Counter interface {
// 	Inc()
// }

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// var testName string = "ts15"

// type prom struct {
//     Scope tally.Scope
// }

// func New() Metrics {
// 	r := promreporter.NewReporter(promreporter.Options{})

// 	scope, _ := tally.NewRootScope(tally.ScopeOptions{
// 		Prefix:         testName,
// 		Tags:           map[string]string{},
// 		CachedReporter: r,
// 		Separator:      promreporter.DefaultSeparator,
// 	}, 1*time.Second)

// 	Scope := scope.Tagged(map[string]string{})

// 	return prom{
// 		Scope: Scope,
// 		// Do not defer closer.Close() here
// 	}
// }

// type counterTally struct {
//     counter tally.Counter
// }

// func (s prom) Counter(name string) Counter {
//     counter := s.Scope.Counter(name)
    
//     return counterTally{
//         counter: counter,
//     }
// }

// func (c counterTally) Inc() {
//     c.counter.Inc(1)
// }

// func main() {

// 	r := promreporter.NewReporter(promreporter.Options{})
// 	m := New()
//     counter := m.Counter("test_counter")

//     go func() {
//         for {
// 			// counter.Inc(1)
//             counter.Inc()
//             time.Sleep(time.Second)
//         }
//     }()

//     http.Handle("/metrics", r.HTTPHandler())
//     fmt.Printf("Serving :8080/metrics\n")
//     fmt.Printf("%v\n", http.ListenAndServe(":8080", nil))
//     select {}
// }