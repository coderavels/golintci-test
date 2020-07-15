package main

import (
	"github.com/heisfullstacked/test/service"
)

func main() {
	test := service.TEST{
		ID: "random",
	}
	print(test)
}
