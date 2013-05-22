//Copyright 2013 Thomson Reuters Global Resources.  All Rights Reserved.  Proprietary and confidential information of TRGR.  Disclosure, use, or reproduction without written authorization of TRGR is prohibited.

package ntlm

import (
	"bytes"
	"crypto/rand"
	"unicode/utf16"
)

// Concatenate two byte slices into a new slice
func concat(ar ...[]byte) []byte {
	return bytes.Join(ar, nil)
}

// Create a 0 initialized slice of bytes
func zeroBytes(length int) []byte {
	return make([]byte, length, length)
}

func randomBytes(length int) []byte {
	randombytes := make([]byte, length)
	_, err := rand.Read(randombytes)
	if err != nil {
	} // TODO: What to do with err here
	return randombytes
}

// Zero pad the input byte slice to the given size
// bytes - input byte slice
// offset - where to start taking the bytes from the input slice
// size - size of the output byte slize
func zeroPaddedBytes(bytes []byte, offset int, size int) []byte {
	newSlice := zeroBytes(size)
	for i := 0; i < size && i+offset < len(bytes); i++ {
		newSlice[i] = bytes[i+offset]
	}
	return newSlice
}

func MacsEqual(slice1, slice2 []byte) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		// bytes between 4 and 7 (inclusive) contains random
		// data that should be ignored while comparing the
		// macs
		if (i < 4 || i > 7) && slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func utf16FromString(s string) []byte {
	encoded := utf16.Encode([]rune(s))
	// TODO: I'm sure there is an easier way to do the conversion from utf16 to bytes
	result := zeroBytes(len(encoded) * 2)
	for i := 0; i < len(encoded); i++ {
		result[i*2] = byte(encoded[i])
		result[i*2+1] = byte(encoded[i] << 8)
	}
	return result
}
