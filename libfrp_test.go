package libfrp

import "testing"

func TestVersion(t *testing.T) {
	t.Log(Version(Frp0282))
	t.Log(Version(Frp0290))
	t.Log(Version(Frp0300))
	t.Log(Version(Frp0312))
	t.Log(Version(Frp0321))
	t.Log(Version(Frp0351))
	t.Log(Version(Frp0362))
	t.Log(Version(Frp0380))
	t.Log(Version(Frp0391))
	t.Log(Version(Frp0420))
}

func TestIsFrpsRunning(t *testing.T) {
	t.Log(IsFrpcRunning(Frp0282))
	t.Log(IsFrpcRunning(Frp0300))
	t.Log(IsFrpcRunning(Frp0312))
	t.Log(IsFrpcRunning(Frp0321))
	t.Log(IsFrpcRunning(Frp0351))
	t.Log(IsFrpcRunning(Frp0362))
	t.Log(IsFrpcRunning(Frp0380))
	t.Log(IsFrpcRunning(Frp0391))
	t.Log(IsFrpcRunning(Frp0420))
}
