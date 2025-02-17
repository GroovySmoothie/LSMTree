package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const DB_FILE_NAME = "db.txt"

var byteOffset = make(map[int]int64)

func WriteValue(id int, value string) {
	file, _ := os.OpenFile(DB_FILE_NAME, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	stat, _ := file.Stat()

	file.WriteString(fmt.Sprintf("%d,%s\n", id, value))

	byteOffset[id] = int64(stat.Size())
}

func ReadValue(id int) string {
	file, _ := os.Open(DB_FILE_NAME)

	offset, present := byteOffset[id]

	if !present {
		return ""
	}

	file.Seek(offset, 0)

	reader := bufio.NewScanner(file)
	reader.Scan()

	line := reader.Text()

	slices := strings.SplitAfterN(line, ",", 2)

	return string(slices[1])
}

func Clear() {
	os.Remove(DB_FILE_NAME)
}

func main() {
	Clear()
	WriteValue(1, "hello world")
	WriteValue(2, "hello world 2")
	WriteValue(3, "hello world 3")
	fmt.Println(ReadValue(1))
	WriteValue(1, "new value, with a delimiter")
	fmt.Println(ReadValue(1))
	WriteValue(1, ",")
	fmt.Println(ReadValue(1))
}
