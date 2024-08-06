package main

import (
	"fmt"
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
	// Get the client IP address
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	// Get the current timestamp in ISO 8601 format
	time := time.Now().UTC().Format(time.RFC3339)

	// Get CPU and Memory usage
	cpuUsage, _ := cpu.Percent(0, false)
	memoryUsage, _ := mem.VirtualMemory()

	return ServerMetrics{
		Host:        ip,
		Timestamp:   time,
		CPUUsage:    fmt.Sprintf("%.2f %%", cpuUsage[0]),
		MemoryUsage: fmt.Sprintf("%.2f %%", memoryUsage.UsedPercent),
	}
}

func formatTimestamp(timestamp string) string {
	return timestamp[11:19] + " (UTC)"
}

func generateServerStatusHTML(metrics ServerMetrics) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang=en>
<head>
	<meta charset=UTF-8>
	<meta name=viewport content=width=device-width, initial-scale=1.0>
	<style>
		@import url('https://fonts.googleapis.com/css?family=Open+Sans&display=swap');
		html {
			position: relative;
			overflow-x: hidden !important;
		}
		* {
			box-sizing: border-box;
		}
		body {
			font-family: 'Open Sans', sans-serif;
			color: #324e63;
		}
		a {
			color: inherit;
			text-decoration: inherit;
		}
		.wrapper {
			width: 100%%;
			width: 100%%;
			height: auto;
			min-height: 90vh;
			padding: 50px 20px;
			padding-top: 100px;
			display: flex;
		}
		.instance-card {
			width: 100%%;
			min-height: 380px;
			margin: auto;
			box-shadow: 12px 12px 2px 1px rgba(13, 28, 39, 0.4);
			background: #fff;
			border-radius: 15px;
			border-width: 1px;
			max-width: 500px;
			position: relative;
			border: thin groove #9c83ff;
		}
		.instance-card__cnt {
			margin-top: 35px;
			text-align: center;
			padding: 0 20px;
			padding-bottom: 40px;
			transition: all .3s;
		}
		.instance-card__name {
			font-weight: 700;
			font-size: 24px;
			color: #6944ff;
			margin-bottom: 15px;
		}
		.instance-card-inf__item {
			padding: 10px 35px;
			min-width: 150px;
		}
		.instance-card-inf__title {
			font-weight: 700;
			font-size: 27px;
			color: #324e63;
		}
		.instance-card-inf__txt {
			font-weight: 500;
			margin-top: 7px;
		}
		.secondary {
			color: #9c83ff;
		}
	</style>
	<title>😈️ pointless-status 😈</title>
</head>
<body>
	<div class=wrapper>
		<div class=instance-card>
			<div class=instance-card__cnt>
				<div class=instance-card__name>😈️ Server is running! 😈️</div>
				<div class=instance-card-inf>
					<div class=instance-card-inf__item>
						<div class=instance-card-inf__txt>Client</div>
						<div class=instance-card-inf__title>%s</div>
					</div>
					<div class=instance-card-inf__item>
						<div class=instance-card-inf__txt>Time</div>
						<div class=instance-card-inf__title>%s</div>
					</div>
					<div class=instance-card-inf__item>
						<div class=instance-card-inf__txt>CPU Usage</div>
						<div class=instance-card-inf__title>%s</div>
					</div>
					<div class=instance-card-inf__item>
						<div class=instance-card-inf__txt>Memory usage</div>
						<div class=instance-card-inf__title>%s</div>
					</div>
				</div>
			</div>
		</div>
</body>
</html>
	`, metrics.Host, formatTimestamp(metrics.Timestamp), metrics.CPUUsage, metrics.MemoryUsage)
}
