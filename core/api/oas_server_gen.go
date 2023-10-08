// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DeleteWebsitesID implements delete-websites-id operation.
	//
	// Delete a website.
	//
	// DELETE /websites/{id}
	DeleteWebsitesID(ctx context.Context, params DeleteWebsitesIDParams) (DeleteWebsitesIDRes, error)
	// GetEventPing implements get-event-ping operation.
	//
	// This is a ping endpoint to determine if the user is unique or not.
	//
	// GET /event/ping
	GetEventPing(ctx context.Context, params GetEventPingParams) (GetEventPingRes, error)
	// GetUsersUserId implements get-users-userId operation.
	//
	// Retrieve the information of the user with the matching user ID.
	//
	// GET /users/{userId}
	GetUsersUserId(ctx context.Context, params GetUsersUserIdParams) (GetUsersUserIdRes, error)
	// GetWebsiteIDSummary implements get-website-id-summary operation.
	//
	// Your GET endpoint.
	//
	// GET /website/{id}/summary
	GetWebsiteIDSummary(ctx context.Context, params GetWebsiteIDSummaryParams) (GetWebsiteIDSummaryRes, error)
	// GetWebsites implements get-websites operation.
	//
	// Get the list of websites.
	//
	// GET /websites
	GetWebsites(ctx context.Context) (GetWebsitesRes, error)
	// GetWebsitesID implements get-websites-id operation.
	//
	// Get website details for an individual website.
	//
	// GET /websites/{id}
	GetWebsitesID(ctx context.Context, params GetWebsitesIDParams) (GetWebsitesIDRes, error)
	// GetWebsitesIDActive implements get-websites-id-active operation.
	//
	// Return the number of active users who triggered a pageview in the past 5 minutes.
	//
	// GET /websites/{id}/active
	GetWebsitesIDActive(ctx context.Context, params GetWebsitesIDActiveParams) (GetWebsitesIDActiveRes, error)
	// PatchUsersUserId implements patch-users-userId operation.
	//
	// Update a user account's details.
	//
	// PATCH /users/{userId}
	PatchUsersUserId(ctx context.Context, req OptUserPatch, params PatchUsersUserIdParams) (PatchUsersUserIdRes, error)
	// PatchWebsitesID implements patch-websites-id operation.
	//
	// Update a website's information.
	//
	// PATCH /websites/{id}
	PatchWebsitesID(ctx context.Context, req OptWebsitePatch, params PatchWebsitesIDParams) (PatchWebsitesIDRes, error)
	// PostAuthLogin implements post-auth-login operation.
	//
	// Login to the service and retrieve a JWT token for authentication.
	//
	// POST /auth/login
	PostAuthLogin(ctx context.Context, req OptPostAuthLoginReq, params PostAuthLoginParams) (PostAuthLoginRes, error)
	// PostAuthRefresh implements post-auth-refresh operation.
	//
	// Login to the service and retrieve a JWT token for authentication.
	//
	// POST /auth/refresh
	PostAuthRefresh(ctx context.Context, req OptPostAuthRefreshReq, params PostAuthRefreshParams) (PostAuthRefreshRes, error)
	// PostEventHit implements post-event-hit operation.
	//
	// Send a hit event to register a user view.
	//
	// POST /event/hit
	PostEventHit(ctx context.Context, req OptEventHit, params PostEventHitParams) (PostEventHitRes, error)
	// PostUser implements post-user operation.
	//
	// Create a new user.
	//
	// POST /users
	PostUser(ctx context.Context, req OptUserCreate) (PostUserRes, error)
	// PostWebsites implements post-websites operation.
	//
	// Add a new website.
	//
	// POST /websites
	PostWebsites(ctx context.Context, req OptWebsiteCreate) (PostWebsitesRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
