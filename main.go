package main

import (
	"TelegramBot/telegram"
	_ "github.com/lib/pq"
	tele "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func main() {

	pref, err := tele.NewBot(tele.Settings{
		Token:  "",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b := telegram.NewBot(pref)
	b.Init()

	b.Bot.Start()

	//b, err := tele.NewBot(pref)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//b.Start()
}

//var mainMenu = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("📩 Написать администратору офиса"),
//		tgbotapi.NewKeyboardButton("🍛 Отправить меню"),
//	),
//)
//
//type MenuList struct {
//	dishes
//}
//
//func main() {
//	var (
//		bot        *tgbotapi.BotAPI
//		err        error
//		updChannel tgbotapi.UpdatesChannel
//		update     tgbotapi.Update
//		updConfig  tgbotapi.UpdateConfig
//	)
//	bot, err = tgbotapi.NewBotAPI(Token)
//	if err != nil {
//		panic("Update channel error: " + err.Error())
//	}
//	updConfig.Timeout = 60
//	updConfig.Limit = 1
//	updConfig.Offset = 0
//
//	updChannel = bot.GetUpdatesChan(updConfig)
//
//	for {
//		update = <-updChannel
//		if update.Message != nil {
//			if update.Message.IsCommand() {
//				cmdText := update.Message.Command()
//				if cmdText == "menu" {
//					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню")
//					msg.ReplyMarkup = mainMenu
//					bot.Send(msg)
//				}
//			} else {
//				fmt.Printf("From: %s ChatID: %v Message: %s\n", update.Message.From.FirstName, update.Message.Chat.ID, update.Message.Text)
//				msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//				bot.Send(msgConfig)
//			}
//		} else {
//			fmt.Printf("Not message: %+v\n", update)
//		}
//
//	}
//	bot.StopReceivingUpdates()
//}

//const (
//	tgBotHost = "api.telegram.org"
//)
//
//func main() {
//	tgClient := telegram.New(tgBotHost, mustToken())
//}
//
//func mustToken() string {
//	token := flag.String(
//		"tg-bot-token",
//		"",
//		"Токен для доступа к телеграм бот",
//	)
//	flag.Parse()
//
//	if *token == "" {
//		log.Fatal("Токен не указан")
//	}
//
//	return *token
//
//}
