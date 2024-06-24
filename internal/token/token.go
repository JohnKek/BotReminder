package token

import (
	"flag"
	"log"
)

func MustToken() string {
	token := flag.String("token-bot-token",
		"",
		"Telegram access token",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is empty")
	}
	return *token
}

func MustHost() string {
	host := flag.String("host-bot-host",
		"",
		"Telegram access host",
	)
	flag.Parse()

	if *host == "" {
		log.Fatal("token is empty")
	}
	return *host
}
