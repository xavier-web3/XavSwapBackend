package main

import (
	"flag"

	"github.com/xavier-web3/XavSwapBackend/src/app"
	"github.com/xavier-web3/XavSwapBackend/src/config"
	"github.com/xavier-web3/XavSwapBackend/src/service/svc"
	"github.com/zeromicro/go-zero/rest/router"
)

const (
	repoRoot          = ""
	defaultConfigPath = "./config/config.toml"
)

func main() {
	conf := flag.String("conf", defaultConfigPath, "conf file path")
	flag.Parse()
	c, err := config.UnmarshalConfig(*conf)
	if err != nil {
		panic(err)

	}
	for _, chain := range c.ChainSupported {
		if chain.ChainID == 0 || chain.Name == "" {
			panic("invalid chain_suffix config")
		}
	}
	serverCtx, err := svc.NewServiceContext(c)
	if err != nil {
		panic(err)

	}
	r := router.NewRouter(serverCtx)
	app, err := app.NewPlatform(c, r, serverCtx)
	if err != nil {
		panic(err)
	}
	app.Start()
}
