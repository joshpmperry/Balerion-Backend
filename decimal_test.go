package main

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

// TestDecimalToThaiBaht covers various scenarios for decimal to Thai text conversion
func TestDecimalToThaiBaht(t *testing.T) {
	testCases := []struct {
		name        string
		input       decimal.Decimal
		expected    string
		expectedErr error
	}{
		// Negative number cases
		{
			name:        "Negative One",
			input:       decimal.NewFromInt(-1),
			expected:    "",
			expectedErr: ErrNegativeNumber,
		},
		{
			name:        "Negative with Decimals",
			input:       decimal.NewFromFloat(-999.99),
			expected:    "",
			expectedErr: ErrNegativeNumber,
		},
		{
			name:        "Negative Zero",
			input:       decimal.NewFromFloat(-0.00),
			expected:    "ศูนย์บาทถ้วน",
			expectedErr: nil,
		},

		// Regular cases
		{
			name:        "Zero",
			input:       decimal.NewFromInt(0),
			expected:    "ศูนย์บาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Zero with decimal",
			input:       decimal.NewFromFloat(0.0),
			expected:    "ศูนย์บาทถ้วน",
			expectedErr: nil,
		},

		// Single digits (1-9)
		{
			name:        "One",
			input:       decimal.NewFromInt(1),
			expected:    "หนึ่งบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Five",
			input:       decimal.NewFromInt(5),
			expected:    "ห้าบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Nine",
			input:       decimal.NewFromInt(9),
			expected:    "เก้าบาทถ้วน",
			expectedErr: nil,
		},

		// Teens (10-19)
		{
			name:        "Ten",
			input:       decimal.NewFromInt(10),
			expected:    "สิบบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Eleven",
			input:       decimal.NewFromInt(11),
			expected:    "สิบเอ็ดบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Fifteen",
			input:       decimal.NewFromInt(15),
			expected:    "สิบห้าบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Nineteen",
			input:       decimal.NewFromInt(19),
			expected:    "สิบเก้าบาทถ้วน",
			expectedErr: nil,
		},

		// Multiples of 10 (20-90)
		{
			name:        "Twenty",
			input:       decimal.NewFromInt(20),
			expected:    "ยี่สิบบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Thirty",
			input:       decimal.NewFromInt(30),
			expected:    "สามสิบบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Ninety",
			input:       decimal.NewFromInt(90),
			expected:    "เก้าสิบบาทถ้วน",
			expectedErr: nil,
		},

		// Two digits with ones (21-99)
		{
			name:        "Twenty One",
			input:       decimal.NewFromInt(21),
			expected:    "ยี่สิบเอ็ดบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Forty Five",
			input:       decimal.NewFromInt(45),
			expected:    "สี่สิบห้าบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Ninety Nine",
			input:       decimal.NewFromInt(99),
			expected:    "เก้าสิบเก้าบาทถ้วน",
			expectedErr: nil,
		},

		// Hundreds
		{
			name:        "One Hundred",
			input:       decimal.NewFromInt(100),
			expected:    "หนึ่งร้อยบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "One Hundred One",
			input:       decimal.NewFromInt(101),
			expected:    "หนึ่งร้อยหนึ่งบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "One Hundred Ten",
			input:       decimal.NewFromInt(110),
			expected:    "หนึ่งร้อยสิบบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "One Hundred Eleven",
			input:       decimal.NewFromInt(111),
			expected:    "หนึ่งร้อยสิบเอ็ดบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Five Hundred Fifty Five",
			input:       decimal.NewFromInt(555),
			expected:    "ห้าร้อยห้าสิบห้าบาทถ้วน",
			expectedErr: nil,
		},

		// Thousands
		{
			name:        "One Thousand",
			input:       decimal.NewFromInt(1000),
			expected:    "หนึ่งพันบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Two Thousand Twenty",
			input:       decimal.NewFromInt(2020),
			expected:    "สองพันยี่สิบบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Nine Thousand Nine Hundred Ninety Nine",
			input:       decimal.NewFromInt(9999),
			expected:    "เก้าพันเก้าร้อยเก้าสิบเก้าบาทถ้วน",
			expectedErr: nil,
		},

		// Ten Thousands
		{
			name:        "Ten Thousand",
			input:       decimal.NewFromInt(10000),
			expected:    "หนึ่งหมื่นบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Fifty Five Thousand Five Hundred Fifty Five",
			input:       decimal.NewFromInt(55555),
			expected:    "ห้าหมื่นห้าพันห้าร้อยห้าสิบห้าบาทถ้วน",
			expectedErr: nil,
		},

		// Hundred Thousands
		{
			name:        "One Hundred Thousand",
			input:       decimal.NewFromInt(100000),
			expected:    "หนึ่งแสนบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Nine Hundred Ninety Nine Tand Nine Hundred Ninety Nine",
			input:       decimal.NewFromInt(999999),
			expected:    "เก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทถ้วน",
			expectedErr: nil,
		},

		// Millions
		{
			name:        "One Million",
			input:       decimal.NewFromInt(1000000),
			expected:    "หนึ่งล้านบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Ten Million",
			input:       decimal.NewFromInt(10000000),
			expected:    "สิบล้านบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "One Hundred Million",
			input:       decimal.NewFromInt(100000000),
			expected:    "หนึ่งร้อยล้านบาทถ้วน",
			expectedErr: nil,
		},

		// Decimal cases
		{
			name:        "Zero Point One",
			input:       decimal.NewFromFloat(0.1),
			expected:    "ศูนย์บาทสิบสตางค์",
			expectedErr: nil,
		},
		{
			name:        "Zero Point Zero One",
			input:       decimal.NewFromFloat(0.01),
			expected:    "ศูนย์บาทหนึ่งสตางค์",
			expectedErr: nil,
		},
		{
			name:        "Zero Point Zero Zero One",
			input:       decimal.NewFromFloat(0.001),
			expected:    "ศูนย์บาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "One Point Two Three",
			input:       decimal.NewFromFloat(1.23),
			expected:    "หนึ่งบาทยี่สิบสามสตางค์",
			expectedErr: nil,
		},
		{
			name:        "One Hundred Point Nine Nine",
			input:       decimal.NewFromFloat(100.99),
			expected:    "หนึ่งร้อยบาทเก้าสิบเก้าสตางค์",
			expectedErr: nil,
		},

		// Edge cases
		{
			name:        "Nine Hundred Ninety Nine Million Nine Hundred Ninety Nine Thousand Nine Hundred Ninety Nine Point Nine Nine",
			input:       decimal.NewFromFloat(999999999.99),
			expected:    "เก้าร้อยเก้าสิบเก้าล้านเก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทเก้าสิบเก้าสตางค์",
			expectedErr: nil,
		},
		{
			name:        "One Billion",
			input:       decimal.NewFromFloat(1000000000),
			expected:    "หนึ่งพันล้านบาทถ้วน",
			expectedErr: nil,
		},

		// Special decimal cases
		{
			name:        "Repeating Decimal",
			input:       decimal.NewFromFloat(1.333333),
			expected:    "หนึ่งบาทสามสิบสามสตางค์",
			expectedErr: nil,
		},
		{
			name:        "Long Decimal",
			input:       decimal.NewFromFloat(1.23456789),
			expected:    "หนึ่งบาทยี่สิบสามสตางค์",
			expectedErr: nil,
		},
		{
			name:        "Half Baht",
			input:       decimal.NewFromFloat(0.50),
			expected:    "ศูนย์บาทห้าสิบสตางค์",
			expectedErr: nil,
		},

		// Special number patterns
		{
			name:        "Repeated Digits",
			input:       decimal.NewFromInt(111111),
			expected:    "หนึ่งแสนหนึ่งหมื่นหนึ่งพันหนึ่งร้อยสิบเอ็ดบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Ascending Digits",
			input:       decimal.NewFromInt(123456),
			expected:    "หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหกบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Descending Digits",
			input:       decimal.NewFromInt(987654),
			expected:    "เก้าแสนแปดหมื่นเจ็ดพันหกร้อยห้าสิบสี่บาทถ้วน",
			expectedErr: nil,
		},

		// Additional edge cases
		{
			name:        "Large Number With Many Zeros",
			input:       decimal.NewFromInt(100000001),
			expected:    "หนึ่งร้อยล้านหนึ่งบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Sequential Zeros",
			input:       decimal.NewFromInt(10001),
			expected:    "หนึ่งหมื่นหนึ่งบาทถ้วน",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := DecimalToThaiBaht(tc.input)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result,
				"Test case %s failed: expected %s, got %s",
				tc.name, tc.expected, result)
		})
	}
}

// TestDecimalToThaiBahtPrecision checks handling of different decimal precisions
func TestDecimalToThaiBahtPrecision(t *testing.T) {
	precisionCases := []struct {
		name        string
		input       decimal.Decimal
		expected    string
		expectedErr error
	}{
		{
			name:        "High precision decimal",
			input:       decimal.NewFromFloat(100.0001),
			expected:    "หนึ่งร้อยบาทถ้วน",
			expectedErr: nil,
		},
		{
			name:        "Rounded precision",
			input:       decimal.NewFromFloat(42.123456),
			expected:    "สี่สิบสองบาทสิบสองสตางค์",
			expectedErr: nil,
		},
	}

	for _, tc := range precisionCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := DecimalToThaiBaht(tc.input)
			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
