package telegram

import (
	tele "gopkg.in/tucnak/telebot.v2"
)

var l []int
var cafeName string

var found = false

func (b *TeleBot) Authentication() {
	getCafe()
	//fmt.Print(s)
	b.Bot.Handle("/start", func(m *tele.Message) {
		if !m.Private() {
			return
		}

		b.Bot.Send(m.Sender, "Пожалуйста авторизуйтесь!", Auth)
	})

	b.Bot.Handle(&BtnAuthCafe, func(m *tele.Message) {
		b.Bot.Send(m.Sender, "Напишите название Кафе")
		b.Bot.Handle(tele.OnText, func(m *tele.Message) {
			for i, j := range cafename {

				if j == m.Text {
					l = append(l, i)
					cafeName = j
					found = true
					b.Bot.Send(m.Sender, "Напишите пароль")
					b.Bot.Handle(tele.OnText, func(m *tele.Message) {
						for i, j := range cafepass {
							if i == l[0] && j == m.Text {
								found = false
								b.Bot.Send(m.Sender, "Вы успешно авторизовались! Выберите команды", menuAdmin)
								b.adminMenu()
							}
						}
						if found == true {
							b.Bot.Send(m.Sender, "Неправильный пароль, попробуйте написать пароль ещё раз!")
						}
					})

				}

			}
			if found == false {
				b.Bot.Send(m.Sender, "Неправильное название кафе, попробуйте написать название ещё раз!")
			}

		})
	})

	b.Bot.Handle(&BtnAuthCompany, func(m *tele.Message) {
		getCompany()
		b.Bot.Send(m.Sender, "Напишите пароль")
		b.Bot.Handle(tele.OnText, func(m *tele.Message) {
			if adminpassword == m.Text {
				b.Bot.Send(m.Sender, "Вы успешно авторизовались!")
			} else {
				b.Bot.Send(m.Sender, "Неправильный пароль, попробуйте написать пароль ещё раз!")
			}

		})
	})

}
