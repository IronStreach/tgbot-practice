package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
	"os"
)

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm DndSpellsBot!"))
}

func main() {
	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	cities := GetCities()

	bot, err := tgbotapi.NewBotAPI("931561769:AAEFSazicKW9Axrr_lYakkTv5S2WSFTUu6E")
	if err != nil {
		panic(err)
	}
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	var row []tgbotapi.InlineKeyboardButton
	btn := tgbotapi.NewInlineKeyboardButtonData("Показать погоду", "weather")
	row = append(row, btn)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	names := GetCitiesName()
	//updates := bot.Li("/" + bot.Token)
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.CallbackQuery != nil {
			callback := update.CallbackQuery.Data
			if callback == "weather" {

				citiesKeyboard := tgbotapi.InlineKeyboardMarkup{}
				var rowCity []tgbotapi.InlineKeyboardButton
				for i := 0; i < len(names); i++ {
					btnCity := tgbotapi.NewInlineKeyboardButtonData(names[i], names[i])
					rowCity = append(rowCity, btnCity)
				}
				citiesKeyboard.InlineKeyboard = append(citiesKeyboard.InlineKeyboard, rowCity)
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Выберите город")
				msg.ReplyMarkup = citiesKeyboard
				bot.Send(msg)

				if update.CallbackQuery != nil {
					var temperature string
					for _, names := range cities {
						if names.city_name == callback {
							temperature = GetTemperature(names.lat, names.lon)
						} else {
							continue
						}
					}
					msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Weather in your city: \n"+callback+": "+temperature+"°C")
					bot.Send(msg)
				}
			} else {
				switch update.Message.Text {
				case "/start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bot doesn't done yet. Please, be patient!")
					msg.ReplyMarkup = keyboard
					bot.Send(msg)
				case "/weather":
					temper, _ := GetWeather(55.011897, 36.462555)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Weather in your city: \nMaloyaroslavets"+": "+temper+"°C")
					bot.Send(msg)
				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't understand you")
					bot.Send(msg)
				}
			}

		}
	}

}