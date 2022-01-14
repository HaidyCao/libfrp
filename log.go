package libfrp

import (
	frp0271Log "github.com/HaidyCao/frp_0271/utils/log"
	frp0282Log "github.com/HaidyCao/frp_0282/utils/log"
	frp0290Log "github.com/HaidyCao/frp_0290/utils/log"
	frp0300Log "github.com/HaidyCao/frp_0300/utils/log"
	frp0312Log "github.com/HaidyCao/frp_0312/utils/log"
	frp0321Log "github.com/HaidyCao/frp_0321/utils/log"
	frp0351Log "github.com/HaidyCao/frp_0351/pkg/util/log"
)

type FrpLogListener interface {
	Log(log string)
	Location() string
}

// type FrpLogListener struct {
// 	name string
// }

// func (l *FrpLogListener) Log(log string) {
// }

func SetFrpLogListener(l FrpLogListener) {
	frp0271Log.AppendListener(l)
	frp0282Log.AppendListener(l)
	frp0290Log.AppendListener(l)
	frp0300Log.AppendListener(l)
	frp0312Log.AppendListener(l)
	frp0321Log.AppendListener(l)
	frp0351Log.AppendListener(l)
}
