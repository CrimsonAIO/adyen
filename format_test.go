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

import "testing"

func TestFormatCardNumber(t *testing.T) {
	test := func(number, expected string) {
		if formatted := FormatCardNumber(number); formatted != expected {
			t.Fatalf("%s should be %s, instead got %s\n", number, expected, formatted)
		}
	}

	test("0123456789012345", "0123 4567 8901 2345")
	test("0123 4567 8901 2345", "0123 4567 8901 2345")
	test("0123 456789012345", "0123 4567 8901 2345")
}

func TestFormatMonthYear(t *testing.T) {
	test := func(m, y int, em, ey string) {
		if fm, fy := FormatMonthYear(m, y); fm != em || fy != ey {
			t.Fatalf("(%d, %d) should be (%s, %s), instead got (%s, %s)\n", m, y, em, ey, fm, fy)
		}
	}

	test(5, 2024, "05", "2024")
	test(12, 2024, "12", "2024")
}
