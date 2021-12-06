package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-gnss/ntrip"
	"github.com/go-gnss/rtcm/rtcm3"
)

func main() {
	caster := flag.String("caster", "https://ntrip.data.gnss.ga.gov.au/ALIC00AUS0", "NTRIP caster address, must include scheme")
	username := flag.String("username", "", "NTRIP username")
	password := flag.String("password", "", "NTRIP password")
	outputJson := flag.Bool("json", false, "output RTCM data as JSON objects instead of printing message latency")
	flag.Parse()

	req, _ := ntrip.NewClientRequest(*caster)
	req.SetBasicAuth(*username, *password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("request failed with error: %s\n", err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("request failed with response code %d\n", resp.StatusCode)
		os.Exit(1)
	}

	scanner := rtcm3.NewScanner(resp.Body)
	for msg, err := scanner.NextMessage(); err == nil; msg, err = scanner.NextMessage() {
		if *outputJson {
			out, _ := json.Marshal(msg)
			fmt.Println(string(out))
			// Don't output latency if outputting json (so stdout is parsable by jq)
			continue
		}

		if obs, ok := msg.(rtcm3.Observation); ok {
			fmt.Println(msg.Number(), time.Now().UTC().Sub(obs.Time()))
		} else {
			fmt.Println(msg.Number(), "N/A")
		}
	}
}
