package email

import (
	"bytes"
	"context"

	"github.com/jordan-wright/email"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/static"
)

// Sender represents interface of email sender
type Sender interface {
	SendInvitationEmail(req *record.Invitation) error
	SendEmailConfirmation(conf *record.EmailConfirmation) error
}

// NewSender creates new sender
func NewSender(ctx context.Context, cfg *config.Config, asset *static.EmailAsset) Sender {
	return &senderImpl{
		ctx:   ctx,
		cfg:   cfg,
		asset: asset,
	}
}

type senderImpl struct {
	ctx   context.Context
	cfg   *config.Config
	asset *static.EmailAsset
}

const (
	subjectPrefix = "[プロラボアカウント]"
)

type invitationEmailData struct {
	RegistrationURL string
}

// SendInvitationEmail sends invitation email
func (s *senderImpl) SendInvitationEmail(inv *record.Invitation) error {
	tmpl, err := s.asset.GetTemplate("invitation.tmpl")
	if err != nil {
		return err
	}

	d := invitationEmailData{
		RegistrationURL: s.cfg.ClientRegistrationURL + "/" + inv.Code,
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, d)
	if err != nil {
		return errors.WithStack(err)
	}

	e := email.NewEmail()
	e.From = s.cfg.EmailFrom
	e.To = []string{inv.Email}
	e.Subject = subjectPrefix + "ユーザー登録"
	e.Text = buf.Bytes()
	err = e.Send(s.cfg.SMTPAddr, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

type emailConfirmationData struct {
	Name            string
	Email           string
	ConfirmationURL string
}

// SendEmailConfirmation sends email confirmation
func (s *senderImpl) SendEmailConfirmation(conf *record.EmailConfirmation) error {
	tmpl, err := s.asset.GetTemplate("email_confirmation.tmpl")
	if err != nil {
		return err
	}

	d := emailConfirmationData{
		Name:            conf.R.User.Name,
		Email:           conf.Email,
		ConfirmationURL: s.cfg.ClientConfirmationURL + "/" + conf.Token,
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, d)
	if err != nil {
		return errors.WithStack(err)
	}

	e := email.NewEmail()
	e.From = s.cfg.EmailFrom
	e.To = []string{conf.Email}
	e.Subject = subjectPrefix + "メールアドレスの確認"
	e.Text = buf.Bytes()
	err = e.Send(s.cfg.SMTPAddr, nil)
	return errors.WithStack(err)
}
