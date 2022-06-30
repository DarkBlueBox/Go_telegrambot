package telegram

import (
	tele "gopkg.in/tucnak/telebot.v2"
)

var Auth = &tele.ReplyMarkup{ResizeReplyKeyboard: true}
var BtnAuthCafe = Auth.Text("Войти как Администратор кафе 🔒")
var BtnAuthCompany = Auth.Text("Войти как Администратор компании 🔒")

var menuAdmin = &tele.ReplyMarkup{ResizeReplyKeyboard: true}
var BtnNewMenu = menuAdmin.Text("Создать новое меню 🍲")
var BtnSendMsg = menuAdmin.Text("Написать сообщение Администратору компании")
var BtnShowMenu = menuAdmin.Text("Посмотреть актуальное меню")

type TeleBot struct {
	Bot  *tele.Bot
	chat *tele.Chat
	user *tele.User
}

func NewBot(bot *tele.Bot) *TeleBot {
	return &TeleBot{Bot: bot}
}

func (b *TeleBot) Init() {
	b.initMenus()
	b.Authentication()
	b.adminMenu()
}

func (b *TeleBot) initMenus() {
	Auth.Reply(
		Auth.Row(BtnAuthCafe),
		Auth.Row(BtnAuthCompany),
	)
	menuAdmin.Reply(
		menuAdmin.Row(BtnNewMenu),
		menuAdmin.Row(BtnShowMenu),
	)

}
