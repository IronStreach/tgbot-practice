package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
	"os"
	"strconv"
)

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm DndSpellsBot!"))
}
func getWeatherData(lat float64, lon float64) string {
	latStr := strconv.FormatFloat(lat, 'f', 2, 64)
	lonStr := strconv.FormatFloat(lon, 'f', 2, 64)
	data := GetWeather("https://api.weather.yandex.ru/v1/forecast?", latStr, lonStr, "6a653901-d939-47c7-8868-db449fd6a7df")
	var temper string
	if data.Fact.Temp > 0 {
		temper = "+" + strconv.Itoa(int(data.Fact.Temp))
	} else {
		temper = strconv.Itoa(int(data.Fact.Temp))
	}
	return temper
}
func main() {
	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	//go getWeatherData()

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
				temper := getWeatherData(55.011897, 36.462555)
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Weather in your city: \nMaloyaroslavets"+": "+temper+"°C")
				bot.Send(msg)
			}
		} else {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bot doesn't done yet. Please, be patient!")
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			case "/weather":
				temper := getWeatherData(55.011897, 36.462555)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Weather in your city: \nMaloyaroslavets"+": "+temper+"°C")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't understand you")
				bot.Send(msg)
			}
		}

	}

}
