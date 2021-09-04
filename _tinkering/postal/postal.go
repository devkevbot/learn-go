package main

import (
	"fmt"
	"strconv"
	"strings"
)

func string_to_zip() {}

func stringToZip(s string) string {

	if len(s) != 27 {
		return "Error: barcode is not 27 digits long."
	}

	if s[0] != '1' || s[len(s)-1] != '1' {
		return "Error: first or last digit in bar code is not 1."
	}

	stringGroups := []string{
		s[1:6],
		s[6:11],
		s[11:16],
		s[16:21],
		s[21:26],
	}

	var zipCode string

	for i, v := range stringGroups {

		var sum int = calculateGroupSum(v)

		if sum == -1 {
			return fmt.Sprintf("Error: invalid zipcode. Group %d doesn't contain exactly two 1s", i)
		}

		zipCode += strconv.Itoa(sum)
	}
	//fmt.Printf("s is %s\n", s)
	return zipCode
}

func calculateGroupSum(s string) int {

	numOnes := strings.Count(s, "1")

	if numOnes != 2 {
		return -1
	}

	var sum int

	for i := 0; i < len(s); i++ {

		num, _ := strconv.Atoi(string(s[i]))

		switch i {
		case 0:
			sum += 7 * num
		case 1:
			sum += 4 * num
		case 2:
			sum += 2 * num
		case 3:
			sum += num
		}
	}

	if sum == 11 {
		return 0
	}

	return sum
}

func zipToString(z string) string {

	if len(z) != 5 {
		return "Not a valid zipcode: not 5 characters long."
	}

	var retStr string

	for i := 0; i < len(z); i++ {

		var val string = convertDigitToString(string(z[i]))

		if val == "Error" {
			return val
		}

		retStr += val
	}

	return "1" + retStr + "1"

}

func convertDigitToString(z string) string {

	switch z {
	case "0":
		return "11000"
	case "1":
		return "00010"
	case "2":
		return "00100"
	case "3":
		return "00110"
	case "4":
		return "01001"
	case "5":
		return "01010"
	case "6":
		return "01100"
	case "7":
		return "10000"
	case "8":
		return "10010"
	case "9":
		return "10100"
	default:
		return "Error"
	}
}

func main() {
	//fmt.Println(stringToZip("111000001100001100101010101"))
	fmt.Println(stringToZip("110100101000101011000010011"))
	fmt.Println(zipToString("99504"))
}
