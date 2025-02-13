package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shopspring/decimal"
)

var thaiNumbers = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var thaiUnits = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}

func convertInteger(num int64) string {
	if num == 0 {
		return thaiNumbers[0]
	}
	result := ""
	digits := []int{}
	for num > 0 {
		digits = append([]int{int(num % 10)}, digits...)
		num /= 10
	}

	n := len(digits)
	for i, digit := range digits {
		if digit == 0 {
			continue
		}
		if digit == 1 && i == n-2 && n > 1 {
			result += "สิบ"
		} else if digit == 2 && i == n-2 && n > 1 {
			result += "ยี่สิบ"
		} else if digit == 1 && i == n-1 && n > 1 && digits[i-1] != 0 {
			result += "เอ็ด"
		} else {
			result += thaiNumbers[digit]
		}

		// Check for out of bounds before accessing thaiUnits
		unitIndex := n - i - 1
		if unitIndex >= 0 && unitIndex < len(thaiUnits) && i < n-1 {
			result += thaiUnits[unitIndex]
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

func DecimalToThaiBaht(value decimal.Decimal) string {
	intPart := value.IntPart()
	fractionalPart := value.Sub(decimal.NewFromInt(intPart)).Mul(decimal.NewFromInt(100)).IntPart()

	result := convertInteger(intPart) + "บาท"
	if fractionalPart == 0 {
		result += "ถ้วน"
	} else {
		result += convertFraction(fractionalPart)
	}
	return result
}

// Used for input files when passing through executable
func processInputFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		num, err := decimal.NewFromString(line)
		if err != nil {
			fmt.Printf("Error converting '%s': %v\n", line, err)
			continue
		}

		thaiText := DecimalToThaiBaht(num)
		fmt.Println(thaiText)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

func main() {
	// Input file number.txt -> change the numbers to your liking :)
	if len(os.Args) > 1 {
		filename := os.Args[1]
		err := processInputFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
}
