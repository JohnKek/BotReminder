package telegram

import "TransactionalOutbox/internal/clients/telegram"

type Processor struct {
	th     *telegram.Client
	offset int // storage
}

func New(client *telegram.Client) {

}
