package operations

import (
	"fmt"
	"github.com/bilgehannal/bbctl/internal/pkg/args"
)

type LoginOperationBuilder struct{}

type LoginOperation struct {
	url      string
	user     string
	password string
}

func (l LoginOperation) Run() {
	fmt.Println("Login with", l.url, l.user, l.password)
}

func (l LoginOperationBuilder) CreateOperation(args args.Args) Operation {
	fmt.Println(args)
	return LoginOperation{
		url:      args.Verb.Object.Value,
		user:     args.Params[0].ValueOfParameter,
		password: args.Params[1].ValueOfParameter,
	}
}
