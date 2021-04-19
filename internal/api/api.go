package api

import "github.com/jamolpe/invitational-generator/internal/invitational"

type API struct {
	service invitational.InvitationalService
}

func New(serv invitational.InvitationalService) *API {
	return &API{service: serv}
}
