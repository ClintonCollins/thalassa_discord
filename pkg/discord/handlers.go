package discord

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/bwmarrin/discordgo"
)

// guildCreate allows us to organize methods for just guild creation.
type guildCreate struct{}

// guildMemberAdd allows us to organize methods just for guild members joining.
type guildMemberAdd struct{}

// messageCreate allows us to organize methods just for message creation.
type messageCreate struct{}

func (s *ShardInstance) guildCreate(dSession *discordgo.Session, guildCreate *discordgo.GuildCreate) {
	log.Info().Msgf("Joined Server: %s", guildCreate.Name)

	serverInfo, err := s.handlers.guildCreate.loadOrCreateDiscordGuildFromDatabase(s.Ctx, s.Db, guildCreate)
	if err != nil {
		log.Error().Err(err).Msg("Error loading or creating server")
		return
	}

	serverInstance := s.handlers.guildCreate.createDiscordGuildInstance(s.Ctx, s.Db, serverInfo, dSession, guildCreate)

	s.addServerInstance(guildCreate.ID, serverInstance)

	if s.SongQueueUpdateCallback != nil {
		s.SongQueueUpdateCallbackMutex.Lock()
		serverInstance.SongQueueUpdateCallbackMutex = s.SongQueueUpdateCallbackMutex
		serverInstance.SongQueueUpdateCallback = s.SongQueueUpdateCallback
		s.SongQueueUpdateCallbackMutex.Unlock()
	}

	s.handlers.guildCreate.startMusicBot(serverInstance, guildCreate)
	s.handlers.guildCreate.createEveryoneRolePermissionsIfNotExist(serverInstance, guildCreate)
	s.handlers.guildCreate.handleMutedRole(serverInstance, guildCreate)
	s.handlers.guildCreate.handleNotifyRole(serverInstance)
	if s.GuildCreatedFunction != nil {
		s.GuildCreatedFunction(dSession, guildCreate)
	}
}

func (s *ShardInstance) guildMemberAdd(dSession *discordgo.Session, guildMemberAdd *discordgo.GuildMemberAdd) {
	s.RLock()
	serverInstance, exists := s.ServerInstances[guildMemberAdd.GuildID]
	if !exists {
		return
	}
	s.RUnlock()
	s.handlers.guildMemberAdd.checkNewUserForMute(serverInstance, guildMemberAdd)
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func (s *ShardInstance) messageCreate(dSession *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	message := messageCreate.Content
	if len(message) <= 0 {
		return
	}
	s.RLock()
	serverInstance, _ := s.ServerInstances[messageCreate.GuildID]
	s.RUnlock()

	// s.handlers.messageCreate.checkForRolePermsSet(serverInstance, messageCreate)

	commandFound, commandName, args := s.parseMessageForCommand(messageCreate.Message, serverInstance)
	if commandFound {
		// Custom commands can override built-in commands.
		start := time.Now()
		if !s.handleCustomCommand(commandName, args, messageCreate.Message, serverInstance) {
			s.handleCommand(commandName, args, messageCreate.Message, serverInstance)
		}
		duration := time.Since(start)
		log.Info().Fields(map[string]interface{}{
			"Username": messageCreate.Author.Username,
			"Command":  commandName,
			"Message":  messageCreate.Content,
		}).Msgf("%s took %v to fully run.", commandName, duration)
	}
}

func (s *ShardInstance) guildMemberUpdate(dSession *discordgo.Session, guildMemberUpdate *discordgo.GuildMemberUpdate) {
}
