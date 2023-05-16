package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/gif"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/invzhi/sorting-visualization/sort"
)

type sorting func([]uint8, int, *gif.GIF)

const (
	minWeight = 5
)

var (
	numCPU                int
	pl                    = fmt.Println
	sortName, filename    string
	weight, height, delay int

	m = map[string]sorting{
		"1": sort.BubbleSort,
		"2": sort.MergeSort,
		"3": sort.HeapSort,
		"4": sort.QuickSort,
		"5": sort.SelectionSort,
		"6": sort.InsertionSort,
		"7": sort.ShellSort,
		"8": sort.RadixSort,
	}
)

// ---------------- Choose Algorithm  -----------------
func readAlgorithm(reader *bufio.Reader) string {
	pl("Choose algorithm: ")
	pl("1. Bubble sort")
	pl("2. Merge sort")
	pl("3. Heap sort")
	pl("4. Quick sort")
	pl("5. Selection sort")
	pl("6. Insertion sort")
	pl("7. Shell sort")
	pl("8. Radix sort")
	pl("9. All")
	pl("0. Exit")
	pl("Enter your choice: ")

	//Read string until newline
	algorithm, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You chose: " + algorithm)
	}
	return strings.TrimSpace(algorithm)
}

// ----------------- Read Delay  -----------------
func readDelay(reader *bufio.Reader) int {
	pl("Enter delay in ms: ")

	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	delay, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You entered: " + n)
	}
	return delay
}

// ----------------- Read Delay  -----------------
func readWeight(reader *bufio.Reader) int {
	pl("Enter Weight (max 256): ")

	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	weight, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You entered: " + n)
	}
	return weight
}

// ----------------- Read Delay  -----------------
func readHeight(reader *bufio.Reader) int {
	pl("Enter Height (max 256): ")

	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	height, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You entered: " + n)
	}
	return height
}

func init() {
	numCPU = runtime.NumCPU()

	const (
		defaultWeight = 256
		defaultHeight = 256
		defaultDelay  = 10

		sortUsage  = "selection, insertion, shell, merge, quick, bubble, radix, all"
		delayUsage = "successive delay times, one per frame, in 100ths of a second"
	)
	flag.StringVar(&sortName, "sorting", "", sortUsage)
	flag.StringVar(&filename, "filename", "", "GIF's filename (default: sorting name)")

	flag.IntVar(&weight, "weight", defaultWeight, "GIF's weight")
	flag.IntVar(&height, "height", defaultHeight, "GIF's height")
	flag.IntVar(&delay, "delay", defaultDelay, delayUsage)
}

func main() {
	sortName = readAlgorithm(bufio.NewReader(os.Stdin))
	delay = readDelay(bufio.NewReader(os.Stdin))
	weight = readWeight(bufio.NewReader(os.Stdin))
	height = readHeight(bufio.NewReader(os.Stdin))
	flag.Parse()

	sortf, isexist := m[sortName]
	if isexist == false && sortName != "all" {
		log.Fatalln("sorting is not existed:", sortName)
	}

	// filename
	if filename == "" {
		filename = sortName
	}
	if strings.HasSuffix(filename, ".gif") == false {
		filename += ".gif"
	}

	// weight, height
	if weight < minWeight {
		log.Fatalln("weight can not less than", minWeight)
	}
	if height <= 0 {
		log.Fatalln("height can not less than 1")
	}

	// delay
	if delay < 0 {
		log.Fatalln("delay can not less than 0")
	}

	// if sorting name is "all", generate all GIF
	if sortName != "all" {
		newGIF(sortf, filename, weight, height, delay)
	} else {
		// no matter about filename
		for name, sortf := range m {
			newGIF(sortf, name+".gif", weight, height, delay)
		}
	}
}
