package app

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

func Notification(message string) {
	token := viper.GetString("telegram_bot_token")
	// 设置消息接收者的 Chat ID
	chatID := viper.GetInt64("chat_id")

	// 初始化 Bot
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// 构造消息的内容，可以包含 Emoji 表情等
	msg := tgbotapi.NewMessage(chatID, message)

	// 发送消息
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
