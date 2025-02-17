package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const DB_FILE_NAME = "db.txt"

func WriteValue(id int, value string) {
	file, _ := os.OpenFile(DB_FILE_NAME, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	file.WriteString(fmt.Sprintf("%d,%s\n", id, value))
}

func ReadValue(id int) string {
	file, _ := os.Open(DB_FILE_NAME)

	scanner := bufio.NewScanner(file)
	var foundValue = ""
	for scanner.Scan() {
		value := scanner.Text()
		if BeginsWithId(value, id) {
			foundValue = value
		}
	}

	return foundValue
}

func BeginsWithId(str string, id int) bool {
	return strings.HasPrefix(str, fmt.Sprintf("%d,", id))
}

func main() {
	WriteValue(1, "hello world")
	WriteValue(2, "hello world 2")
	WriteValue(3, "hello world 3")
	fmt.Println(ReadValue(1))
}
