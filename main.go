package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/RedaBourma/task/cmd"
	"task/cmd"
)

func main() {
	home, _ := homedir.Dir()
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.exit(1)
	}
}
