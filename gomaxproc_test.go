package go_routine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxProc(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread ", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total GoRuntine ", totalGoRoutine)

	group.Wait()
}

func TestChangeThreadNum(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU ", totalCpu)

	runtime.GOMAXPROCS(8)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread ", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total GoRuntine ", totalGoRoutine)

	group.Wait()
}
