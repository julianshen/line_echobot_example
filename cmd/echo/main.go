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

	if e != nil {
		log.Fatal(e)
	}

	bot.OnText(func(context *lbotx.BotContext, text string) (bool, error) {
		context.Messages.AddTextMessage(text)

		return true, nil
	})

	http.Handle("/callback", bot)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
