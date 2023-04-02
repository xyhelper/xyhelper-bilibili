package main

import (
	"fmt"

	"xyhelper-bilibili/queue"

	bili "github.com/FishZe/go-bili-chat"
	handle "github.com/FishZe/go-bili-chat/handler"
	"github.com/gogf/gf/v2/util/gconv"
)

func main() {
	go queue.QueueAnswer()
	// 新建一个命令处理器
	h := bili.GetNewHandler()
	// 注册一个处理，将该直播间的弹幕消息绑定到这个函数
	h.AddOption(handle.CmdDanmuMsg, 11690044, func(event handle.MsgEvent) {
		// 打印出弹幕消息
		// fmt.Printf("[%v] %v: %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Content)
		// ctx := gctx.New()
		ask := &queue.Ask{
			Msg:   event.DanMuMsg.Data.Content,
			Asker: event.DanMuMsg.Data.Sender.Name,
		}
		askstring := gconv.String(ask)
		queue.Queue.Push(askstring)
		// go gpt.AskGPT(ctx, event.DanMuMsg.Data.Content)
	})
	// 连接到直播间
	err := h.AddRoom(11690044)
	if err != nil {
		fmt.Println(err)
	}
	// 启动处理器
	h.Run()
}
