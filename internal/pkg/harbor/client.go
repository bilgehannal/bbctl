package harbor

import (
	"github.com/bilgehannal/harbctl/internal/pkg/config"
	"github.com/bilgehannal/harbctl/internal/pkg/errors"
	"github.com/mittwald/goharbor-client/v5/apiv2"
)

func GetHTTPClient(ctx config.Context) *apiv2.RESTClient {
	harborClient, err := apiv2.NewRESTClientForHost(ctx.GetApiUrl(), ctx.User, ctx.DecodePassword(), nil)
	if err != nil {
		errors.FatalPanic("Client could not be generated!", err)
	}
	return harborClient
}
