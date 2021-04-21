package api

import (
	"github.com/jamolpe/invitational-generator/internal/invitational"
	"github.com/jamolpe/invitational-generator/internal/mailer"
)

type InvitationRequest struct {
	Invitation invitational.Invitation `json:"invitation" bson:"invitation"`
	ClientData mailer.MailClient       `json:"clientData" bson:"clientData"`
}
