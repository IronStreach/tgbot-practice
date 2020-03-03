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

func main() {
	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	data := GetWeather("https://api.weather.yandex.ru/v1/forecast?lat=55.011897&lon=36.462555&extra=true", "6a653901-d939-47c7-8868-db449fd6a7df")
	var temper string
	city := "Maloyaroslavets"
	if  data.Fact.Temp > 0 {
		temper = "+" + strconv.Itoa(int(data.Fact.Temp))
	} else {
		temper = strconv.Itoa(int(data.Fact.Temp))
	}

	keyboard := tgbotapi.InlineKeyboardMarkup{}
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData("Показать погоду", "/weather")
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	bot, err := tgbotapi.NewBotAPI("931561769:AAEFSazicKW9Axrr_lYakkTv5S2WSFTUu6E")
	if err != nil {
		panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bot doesn't done yet. Please, be patient!")
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		case "/weather":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Weather in your city: \n" + city + ": " + temper + "°C")
			bot.Send(msg)
		}

	}

}
