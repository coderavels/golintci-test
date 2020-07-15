package main

import (
	"github.com/heisfullstacked/golintci-test/service"
)

func main() {
	test := service.ShouldBeLint{
		ID: "random",
	}
	print(test)
}

//type change string
