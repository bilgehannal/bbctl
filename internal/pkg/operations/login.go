package operations

import (
	"context"
	"fmt"
	"github.com/bilgehannal/harbctl/internal/pkg/args"
	"github.com/bilgehannal/harbctl/internal/pkg/config"
	"github.com/bilgehannal/harbctl/internal/pkg/errors"
	"github.com/bilgehannal/harbctl/internal/pkg/harbor"
)

type LoginOperationBuilder struct{}

type LoginOperation struct {
	url      string
	user     string
	password string
}

func (l LoginOperation) Run() {
	ctx := context.Background()
	harborCtx := config.ContextBuilder{}.CreateContext(l.url, l.user, l.password)
	harborClient := harbor.GetHTTPClient(harborCtx)

	userInfo, err := harborClient.GetCurrentUserInfo(ctx)
	if err != nil {
		errors.FatalPanic("Login Failed!", err)
	}

	defaultConfig := config.GetDefaultConfig()
	err = defaultConfig.SetDefaultConfig(harborCtx, userInfo)
	if err != nil {
		errors.FatalPanic("Default conf cannot be set...", err)
	}
	fmt.Println("Login Successful...")
}

func (l LoginOperationBuilder) CreateOperation(argsInstance args.Args) Operation {
	userParameter, passwordParameter := getUserAndPassword(argsInstance)
	return LoginOperation{
		url:      argsInstance.Verb.Object.Value,
		user:     userParameter.ValueOfParameter,
		password: passwordParameter.ValueOfParameter,
	}
}

func getUserAndPassword(argsInstance args.Args) (args.Parameter, args.Parameter) {
	userParameter, err := argsInstance.GetParameter(args.GetParameterTypes().User)
	if err != nil {
		errors.FatalPanic("User parameter could not be found", err)
	}
	passwordParameter, err := argsInstance.GetParameter(args.GetParameterTypes().Password)
	if err != nil {
		errors.FatalPanic("Password parameter could not be found", err)
	}
	return userParameter, passwordParameter
}
