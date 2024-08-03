package mux

import (
	"JerryBot/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
)

type record struct {
	Word     string
	Language string
}

type profanityStruct struct {
	Records []record
}

func (m *Mux) GPT(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {
	resp, err := http.Get("https://raw.githubusercontent.com/turalus/encycloDB/master/Dirty%20Words/DirtyWords.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	pResp := &profanityStruct{}
	if err = json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		fmt.Println(err)
	}

	client := openai.NewClient(config.OpenApiKey)

	response, err := getResponse(client, ctx.Content)
	if err != nil {
		fmt.Println(err)
	}

	embed := &discordgo.MessageEmbed{}
	embed.Title = "Jerry Says..."
	embed.Color = 16766720

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	embed.Thumbnail.URL = "https://cdn.discordapp.com/avatars/1249943053226217593/8978a12ea9467ceebfbb18c367b7bd5a.webp?size=128"

	if success, errMsg := moderationTest(response, pResp); success {
		embed.Description = response
	} else {
		embed.Description = errMsg
	}

	_, err = ds.ChannelMessageSendEmbed(dm.ChannelID, embed)

	if err != nil {
		fmt.Println(err)
	}

}

func getResponse(client *openai.Client, question string) (response string, err error) {
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: question,
			},
		},
		MaxTokens:   3000,
		Temperature: 0.5,
	})
	if err != nil {
		return "", err
	}

	response = resp.Choices[0].Message.Content
	response = strings.TrimLeft(response, "\n")

	return response, nil

}

func moderationTest(str string, pResp *profanityStruct) (bool, string) {
	for _, record := range pResp.Records {
		if strings.Contains(" "+strings.ToLower(str)+" ", " "+record.Word+" ") && record.Language == "en" {
			return false, "Whoa there buddy! Let's not say bad words"
		}
		if strings.Contains(strings.ToLower(str), "@") {
			return false, "Whoa there buddy! Let's not do a ping okay?"
		}
	}
	return true, ""
}
