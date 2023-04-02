package gpt

import (
	"log"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/xyhelper/chatgpt-go"
)

// AskGPT ask gpt
func AskGPT(ctx g.Ctx, msg string) error {
	token := uuid.New().String()
	cli := chatgpt.NewClient(
		// chatgpt.WithDebug(true),
		chatgpt.WithTimeout(120*time.Second),
		chatgpt.WithAccessToken(token),
		chatgpt.WithBaseURI("https://freechat.lidong.xin"),
	)
	conversationID := ""
	parentMessage := ""
	stream, err := cli.GetChatStream(msg, conversationID, parentMessage)
	if err != nil {
		log.Fatalf("get chat stream failed: %v\n", err)
	}

	var answer string
	for text := range stream.Stream {
		// log.Printf("stream text: %s\n", text.Content)
		print(strings.Replace(text.Content, answer, "", 1))

		answer = text.Content

	}

	if stream.Err != nil {
		print("stream closed with error: %v\n", stream.Err)
	}
	// 输出换行
	println()

	return nil
}
