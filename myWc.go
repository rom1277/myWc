package main

import (
	"flag"
	"fmt"
	"sync"
)

// -l для подсчета строк, -m для подсчета символов и -w для подсчета слов.
func main() {
	var l, m, w bool
	lFlag := flag.Bool("l", false, "Print symbolic links")
	mFlag := flag.Bool("m", false, "Print directories")
	wFlag := flag.Bool("w", false, "Print files")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Enter path or paths")
		return
	}
	if *lFlag {
		l = true
	}
	if *mFlag {
		m = true
	}
	if *wFlag {
		w = true
	}

	if ((m && l) || (m && w) || (l && w)) || !m && !l && !w {
		fmt.Println("Error: use one flag.")
		return
	}

	var wg sync.WaitGroup
	var path string
	for i := 0; i < flag.NArg(); i++ {
		path = flag.Arg(i)
		wg.Add(1)
		count, err := wc(path, l, m, w, &wg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(count, path)
		}
	}
	wg.Wait()
}
