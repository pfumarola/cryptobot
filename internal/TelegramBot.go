package internal

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

type TelegramBot struct {
	bot        *gotgbot.Bot
	dispatcher *ext.Dispatcher
	updater    *ext.Updater
}

func NewTelegramBot(token string) *TelegramBot {

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		BotClient: &gotgbot.BaseBotClient{
			Client: http.Client{},
			DefaultRequestOpts: &gotgbot.RequestOpts{
				Timeout: gotgbot.DefaultTimeout, // Customise the default request timeout here
				APIURL:  gotgbot.DefaultAPIURL,  // As well as the Default API URL here (in case of using local bot API servers)
			},
		},
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	// Add echo handler to reply to all text messages.
	dispatcher.AddHandler(handlers.NewMessage(message.Text, echo))

	return &TelegramBot{
		bot:        b,
		dispatcher: dispatcher,
		updater:    updater,
	}
}

func (t *TelegramBot) Start() {

	// Start receiving updates.
	err := t.updater.StartPolling(t.bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})

	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", t.bot.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	t.updater.Idle()
}

func (t *TelegramBot) SendMessage(chatID string, message string) {
	intChatID, err1 := strconv.ParseInt(chatID, 10, 64)

	if err1 != nil {
		fmt.Println(err1)
	}

	_, err := t.bot.SendMessage(intChatID, message, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// echo replies to a messages with its own contents.
func echo(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, ctx.EffectiveMessage.Text, nil)
	if err != nil {
		return fmt.Errorf("failed to echo message: %w", err)
	}
	return nil
}
