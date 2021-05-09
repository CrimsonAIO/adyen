/*
 * MIT License
 *
 * Copyright (C) 2021 Crimson Technologies LLC. All rights reserved.
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
	"testing"
)

// testPublicKey is a public key to run tests with.
const testPublicKey = "A237060180D24CDEF3E4E27D828BDB6A13E12C6959820770D7F2C1671DD0AEF4729670C20C6C5967C664D18955058B69549FBE8BF3609EF64832D7C033008A818700A9B0458641C5824F5FCBB9FF83D5A83EBDF079E73B81ACA9CA52FDBCAD7CD9D6A337A4511759FA21E34CD166B9BABD512DB7B2293C0FE48B97CAB3DE8F6F1A8E49C08D23A98E986B8A995A8F382220F06338622631435736FA064AEAC5BD223BAF42AF2B66F1FEA34EF3C297F09C10B364B994EA287A5602ACF153D0B4B09A604B987397684D19DBC5E6FE7E4FFE72390D28D6E21CA3391FA3CAADAD80A729FEF4823F6BE9711D4D51BF4DFCB6A3607686B34ACCE18329D415350FD0654D"

// TestEncryption encrypts a card number.
func TestEncryption(t *testing.T) {
	c, err := NewClient(testPublicKey)
	if err != nil {
		panic(err)
	}

	// encrypt
	encrypted, err := c.Encrypt(Version1_18, "number", "5246510090348187")
	if err != nil {
		panic(err)
	}

	t.Log(encrypted)
}

// BenchmarkEncryption benchmarks how long it takes to encrypt a card number, re-using the same client.
func BenchmarkEncryption(b *testing.B) {
	// don't time while creating client.
	b.StopTimer()
	c, err := NewClient(testPublicKey)
	if err != nil {
		panic(err)
	}

	// start benchmark.
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err = c.Encrypt(Version1_18, "number", "5246510090348187"); err != nil {
			panic(err)
		}
	}
}
