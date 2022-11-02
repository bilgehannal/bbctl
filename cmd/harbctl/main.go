package main

import (
	"github.com/bilgehannal/harbctl/internal/pkg/args"
	"github.com/bilgehannal/harbctl/internal/pkg/operations"
	"os"
)

func main() {
	projectArgs := args.GetArgs(os.Args[1:])
	o := operations.GetOperation(projectArgs)
	o.Run()
}
