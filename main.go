package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type Arguments map[string]string

func Perform(args Arguments, writer io.Writer) error {

	if args["operation"] == "" {
		return errors.New("-operation flag has to be specified")
	}
	if args["fileName"] == "" {
		return errors.New("-fileName flag has to be specified")
	}

	switch args["operation"] {
	case "add":
		return add(args, writer)
	case "list":
		return list(args, writer)
	case "findById":
		return findById(args, writer)
	case "remove":
		return remove(args, writer)
	default:
		return fmt.Errorf("Operation %s not allowed!", args["operation"])
	}
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	operation := *flag.String("operation", "", "Available operations: add, list, findById, remove")
	fileName := *flag.String("fileName", "", "Destination file: users.json")
	id := *flag.String("id", "", "[optional] record id")
	item := *flag.String("item", "", "[oprional] Valid json item {«id»: \"1\", «email»: «email@test.com», «age»: 23}")

	return Arguments{"operation": operation, "item": item, "fileName": fileName, "id": id}
}
