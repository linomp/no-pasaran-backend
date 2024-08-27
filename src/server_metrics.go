package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type ServerMetrics struct {
	Host        string
	Timestamp   string
	CPUUsage    string
	MemoryUsage string
}

func getStatus(r *http.Request) ServerMetrics {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	formattedTime := time.Now().UTC().Format(time.RFC3339)[11:19] + " (UTC)"
	cpuUsage, _ := cpu.Percent(0, false)
	memoryUsage, _ := mem.VirtualMemory()

	return ServerMetrics{
		Host:        ip,
		Timestamp:   formattedTime,
		CPUUsage:    fmt.Sprintf("%.2f %%", cpuUsage[0]),
		MemoryUsage: fmt.Sprintf("%.2f %%", memoryUsage.UsedPercent),
	}
}

func generateServerStatusHTML(metrics ServerMetrics) (string, error) {
	t, err := template.ParseFiles("./templates/server_metrics.html")
	if err != nil {
		return "<div>Error executing server_metrics template!</div>", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, metrics)
	if err != nil {
		return "<div>Error executing server_metrics template!</div>", err
	}

	return tpl.String(), nil
}
