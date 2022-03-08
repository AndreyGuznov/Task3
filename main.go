package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var command = map[int]string{}

func main() {
	var result string
	forCommand()
	createFile()
	for i := 2; i < howmanyRanges()-1; i++ {
		result += fmt.Sprintf("%d\n", findRanges(command, i))
	}
	err := ioutil.WriteFile(string(command[0]), []byte(result), 0777)
	if err != nil {
		fmt.Printf("Error of creating file")
	}
	// fmt.Println(result)
}

func forCommand() {
	var miniComand string
	fmt.Println("Greeting You, Example of correct command: find_primes--file testfile.txt--timeout 10--range 1:10--range 400:500..(more ranges) enter end if your command full")
	fmt.Printf("find_primes--file ")
	fmt.Scan(&miniComand)
	command[0] = miniComand
	if !(strings.Contains(string(command[0]), ".txt")) {
		fmt.Println("You must enter filename.txt PLEASE enter end and try another time")
		return
	}
	fmt.Printf("--timeout ")
	for i := 1; i < 20; i++ {
		fmt.Scan(&miniComand)
		inc := 0
		inc++
		command[i] = miniComand
		if miniComand == "" || miniComand == "end" {
			break
		}
		fmt.Printf("--range ")
	}

}

func createFile() {
	if strings.Contains(string(command[0]), ".txt") {
		file, err := os.Create(command[0])
		if err != nil {
			fmt.Printf("Error of creating file")
			return
		}
		defer file.Close()
	}
}

func findRanges(command map[int]string, i int) []int {
	str := command[i]
	var integ [2]int
	r, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		fmt.Println("Uncorrect range, try to restart programm")
	}
	matches := r.FindAllString(string(str), -1)
	for i, m := range matches {
		num, err := strconv.Atoi(m)
		if err != nil {
			fmt.Println(err)
		}
		integ[i] = num
	}
	slice := make([]int, 0)
	// fmt.Println(integ)
	if integ[0] > integ[1] {
		fmt.Println("Uncorrect range, try to restart programm")
	}

	arr := make([]bool, integ[1])
	for i := 2; i <= int(math.Sqrt(float64(integ[1])+1)); i++ {
		if arr[i] == false {
			for j := i * i; j < integ[1]; j += i {
				arr[j] = true
			}
		}
	}
	var primes []int

	for i, isComposite := range arr {
		if i > 1 && !isComposite {
			primes = append(primes, i)
		}
	}
	for _, val := range primes {
		if val >= integ[0] {
			slice = append(slice, val)
		}
	}

	return slice
}
func howmanyRanges() int {
	res := 0
	for range command {
		res++
	}
	return res
}

// fmt.Println("Uncorrect range, try to restart programm")
