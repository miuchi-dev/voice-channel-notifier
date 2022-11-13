package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("Error creating discord session:", err)
		return
	}

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	s.AddHandler(onReady)
	s.AddHandler(onVoiceStateUpdate)

	err = s.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	defer s.Close()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Logged in as ", event.User.Username, "(", event.User.ID, ")")
}

func onVoiceStateUpdate(s *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	if event.ChannelID == "" && event.BeforeUpdate != nil {
		s.ChannelMessageSend(os.Getenv("TEXT_CHANNEL_ID"), fmt.Sprint(event.Member.User.Username, "が <#", event.BeforeUpdate.ChannelID, "> から退出しました"))
	} else if event.ChannelID != "" && event.BeforeUpdate == nil {
		s.ChannelMessageSend(os.Getenv("TEXT_CHANNEL_ID"), fmt.Sprint(event.Member.User.Username, "が <#", event.ChannelID, "> に参加しました"))
	}
}
