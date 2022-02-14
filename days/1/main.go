package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	extended = flag.Bool("x", false, "solve for problem extension")
)

func readDepths(in io.Reader) (depths []int) {
	reader := bufio.NewScanner(in)
	for reader.Scan() {
		ln := reader.Text()
		v, err := strconv.Atoi(ln)
		if err != nil {
			panic(err)
		}
		depths = append(depths, v)
	}
	if err := reader.Err(); err != nil {
		panic(err)
	}
	return
}

func transformSumWindowed(depths []int, sampleSize uint) (windows []int) {
	for i := 0; i <= len(depths)-int(sampleSize); i++ {
		sum := 0
		for j := i; j < i+int(sampleSize); j++ {
			sum += depths[j]
		}
		windows = append(windows, sum)
	}
	return
}

func main() {
	flag.Parse()
	depths := readDepths(os.Stdin)
	if *extended {
		depths = transformSumWindowed(depths, 3)
	}
	increments := countIncreaseInDepths(depths)
	fmt.Println(increments)
}

func countIncreaseInDepths(depths []int) (count int) {
	if len(depths) <= 1 {
		return
	}
	for i := 1; i < len(depths); i++ {
		c := depths[i]
		p := depths[i-1]
		if p < c {
			count++
		}
	}
	return
}
