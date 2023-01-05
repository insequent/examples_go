package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func checkErr(err error) {
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	dir, err := os.MkdirTemp(os.TempDir(), "example_go")
	checkErr(err)

	// Using os
	basicFile := path.Join(dir, "basic")

	f, err := os.Create(basicFile) // Creates or Opens existing and truncates
	checkErr(err)

	str := "This is testing using 'os'"
	_, err = f.WriteString(str)
	checkErr(err)

	err = f.Close()
	checkErr(err)

	f, err = os.Open(basicFile)
	checkErr(err)

	bytes := make([]byte, len(str))
	_, err = f.Read(bytes)
	checkErr(err)
	fmt.Println("Basic result:", string(bytes))

	// Using ioutil (easier for the most part)
	easyFile := path.Join(dir, "easy")

	str = "This is testing 'ioutil'"
	err = ioutil.WriteFile(easyFile, []byte(str), 0666)
	checkErr(err)

	bytes, err = ioutil.ReadFile(easyFile)
	checkErr(err)
	fmt.Println("Easy result:", string(bytes))
}
