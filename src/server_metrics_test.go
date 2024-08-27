package main

import (
	"log"
	"strings"
	"testing"
)

func TestGenerateServerStatusHTML(t *testing.T) {
	metrics := ServerMetrics{
		Host:        "localhost",
		Timestamp:   "2024-08-27T12:34:56Z",
		CPUUsage:    "15%",
		MemoryUsage: "45%",
	}

	output, err := generateServerStatusHTML(metrics)
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(output, metrics.Host) {
		t.Errorf("Output does not contain Host: %s", metrics.Host)
	}

	if !strings.Contains(output, metrics.Timestamp) {
		t.Errorf("Output does not contain Timestamp: %s", metrics.Timestamp)
	}

	if !strings.Contains(output, metrics.CPUUsage) {
		t.Errorf("Output does not contain CPUUsage: %s", metrics.CPUUsage)
	}

	if !strings.Contains(output, metrics.MemoryUsage) {
		t.Errorf("Output does not contain MemoryUsage: %s", metrics.MemoryUsage)
	}
}
