package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/DennisPing/cs6650-twinder-a2/client"
	"github.com/DennisPing/cs6650-twinder-a2/datagen"
	"github.com/DennisPing/cs6650-twinder-a2/lib/logger"
	"github.com/montanaflynn/stats"
)

const (
	maxWorkers  = 100
	numRequests = 500_000
)

func main() {
	serverURL := os.Getenv("SERVER_URL")
	if serverURL == "" {
		logger.Logger.Fatal().Msg("SERVER_URL env variable not set")
	}

	port := os.Getenv("PORT") // Set the PORT to 8081 for local testing
	if port == "" {
		port = "8081" // Running in the cloud
	}

	// Health check endpoint
	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		addr := fmt.Sprintf(":%s", port)
		logger.Logger.Fatal().Msg(http.ListenAndServe(addr, nil).Error())
	}()

	ctx := context.Background()

	// Populate the task queue with tasks (token)
	taskQueue := make(chan struct{}, numRequests)
	for i := 0; i < numRequests; i++ {
		taskQueue <- struct{}{}
	}
	close(taskQueue) // Close the queue. Nothing is ever being put into the queue.

	var wg sync.WaitGroup

	responseTimes := make([][]time.Duration, maxWorkers)

	// Create a shared transport since we have concurrent clients
	sharedTransport := &http.Transport{
		MaxIdleConns:        maxWorkers,
		MaxIdleConnsPerHost: maxWorkers,
		IdleConnTimeout:     60 * time.Second,
	}

	workerPool := make([]*client.ApiClient, maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		workerPool[i] = client.NewApiClient(sharedTransport, serverURL)
	}

	log.Logger.Info().Msgf("Using %d goroutines", maxWorkers)
	log.Logger.Info().Msgf("Starting %d requests...", numRequests)
	startTime := time.Now()

	// Activate workers
	for i := 0; i < len(workerPool); i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			apiClient := workerPool[workerId]

			// Do tasks until taskQueue is empty. Then all workers can go home.
			for range taskQueue {
				ctxWithTimeout, cancelCtx := context.WithTimeout(ctx, 10*time.Second)
				direction := datagen.RandDirection(apiClient.Rng)
				t0 := time.Now()
				apiClient.SwipeLeftOrRight(ctxWithTimeout, direction) // The actual HTTP request
				t1 := time.Since(t0)
				cancelCtx()
				responseTimes[workerId] = append(responseTimes[workerId], t1) // Thread safe
			}
		}(i)
	}
	wg.Wait()

	duration := time.Since(startTime)

	// Calculate metrics
	var successCount uint64
	var errorCount uint64
	for _, worker := range workerPool {
		successCount += worker.SuccessCount
		errorCount += worker.ErrorCount
	}
	throughput := float64(successCount) / duration.Seconds()

	log.Logger.Info().Msgf("Success count: %d", successCount)
	log.Logger.Info().Msgf("Error count: %d", errorCount)
	log.Logger.Info().Msgf("Total run time: %v", duration)
	log.Logger.Info().Msgf("Throughput: %.2f req/sec", throughput)

	allResponseTimes := make([]float64, 0, numRequests)
	for _, slice := range responseTimes { // Convert all time.Duration to float64
		for _, rt := range slice {
			rtFloat := float64(rt.Milliseconds())
			allResponseTimes = append(allResponseTimes, rtFloat)
		}
	}
	mean, _ := stats.Mean(allResponseTimes)
	median, _ := stats.Median(allResponseTimes)
	p99, _ := stats.Percentile(allResponseTimes, 99)
	min, _ := stats.Min(allResponseTimes)
	max, _ := stats.Max(allResponseTimes)

	log.Logger.Info().Msgf("Mean response time: %.2f ms", mean)
	log.Logger.Info().Msgf("Median response time: %.2f ms", median)
	log.Logger.Info().Msgf("P99 response time: %.2f ms", p99)
	log.Logger.Info().Msgf("Min response time: %.2f ms", min)
	log.Logger.Info().Msgf("Max response time: %.2f ms", max)
}
