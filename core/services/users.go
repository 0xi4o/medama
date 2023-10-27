package services

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/medama-io/medama/api"
	"github.com/medama-io/medama/model"
	"go.jetpack.io/typeid"
)

func (h *Handler) GetUser(ctx context.Context, params api.GetUserParams) (api.GetUserRes, error) {
	userId, ok := ctx.Value(model.ContextKeyUserID).(string)
	if !ok {
		return ErrUnauthorised(model.ErrSessionNotFound), nil
	}

	user, err := h.db.GetUser(ctx, userId)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return ErrNotFound(err), nil
		}

		return nil, err
	}

	return &api.UserGet{
		Email:       user.Email,
		Language:    api.UserGetLanguage(user.Language),
		DateCreated: user.DateCreated,
		DateUpdated: user.DateUpdated,
	}, nil
}

func (h *Handler) PatchUser(ctx context.Context, req api.OptUserPatch, params api.PatchUserParams) (api.PatchUserRes, error) {
	// Check user exists
	userId, ok := ctx.Value(model.ContextKeyUserID).(string)
	if !ok {
		return ErrUnauthorised(model.ErrSessionNotFound), nil
	}

	user, err := h.db.GetUser(ctx, userId)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return ErrNotFound(err), nil
		}

		return nil, err
	}

	// Update values
	dateUpdated := time.Now().Unix()
	email := req.Value.Email.Value
	if email != "" {
		user.Email = email
		err = h.db.UpdateUserEmail(ctx, user.ID, email, dateUpdated)
		if err != nil {
			return nil, err
		}
	}

	password := req.Value.Password.Value
	if password != "" {
		pwdHash, err := h.auth.HashPassword(password)
		if err != nil {
			return nil, err
		}

		err = h.db.UpdateUserPassword(ctx, user.ID, pwdHash, dateUpdated)
		if err != nil {
			return nil, err
		}
	}

	return &api.UserGet{
		Email:       user.Email,
		Language:    api.UserGetLanguage(user.Language),
		DateCreated: user.DateCreated,
		DateUpdated: user.DateUpdated,
	}, nil
}

func (h *Handler) PostUser(ctx context.Context, req api.OptUserCreate) (api.PostUserRes, error) {
	// Do not allow user account creation if number of users exceeds 1
	/* count, err := h.db.GetUserCount(ctx)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		slog.WarnContext(ctx, "maximum number of users reached when creating user", slog.Int("count", count))
		return ErrForbidden(model.ErrUserMax), nil
	} */

	// UUIDv7 id generation
	typeid, err := typeid.New("user")
	if err != nil {
		return nil, err
	}
	id := typeid.String()
	email := req.Value.Email

	// Validate language as an accepted enum
	err = req.Value.Language.Value.Validate()
	if err != nil {
		return nil, err
	}
	language := string(req.Value.Language.Value)

	dateCreated := time.Now().Unix()
	dateUpdated := dateCreated

	attributes := []slog.Attr{
		slog.String("id", id),
		slog.String("email", email),
		slog.String("language", language),
		slog.Int64("date_created", dateCreated),
		slog.Int64("date_updated", dateUpdated),
	}
	slog.LogAttrs(ctx, slog.LevelDebug, "creating user", attributes...)

	// Hash password
	pwdHash, err := h.auth.HashPassword(req.Value.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:          id,
		Email:       email,
		Password:    pwdHash,
		Language:    language,
		DateCreated: dateCreated,
		DateUpdated: dateUpdated,
	}

	err = h.db.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, model.ErrUserExists) {
			slog.WarnContext(ctx, "user already exists when creating user", slog.String("id", id))
			return ErrConflict(err), nil
		}

		return nil, err
	}

	// Add session to cache
	cookie, err := h.auth.CreateSession(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &api.UserGetHeaders{
		SetCookie: api.NewOptString(cookie.String()),
		Response: api.UserGet{
			Email:       user.Email,
			Language:    api.UserGetLanguage(user.Language),
			DateCreated: dateCreated,
			DateUpdated: dateUpdated,
		},
	}, nil
}
