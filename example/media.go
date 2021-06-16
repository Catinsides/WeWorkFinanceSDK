package main

import (
	"bytes"
	"fmt"
	"github.com/NICEXAI/WeWorkFinanceSDK"
	"io/ioutil"
	"os"
)

// 参数1: sdkfileid, 参数2: filePath

func main() {
	corpID := ""
	corpSecret := ""
	rsaPrivateKey := ``

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Print("缺少必要参数")
		return
	}

	sdkfileid := args[0]
	filePath := args[1]

	//初始化客户端
	client, err := WeWorkFinanceSDK.NewClient(corpID, corpSecret, rsaPrivateKey)
	if err != nil {
		fmt.Printf("SDK 初始化失败：%v \n", err)
		return
	}

	isFinish := false
	buffer := bytes.Buffer{}
	for !isFinish {
		//获取媒体数据
		mediaData, err := client.GetMediaData("", sdkfileid, "", "", 5)
		if err != nil {
			fmt.Printf("媒体数据拉取失败：%v \n", err)
			return
		}
		buffer.Write(mediaData.Data)
		if mediaData.IsFinish {
			isFinish = mediaData.IsFinish
		}
	}

	err = ioutil.WriteFile(filePath, buffer.Bytes(), 0666)
	if err != nil {
		fmt.Printf("文件存储失败：%v \n", err)
		return
	}
}
