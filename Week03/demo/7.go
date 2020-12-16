package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func operation1(ctx context.Context) error {
	// 让我们假设这个操作会因为某种原因失败
	// 我们使用time.Sleep来模拟一个资源密集型操作
	time.Sleep(5 * time.Second)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	// 我们使用在前面HTTP服务器例子里使用过的类型模式
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("done")
	}
}

func main() {
	// 新建一个上下文
	ctx := context.Background()
	// 在初始上下文的基础上创建一个有取消功能的上下文
	ctx, cancel := context.WithCancel(ctx)
	// 在不同的goroutine中运行operation2
	go func() {
		operation2(ctx)
	}()
	
	err := operation1(ctx)
	// 如果这个操作返回错误，取消所有使用相同上下文的操作
	if err != nil {
		cancel()
	}
	
	select {
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}