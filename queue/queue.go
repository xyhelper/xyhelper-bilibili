package queue

import (
	"fmt"
	"xyhelper-bilibili/gpt"

	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var Queue = gqueue.New()

type Ask struct {
	Msg   string
	Asker string
}

// Talk 会话
type Talk struct {
	ConversationID string `json:"conversation_id"` // 会话ID
	ParentMessage  string `json:"parent_message"`  // 父消息
}

// Talks 会话列表
var Talks = make(map[string]Talk)

func QueueAnswer() {
	for {
		if v := Queue.Pop(); v != nil {
			fmt.Println(" Pop:", v)
			g.Dump(Talks)
			ask := &Ask{}
			gconv.Struct(v, ask)
			ctx := gctx.New()
			if _, ok := Talks[ask.Asker]; !ok {
				// 不存在
				Talks[ask.Asker] = Talk{}
			}

			println("-----------------------------")
			println("回答用户:", ask.Asker, "的问题:", ask.Msg, "")
			conversation_id, parent_message, err := gpt.AskGPT(ctx, ask.Msg+"\n", ask.Asker, Talks[ask.Asker].ConversationID, Talks[ask.Asker].ParentMessage)
			if err != nil {
				println("回答失败:", err)
			} else {
				Talks[ask.Asker] = Talk{
					ConversationID: conversation_id,
					ParentMessage:  parent_message,
				}
			}
			println("-----------------------------")

		} else {
			break
		}
	}

}
