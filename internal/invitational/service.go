package invitational

import "github.com/jamolpe/invitational-generator/internal/mailer"

type InvitationalService interface {
	CreateInvitation(invitation Invitation, client mailer.MailClient) (bool, error)
	GetSentInvitations() (*[]Invitation, error)
}

type InvitationService struct {
	mailer mailer.Mailer
	repo   InvitationalRepository
}

func New(repo InvitationalRepository) InvitationalService {
	mailer := mailer.New()
	return &InvitationService{repo: repo, mailer: mailer}
}

func (inv *InvitationService) CreateInvitation(invitation Invitation, client mailer.MailClient) (bool, error) {
	saved, err := inv.repo.SaveInvitation(invitation)
	errMail := inv.mailer.Send("./templates/template.html", map[string]string{"username": "user name"}, mailer.MailClient{To: []string{invitation.Email}, Subject: "invitation", Body: ""})
	if err != nil || errMail != nil {
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
