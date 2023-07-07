package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

var (
	client = &http.Client{}
)

func ChannelNotification(info string, path string, fileSize float64) {
	volumePath := viper.GetString("volume_path")
	mediaDir := strings.Replace(path, volumePath, "", 1)
	telegramBotToken := viper.GetString("telegram_bot_token")
	channelId := viper.GetString("channel_id")
	apiToken := viper.GetString("api_token")
	webApi := viper.GetString("web_api")
	trimmedPath := strings.ToLower(strings.Trim(mediaDir, "/"))

	mediaData := viper.GetStringMap("data")

	resource, ok := mediaData[trimmedPath].(map[string]interface{})
	if !ok {
		return
	}

	if _, ok = resource["resource_id"]; !ok {
		return
	}

	payload := make(map[string]interface{})
	payload["chat_id"] = channelId
	payload["telegram_bot_token"] = telegramBotToken
	payload["remark"] = fmt.Sprintf(resource["remark_format"].(string), info)
	payload["resource_id"] = resource["resource_id"]
	payload["share_size"] = fileSize

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", webApi, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		// JSON 解析失败
		return
	}

	fmt.Println(response)
}
