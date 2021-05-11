package invitational

import (
	"math/rand"
	"strings"

	"github.com/jamolpe/invitational-generator/internal/mailer"
)

type InvitationalService interface {
	CreateInvitation(invitation Invitation, client mailer.MailClient) bool
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func createCode() string {
	code := randStringBytesMaskImprSrcSB(10)
	return code
}

func (inv *InvitationService) CreateInvitation(invitation Invitation, client mailer.MailClient) bool {
	code := createCode()
	invitation.Code = code
	go inv.repo.SaveInvitation(invitation)
	go inv.mailer.Send("./templates/template.html", map[string]string{"email": invitation.Email, "code": invitation.Code}, mailer.MailClient{To: []string{invitation.Email}, Subject: "invitation", Body: ""})
	return true
}

func (inv *InvitationService) GetSentInvitations() (*[]Invitation, error) {
	invitations, err := inv.repo.GetInvitations()
	if err != nil {
		return nil, err
	}
	return &invitations, nil
}
