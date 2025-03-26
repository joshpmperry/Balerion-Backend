package main

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

// error handling for negative numbers (Not sure what to do when there is a negative number)
var ErrNegativeNumber = errors.New("negative numbers are not supported")

var thaiNumbers = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var thaiUnits = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}

func convertInteger(num int64) string {
	if num == 0 {
		return thaiNumbers[0]
	}
	// Split number into millions and less than millions (remainder)
	millions := num / 1000000
	remainder := num % 1000000

	result := ""
	if millions > 0 {
		result += lessThanMillionConvertion(millions) + "ล้าน"
	}
	if remainder > 0 {
		result += lessThanMillionConvertion(remainder)
	}

	return result
}

// Handles numbers < 1 million
func lessThanMillionConvertion(num int64) string {
	if num == 0 {
		return ""
	}

	digits := []int{}
	n := num
	for n > 0 {
		digits = append([]int{int(n % 10)}, digits...)
		n /= 10
	}

	result := ""
	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		position := len(digits) - i - 1

		if digit == 0 {
			continue
		}
		// Special case for สิบ and ยี่สิบ
		if position == 1 {
			if digit == 1 {
				result += "สิบ"
				continue
			}
			if digit == 2 {
				result += "ยี่สิบ"
				continue
			}
		}

		// Special case for เอ็ด
		if position == 0 && digit == 1 && len(digits) > 1 && digits[len(digits)-2] != 0 {
			result += "เอ็ด"
		} else {
			result += thaiNumbers[digit]
		}

		if position > 0 {
			result += thaiUnits[position]
		}
	}

	return result
}

func convertFraction(num int64) string {
	if num == 0 {
		return ""
	}
	if num < 10 {
		return thaiNumbers[num] + "สตางค์"
	}
	return convertInteger(num) + "สตางค์"
}

func DecimalToThaiBaht(value decimal.Decimal) (string, error) {
	if value.IsNegative() {
		return "", ErrNegativeNumber
	}

	intPart := value.IntPart()
	fractionalPart := value.Sub(decimal.NewFromInt(intPart)).Mul(decimal.NewFromInt(100)).IntPart()

	result := convertInteger(intPart) + "บาท"
	if fractionalPart == 0 {
		result += "ถ้วน"
	} else {
		result += convertFraction(fractionalPart)
	}
	return result, nil
}

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(999999999.99), // Change this number to whatever you'd like
	}

	for _, input := range inputs {
		result, err := DecimalToThaiBaht(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("%v: %s\n", input, result)
	}
}

// Used for input files when passing through executable
// func processInputFile(filename string) error {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return fmt.Errorf("error opening file: %w", err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		line = strings.TrimSpace(line)
// 		if line == "" {
// 			continue
// 		}

// 		num, err := decimal.NewFromString(line)
// 		if err != nil {
// 			fmt.Printf("Error converting '%s': %v\n", line, err)
// 			continue
// 		}

// 		thaiText := DecimalToThaiBaht(num)
// 		fmt.Println(thaiText)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return fmt.Errorf("error reading file: %w", err)
// 	}

// 	return nil
// }

// func main() {
// 	// Input file number.txt -> change the numbers to your liking :)
// 	if len(os.Args) > 1 {
// 		filename := os.Args[1]
// 		err := processInputFile(filename)
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			os.Exit(1)
// 		}
// 	} else {
// 		fmt.Println("Usage: go run main.go <input_file>")
// 		os.Exit(1)
// 	}
// }
