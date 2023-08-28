package main

import (
	"encoding/binary"
	"math"
	"time"
)

// CounterProvider is a function that returns a uint64. It is used to dynamically provide
// the counter value to the totp function, e.g. the current timestamp or a fixed int counter value.
type CounterProvider interface {
	provide() uint64
}

type CurrentTimestampProvider struct {
	step int 
}

func (c *CurrentTimestampProvider) provide() uint64 {
	var timestamp int64 = time.Now().Unix()
	var counter uint64 = uint64(math.Floor(float64(timestamp) / float64(c.step)))
	return counter
}


// get_totp_code calculates the totp code of the key given the digit parameter.
// It uses the hotp algorithm with the value provided by to counter_provider function as a counter.
func get_totp_code(key []byte, digits int, provider CounterProvider) (int, error) {
		var counter uint64 = provider.provide()
		c := make([]byte, 8)
		binary.BigEndian.PutUint64(c, counter)
		secretBytes, err := decode_base32(key)
		if err != nil {
			return -1, err
		}
		hs := hmac_sha_1([]byte(secretBytes), c) 

		pNum,_ := dynamic_truncation(hs)

		code := hotp_value(pNum, digits)
		return code, nil
}