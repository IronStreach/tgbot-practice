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
	var cities = GetCityName()
	var names = make([]string, len(cities))
	for i := 0; i < len(cities); i++ {
		names[i] = cities[i].city_name
	}
	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

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

	updates := bot.ListenForWebhook("/" + bot.Token)
	for update := range updates {
		if update.CallbackQuery != nil {
			callback := update.CallbackQuery.Data
			if callback == "weather" {
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Выберите город")
				bot.Send(msg)
				citiesKeyboard := tgbotapi.InlineKeyboardMarkup{}
				var row []tgbotapi.InlineKeyboardButton
				for i := 0; i < len(names); i++ {
					btn := tgbotapi.NewInlineKeyboardButtonData(names[i], names[i])
					row = append(row, btn)
				}
				citiesKeyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
				updates := bot.ListenForWebhook("/" + bot.Token)
				for update := range updates {
					if update.CallbackQuery != nil {
						callback := update.CallbackQuery.Data
						if callback == "Малоярославец" {
							temper, _ := GetWeather(cities[0].lat, cities[0].lon)
							msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Weather in your city: \nMaloyaroslavets"+": "+temper+"°C")
							bot.Send(msg)
						}
						}
					}
				//temper, _ := GetWeather(55.011897, 36.462555)
				//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Weather in your city: \nMaloyaroslavets"+": "+temper+"°C")
				//bot.Send(msg)
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
