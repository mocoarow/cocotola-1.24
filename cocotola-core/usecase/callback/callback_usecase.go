package callback

import (
	"context"
	"log/slog"

	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type Callback struct {
	logger *slog.Logger
}

func NewCallback() *Callback {
	return &Callback{
		logger: slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CallbackUsecase"))}
}

func (u *Callback) OnAddAppUser(ctx context.Context, appUserID *mbuserdomain.AppUserID) error {
	u.logger.InfoContext(ctx, "OnAddAppUser", slog.Int("app_user_id", appUserID.Int()))
	return nil
}
