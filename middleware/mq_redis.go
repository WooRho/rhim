package middleware

import (
	"context"
	"fmt"
)

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish 。。。。", msg)
	err = SvcCtx.Redis.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := SvcCtx.Redis.Subscribe(ctx, channel)
	fmt.Println("Subscribe 。。。。", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer fmt.Println("Subscribe 。。。。", msg.Payload)
	return msg.Payload, err
}
