package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var extended = flag.Bool("x", false, "solve for problem extension")

func readDiagnostics(r io.Reader) (data [][]byte) {
	scn := bufio.NewScanner(r)
	for scn.Scan() {
		row := []byte{}
		ln := strings.TrimSpace(scn.Text())
		for _, col := range strings.Split(ln, "") {
			v, err := strconv.Atoi(col)
			if err != nil {
				panic(err)
			}
			row = append(row, byte(v))
		}
		data = append(data, row)
	}
	if err := scn.Err(); err != nil {
		panic(err)
	}
	return data
}

// invert the structure over it's diagnoal, converting rows to columns
func transpose(data [][]byte) (rv [][]byte) {
	for y := 0; y < len(data[0]); y++ {
		row := []byte{}
		for x := 0; x < len(data); x++ {
			row = append(row, data[x][y])
		}
		rv = append(rv, row)
	}
	return rv
}

func mostCommonBit(v []byte) byte {
	var counter = map[byte]int{}
	for _, el := range v {
		counter[el] = counter[el] + 1
	}
	var (
		mcb = byte(0)
		max = 0
	)

	// accouting for equal distribution
	if counter[0] == counter[1] {
		return byte(1)
	}
	for bit, count := range counter {
		if count > max {
			max = count
			mcb = bit
		}
	}
	return mcb
}

func leastCommonBit(v []byte) byte {
	return 1 ^ mostCommonBit(v)
}

func decimal(vec []byte) (d int) {
	for i := 0; i < len(vec); i++ {
		offset := len(vec) - i - 1
		d = d | int(vec[i])<<offset
	}
	return
}

func newRateComputer(criteria func([]byte) byte) func(data [][]byte) int {
	return func(data [][]byte) int {
		data = transpose(data)
		rawRate := []byte{}
		for _, col := range data {
			b := criteria(col)
			rawRate = append(rawRate, b)
		}
		return decimal(rawRate)
	}
}

func newLifeSupportRating(criteria func([]byte) byte) func([][]byte) int {
	return func(data [][]byte) int {
		for i := 0; len(data) != 1; i++ {
			bitCriteria := criteria(transpose(data)[i])
			filtered := [][]byte{}
			for _, row := range data {
				if row[i] == bitCriteria {
					filtered = append(filtered, row)
				}
			}
			data = filtered
		}
		return decimal(data[0])
	}
}

var (
	gammaRate   = newRateComputer(mostCommonBit)
	epsilonRate = newRateComputer(leastCommonBit)
	o2Rating    = newLifeSupportRating(mostCommonBit)
	co2Rating   = newLifeSupportRating(leastCommonBit)
)

func main() {
	flag.Parse()
	data := readDiagnostics(os.Stdin)
	if *extended {
		fmt.Println(co2Rating(data) * o2Rating(data))
	} else {
		fmt.Println(gammaRate(data) * epsilonRate(data))
	}
}
