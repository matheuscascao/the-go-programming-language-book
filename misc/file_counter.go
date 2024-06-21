package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exePath := filepath.Dir(ex)
	fmt.Println(exePath)

	files, err := os.ReadDir(exePath)
	if err != nil {
		panic(err)
	}
	
	mapFileCount := make(map[string]int)

	for _, file := range files {
		// fmt.Println(file.Name(), file.IsDir())
		fileName := file.Name()
		splitedFileName := strings.Split(fileName, ".")
		fileExtension := splitedFileName[len(splitedFileName)-1]

		mapFileCount[fileExtension] = mapFileCount[fileExtension]+1
	}

	fmt.Println(mapFileCount)
}