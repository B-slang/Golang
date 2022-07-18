package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() //interpret as a string
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("Here it is : %d", 2022-input)
}
