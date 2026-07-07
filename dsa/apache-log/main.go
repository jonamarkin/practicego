package main

/**
 *Apache Log Parser
 Given a list of logs with IP addresses in the following format.
lines = ["10.0.0.1 - GET 2020-08-24", "10.0.0.1 - GET 2020-08-24", "10.0.0.2 - GET 2020-08-20"]

Return the most frequent IP address from the logs.
The retuned IP address value must be in a string format.
If multiple IP addresses have the count equal to max count,
then return the address as a comma-separated string with IP addresses in sorted order.

**/

import (
	"fmt"
	"sort"
	"strings"
)

func MostFrequentIP(logs []string) string {
	ipCount := make(map[string]int)

	for _, log := range logs {
		parts := strings.Split(log, " ")
		if len(parts) > 0 {
			ip := parts[0]
			ipCount[ip]++
		}
	}

	maxCount := 0
	for _, count := range ipCount {
		if count > maxCount {
			maxCount = count
		}
	}

	var mostFrequentIPs []string
	for ip, count := range ipCount {
		if count == maxCount {
			mostFrequentIPs = append(mostFrequentIPs, ip)
		}
	}

	sort.Strings(mostFrequentIPs)
	return strings.Join(mostFrequentIPs, ",")
}

func main() {
	logs := []string{
		"10.0.0.1 - GET 2020-08-24",
		"10.0.0.1 - GET 2020-08-24",
		"10.0.0.2 - GET 2020-08-20",
	}
	result := MostFrequentIP(logs)
	fmt.Println(result)
}
