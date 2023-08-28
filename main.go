package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

func main() {
	secretKey := "12345678901234567890"
	var timestamp int64 = time.Now().Unix()
	var counter uint64 = uint64(math.Floor(float64(timestamp) / 30.0))
	counter = 0
	digits := 6
	c := make([]byte, 8)
	binary.BigEndian.PutUint64(c, counter)
	hs := hmac_sha_1([]byte(secretKey), c) 

	pNum,_ := dynamic_truncation(hs)

	code := hotp_value(pNum, digits) 
	fmt.Printf("Step 1: Hash: %040X or %v\n", hs, hs)
	fmt.Printf("Step 2: Dynamic Truncation: %08X\n", pNum)
	fmt.Printf("Step 3: OTP Code: %06d\n", code)	
}