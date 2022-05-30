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
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"math/big"
	"time"
)

// GenerationTimeFunc is a function responsible for returning the time that
// a payload was generated at.
type GenerationTimeFunc func() time.Time

// An Encrypter encrypts content into the Adyen format
// using an RSA public key and AES-256.
type Encrypter struct {
	// pubKey is the RSA public key to use to seal key.
	pubKey *rsa.PublicKey

	// key and block are used for AES encryption.
	// Both are set by Reset and should not be written to
	// from anywhere else.
	key   [32]byte
	block cipher.Block

	// Version is the Adyen version that this Encrypter will
	// seal plaintext for.
	Version string

	// GetGenerationTime gets the time.Time to use for the
	// required "generationtime" JSON field. The default is
	// time.Now.
	//
	// This may be modified by the caller to return custom times
	// that differ from the default.
	GetGenerationTime GenerationTimeFunc
}

// Reset resets the AES key and cipher block for the encrypter to use.
//
// If err != nil, the Encrypter is not safe to use.
func (enc *Encrypter) Reset() (err error) {
	if _, err = rand.Read(enc.key[:]); err != nil {
		return
	}

	enc.block, err = aes.NewCipher(enc.key[:])
	return
}

// NewEncrypter creates a new Encrypter with the given version and RSA public key.
//
// Calls to Encrypter.EncryptPlaintext will panic if pubKey == nil.
func NewEncrypter(version string, pubKey *rsa.PublicKey) (enc *Encrypter, err error) {
	enc = &Encrypter{pubKey: pubKey}
	enc.Version = version
	enc.GetGenerationTime = time.Now
	err = enc.Reset()
	return
}

// PubKeyFromBytes creates a new RSA public key from b with the optional public exponent.
func PubKeyFromBytes(b []byte, publicExponent ...int) *rsa.PublicKey {
	key := new(rsa.PublicKey)
	key.N = new(big.Int).SetBytes(b)

	if len(publicExponent) == 0 {
		key.E = 65537
	} else {
		key.E = publicExponent[0]
	}

	return key
}
