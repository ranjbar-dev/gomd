package main

import (
	"fmt"
	"os"

	"github.com/ranjbar-dev/gomd/structparser"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: gomd <action> ...")
		return
	}

	action := os.Args[1]

	if action == "parse-structs" {

		if len(os.Args) < 4 {

			fmt.Println("Usage: gomd parse-structs <path_to_go_folder> <path_to_output_folder>")
			return
		}

		err := structparser.ParseFolder(os.Args[2], os.Args[3])
		if err != nil {

			fmt.Println("Error parsing folder:", err)
		}

	} else {

		fmt.Println("Unknown action:", action)
	}
}
