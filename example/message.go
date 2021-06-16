package main

import (
	"fmt"
	"github.com/NICEXAI/WeWorkFinanceSDK"
	"os"
	"encoding/json"
	"strconv"
)

// 参数1: seq, 参数2: limit

func main() {
	corpID := ""
	corpSecret := ""
	rsaPrivateKey := ``

	args := os.Args[1:]
	seqVal := "0"
	limitVal := "100"

	if len(args) > 0 {
		seqVal = args[0]
		limitVal = args[1]
	}

	seq, _ := strconv.ParseUint(seqVal, 10, 64)
	limit, _ := strconv.ParseUint(limitVal, 10, 64)

	//初始化客户端
	client, err := WeWorkFinanceSDK.NewClient(corpID, corpSecret, rsaPrivateKey)
	if err != nil {
		msg := fmt.Sprintf("SDK 初始化失败：%v \n", err)
		fmt.Fprintf(os.Stdout, msg)
		return
	}

	//同步消息
	chatDataList, err := client.GetChatData(seq, limit, "", "", 3)
	if err != nil {
		msg := fmt.Sprintf("消息同步失败：%v \n", err)
		fmt.Fprintf(os.Stdout, msg)
		return
	}

	for _, chatData := range chatDataList {
		//消息解密
		chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
		if err != nil {
			msg := fmt.Sprintf("消息解密失败：%v \n", err)
			fmt.Fprintf(os.Stdout, msg)
			return
		}

		chatInfoJson, _ := json.Marshal(chatInfo)
		chatString := fmt.Sprintf("%s", chatInfoJson)

		type Out struct {
			ChatInfo string `json:"chatInfo"`
			OriginData string `json:"originData"`
		}

		originData := chatInfo.GetOriginMessage()
		originJson, _ := json.Marshal(originData)
		originString := fmt.Sprintf("%s", originJson)

		out := Out{chatString, originString}
		outputJSON, _ := json.Marshal(out)
		outputString := fmt.Sprintf("%s", outputJSON)

		fmt.Fprintf(os.Stdout, outputString)
	}
}
