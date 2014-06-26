package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	verbose := flag.Bool("verbose", false, "verbose output")
	rate := flag.Duration("wait", 10 * time.Millisecond, "sampling interval")
	flag.Parse()
	start_time := time.Now()
	previous := start_time
	samples := 0
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			elapsed := time.Now().Sub(start_time)
			rate := elapsed / time.Duration(samples)
			fmt.Printf("%v: samples = %v, elapsed = %v, rate = %v\n", sig, samples, elapsed, rate)
		}
	}()
	for {
		sleeptime = *rate
		time.Sleep(sleeptime)
		now := time.Now()
		diff := now.Sub(previous)
		previous = now
		samples += 1
		if *verbose {
			fmt.Printf("%s: diff = %s\n", now, diff)
		}
		if diff < 0 {
			fmt.Printf("%s: diff = %s\n", now, diff)
			break
		}
	}
}
