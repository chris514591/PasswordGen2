package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[]|:;<>,.?/~")
	password := make([]rune, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
		time.Sleep(100 * time.Millisecond)
	}
	return string(password)
}

var (
	length = flag.Int("length", 10, "password length")
	count  = flag.Int("count", 3, "number of passwords to generate")
)

func main() {
	flag.Parse()

	if *length <= 0 {
		log.Fatalf("invalid password length: %d", *length)
	}

	if *length == 10 {
		for i := 0; i < *count; i++ {
			password := generatePassword(*length)
			fmt.Println(password)
		}
	} else {
		password := generatePassword(*length)
		fmt.Println(password)
	}
}

func init() {
	logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
}
