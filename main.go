package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

const trackingUserID = "user ID"

func main() {
	client, _ := discordgo.New("bot token or user token")
	lastStatus := "-"
	client.AddHandler(func (client *discordgo.Session, user *discordgo.PresenceUpdate) {
		if user.User.ID == trackingUserID {
			if string(user.Presence.Status) != lastStatus {
				lastStatus = string(user.Presence.Status)
				channel, _ := client.UserChannelCreate(user.User.ID)
				client.ChannelMessageSend(channel.ID, "User is now "+string(user.Presence.Status))
			}
		}
	})
	client.Identify.Intents = discordgo.IntentsGuildPresences
	client.Open()
	fmt.Println("Tracker is now working!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	client.Close()
}
