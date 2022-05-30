/*
 * MIT License
 *
 * Copyright (C) 2022 Crimson Technologies, LLC. All rights reserved.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package adyen

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// FormatCardNumber formats the given card number into the Adyen format.
// Numbers less than 15 digits (excluding white space) are ignored.
//
// Examples:
//
// 0123456789012345 -> 0123 4567 8901 2345
//
// 0123 4567 8901 2345 -> (no change)
//
// 0123 456789012345 -> 0123 4567 8901 2345
func FormatCardNumber(number string) string {
	if cnt := strings.Count(number, " "); cnt == 4 {
		// we assume the number is already formatted if there are exactly 4 spaces.
		return number
	} else if cnt > 0 {
		// else if there was at least 1 space, replace them all.
		number = strings.ReplaceAll(number, " ", "")
	}
	// ignore if the number is less than 15 digits.
	if len(number) < 15 {
		return number
	}

	return number[:4] + " " + number[4:8] + " " + number[8:12] + " " + number[12:]
}

// FormatMonthYear formats a card expiry month and year into the Adyen format.
// It is assumed that the given year is the fully-qualified year,
// like "2020" (instead of "20".)
//
// Examples:
//
// 5, 2024 -> "05", "2024"
//
// 12, 2024 -> "12", "2024"
func FormatMonthYear[T time.Month | int](month T, year int) (string, string) {
	return fmt.Sprintf("%02d", month), strconv.Itoa(year)
}
