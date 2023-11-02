package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/ClintonCollins/thalassa_discord/pkg/api"
	"github.com/ClintonCollins/thalassa_discord/pkg/commands/example"
	"github.com/ClintonCollins/thalassa_discord/pkg/commands/general"
	"github.com/ClintonCollins/thalassa_discord/pkg/commands/lookup"
	"github.com/ClintonCollins/thalassa_discord/pkg/commands/moderation"
	"github.com/ClintonCollins/thalassa_discord/pkg/commands/music"
	"github.com/ClintonCollins/thalassa_discord/pkg/commands/random"
	"github.com/ClintonCollins/thalassa_discord/pkg/discord"
)

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()
	discordInstance, err := discord.NewInstance(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Register all created commands.
	moderation.RegisterCommands(discordInstance)
	lookup.RegisterCommands(discordInstance)
	music.RegisterCommands(discordInstance)
	random.RegisterCommands(discordInstance)
	example.RegisterCommands(discordInstance)
	general.RegisterCommands(discordInstance)

	if discordInstance.BotConfig.EnableAPI {
		apiInstance := api.New(discordInstance, discordInstance.BotConfig.APIHost, discordInstance.BotConfig.APIPort)
		discordInstance.SongQueueUpdateCallbackMutex.Lock()
		discordInstance.SongQueueUpdateCallback = apiInstance.SongQueueEventUpdate
		discordInstance.SongQueueUpdateCallbackMutex.Unlock()
		apiInstance.Start(discordInstance.Ctx)
	}

	go func() {
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			log.Println(err)
		}
	}()

	discordInstance.Start()
}
