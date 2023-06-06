package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	sumFlag := flag.Bool("s", false, "Perform Sum")
	mulFlag := flag.Bool("m", false, "Perform Multiplication")
	divFlag := flag.Bool("d", false, "Perform Division")
	flag.Parse()

	if *sumFlag {
		args := flag.Args()

		mininumNumbersError(args)

		var nums []float64
		for _, v := range args {
			n, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println("Error: the values passed are not integers")
			}

			nums = append(nums, n)
		}

		var result float64
		for _, v := range nums {
			result += v
		}

		fmt.Printf("The sum result is: %.2f\n", result)
	}

	if *mulFlag {
		args := flag.Args()

		mininumNumbersError(args)

		num1, err1 := strconv.ParseFloat(args[0], 64)
		num2, err2 := strconv.ParseFloat(args[1], 64)

		if err1 != nil || err2 != nil {
			log.Fatal("Error: the values passed are not integers")
		}

		result := num1 * num2

		fmt.Printf("The multiplication result is: %.2f", result)
	}

	if *divFlag {
		args := flag.Args()

		mininumNumbersError(args)

		num1, err1 := strconv.ParseFloat(args[0], 64)
		num2, err2 := strconv.ParseFloat(args[1], 64)

		if err1 != nil || err2 != nil {
			log.Fatal("Error: the values passed are not integers")
		}

		result := num1 / num2

		fmt.Printf("The divsion result is: %.2f", result)
	}

	if !*sumFlag && !*mulFlag && !*divFlag {
		missingCommandError()
	}
}

func missingCommandError() {
	log.Fatal("You didn't select any option.")
	log.Fatal("use '-s' to execute a sum")
	log.Fatal("use '-m' to execute a multiplication")
}

func mininumNumbersError(arr []string) {
	if len(arr) < 2 {
		log.Fatal("Error: it is necessary at least two numbers")
		return
	}
}
