package harbor

import (
	"context"
	"fmt"
	"github.com/bilgehannal/harbctl/internal/pkg/config"
	"github.com/bilgehannal/harbctl/internal/pkg/utils"
)

func GetListOfProjects() {
	ctx := context.Background()
	harborCtx := config.ContextBuilder{}.CreateContext("https://harbor.ngntest.sahibindenlocal.net", "bilgehan.nal", "xxxx")
	harborClient := GetHTTPClient(harborCtx)
	projects, err := harborClient.ListProjects(ctx, "")
	projectsString := ""
	for _, project := range projects {
		projectsString += project.Name + "\n"
	}
	fmt.Println(projectsString)
	projectStr := utils.PrettyPrintJson(projects[0])
	fmt.Println(projectStr)
	fmt.Println(err)
	if err != nil {
		return
	}
}
