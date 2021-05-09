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
	crand "crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"math/big"
)

// client represents unique data, such as AESKey, that should persist per session.
type client struct {
	// SiteKey is the website's public RSA key used for encrypting AESKey.
	SiteKey *rsa.PublicKey

	// AESKey is a unique AES key used for encrypting content.
	AESKey []byte

	// AESNonce is the nonce used when encrypting content with AES.
	AESNonce []byte
}

// NewClient creates a new client with the specified hex encoded site key.
func NewClient(siteKey string) (*client, error) {
	result := new(client)

	// decode site key from hex
	if decoded, err := hex.DecodeString(siteKey); err != nil {
		return nil, err
	} else {
		// decode key from adyen format
		result.SiteKey = new(rsa.PublicKey)
		result.SiteKey.N = new(big.Int).SetBytes(decoded)
		result.SiteKey.E = 65537
	}

	// generate random key
	result.AESKey = make([]byte, 32)
	if _, err := crand.Read(result.AESKey); err != nil {
		return nil, err
	}

	// generate random nonce
	result.AESNonce = make([]byte, 12)
	if _, err := crand.Read(result.AESNonce); err != nil {
		return nil, err
	}

	return result, nil
}
