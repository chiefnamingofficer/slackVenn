package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func getChannelMembers(api *slack.Client, channelID string) ([]string, error) {
	var members []string
	cursor := ""
	for {
		params := &slack.GetUsersInConversationParameters{
			ChannelID: channelID,
			Cursor:    cursor,
			Limit:     1000,
		}
		ids, nextCursor, err := api.GetUsersInConversation(params)
		if err != nil {
			return nil, err
		}
		members = append(members, ids...)
		if nextCursor == "" {
			break
		}
		cursor = nextCursor
	}
	return members, nil
}

func userIDToName(api *slack.Client, ids []string) map[string]string {
	names := make(map[string]string)
	for _, id := range ids {
		user, err := api.GetUserInfo(id)
		if err == nil {
			names[id] = user.Name
		} else {
			names[id] = "unknown"
		}
	}
	return names
}

func difference(a, b []string) []string {
	m := make(map[string]bool)
	for _, item := range b {
		m[item] = true
	}
	var diff []string
	for _, item := range a {
		if !m[item] {
			diff = append(diff, item)
		}
	}
	return diff
}

func intersection(a, b []string) []string {
	m := make(map[string]bool)
	var inter []string
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		if m[item] {
			inter = append(inter, item)
		}
	}
	return inter
}

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("SLACK_TOKEN env var is required")
	}

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <channel_id_1> <channel_id_2>\n\nslackVenn - Slack channel membership analyzer\nCreates Venn diagrams of user overlaps between channels", os.Args[0])
	}

	channelA := os.Args[1]
	channelB := os.Args[2]

	api := slack.New(token)

	fmt.Printf("📊 slackVenn: Analyzing channel membership overlap...\n")
	fmt.Printf("🔍 Channel A: %s\n", channelA)
	fmt.Printf("🔍 Channel B: %s\n", channelB)
	fmt.Println()

	membersA, err := getChannelMembers(api, channelA)
	if err != nil {
		log.Fatalf("Error getting members of channel A: %v", err)
	}

	membersB, err := getChannelMembers(api, channelB)
	if err != nil {
		log.Fatalf("Error getting members of channel B: %v", err)
	}

	allIDs := append(membersA, membersB...)
	usernames := userIDToName(api, allIDs)

	onlyA := difference(membersA, membersB)
	onlyB := difference(membersB, membersA)
	common := intersection(membersA, membersB)

	fmt.Printf("📈 Analysis Results:\n")
	fmt.Printf("   Channel A: %d members\n", len(membersA))
	fmt.Printf("   Channel B: %d members\n", len(membersB))
	fmt.Printf("   Overlap: %d members\n", len(common))
	fmt.Println()

	fmt.Println("🟢 Users in BOTH channels:")
	for _, id := range common {
		fmt.Println(" -", usernames[id])
	}

	fmt.Println("\n🔵 Users ONLY in Channel A:")
	for _, id := range onlyA {
		fmt.Println(" -", usernames[id])
	}

	fmt.Println("\n🟣 Users ONLY in Channel B:")
	for _, id := range onlyB {
		fmt.Println(" -", usernames[id])
	}
} 