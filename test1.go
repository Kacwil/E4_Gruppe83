package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	maxMissedMessages = 3
)

type Message struct {
	Number int
	Alive  bool
}

func primary(numberChannel chan Message, exitChannel chan struct{}) {
	counter := 1
	missedMessages := 0

	for {
		select {
		case <-exitChannel:
			return
		default:
		}

		message := Message{Number: counter, Alive: true}
		numberChannel <- message

		time.Sleep(time.Second) // Simulate some work

		select {
		case <-numberChannel:
			// Backup is still alive
			missedMessages = 0
		default:
			// Backup missed a message
			missedMessages++
			if missedMessages >= maxMissedMessages {
				fmt.Println("Primary died! Promoting backup to primary.")
				return
			}
		}

		counter++
	}
}

func backup(numberChannel chan Message, exitChannel chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		message := <-numberChannel

		if !message.Alive {
			fmt.Println("Primary died! Becoming the new primary.")
			go primary(numberChannel, exitChannel)
			return
		}

		fmt.Println("Backup received:", message.Number)
		time.Sleep(time.Second) // Simulate some work
	}
}

func main() {
	numberChannel := make(chan Message)
	exitChannel := make(chan struct{})

	var wg sync.WaitGroup

	// Start the primary process
	wg.Add(1)
	go primary(numberChannel, exitChannel)

	// Start the backup process
	wg.Add(1)
	go backup(numberChannel, exitChannel, &wg)

	// Handle SIGINT (Ctrl+C) to gracefully exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		<-sigCh
		close(exitChannel)
		wg.Wait()
		os.Exit(0)
	}()

	wg.Wait()
}
