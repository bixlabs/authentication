package in_memory

import (
	"github.com/bixlabs/authentication/authenticator/structures"
)

type DummySender struct {
}

func (DummySender) SendEmailPasswordRequest(user structures.User, code string) error {
	return nil
}
