package middleware

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"net/http"
	"time"
)

const (
	ServiceName = "service"
	FullTime    = "duration"
	URL         = "url"
	Method      = "method"
	StatusCode  = "code"
)

type writer struct {
	http.ResponseWriter
	statusCode int
}

func NewWriter(w http.ResponseWriter) *writer {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &writer{w, http.StatusOK}
}

func (w *writer) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

type MetricsMiddleware struct {
	metric    *prometheus.GaugeVec
	counter   *prometheus.CounterVec   //количество ошибок
	durations *prometheus.HistogramVec //сколько выполняются различные запросы
	errors    *prometheus.CounterVec
	name      string
}

func NewMetricsMiddleware() *MetricsMiddleware {
	return &MetricsMiddleware{}
}

func (m *MetricsMiddleware) Register(name string) {

	m.name = name
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: fmt.Sprintf("SLO for service %s", name),
		},
		[]string{
			ServiceName, URL, Method, StatusCode,
		})

	m.metric = gauge

	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hits",
			Help: "Number of all errors.",
		}, []string{URL})
	m.counter = counter

	hist := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "durations_stats",
		Help:    "durations_stats",
		Buckets: prometheus.LinearBuckets(0, 0.01, 50),
	}, []string{URL})
	m.durations = hist

	errs := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "errors_hits",
		Help: "Number of all errors.",
	}, []string{URL})

	m.errors = errs
	rand.Seed(time.Now().Unix())
	prometheus.MustRegister(m.metric)
	prometheus.MustRegister(m.counter)
	prometheus.MustRegister(m.durations)
	prometheus.MustRegister(m.errors)
}

func (m *MetricsMiddleware) LogMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		start := time.Now()

		wrapper := NewWriter(w)

		next.ServeHTTP(wrapper, r.WithContext(ctx))

		tm := time.Since(start)
		m.metric.With(prometheus.Labels{
			ServiceName: m.name,
			URL:         r.URL.Path,
			Method:      r.Method,
			StatusCode:  fmt.Sprintf("%d", wrapper.statusCode),
		}).Inc()

		m.durations.With(prometheus.Labels{URL: r.URL.Path}).Observe(tm.Seconds())

		if wrapper.statusCode != http.StatusOK {
			m.errors.With(prometheus.Labels{URL: r.URL.Path}).Inc()
		}
		m.counter.With(prometheus.Labels{URL: r.URL.Path}).Inc()
	})
}
