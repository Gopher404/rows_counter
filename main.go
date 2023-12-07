package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var (
	CountOfFiles int
	CountOfRows  int
	Extension    = "go"
)

func main() {
	fmt.Print(">>>")
	if _, err := fmt.Scan(&Extension); err != nil {
		panic(err)
	}
	startTime := time.Now()
	fmt.Println("running")
	ScanDir("./")
	executingTime := time.Since(startTime)
	fmt.Println("files: ", CountOfFiles)
	fmt.Println("rows:  ", CountOfRows)
	fmt.Printf("Executing for %f seconds\n", executingTime.Seconds())
	fmt.Println("Press Ctrl+c to exit")

	for {
		time.Sleep(time.Hour)
	}
}

func ScanDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			ScanDir(path + name + "/")
		} else {
			splitName := strings.Split(name, ".")
			if len(splitName) >= 2 {
				ext := splitName[len(splitName)-1]
				if ext == Extension {
					CountOfFiles++
					ParseFile(path + name)
				}
			}
		}
	}
}

func ParseFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	CountOfRows += strings.Count(string(data), "\n") + 1
}
