package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(analyticsChan <-chan *slacker.CommandEvent) {
	for event := range analyticsChan {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06B1QXGA15-6383750574150-7f9070cd751807f9429996849d61b24636ea56443bc3507a4a618209aafd7f00")
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6387454136293-6375854698087-7ZyXE4hUZvH9pn84R4Ic94qE")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
w
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 2020"},
		Handler: func(botContext slacker.BotContext, request slacker.Request, writer slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("Error while converting string to int")
			}
			age := 2023 - yob
			r := fmt.Sprintf("age is %d", age)
			writer.Reply(r)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}
