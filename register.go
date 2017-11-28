package cntk

import (
	"os"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/rai-project/dlframework"
	"github.com/rai-project/dlframework/framework"
)

// FrameworkManifest ...
var FrameworkManifest = dlframework.FrameworkManifest{
	Name:    "CNTK",
	Version: "2.3",
	Container: map[string]*dlframework.ContainerHardware{
		"amd64": {
			Cpu: "raiproject/carml-cntk:amd64-cpu",
			Gpu: "raiproject/carml-cntk:amd64-gpu",
		},
		"ppc64le": {
			Cpu: "raiproject/carml-cntk:ppc64le-gpu",
			Gpu: "raiproject/carml-cntk:ppc64le-gpu",
		},
	},
}

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}

func Register() {
	if supportedSystem {
		framework.Register(FrameworkManifest, assetFS())
	}
}
