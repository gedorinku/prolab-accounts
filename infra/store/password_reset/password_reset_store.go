package resetstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"golang.org/x/crypto/bcrypt"

	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/infra/store"
	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/sqlutil"
)

type passwordResetStoreImpl struct {
	ctx context.Context
	db  *sqlutil.DB
}

// NewPasswordResetStore returns new password reset store.
func NewPasswordResetStore(ctx context.Context, db *sqlutil.DB) store.PasswordResetStore {
	return &passwordResetStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

const (
	tokenLength = 32
	lifeTime    = 30 * time.Minute
)

func (s *passwordResetStoreImpl) CreateConfirmation(userID model.UserID, email string) (r *record.PasswordReset, token string, err error) {
	token, err = model.GenerateSecureToken(tokenLength)
	if err != nil {
		return nil, "", err
	}

	d, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.WithStack(err)
	}

	r = &record.PasswordReset{
		TokenDigest: string(d),
		Email:       email,
		UserID:      int64(userID),
	}
	err = s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		_, err := record.PasswordResets(record.PasswordResetWhere.Email.EQ(email)).DeleteAll(ctx, tx)
		if err != nil {
			return errors.WithStack(err)
		}

		err = r.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return errors.WithStack(err)
		}

		err = r.L.LoadUser(ctx, tx, true, r, nil)
		return errors.WithStack(err)
	})
	if err != nil {
		return nil, "", err
	}

	return r, token, nil
}

func (s *passwordResetStoreImpl) GetConfirmation(email, token string) (*record.PasswordReset, error) {
	r, err := record.PasswordResets(record.PasswordResetWhere.Email.EQ(email)).One(s.ctx, s.db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if lifeTime < time.Since(r.CreatedAt) {
		return nil, sql.ErrNoRows
	}

	err = bcrypt.CompareHashAndPassword([]byte(r.TokenDigest), []byte(token))
	if err != nil {
		return nil, sql.ErrNoRows
	}

	err = r.L.LoadUser(s.ctx, s.db, true, r, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

func (s *passwordResetStoreImpl) UpdatePassword(reset *record.PasswordReset, password string) error {
	d, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithStack(err)
	}

	err = s.db.Watch(s.ctx, func(ctx context.Context, tx *sql.Tx) error {
		u := reset.R.User
		u.PasswordDigest = string(d)
		_, err := u.Update(ctx, tx, boil.Whitelist("password_digest", "updated_at"))
		if err != nil {
			return errors.WithStack(err)
		}

		_, err = reset.Delete(ctx, tx)
		return errors.WithStack(err)
	})

	return err
}
