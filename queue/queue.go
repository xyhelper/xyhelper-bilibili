package queue

import (
	"fmt"
	"xyhelper-bilibili/gpt"

	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var Queue = gqueue.New()

type Ask struct {
	Msg   string
	Asker string
}

func QueueAnswer() {
	for {
		if v := Queue.Pop(); v != nil {
			fmt.Println(" Pop:", v)
			ask := &Ask{}
			gconv.Struct(v, ask)
			ctx := gctx.New()
			println("-----------------------------")
			println("回答用户:", ask.Asker, "的问题:", ask.Msg, "")
			gpt.AskGPT(ctx, ask.Msg+"\n")
			println("-----------------------------")

		} else {
			break
		}
	}

}
