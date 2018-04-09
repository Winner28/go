package multithreading

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type words struct {
	sync.Mutex
	found map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func tallyWords(fileName string, word *words) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		w := strings.ToLower(scanner.Text())
		word.add(w)
	}
	return scanner.Err()
}

func (w *words) add(word string) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = 1
		return
	}
	w.found[word] = count + 1
}

// RunRaceApp app
func RunRaceApp() {
	var wg sync.WaitGroup
	words := newWords()
	for _, file := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, words); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(file)
	}
	wg.Wait()
}

// RunSimpleEcho runs se
func RunSimpleEcho() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("Main rutine ends!")
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
