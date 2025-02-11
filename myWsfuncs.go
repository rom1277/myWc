package main

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

func ws(path string, l, m, w bool, wg *sync.WaitGroup) (int, error) {
	defer wg.Done()
	var count int
	openfile, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer openfile.Close()
	scanner := bufio.NewScanner(openfile)
	var line string
	switch {
	case l:
		for scanner.Scan() {
			count += 1
		}
	case m:
		{
			for scanner.Scan() {
				line = scanner.Text()
				count += len(line)
			}
		}
	case w:
		{
			for scanner.Scan() {
				line = scanner.Text()

				count += len(strings.Fields(line))
			}
		}
	}
	return count, nil
}
