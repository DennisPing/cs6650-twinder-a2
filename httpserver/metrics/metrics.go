package metrics

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/DennisPing/cs6650-twinder-a2/lib/models"
)

type Metrics struct {
	datasetName string
	apiToken    string
	ingestUrl   string
	throughput  uint64
	mutex       sync.Mutex
}

func NewMetrics() (*Metrics, error) {
	datasetName := os.Getenv("AXIOM_DATASET")
	apiToken := os.Getenv("AXIOM_API_TOKEN")
	ingestUrl := "https://api.axiom.co/v1/datasets/%s/ingest"

	if datasetName == "" || apiToken == "" {
		return nil, errors.New("you forgot to set the RABBITMQ env variables")
	}
	return &Metrics{
		datasetName: datasetName,
		apiToken:    apiToken,
		ingestUrl:   ingestUrl,
	}, nil
}

// Increment the throughput count
func (m *Metrics) IncrementThroughput() {
	m.mutex.Lock()
	m.throughput++
	m.mutex.Unlock()
}

// Return the throughput and reset the count
func (m *Metrics) GetThroughput() uint64 {
	m.mutex.Lock()
	throughput := m.throughput
	m.throughput = 0
	m.mutex.Unlock()
	return throughput
}

// Send the metrics over to Axiom
func (m *Metrics) SendMetrics() error {
	throughput := m.GetThroughput()
	payload := models.AxiomPayload{
		Time:       time.Now().Format(time.RFC3339Nano),
		Throughput: throughput,
	}

	jsonPayload, err := json.Marshal([]models.AxiomPayload{payload})
	if err != nil {
		return err
	}

	url := fmt.Sprintf(m.ingestUrl, m.datasetName)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+m.apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("%s", resp.Status)
	}

	return nil
}
