package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func generatePassword(passLength int) string {
	rand.Seed(time.Now().UnixNano())
	characters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[]|:;<>,.?/~")
	password := make([]rune, passLength)
	for i := range password {
		password[i] = characters[rand.Intn(len(characters))]
		time.Sleep(1 * time.Millisecond)
	}
	return string(password)
}

var (
	length = flag.Int("length", 12, "password length")
	count  = flag.Int("count", 1, "number of passwords to generate")
	user   = flag.String("user", "", "user name")
)

func main() {
	logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	flag.Parse()

	if *length <= 0 {
		log.Fatalf("invalid password length: %d", *length)
	}

	if *count <= 0 {
		log.Fatalf("invalid password count: %d", *count)
	}

	if *user == "" {
		*user = fmt.Sprintf("user_%d", time.Now().Unix())
	}

	f, err := os.OpenFile("test123.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if *length == 12 {
		for i := 0; i < *count; i++ {
			password := generatePassword(*length)
			line := fmt.Sprintf("%s: %s\n", *user, password)
			if _, err := f.WriteString(line); err != nil {
				log.Fatal(err)
			}
			fmt.Println(password)
		}
	} else {
		password := generatePassword(*length)
		line := fmt.Sprintf("%s: %s\n", *user, password)
		if _, err := f.WriteString(line); err != nil {
			log.Fatal(err)
		}
		fmt.Println(password)
	}
}
