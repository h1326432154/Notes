package main

import (
	"sync"
	"testing"
	"time"
)

func TestConter(t *testing.T) {
	conter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			conter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("conter = %d", conter)
}

func TestConterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	conter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			conter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("conter = %d", conter)
}

func TestConterWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	conter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			conter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("conter = %d", conter)
}
