# Balerion-Backend

## Thai Baht Converter

### **Description**

Develop a function to convert a decimal value to Thai text with a "baht" currency suffix. Please take into consideration how this code should be integrated into a service.

Requirements:
1. The input will be a decimal value using the github.com/shopspring/decimal package.
2. The output must be a string in Thai text format.
3. If the value has no fractional part, append the suffix "ถ้วน" to the result.
4. If the value has a fractional part, convert the fractional part into Thai text representing "สตางค์".


### Thought Process and Design

The development of this involved the following steps:

**Understanding Thai Number Representation:** The core challenge was to accurately represent numbers in Thai text, considering the specific rules for units (tens, hundreds, thousands, etc.) and special cases like "เอ็ด" and "ยี่สิบ".

**Decimal Handling:**  Using the `shopspring/decimal` package was crucial to handle decimal numbers precisely, avoiding the floating-point inaccuracies that can occur with Go's built-in `float` types.

**Integer and Fractional Part Conversion:** The conversion process was divided into two functions: `converInteger` for the integer part (Baht) and `convertFraction` for the fractional part (Satang).


> Functions
**`converInteger` Function:** This function iterates through the digits of the integer part, applying the rules for Thai number representation.  Special attention was given to handling the cases of "สิบ" (ten), "ยี่สิบ" (twenty), and "เอ็ด" (one).  The `thaiUnits` slice stores the Thai words for units like "ร้อย" (hundred), "พัน" (thousand), etc., and the code dynamically adds these units based on the digit's position.

**`convertFraction` Function:** This function handles the fractional part. It multiplies the fractional part by 100 to represent it as an integer (Satang) and then uses `converInteger` to convert it to Thai text.

**`DecimalToThaiBaht` Function:** This function combines the converted integer and fractional parts, adding "บาท" (Baht) and "ถ้วน" (even) or "สตางค์" (Satang) as needed.

> File Input for large amounts of data without having to rebuild the go executable

**File Input:** To process multiple numbers at once to read the file line by line, converting each line to a decimal number and then to Thai text.

> Problems encountered

**Index Out of Range Fix:** A critical bug was identified and fixed.  The original `converInteger` function could attempt to access elements outside the bounds of the `thaiUnits` slice.  A check was added to prevent this, ensuring the program's stability.

### Running the file with number.txt

> Running with text file
```bash
go run main.go numbers.txt
```

OR

> Running with build then passing text file
```bash
go build
```
THEN
```bash
./backend-takehome numbers.txt
```

This code is submitted as part of a take-home assignment for the application process at Balerion. It may not be used, modified, or redistributed without the express permission of the applicant.
