package storage

import (
	"fmt"
	"strconv"
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
	ch      chan bool
	i       int
}

// Writer struct represents writer
type Writer struct {
	storage *Storage
	ch      chan bool
}

func (reader *Reader) readingMessage() {
	for {
		if gotMsg := <-reader.ch; !gotMsg {
			fmt.Println(reader.name + " - " + "Waiting for a message...")
		} else {
			msg := reader.storage.getMessage()
			fmt.Println(reader.name + ": " + msg)
			reader.i++
		}
	}
}

func (w *Writer) writingMessage(msg string) {
	i := 0
	for {
		w.ch <- false
		fmt.Println("Writer writes...")
		time.Sleep(1500 * time.Millisecond)
		msg = " :a new one: " + strconv.Itoa(i)
		w.storage.setMessage(msg)
		w.ch <- true
		i++
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
	fmt.Println("Starting")
	ch := make(chan bool)
	//ch <- false
	storage := &Storage{}
	reader1 := &Reader{storage: storage, name: "Reader-1", ch: ch}
	reader2 := &Reader{storage: storage, name: "Reader-2", ch: ch}
	writer := &Writer{storage: storage, ch: ch}
	msg := "You got a new message"
	go reader1.readingMessage()
	go reader2.readingMessage()
	go writer.writingMessage(msg)

	time.Sleep(10 * time.Second)
	readers := []Reader{*reader1, *reader2}
	for _, value := range readers {
		fmt.Println(value.name + " count:" + strconv.Itoa(value.i))
	}
	fmt.Println("Main goroutine ends!")
}
