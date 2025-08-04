package main

import (
	"cachet/internal/store"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Cachet Benchmark Suite ===\n")

	s := store.NewMemoryStore()

	benchmarkBasicOperations(s)
	benchmarkStringOperations(s)
	benchmarkConcurrency(s)
	benchmarkMemoryUsage(s)
	benchmarkLargeDataset(s)

	fmt.Println("\n=== Benchmark Complete ===")
}

func benchmarkBasicOperations(s store.Store) {
	fmt.Println("--- Basic Operations Benchmark ---")

	s.Clear()

	start := time.Now()
	for i := 0; i < 100000; i++ {
		key := fmt.Sprintf("key:%d", i)
		value := fmt.Sprintf("value:%d", i)
		s.Set(key, value)
	}
	setDuration := time.Since(start)

	start = time.Now()
	hits := 0
	for i := 0; i < 100000; i++ {
		key := fmt.Sprintf("key:%d", i)
		_, exists := s.Get(key)
		if exists {
			hits++
		}
	}
	getDuration := time.Since(start)

	start = time.Now()
	existsCount := 0
	for i := 0; i < 100000; i++ {
		key := fmt.Sprintf("key:%d", i)
		if s.Exists(key) {
			existsCount++
		}
	}
	existsDuration := time.Since(start)

	start = time.Now()
	deleted := 0
	for i := 0; i < 50000; i++ { // Delete half
		key := fmt.Sprintf("key:%d", i)
		if s.Delete(key) {
			deleted++
		}
	}
	deleteDuration := time.Since(start)

	fmt.Printf("SET 100k items:    %v (%.0f ops/sec)\n",
		setDuration, 100000.0/setDuration.Seconds())
	fmt.Printf("GET 100k items:    %v (%.0f ops/sec, %d hits)\n",
		getDuration, 100000.0/getDuration.Seconds(), hits)
	fmt.Printf("EXISTS 100k items: %v (%.0f ops/sec, %d found)\n",
		existsDuration, 100000.0/existsDuration.Seconds(), existsCount)
	fmt.Printf("DELETE 50k items:  %v (%.0f ops/sec, %d deleted)\n",
		deleteDuration, 50000.0/deleteDuration.Seconds(), deleted)
	fmt.Printf("Final store size:  %d items\n\n", s.Size())
}

func benchmarkStringOperations(s store.Store) {
	fmt.Println("--- String Operations Benchmark ---")

	s.Clear()

	start := time.Now()
	for i := 0; i < 50000; i++ {
		key := "counter"
		value, exists := s.Get(key)
		var num int64 = 0
		if exists {
			num, _ = strconv.ParseInt(value, 10, 64)
		}
		num++
		s.Set(key, strconv.FormatInt(num, 10))
	}
	incrDuration := time.Since(start)

	s.Set("text", "Hello")
	start = time.Now()
	for i := 0; i < 10000; i++ {
		value, _ := s.Get("text")
		newValue := value + "!"
		s.Set("text", newValue)
	}
	appendDuration := time.Since(start)

	finalValue, _ := s.Get("text")
	finalCounter, _ := s.Get("counter")

	fmt.Printf("INCR 50k times:     %v (%.0f ops/sec)\n",
		incrDuration, 50000.0/incrDuration.Seconds())
	fmt.Printf("APPEND 10k times:   %v (%.0f ops/sec)\n",
		appendDuration, 10000.0/appendDuration.Seconds())
	fmt.Printf("Final counter:      %s\n", finalCounter)
	fmt.Printf("Final text length:  %d chars\n\n", len(finalValue))
}

func benchmarkConcurrency(s store.Store) {
	fmt.Println("--- Concurrency Benchmark ---")

	s.Clear()

	numGoroutines := 10
	opsPerGoroutine := 10000

	start := time.Now()
	var wg sync.WaitGroup

	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for i := 0; i < opsPerGoroutine; i++ {
				key := fmt.Sprintf("g%d:key:%d", goroutineID, i)
				value := fmt.Sprintf("g%d:value:%d", goroutineID, i)
				s.Set(key, value)
			}
		}(g)
	}
	wg.Wait()
	concurrentSetDuration := time.Since(start)

	start = time.Now()
	totalReads := 0
	var readMutex sync.Mutex

	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			localReads := 0
			for i := 0; i < opsPerGoroutine; i++ {
				key := fmt.Sprintf("g%d:key:%d", goroutineID, rand.Intn(opsPerGoroutine))
				_, exists := s.Get(key)
				if exists {
					localReads++
				}
			}
			readMutex.Lock()
			totalReads += localReads
			readMutex.Unlock()
		}(g)
	}
	wg.Wait()
	concurrentGetDuration := time.Since(start)

	totalOps := numGoroutines * opsPerGoroutine
	fmt.Printf("Concurrent SET (%d goroutines): %v (%.0f ops/sec)\n",
		numGoroutines, concurrentSetDuration, float64(totalOps)/concurrentSetDuration.Seconds())
	fmt.Printf("Concurrent GET (%d goroutines): %v (%.0f ops/sec, %d hits)\n",
		numGoroutines, concurrentGetDuration, float64(totalOps)/concurrentGetDuration.Seconds(), totalReads)
	fmt.Printf("Store size after concurrent ops: %d\n\n", s.Size())
}

func benchmarkMemoryUsage(s store.Store) {
	fmt.Println("--- Memory Usage Benchmark ---")

	s.Clear()
	runtime.GC()

	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)

	numItems := 1000000
	for i := 0; i < numItems; i++ {
		key := fmt.Sprintf("memory_test_key_%d", i)
		value := fmt.Sprintf("memory_test_value_%d_with_some_longer_content_to_test_memory_usage", i)
		s.Set(key, value)
	}

	runtime.GC()
	runtime.ReadMemStats(&m2)

	memoryUsed := m2.Alloc - m1.Alloc
	memoryPerItem := float64(memoryUsed) / float64(numItems)

	fmt.Printf("Items stored:       %d\n", numItems)
	fmt.Printf("Memory used:        %d bytes (%.2f MB)\n", memoryUsed, float64(memoryUsed)/(1024*1024))
	fmt.Printf("Memory per item:    %.2f bytes\n", memoryPerItem)
	fmt.Printf("Store size:         %d\n\n", s.Size())
}

func benchmarkLargeDataset(s store.Store) {
	fmt.Println("--- Large Dataset Benchmark ---")

	s.Clear()

	dataSizes := []struct {
		name  string
		count int
		size  int
	}{
		{"Small values", 100000, 10},
		{"Medium values", 50000, 100},
		{"Large values", 10000, 1000},
	}

	for _, test := range dataSizes {
		s.Clear()

		valueTemplate := strings.Repeat("x", test.size)

		start := time.Now()
		for i := 0; i < test.count; i++ {
			key := fmt.Sprintf("large_key_%d", i)
			value := fmt.Sprintf("%s_%d", valueTemplate, i)
			s.Set(key, value)
		}
		setDuration := time.Since(start)

		start = time.Now()
		for i := 0; i < test.count; i++ {
			key := fmt.Sprintf("large_key_%d", i)
			s.Get(key)
		}
		getDuration := time.Since(start)

		start = time.Now()
		for i := 0; i < test.count/10; i++ {
			randomKey := fmt.Sprintf("large_key_%d", rand.Intn(test.count))
			s.Get(randomKey)
		}
		randomDuration := time.Since(start)

		fmt.Printf("%s (%d items, %d bytes each):\n", test.name, test.count, test.size)
		fmt.Printf("  SET: %v (%.0f ops/sec)\n", setDuration, float64(test.count)/setDuration.Seconds())
		fmt.Printf("  GET: %v (%.0f ops/sec)\n", getDuration, float64(test.count)/getDuration.Seconds())
		fmt.Printf("  Random GET: %v (%.0f ops/sec)\n", randomDuration, float64(test.count/10)/randomDuration.Seconds())
		fmt.Printf("  Store size: %d\n\n", s.Size())
	}
}
