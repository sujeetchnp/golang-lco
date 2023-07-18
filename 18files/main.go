package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Demo golang files")

	contents := "This needs to go in a file -  Golang Tutorial"

	file, err := os.Create("./mylcogofile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, contents)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
