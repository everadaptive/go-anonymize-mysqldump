package pkg

import (
	"io"
	"sync"
)

func SetupAndProcessInput(config Config, input io.Reader) chan chan string {
	var wg sync.WaitGroup
	lines := make(chan chan string, 10)

	wg.Add(1)
	go processInput(&wg, input, lines, config)

	go func() {
		wg.Wait()
		close(lines)
	}()

	return lines
}
