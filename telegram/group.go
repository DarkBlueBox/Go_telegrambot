package telegram

import tele "gopkg.in/tucnak/telebot.v2"

func (b *TeleBot) addToGroup() {
	b.Bot.Handle(tele.OnAddedToGroup, func(m *tele.Message) {
		b.Bot.Send(m.Chat, "Привет всем!")

	})
}
