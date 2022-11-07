package operations

import (
	"github.com/bilgehannal/harbctl/internal/pkg/args"
	"github.com/bilgehannal/harbctl/internal/pkg/harbor"
)

type GettingOperationBuilder struct{}

type GettingOperation struct{}

func (e GettingOperation) Run() {
	//ctx := context.Background()
	//harborCtx := config.ContextBuilder{}.CreateContext(l.url, l.user, l.password)
	//harborClient := harbor.GetHTTPClient(harborCtx)
	//x := ListPro
	harbor.GetListOfProjects()

}

func (e GettingOperationBuilder) CreateOperation(args args.Args) Operation {
	return GettingOperation{}
}
