package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	isPublishing bool
	stopCh       chan struct{}
	mu           sync.Mutex
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go handleUserInput()
	<-sigs
	fmt.Println("exiting")
}
