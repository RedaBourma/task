package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"path/filpath"
)

func main() {
	home, _ := homedir.Dir()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.exit(1)
	}
}
