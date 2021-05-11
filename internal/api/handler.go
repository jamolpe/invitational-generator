package api

import (
	"github.com/labstack/echo"
)

func (api *API) GetInvitations(c echo.Context) error {
	invitations, err := api.service.GetSentInvitations()
	if err != nil {
		return c.JSON(500, "error saving")
	}
	return c.JSON(200, invitations)
}

func (api *API) CreateInvitation(c echo.Context) error {
	invitationRequest := InvitationRequest{}
	err := c.Bind(&invitationRequest)
	if err == nil {
		created := api.service.CreateInvitation(invitationRequest.Invitation, invitationRequest.ClientData)
		if created {
			return c.JSON(200, true)
		}
	}
	return c.JSON(500, "invitation not created")
}
