package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0 cannot be a divisor")
	}

	return a / b, nil
}
