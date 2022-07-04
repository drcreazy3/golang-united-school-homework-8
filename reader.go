package main

import (
	"bufio"
	"encoding/json"
	"os"
)

func openFile(file string) (*os.File, error) {
	return os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
}

func readItems(f *os.File) ([]Item, error) {
	var items []Item
	scanner := bufio.NewScanner(f)

	str := ""
	for scanner.Scan() {
		str += scanner.Text()
	}

	if str == "" {
		return items, nil
	}

	if err := scanner.Err(); err != nil {
		return items, err
	}

	err := json.Unmarshal([]byte(str), &items)

	if err != nil {
		return items, err
	}

	return items, nil
}
