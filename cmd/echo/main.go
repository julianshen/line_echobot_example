package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julianshen/lbotx"
)

func main() {
	bot, e := lbotx.NewBot(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)

	log.Println("bot created")

	if e != nil {
		log.Fatal(e)
	}

	log.Println("bot OnText")
	bot.OnText(func(context *lbotx.BotContext, text string) (bool, error) {
		context.Messages.AddTextMessage(text)

		return true, nil
	})

	log.Println("register")
	http.Handle("/callback", bot)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
