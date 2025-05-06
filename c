package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
ch1, err := readCSV("file.csv")
if err != nil {
panic(fmt.Errorf("Could not read file 1 %v\n", err))
}

ch2, err := readCSV("file2.csv")
if err != nil {

}
}
