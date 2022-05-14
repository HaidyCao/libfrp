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
	frp0351 "github.com/HaidyCao/frp_0351/cmd/frp"
	frp0362 "github.com/HaidyCao/frp_0362/cmd/frp"
	frp0380 "github.com/HaidyCao/frp_0380/cmd/frp"
	frp0391 "github.com/HaidyCao/frp_0390/cmd/frp"
	"github.com/fatedier/golib/crypto"
)

var (
	Frp0391 = "0.39.1"
	Frp0380 = "0.38.0"
	Frp0362 = "0.36.2"
	Frp0351 = "0.35.1"
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
	case Frp0351:
		return frp0351.RunFrpc(cfgFilePath)
	case Frp0362:
		return frp0362.RunFrpc(cfgFilePath)
	case Frp0380:
		return frp0380.RunFrpc(cfgFilePath)
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
	case Frp0351:
		return frp0351.StopFrpc()
	case Frp0362:
		return frp0362.StopFrpc()
	case Frp0380:
		return frp0380.StopFrpc()
	case Frp0391:
		return frp0391.StopFrpc()
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
	case Frp0351:
		return frp0351.IsFrpcRunning()
	case Frp0362:
		return frp0362.IsFrpcRunning()
	case Frp0380:
		return frp0380.IsFrpcRunning()
	case Frp0391:
		return frp0391.IsFrpcRunning()
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
	case Frp0351:
		return frp0351.RunFrps(cfgFilePath)
	case Frp0362:
		return frp0362.RunFrps(cfgFilePath)
	case Frp0380:
		return frp0380.RunFrps(cfgFilePath)
	case Frp0391:
		return frp0391.RunFrps(cfgFilePath)
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
	case Frp0351:
		return frp0351.StopFrps()
	case Frp0362:
		return frp0362.StopFrps()
	case Frp0380:
		return frp0380.StopFrps()
	case Frp0391:
		return frp0391.StopFrps()
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
	case Frp0351:
		return frp0351.IsFrpsRunning()
	case Frp0362:
		return frp0362.IsFrpsRunning()
	case Frp0380:
		return frp0380.IsFrpsRunning()
	case Frp0391:
		return frp0391.IsFrpsRunning()
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
	case Frp0351:
		return frp0351.Version()
	case Frp0362:
		return frp0362.Version()
	case Frp0380:
		return frp0380.Version()
	case Frp0391:
		return frp0391.Version()
	default:
		return "bad version"
	}
}
