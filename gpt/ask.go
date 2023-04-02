package gpt

import (
	"log"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/xyhelper/chatgpt-go"
)

// AskGPT ask gpt
func AskGPT(ctx g.Ctx, msg string, token string, conversationID string, parentMessage string) (conversation_id string, parent_message string, err error) {
	// token := uuid.New().String()
	cli := chatgpt.NewClient(
		// chatgpt.WithDebug(true),
		chatgpt.WithTimeout(120*time.Second),
		chatgpt.WithAccessToken(token),
		chatgpt.WithBaseURI("https://freechat.lidong.xin"),
	)
	stream, err := cli.GetChatStream(msg, conversationID, parentMessage)
	if err != nil {
		log.Fatalf("get chat stream failed: %v\n", err)
	}

	var answer string
	for text := range stream.Stream {
		// log.Printf("stream text: %s\n", text.Content)
		print(strings.Replace(text.Content, answer, "", 1))

		answer = text.Content

		conversation_id = text.ConversationID
		parent_message = text.MessageID

	}

	if stream.Err != nil {
		print("stream closed with error: %v\n", stream.Err)
	}
	// 输出换行
	println()
	println("conversation_id:", conversation_id)
	println("parent_message:", parent_message)
	return
}
