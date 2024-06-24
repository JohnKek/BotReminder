package cmd

import (
	"TransactionalOutbox/internal/clients/telegram"
	"TransactionalOutbox/internal/token"
)

func main() {
	tgCleint := telegram.New(token.MustHost(), token.MustToken())
}
