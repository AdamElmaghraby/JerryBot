package main

import (
	"JerryBot/config"
	"JerryBot/mux"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	bot    *discordgo.Session
	Router *mux.Mux
)

func init() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
	}

	bot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err)
	}

	err = bot.Open()

	if err != nil {
		fmt.Println(err)
	}

	Router = mux.New()
	Router.Prefix = config.BotPrefix

	bot.AddHandler(Router.OnMessageCreate)

	Router.Route("ping", "Ping command that reuturns latency", Router.Ping)
	Router.Route("embed", "Returns an embed", Router.Embed)
	Router.Route("chat", "Chats with Jerry using OpenAI", Router.GPT)

}

func main() {

	fmt.Println("Bot is now running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}

// func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	if m.Author.ID == BotID {
// 		return
// 	}

// 	if m.Content == config.BotPrefix+"ping" {
// 		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!"+(time.Since(m.Timestamp)).String())

// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}

// 	if m.Content == config.BotPrefix+"embed" {
// 		embed := &discordgo.MessageEmbed{}
// 		embed.Title = "Jerry Says مرحبًا"
// 		embed.Color = 1752220
// 		embed.Description = "This translates to give me your family"

// 		embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
// 		embed.Thumbnail.URL = "https://www.marthastewart.com/thmb/gCXKR-31DYnpsLi7uUj0S4zyfqc=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/happy-labrador-retriever-getty-0322-2000-eb585d9e672e47da8b1b7e9d3215a5cb.jpg"

// 		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)

// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}
// }
