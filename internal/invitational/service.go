package invitational

type InvitationalService interface {
	CreateInvitation(invitation Invitation) (bool, error)
	GetSentInvitations() (*[]Invitation, error)
}

type InvitationService struct {
	repo InvitationalRepository
}

func New(repo InvitationalRepository) InvitationalService {
	return &InvitationService{repo}
}

func (inv *InvitationService) CreateInvitation(invitation Invitation) (bool, error) {
	saved, err := inv.repo.SaveInvitation(invitation)
	if err != nil {
		return false, err
	}
	return saved, nil
}

func (inv *InvitationService) GetSentInvitations() (*[]Invitation, error) {
	invitations, err := inv.repo.GetInvitations()
	if err != nil {
		return nil, err
	}
	return &invitations, nil
}
