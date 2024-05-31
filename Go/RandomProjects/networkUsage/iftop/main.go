package main

import (
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"strings"
)

type Mode int

const (
	Basic Mode = iota
	Ip
)

const interval = 2

type NetworkStats struct {
	Host    string
	TxRx    string
	Last2s  string
	Last10s string
	Last40s string
}

type ConvertedStats struct {
	Host    net.IP
	TxRx    string
	Last2s  float64
	Last10s float64
	Last40s float64
}

func main() {
	mode := Ip

	for {
		stats, err := retrieveNetworkStatistics()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		Cstats := toConvertedStats(stats)

		if mode == Ip {
			for _, stat := range stats {
				printIpStats(stat)
			}
		}
		PrintStats(Cstats, "Kb") // Change "Kb" to any desired unit
	}
}

func retrieveNetworkStatistics() ([]NetworkStats, error) {
	cmd := exec.Command("sudo", "iftop", "-t", "-s", strconv.Itoa(interval), "-n", "-N")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error retrieving network statistics: %v", err)
	}

	stats := parseIftopOutput(string(out))
	return stats, nil
}

func parseIftopOutput(output string) []NetworkStats {
	lines := strings.Split(output, "\n")

	var stats []NetworkStats
	for _, line := range lines {
		fields := strings.Fields(line)

		//println(len(fields))
		//fmt.Printf("%+v\n", fields)

		switch {
		case len(fields) <= 5: //cumulative
			continue
		case len(fields) == 6:
			if parsedIP := net.ParseIP(fields[0]); parsedIP == nil {
				//Peak or total rate
				continue
			}
			fields = fields[:len(fields)-1]
			//fmt.Printf("%+v\n", fields)
		case len(fields) == 7:
			if parsedIP := net.ParseIP(fields[1]); parsedIP == nil {
				continue
			}
			fields = fields[1 : len(fields)-1]

		case len(fields) == 8:
			//Total send and receive rate
			//fmt.Printf("%+v\n", fields)
			continue
		case len(fields) > 8:
			continue
		}

		if fields[2] == "0b" && fields[3] == "0b" && fields[4] == "0b" {
			continue
		}

		txrx := "Sent"
		if fields[1] == "<=" {
			txrx = "Received"
		}

		stat := NetworkStats{
			Host:    fields[0],
			TxRx:    txrx,
			Last2s:  fields[2],
			Last10s: fields[3],
			Last40s: fields[4],
		}

		stats = append(stats, stat)
	}

	return stats
}

func toConvertedStats(stats []NetworkStats) []ConvertedStats {
	convertedStats := make([]ConvertedStats, len(stats))

	for i, stat := range stats {
		last2s := parseSizeBits(stat.Last2s)
		last10s := parseSizeBits(stat.Last10s)
		last40s := parseSizeBits(stat.Last40s)

		convertedStat := ConvertedStats{
			Host:    net.ParseIP(stat.Host),
			TxRx:    stat.TxRx,
			Last2s:  last2s,
			Last10s: last10s,
			Last40s: last40s,
		}

		convertedStats[i] = convertedStat
	}

	return convertedStats
}

func parseSizeBits(value string) float64 {
	var parsedValue float64

	switch {
	case value[len(value)-2:] == "Kb":
		pFloat, err := strconv.ParseFloat(value[:len(value)-2], 64)
		if err != nil {
			println("KB error: ", err.Error())
			return 0
		}
		parsedValue = float64(pFloat) * 1000 // Convert kilobits to bits
	case value[len(value)-2:] == "Mb":
		pFloat, err := strconv.ParseFloat(value[:len(value)-2], 64)
		if err != nil {
			println("error: ", err.Error())
			return 0
		}
		parsedValue = pFloat * (1000 * 1000) // Convert megabits to bits
	case value[len(value)-2:] == "Gb":
		pFloat, err := strconv.ParseFloat(value[:len(value)-2], 64)
		if err != nil {
			println("error: ", err.Error())
			return 0
		}
		parsedValue = pFloat * (1000 * 1000 * 1000) // Convert gigabits to bits
	case value[len(value)-1:] == "b":
		pFloat, err := strconv.ParseFloat(value[:len(value)-1], 64)
		if err != nil {
			println("error: ", err.Error())
			return 0
		}
		parsedValue = pFloat
	case value[len(value)-1:] == "B":
		pFloat, err := strconv.ParseFloat(value[:len(value)-1], 64)
		if err != nil {
			println("error: ", err.Error())
			return 0
		}
		parsedValue = pFloat * 8 // Convert bytes to bits
	}

	return parsedValue
}

func PrintStats(stats []ConvertedStats, unit string) {

	totalR2s, totalR10s, totalR40s := 0.0, 0.0, 0.0
	totalS2s, totalS10s, totalS40s := 0.0, 0.0, 0.0

	multiplier := getMultiplier(unit)

	for _, stat := range stats {
		switch stat.TxRx {
		case "Sent":
			totalS2s += stat.Last2s * multiplier
			totalS10s += stat.Last10s * multiplier
			totalS40s += stat.Last40s * multiplier
		case "Received":
			totalR2s += stat.Last2s * multiplier
			totalR10s += stat.Last10s * multiplier
			totalR40s += stat.Last40s * multiplier
		}

	}

	switch {
	case interval < 10:
		//Print the totals after processing all stats
		fmt.Printf("%-10s %9s\n", "Period:", "2s")
		fmt.Printf("%-10s %7.2f%s\n", "Sent:", totalS2s, unit)
		fmt.Printf("%-10s %7.2f%s\n", "Received:", totalR2s, unit)
	case interval >= 10 && interval < 40:
		//Print the totals after processing all stats
		fmt.Printf("%-10s %9s %9s\n", "Period:", "2s", "10s")
		fmt.Printf("%-10s %7.2f%s %7.2f%s\n",
			"Sent:", totalS2s, unit, totalS10s, unit)
		fmt.Printf("%-10s %7.2f%s %7.2f%s\n",
			"Received:", totalR2s, unit, totalR10s, unit)
	case interval >= 40:
		fmt.Printf("%-10s %9s %9s %9s\n", "Period:", "2s", "10s", "40s")
		fmt.Printf("%-10s %7.2f%s %7.2f%s %7.2f%s\n",
			"Sent:", totalS2s, unit, totalS10s, unit, totalS40s, unit)
		fmt.Printf("%-10s %7.2f%s %7.2f%s %7.2f%s\n",
			"Received:", totalR2s, unit, totalR10s, unit, totalR40s, unit)
	}
	println()
}

func getMultiplier(unit string) float64 {
	multiplier := 1.0
	switch unit {
	case "b":
		multiplier = 1.0
	case "B":
		multiplier = .125 // Convert bits to Bytes
	case "Kb":
		multiplier = .001 // Convert bits to kilobits
	case "Mb":
		multiplier = 0.000001 // Convert bits to megabits
	case "Gb":
		multiplier = 0.000000001 // Convert bits to gigabits
	}
	return multiplier
}

func printIpStats(stat NetworkStats) {
	fmt.Printf("%-10s %9s\n", "Host:", stat.Host)
	name, _ := resolveDNSName(stat.Host)
	fmt.Printf("%-10s %9s\n", "Name:", name)
	fmt.Printf("%-10s %9s\n", stat.TxRx, stat.Last2s)

}

func resolveDNSName(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if err != nil {
		return "", err
	}
	// If multiple DNS names are associated with the IP address,
	// only the first one is returned.
	return names[0], nil
}
