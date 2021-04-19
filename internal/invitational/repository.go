package invitational

type InvitationalRepository interface {
	SaveInvitation(invitation Invitation) (bool, error)
	GetInvitations() ([]Invitation, error)
}
