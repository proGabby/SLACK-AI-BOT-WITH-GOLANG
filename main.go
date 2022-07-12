package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

// var  wolframClient *wolfram.Client

func displayCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	//read .env file to ENV process
	godotenv.Load(".env")

	//get env tokens from .env
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	slackAppToken := os.Getenv("SLACK_APP_TOKEN")
	witAI_Token, isWitToken := os.LookupEnv("WIT_AI_TOKEN")
	wolframAppId, isWolramId := os.LookupEnv("WOLFRAM_APP_ID")

	//check if wittoken is available
	if !isWitToken {
		log.Fatal("No variable found for wit-ai-token")
	}

	//check if wolframId is available
	if !isWolramId {
		log.Fatal("No variable found for wit-ai-token")
	}

	//create a new client using the Slack API
	bot := slacker.NewClient(botToken, slackAppToken)

	//create a Wit.ai client for default API version
	witClient := witai.NewClient(witAI_Token)
	//create a wolfram client using the app id
	wolframClient := &wolfram.Client{AppID: wolframAppId}

	//printcomments here
	go displayCommandEvents(bot.CommandEvents())

	bot.Command(" <msg>", &slacker.CommandDefinition{
		Description: "direct my question to wolfram",
		Example:     " where is Nigeria?",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			//get message from request
			msgQuery := request.Param("msg")

			//get the message to witai
			witResponse, err := witClient.Parse(&witai.MessageRequest{
				Query: msgQuery,
			})
			//check for error
			if err != nil {
				log.Fatal(err)
			}
			//format the witresponse to json
			jsonData, _ := json.MarshalIndent(witResponse, "", "   ")

			//stringify the data
			dataAsString := string(jsonData[:])

			//get the desire value from the json using gjson package
			gjsonResult := gjson.Get(dataAsString, "entities.with$wolfram_search_query:wolfram_search_query.0.value")
			//put the gjson result in a string format
			question := gjsonResult.String()
			//query the question from wolfram
			wolframRespond, err := wolframClient.GetSpokentAnswerQuery(question, wolfram.Metric, 1000)

			//check for error
			if err != nil {
				fmt.Println(err)
				fmt.Println("there is an error fetching a answer to your response")
			}
			//send back response to slack
			response.Reply(wolframRespond)

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}

}
