package libfrp

import (
	"fmt"
	"runtime/debug"

	frp0271 "github.com/HaidyCao/frp_0271/cmd/frp"
	frp0282 "github.com/HaidyCao/frp_0282/cmd/frp"
	frp0290 "github.com/HaidyCao/frp_0290/cmd/frp"
	frp0300 "github.com/HaidyCao/frp_0300/cmd/frp"
	frp0312 "github.com/HaidyCao/frp_0312/cmd/frp"
	frp0321 "github.com/HaidyCao/frp_0321/cmd/frp"
	frp0341 "github.com/HaidyCao/frp_0341/cmd/frp"
	"github.com/fatedier/golib/crypto"
)

var (
	Frp0341 = "0.34.1"
	Frp0321 = "0.32.1"
	Frp0312 = "0.31.2"
	Frp0300 = "0.30.0"
	Frp0290 = "0.29.0"
	Frp0282 = "0.28.2"
	Frp0271 = "0.27.1"
)

func RunFrpc(cfgFilePath, version string) (err error) {
	crypto.DefaultSalt = "frp"

	switch version {
	case Frp0282:
		return frp0282.RunFrpc(cfgFilePath)
	case Frp0271:
		return frp0271.RunFrpc(cfgFilePath)
	case Frp0290:
		return frp0290.RunFrpc(cfgFilePath)
	case Frp0300:
		return frp0300.RunFrpc(cfgFilePath)
	case Frp0312:
		return frp0312.RunFrpc(cfgFilePath)
	case Frp0321:
		return frp0321.RunFrpc(cfgFilePath)
	case Frp0341:
		return frp0341.RunFrpc(cfgFilePath)
	default:
		return fmt.Errorf("bad version: %s", version)
	}
}

func StopFrpc(version string) (err error) {
	defer debug.FreeOSMemory()
	switch version {
	case Frp0282:
		return frp0282.StopFrpc()
	case Frp0271:
		return frp0271.StopFrpc()
	case Frp0290:
		return frp0290.StopFrpc()
	case Frp0300:
		return frp0300.StopFrpc()
	case Frp0312:
		return frp0312.StopFrpc()
	case Frp0321:
		return frp0321.StopFrpc()
	case Frp0341:
		return frp0341.StopFrpc()
	default:
		return fmt.Errorf("bad version: %s", version)
	}
}

func IsFrpcRunning(version string) bool {
	switch version {
	case Frp0282:
		return frp0282.IsFrpcRunning()
	case Frp0271:
		return frp0271.IsFrpcRunning()
	case Frp0290:
		return frp0290.IsFrpcRunning()
	case Frp0300:
		return frp0300.IsFrpcRunning()
	case Frp0312:
		return frp0312.IsFrpcRunning()
	case Frp0321:
		return frp0321.IsFrpcRunning()
	case Frp0341:
		return frp0341.IsFrpcRunning()
	default:
		return false
	}
}

func RunFrps(cfgFilePath string, version string) (err error) {
	crypto.DefaultSalt = "frp"
	switch version {
	case Frp0282:
		return frp0282.RunFrps(cfgFilePath)
	case Frp0271:
		return frp0271.RunFrps(cfgFilePath)
	case Frp0290:
		return frp0290.RunFrps(cfgFilePath)
	case Frp0300:
		return frp0300.RunFrps(cfgFilePath)
	case Frp0312:
		return frp0312.RunFrps(cfgFilePath)
	case Frp0321:
		return frp0321.RunFrps(cfgFilePath)
	case Frp0341:
		return frp0341.RunFrps(cfgFilePath)
	default:
		return fmt.Errorf("bad version: %s", version)
	}
}

// StopFrps 停止frps服务
func StopFrps(version string) error {
	switch version {
	case Frp0282:
		return frp0282.StopFrps()
	case Frp0271:
		return frp0271.StopFrps()
	case Frp0290:
		return frp0290.StopFrps()
	case Frp0300:
		return frp0300.StopFrps()
	case Frp0312:
		return frp0312.StopFrps()
	case Frp0321:
		return frp0321.StopFrps()
	case Frp0341:
		return frp0341.StopFrps()
	default:
		return fmt.Errorf("bad version: %s", version)
	}
}

// IsFrpsRunning 是否还在运行
func IsFrpsRunning(version string) bool {
	switch version {
	case Frp0282:
		return frp0282.IsFrpsRunning()
	case Frp0271:
		return frp0271.IsFrpsRunning()
	case Frp0290:
		return frp0290.IsFrpsRunning()
	case Frp0300:
		return frp0300.IsFrpsRunning()
	case Frp0312:
		return frp0312.IsFrpsRunning()
	case Frp0321:
		return frp0321.IsFrpsRunning()
	case Frp0341:
		return frp0341.IsFrpsRunning()
	default:
		return false
	}
}

// Version frp 版本号
func Version(version string) string {
	switch version {
	case Frp0282:
		return frp0282.Version()
	case Frp0271:
		return frp0271.Version()
	case Frp0290:
		return frp0290.Version()
	case Frp0300:
		return frp0300.Version()
	case Frp0312:
		return frp0312.Version()
	case Frp0321:
		return frp0321.Version()
	case Frp0341:
		return frp0341.Version()
	default:
		return "bad version"
	}
}
