package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
)

// decode_base32 decodes the input bytes that are encoded in base32.
func decode_base32(input []byte) ([]byte, error) {
	return base32.StdEncoding.DecodeString(string(input))
}

// hmac_sha_1 calculates the mac of value using sha1 and key.
func hmac_sha_1(key, value []byte) []byte{
	h := hmac.New(sha1.New, key)
	h.Write(value)
	return h.Sum(nil)
}

// dynamic_truncation performs dynamic truncation on the input bytes.
func dynamic_truncation(input []byte) (int, error) {
	if len(input) != 20 {
		return -1, fmt.Errorf("input must be exactly 20 bytes long, but is %d bytes long", len(input))
	}
	offset_bits := input[19] & 0xF // Extract last 4 bits of the last byte
	p := input[offset_bits:offset_bits+4]
	pNum := int(binary.BigEndian.Uint32(p))
	//TODO: check if plain int can be avoided
	return pNum & (0xFFFFFFFF >> 1), nil
}

// hotp_value returns the last digit digits off of value, typically 6 digits in hotp.
func hotp_value(value, digit int) int {
	return value % int(math.Pow10(digit))
}
