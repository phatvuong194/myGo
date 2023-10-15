package main

import (
	"fmt"

	"github.com/phatvuong194/myGo/helpers"
	"github.com/phatvuong194/myGo/loggers"
)

func main() {
	random := helpers.RandomNumber(5)
	fmt.Println(random)

	loggers.LogInfor("main run")
}
