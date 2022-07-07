package main

import (
	"sync"
	"time"
)

var m = make(map[string]int)
var mutex = &sync.RWMutex{}

// 1) Use RWMutex
func Write() {
	for {
		mutex.Lock()
		m["b"] = 2
		mutex.Unlock()
	}
}

// 2) Use sync Map

func WriteSM() {
	wm := sync.Map{}
	wm.Store("c", 3)
}

func main() {
	go Write()
	go WriteSM()
	time.Sleep(1 * time.Second)
}
