package api

import (
	"github.com/jamolpe/invitational-generator/internal/invitational"
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
	invitationRequest := invitational.Invitation{}
	err := c.Bind(&invitationRequest)
	if err != nil {
		created, err := api.service.CreateInvitation(invitationRequest)
		if err != nil {
			return c.JSON(500, "error creating invitation")
		}
		if created {
			return c.JSON(200, true)
		}
	}
	return c.JSON(500, "invitation not created")
}
