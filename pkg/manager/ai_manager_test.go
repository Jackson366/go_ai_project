package manager

import (
	"context"
	"fmt"
	"github.com/yankeguo/zhipu"
	"testing"
	"time"
)

const (
	GenerateQuestionPrompt = "你是一位严谨的出题专家，我会给你如下信息：\n" +
		"```\n" +
		"应用名称，\n" +
		"【【【应用描述】】】，\n" +
		"应用类别，\n" +
		"要生成的题目数，\n" +
		"每个题目的选项数\n" +
		"```\n" +
		"\n" +
		"请你根据上述信息，按照以下步骤来出题：\n" +
		"1. 要求：题目和选项尽可能地短，题目不要包含序号，每题的选项数以我提供的为主，题目不能重复\n" +
		"2. 严格按照下面的 json 格式输出题目和选项\n" +
		"```\n" +
		"[{\"options\":[{\"value\":\"选项内容\",\"key\":\"A\"},{\"value\":\"\",\"key\":\"B\"}],\"title\":\"题目标题\"}]\n" +
		"```\n" +
		"title 是题目，options 是选项，每个选项的 key 按照英文字母序（比如 A、B、C、D）以此类推，value 是选项内容\n" +
		"3. 检查题目是否包含序号，若包含序号则去除序号\n" +
		"4. 返回的题目列表格式必须为 JSON 数组"
)

func TestChat(t *testing.T) {
	client, _ := zhipu.NewClient(zhipu.WithAPIKey("17045bbd4e5db9bc76b98c5557c6cbcd.UvQCu5rFjOFp3fCq"))
	service := client.ChatCompletion("glm-4").
		AddMessage(zhipu.ChatCompletionMessage{
			Role:    "system",
			Content: "你是家教助手，你叫小王",
		}, zhipu.ChatCompletionMessage{
			Role:    "user",
			Content: "你叫什么名字",
		})

	res, err := service.Do(context.Background())

	if err != nil {
		zhipu.GetAPIErrorCode(err) // get the API error code
	} else {
		fmt.Println(res.Choices[0].Message.Content)
		t.Log(res.Choices[0].Message.Content)
	}
}

func TestChatGLM(t *testing.T) {
	glm, _ := GetChatGLM()
	service := glm.client.ChatCompletion("glm-4").
		AddMessage(zhipu.ChatCompletionMessage{
			Role:    "system",
			Content: "你是家教助手，你叫小王",
		}, zhipu.ChatCompletionMessage{
			Role:    "user",
			Content: "你叫什么名字",
		})

	res, err := service.Do(context.Background())

	if err != nil {
		zhipu.GetAPIErrorCode(err) // get the API error code
	} else {
		fmt.Println(res.Choices[0].Message.Content)
		t.Log(res.Choices[0].Message.Content)
	}
}

func TestChatGLM_DoRequest(t *testing.T) {
	glm, _ := GetChatGLM()
	res := glm.DoSyncStableRequest(GenerateQuestionPrompt, "应用名称：快车\n 应用描述：一款提供快速叫车服务的应用\n 应用类别：交通出行\n 要生成的题目数：3\n 每个题目的选项数：4")
	time.Sleep(5 * time.Second)
	t.Log(res)

}
