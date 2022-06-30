package telegram

import (
	"fmt"
	tele "gopkg.in/tucnak/telebot.v2"
)

var f []string
var s []string
var a []string
var d []string
var save = false
var i = 0

func (b *TeleBot) adminMenu() {
	b.Bot.Handle(&BtnNewMenu, func(m *tele.Message) {
		if !m.Private() {
			return
		}
		for i == 0 {
			i += 1
			b.Bot.Send(m.Sender, "Введите первые блюда через запятую")
			b.Bot.Handle(tele.OnText, func(m *tele.Message) {
				f = append(f, m.Text)
				b.Bot.Send(m.Sender, "Введите вторые блюда через запятую")
				b.Bot.Handle(tele.OnText, func(m *tele.Message) {
					save = false
					s = append(s, m.Text)
					b.Bot.Send(m.Sender, "Введите апитайзеры через запятую ")
					b.Bot.Handle(tele.OnText, func(m *tele.Message) {
						save = true
						a = append(a, m.Text)
						b.Bot.Send(m.Sender, "Введите напитоки через запятую ")
						b.Bot.Handle(tele.OnText, func(m *tele.Message) {
							d = append(d, m.Text)
							b.Bot.Send(m.Sender, "Меню сохранено!", menuAdmin)
							createMenu(f, s, a, d)
						})

					})

				})

			})

		}

	})
	b.Bot.Handle(&BtnShowMenu, func(m *tele.Message) {
		showMenu()
		b.Bot.Send(m.Sender, "Вот меню ", menuAdmin)
		fmt.Println(firstDish, secondDish, apitizerDish, drink)

	})
}
