package main

import (
	"github.com/cloudquery/cq-provider-github/resources"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	p := resources.Provider()
	serve.Serve(&serve.Options{
		Name:                p.Name,
		Provider:            p,
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
