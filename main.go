package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var digits int
	var step int
	flag.IntVar(&digits, "d", 6, "the number of digits in the code")
	flag.IntVar(&step, "s", 30, "the time interval in seconds of the service in which the counter increases")
	flag.Parse()

	if secret := flag.Arg(0); secret != "" {
		code, err := get_totp_code([]byte(secret), digits, &CurrentTimestampProvider{step: step})
		if err != nil {
			log.Fatal("An error occured while parsing the key. Check if it was entered correctly!")
		}
		fmt.Printf("The code for the service is: %0*d",digits, code)
	}
}