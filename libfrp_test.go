package libfrp

import "testing"

func TestVersion(t *testing.T) {
	t.Log(Version(Frp0271))
	t.Log(Version(Frp0282))
	t.Log(Version(Frp0290))
	t.Log(Version(Frp0300))
	t.Log(Version(Frp0312))
	t.Log(Version(Frp0321))
	t.Log(Version(Frp0341))
}

func TestIsFrpsRunning(t *testing.T) {
	t.Log(IsFrpcRunning(Frp0271))
	t.Log(IsFrpcRunning(Frp0282))
	t.Log(IsFrpcRunning(Frp0300))
	t.Log(IsFrpcRunning(Frp0312))
}