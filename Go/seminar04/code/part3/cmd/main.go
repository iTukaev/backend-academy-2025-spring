package main

import (
	"go.uber.org/fx"
)

func main() {
	fx.New(BuildApp()).Run()
}
