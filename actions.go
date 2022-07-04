package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

type Item struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func add(args Arguments, writer io.Writer) error {
	if args["item"] == "" {
		return errors.New("-item flag has to be specified")
	}

	records, err := getRecords(args["fileName"])

	if err != nil {
		return err
	}

	var items []Item
	err = json.Unmarshal([]byte(args["item"]), &items)

	if err != nil {
		return err
	}

	records = append(records, items...)

	data, err := serialize(records)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(args["fileName"], data, 0755)
	if err != nil {
		panic(err)
	}

	return nil
}

func list(args Arguments, writer io.Writer) error {
	records, err := getRecords(args["fileName"])

	if err != nil {
		return err
	}
	data, err := serialize(records)
	if err != nil {
		return err
	}

	writer.Write(data)

	return nil
}

func findById(args Arguments, writer io.Writer) error {
	if args["id"] == "" {
		return errors.New("-id flag has to be specified")
	}

	return nil
}

func remove(args Arguments, writer io.Writer) error {
	return nil
}

func getRecords(filename string) ([]Item, error) {
	var items []Item
	f, err := openFile(filename)
	if err != nil {
		return items, err
	}

	defer f.Close()
	return readItems(f)
}

func serialize(records []Item) ([]byte, error) {
	return json.Marshal(records)
}
