package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

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
	user   = flag.String("user", "", "username")
)

func loadConfig() (Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func main() {
	logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
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

	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS credentials (
		userid INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255),
		password VARCHAR(255)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	if *length == 12 {
		for i := 0; i < *count; i++ {
			password := generatePassword(*length)
			if _, err := db.Exec("INSERT INTO credentials (username, password) VALUES (?, ?)", *user, password); err != nil {
				log.Fatal(err)
			}
			fmt.Println(password)
		}
	} else {
		password := generatePassword(*length)
		if _, err := db.Exec("INSERT INTO credentials (username, password) VALUES (?, ?)", *user, password); err != nil {
			log.Fatal(err)
		}
		fmt.Println(password)
	}
}
