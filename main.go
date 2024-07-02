package main

import (
	"flag"
	"log"
)

func main() {
	token := mustToken()

	// tgClient = telegram.New(Token)

	// fetcher  = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() (string, error) {
	// bot -tg-bot-token 'my token'
	token := flag.String(
		"token-bot-token",
		"",
		"Token for access to telegram",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
}
