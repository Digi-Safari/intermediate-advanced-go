package main

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
)

const token = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTc0MTI5OTg0NCwiaWF0IjoxNzQxMjk2ODQ0LCJyb2xlcyI6WyJhZG1pbiIsInVzZXIiXX0.p31JkHiZJrGzWJyzO7DObkkk4SkkpVhA77sZNb0WQM6143gKGva7ODc4mzIYndPp_zpDQFc3hNvCaz7sWjUGmB2_mpskgBlCWz0CIRiC8nx4ZTfRsbWFvbZpiC9Uc9OFCOXvqWcWFJ5LWaEXDTXJC-y85F2L4jeG2JgdzGDqfzqegEiwLvpvGoEVjINCGBPZqQLV__L7wV5en1CIbskr8_Y_r1TEBHs9haPWXsLM97tYJZaAHyp5k2TUGg_IqYLspWTWTQaaAKOgmIPv8EyfQ_k2ry3PRbU0i5H6RJp170kJz-WibHpvlHbu6rr1YuESvxtysHfOQFns2lWlZUotsw`

func main() {
	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		// If there's an error reading the file, print an error message and stop execution
		log.Fatalln("not able to read pem file")
	}

	// Parse the read public key to RSA public key format
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)
	if err != nil {
		// If there's an error parsing the public key, log the error and stop execution
		log.Fatalln(err)
	}

}
