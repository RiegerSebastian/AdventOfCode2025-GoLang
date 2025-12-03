package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
var wheelPosition int = 50
var zeroCounter int = 0

func main() {
	file, err := os.Open("input_Zahlenrad.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var line = scanner.Text()
		fmt.Println(line)
		if strings.HasPrefix(line, "R") {
			var numline = strings.Replace(line, "R", "", -1)
			var num int
			num, err = strconv.Atoi(numline)
			spinRight(num)
		} else if strings.HasPrefix(line, "L") {
			var numline = strings.Replace(line, "L", "", -1)
			var num int
			num, err = strconv.Atoi(numline)
			spinLeft(num)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ZeroCounter, %d\n", zeroCounter)
}

func spinRight(spinningCount int) {
	for range spinningCount {
		if wheelPosition != 99 {
			wheelPosition = wheelPosition + 1
		} else {
			wheelPosition = 0
			zeroCounter++
		}
	}
	fmt.Printf("Spin to the right, %d\n", wheelPosition)
}

func spinLeft(spinningCount int) {
	// for the first entry the 0 should not be counted as it was already counted by the last spin
	if wheelPosition == 0 {
		wheelPosition = 99
		spinningCount--
	}
	for range spinningCount {
		if wheelPosition != 0 {
			wheelPosition = wheelPosition - 1
		} else {
			wheelPosition = 99
			zeroCounter++
		}
	}
	// when after the last spine we land on 0, we should count it
	if wheelPosition == 0 {
		zeroCounter++
	}
	fmt.Printf("Spin to the left, %d\n", wheelPosition)
}