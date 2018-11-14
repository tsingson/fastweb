package utils

import "fmt"

// SyncPrinter is a thread-safe printer.
// It will output given queue without
type SyncPrinter struct {
	Queue chan string
	Close chan bool
}

// NewSyncPrinter initialize a new SyncPrinter
func NewSyncPrinter() (*SyncPrinter, error) {
	return &SyncPrinter{
		Queue: make(chan string, 64),
		Close: make(chan bool, 1),
	}, nil
}

// Run a SyncPrinter instance
func (printer *SyncPrinter) Run() {
	for {
		select {
		case s := <-printer.Queue:
			fmt.Println(s)
		case <-printer.Close:
			close(printer.Close)
			close(printer.Queue)
			return
		}
	}
}
