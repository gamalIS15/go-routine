package go_routine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(3 * time.Second)
		channel <- "eko kurniawan"
		fmt.Println("Selesai Mengirim Data")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(7 * time.Second)
}
