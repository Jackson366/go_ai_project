package manager

import (
	"context"
	"fmt"
	"github.com/yankeguo/zhipu"
	"goAiproject/global"
	"sync"
)

const (
	// StableTemperature 稳定答案
	StableTemperature = 0.05
	// UnstableTemperature 非稳定答案
	UnstableTemperature = 0.99
	// ModelVersion 模型版本
	ModelVersion = "glm-4"
)

var (
	instance *ChatGLM
	once     sync.Once
)

type ChatGLM struct {
	client *zhipu.Client
}

func GetChatGLM() (*ChatGLM, error) {
	once.Do(func() {
		client, err := zhipu.NewClient(zhipu.WithAPIKey("17045bbd4e5db9bc76b98c5557c6cbcd.UvQCu5rFjOFp3fCq"))
		if err != nil {
			global.Logger.Fatalf("Failed to create zhipu client: %v", err)
			return
		}
		instance = &ChatGLM{
			client: client,
		}
	})
	if instance == nil || instance.client == nil {
		return nil, fmt.Errorf("failed to create ChatGLM instance or zhipu client is nil")
	}
	return instance, nil
}

func (c ChatGLM) DoSyncStableRequest(systemMessage, userMessage string) string {
	res, _ := c.DoRequest(systemMessage, userMessage, StableTemperature)
	return res
}

func (c ChatGLM) DoSyncUnStableRequest(systemMessage, userMessage string) string {
	res, _ := c.DoRequest(systemMessage, userMessage, UnstableTemperature)
	return res
}

func (c ChatGLM) DoRequest(systemMessage, userMessage string, temperature float64) (string, error) {
	service := c.client.ChatCompletion(ModelVersion).
		AddMessage(zhipu.ChatCompletionMessage{
			Role:    "system",
			Content: systemMessage,
		}, zhipu.ChatCompletionMessage{
			Role:    "user",
			Content: userMessage,
		}).SetTemperature(temperature).SetMaxTokens(4095)
	// 上下文空白，可否修改
	res, err := service.Do(context.Background())
	if err != nil {
		// 处理错误
		zhipu.GetAPIErrorCode(err) // get the API error code
	} else {
		println(res.Choices[0].Message.Content)
	}
	return res.Choices[0].Message.Content, nil
}
