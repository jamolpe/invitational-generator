package invitational

type InvitationalRepository interface {
	SaveInvitation(invitation Invitation)
	GetInvitations() ([]Invitation, error)
}
