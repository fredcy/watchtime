package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	previous := time.Now()
	for {
		sleeptime := time.Duration(rand.Int31n(10)) * time.Millisecond
		time.Sleep(sleeptime)
		now := time.Now()
		diff := now.Sub(previous)
		previous = now
		//fmt.Printf("%s: diff = %s\n", now, diff)
		if diff < 0 {
			fmt.Printf("%s: diff = %s\n", now, diff)
			break
		}
	}
}
