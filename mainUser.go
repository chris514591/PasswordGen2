package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
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

func main() {
	logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)

	reader := bufio.NewReader(os.Stdin)

	for {
		var length int
		fmt.Print("Enter password length (default 10): ")
		_, err := fmt.Scanf("%d", &length)
		if err != nil {
			length = 10
		}

		if length <= 0 {
			log.Fatalf("invalid password length: %d", length)
		}

		if length == 10 {
			password := generatePassword(length)
			fmt.Println(password)
			err1 := clipboard.WriteAll(password)
			if err1 != nil {
				fmt.Println("Failed to copy to clipboard:", err1)
			}
			fmt.Println("Copied last generated password to clipboard!")
		} else {
			password := generatePassword(length)
			fmt.Println(password)
			err1 := clipboard.WriteAll(password)
			if err1 != nil {
				fmt.Println("Failed to copy to clipboard:", err1)
			} else {
				fmt.Println("Copied to clipboard!")
			}
		}

		fmt.Println("Generate another password? (y/n): ")
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		answer = strings.TrimSpace(strings.ToLower(answer))
		if answer == "n" {
			break
		}
	}
}
