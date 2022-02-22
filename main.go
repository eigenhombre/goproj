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
	check(err)

	err = os.WriteFile("main.go", []byte(MAINTXT), 0644)
	check(err)

	fmt.Println(string(modOut))
	fmt.Println("OK (" + proj + ")")
}
