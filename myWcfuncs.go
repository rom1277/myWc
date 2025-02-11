package main

import (
	"bufio"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func wc(path string, l, m, w bool, wg *sync.WaitGroup) (int, error) {
	defer wg.Done()
	var count int
	openfile, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer openfile.Close()
	scanner := bufio.NewScanner(openfile)

	switch {
	case l:
		for scanner.Scan() {
			count += 1
		}
	case m:
		{
			var line string
			for scanner.Scan() {
				line = scanner.Text()
				count += utf8.RuneCountInString(line)
			}
		}
	case w:
		{
			var line string
			for scanner.Scan() {
				line = scanner.Text()

				count += len(strings.Fields(line))
			}
		}
	}
	return count, nil
}
