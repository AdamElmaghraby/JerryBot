package mux

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (m *Mux) Embed(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	embed := &discordgo.MessageEmbed{}
	embed.Title = "Jerry Says مرحبًا"
	embed.Color = 1752220
	embed.Description = "This translates to give me your meat"

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	embed.Thumbnail.URL = "https://www.marthastewart.com/thmb/gCXKR-31DYnpsLi7uUj0S4zyfqc=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/happy-labrador-retriever-getty-0322-2000-eb585d9e672e47da8b1b7e9d3215a5cb.jpg"

	_, err := ds.ChannelMessageSendEmbed(dm.ChannelID, embed)

	if err != nil {
		fmt.Println(err)
	}
}
