package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/net"
)

type TransmitRates struct {
	Rate2s  uint64 // Transmit rate observed in the last 2 seconds
	Rate10s uint64 // Transmit rate observed in the last 10 seconds
	Rate40s uint64 // Transmit rate observed in the last 40 seconds
}

func formatRate(rate uint64, unit string) string {
	switch unit {
	case "b":
		return fmt.Sprintf("%d bits/s", rate)
	case "Kb":
		return fmt.Sprintf("%.2f Kb/s", float64(rate)/1024)
	case "Mb":
		return fmt.Sprintf("%.2f Mb/s", float64(rate)/1024/1024)
	case "Gb":
		return fmt.Sprintf("%.2f Gb/s", float64(rate)/1024/1024/1024)
	default:
		return "Unknown unit"
	}
}

func printTransmitRates(unit string, transmitRates TransmitRates) {
	fmt.Println("Transmit rates observed:")
	fmt.Printf("Last 2 seconds:  %s\n", formatRate(transmitRates.Rate2s, unit))
	fmt.Printf("Last 10 seconds: %s\n", formatRate(transmitRates.Rate10s, unit))
	fmt.Printf("Last 40 seconds: %s\n", formatRate(transmitRates.Rate40s, unit))
}

func main() {
	unit := "Kb"

	// Set up channels to signal when to print the transmit rates
	print2s := time.Tick(2 * time.Second)
	print10s := time.Tick(10 * time.Second)
	print40s := time.Tick(40 * time.Second)

	var transmitRates TransmitRates
	var netStatsBefore []net.IOCountersStat

	for {
		// Retrieve network statistics after sleeping
		netStatsAfter, err := net.IOCounters(true)
		if err != nil {
			fmt.Println("Error retrieving network statistics after:", err)
			continue
		}

		// Calculate transmit rates observed in the last 2 seconds
		transmitRates.Rate2s, err = calculateRate(netStatsAfter, netStatsBefore, 2*time.Second)
		if err != nil {
			fmt.Println("Error calculating transmission rate:", err)
		}

		// Calculate transmit rates observed in the last 10 seconds
		transmitRates.Rate10s, err = calculateRate(netStatsAfter, netStatsBefore, 10*time.Second)
		if err != nil {
			fmt.Println("Error calculating transmission rate:", err)
		}

		// Calculate transmit rates observed in the last 40 seconds
		transmitRates.Rate40s, err = calculateRate(netStatsAfter, netStatsBefore, 40*time.Second)
		if err != nil {
			fmt.Println("Error calculating transmission rate:", err)
		}

		// Print transmit rates observed every 2 seconds
		select {
		case <-print2s:
			printTransmitRates(unit, transmitRates)
		default:
		}

		// Print transmit rates observed every 10 seconds
		select {
		case <-print10s:
			printTransmitRates(unit, transmitRates)
		default:
		}

		// Print transmit rates observed every 40 seconds
		select {
		case <-print40s:
			printTransmitRates(unit, transmitRates)
		default:
		}

		// Update netStatsBefore for the next iteration
		netStatsBefore = netStatsAfter
	}
}

func calculateRate(statsNow []net.IOCountersStat, statsBefore []net.IOCountersStat, duration time.Duration) (uint64, error) {
	var bytesSentNow, bytesSentBefore uint64
	for _, stat := range statsNow {
		bytesSentNow += stat.BytesSent
	}
	for _, stat := range statsBefore {
		bytesSentBefore += stat.BytesSent
	}

	// Calculate the transmission rate for the given duration
	transmitRate := (bytesSentNow - bytesSentBefore) * 8 / uint64(duration.Seconds())
	return transmitRate, nil
}
