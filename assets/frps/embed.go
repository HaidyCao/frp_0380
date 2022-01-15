package frpc

import (
	"embed"

	"github.com/HaidyCao/frp_0380/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
