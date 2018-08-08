package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/subosito/gotenv"
)

// BotID globally stores the bot's ID
var BotID string

func init() { gotenv.Load() }

func main() {
	// Create discord session
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("Error while creating discord session:\n", err)
		return
	}

	// Get user info for the bot
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("Error while getting user info:\n", err)
		return
	}
	// Get the id of the bot
	BotID = u.ID

	dg.AddHandler(MessageHandler)

	// Open the websocket and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("Error while openning the websocket:\n", err)
		return
	}

	fmt.Println("Discord Bot is running!")

	defer dg.Close()
	<-make(chan struct{})
}
