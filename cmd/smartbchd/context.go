package main

import (
	"os"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/superbch/superbch/param"
)

var (
	DefaultNodeHome = os.ExpandEnv("$HOME/.superbchd")
)

type Context struct {
	Config *param.ChainConfig
	Logger log.Logger
}

func NewDefaultContext() *Context {
	return NewContext(
		param.DefaultConfig(),
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)),
	)
}

func NewContext(config *param.ChainConfig, logger log.Logger) *Context {
	return &Context{config, logger}
}
