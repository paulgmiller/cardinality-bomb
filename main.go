package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a new histogram with a large number of labels
	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "test_high_cardinality_histogram",
			Help:    "A histogram metric with high cardinality for testing",
			Buckets: prometheus.LinearBuckets(0, 10, 10),
		},
		[]string{"label1", "label2", "label3"},
	)

	// Register the histogram with Prometheus
	prometheus.MustRegister(histogram)

	// Generate random label values and observe data
	go func() {
		for {
			// Generate random label values
			label1 := fmt.Sprintf("value1-%d", rand.Intn(1000)) // 1000 unique values
			label2 := fmt.Sprintf("value2-%d", rand.Intn(1000)) // 1000 unique values
			label3 := fmt.Sprintf("value3-%d", rand.Intn(1000)) // 1000 unique values

			// Add a random observation to the histogram
			histogram.WithLabelValues(label1, label2, label3).Observe(float64(rand.Intn(100)))

			// Sleep to simulate normal metric updates
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Expose the metrics at /metrics
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
