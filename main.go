package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	PROJBASE = "github.com/eigenhombre/"
	MAINTXT  = `package main
import (
	"fmt"
)

func main() {
	fmt.Println("OK")
}
`
	TESTTXT = `package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainChangeMyName(t *testing.T) {
	var tests = []struct {
		input  int
		output int
	}{
		{1, 2},
		{2, 4},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, 2*test.input)
	}
}
`
)

// func projName(r *rand.Rand) string {
// 	return fmt.Sprintf(
// 		"project%d",
// 		r.Intn(1000))
// }

func projDir(proj string) string {
	return os.Getenv("GOPATH") + "/" + proj
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: goproj <project-name>")
		os.Exit(1)
	}

	proj := os.Args[1]
	dir := projDir(proj)
	os.MkdirAll(dir, os.ModePerm)
	err := os.Chdir(dir)
	check(err)

	modCmd := exec.Command("go", "mod", "init", PROJBASE+proj)
	modOut, err := modCmd.CombinedOutput()
	fmt.Println(string(modOut))
	check(err)

	err = os.WriteFile("main.go", []byte(MAINTXT), 0644)
	check(err)

	err = os.WriteFile("main_test.go", []byte(TESTTXT), 0644)
	check(err)

	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyOut, err := tidyCmd.CombinedOutput()
	check(err)
	fmt.Println(string(tidyOut))

	fmt.Println("OK (" + proj + ")")
}
