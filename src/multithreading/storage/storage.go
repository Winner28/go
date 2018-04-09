package storage

import (
	"fmt"
	"sync"
	"time"
)

// Storage struct represents storage
type Storage struct {
	sync.Mutex
	message string
}

// Reader struct represents reader
type Reader struct {
	storage *Storage
	name    string
}

// Writer struct represents writer
type Writer struct {
	storage *Storage
}

func (reader *Reader) readingMessage() {
	for {
		time.Sleep(500 * time.Millisecond)
		msg := reader.storage.getMessage()
		fmt.Println(reader.name + ": " + msg)

	}
}

func (s Storage) isEmpty() bool {
	return len(s.message) == 0
}

func (w *Writer) writingMessage(msg string) {
	for {
		time.Sleep(time.Second)
		w.storage.setMessage(msg)
	}
}

func (s *Storage) setMessage(msg string) {
	s.message = msg
}

func (s *Storage) getMessage() string {
	return s.message
}

// RunStorageProg runs storage prog
func RunStorageProg() {
	storage := &Storage{}
	reader := &Reader{storage: storage, name: "Writer-1"}
	writer := &Writer{storage}
	msg := "new message"

	go reader.readingMessage()
	go writer.writingMessage(msg)
	time.Sleep(5 * time.Second)
	time.Sleep(65 * time.Second)
	fmt.Println("Main goroutine ends!")
}
