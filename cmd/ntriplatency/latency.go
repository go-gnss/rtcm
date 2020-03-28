package main

import (
	"flag"
	"fmt"
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
	"github.com/umeat/go-ntrip/ntrip"
	"time"
)

func main() {
	caster := flag.String("caster", "http://auscors.ga.gov.au:2101/ALIC7", "NTRIP caster mountpoint to stream from")
	username := flag.String("username", "", "NTRIP username")
	password := flag.String("password", "", "NTRIP password")
	flag.Parse()

	client, err := ntrip.NewClient(*caster)
	client.SetBasicAuth(*username, *password)
	resp, err := client.Connect()
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode, err)
	}

	scanner := rtcm3.NewScanner(resp.Body)
	for msg, err := scanner.Next(); err == nil; msg, err = scanner.Next() {
		if obs, ok := msg.(rtcm3.Observable); ok {
			fmt.Println(msg.Number(), time.Now().UTC().Sub(obs.Time()))
		} else {
			fmt.Println(msg.Number())
		}
		fmt.Printf("%+v\n\n", msg)
	}
}
