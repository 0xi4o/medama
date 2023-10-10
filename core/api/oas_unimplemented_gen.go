// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// DeleteWebsitesID implements delete-websites-id operation.
//
// Delete a website.
//
// DELETE /websites/{hostname}
func (UnimplementedHandler) DeleteWebsitesID(ctx context.Context, params DeleteWebsitesIDParams) (r DeleteWebsitesIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetEventPing implements get-event-ping operation.
//
// This is a ping endpoint to determine if the user is unique or not.
//
// GET /event/ping
func (UnimplementedHandler) GetEventPing(ctx context.Context, params GetEventPingParams) (r GetEventPingRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUsersUserId implements get-users-userId operation.
//
// Retrieve the information of the user with the matching user ID.
//
// GET /users/{userId}
func (UnimplementedHandler) GetUsersUserId(ctx context.Context, params GetUsersUserIdParams) (r GetUsersUserIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWebsiteIDSummary implements get-website-id-summary operation.
//
// Get a summary of the website's stats.
//
// GET /website/{hostname}/summary
func (UnimplementedHandler) GetWebsiteIDSummary(ctx context.Context, params GetWebsiteIDSummaryParams) (r GetWebsiteIDSummaryRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWebsites implements get-websites operation.
//
// Get a list of all websites from the user.
//
// GET /websites
func (UnimplementedHandler) GetWebsites(ctx context.Context, params GetWebsitesParams) (r GetWebsitesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWebsitesID implements get-websites-id operation.
//
// Get website details for an individual website.
//
// GET /websites/{hostname}
func (UnimplementedHandler) GetWebsitesID(ctx context.Context, params GetWebsitesIDParams) (r GetWebsitesIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWebsitesIDActive implements get-websites-id-active operation.
//
// Return the number of active users who triggered a pageview in the past 5 minutes.
//
// GET /websites/{hostname}/active
func (UnimplementedHandler) GetWebsitesIDActive(ctx context.Context, params GetWebsitesIDActiveParams) (r GetWebsitesIDActiveRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PatchUsersUserId implements patch-users-userId operation.
//
// Update a user account's details.
//
// PATCH /users/{userId}
func (UnimplementedHandler) PatchUsersUserId(ctx context.Context, req OptUserPatch, params PatchUsersUserIdParams) (r PatchUsersUserIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PatchWebsitesID implements patch-websites-id operation.
//
// Update a website's information.
//
// PATCH /websites/{hostname}
func (UnimplementedHandler) PatchWebsitesID(ctx context.Context, req OptWebsitePatch, params PatchWebsitesIDParams) (r PatchWebsitesIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PostAuthLogin implements post-auth-login operation.
//
// Login to the service and retrieve a session token for authentication.
//
// POST /auth/login
func (UnimplementedHandler) PostAuthLogin(ctx context.Context, req OptPostAuthLoginReq) (r PostAuthLoginRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PostEventHit implements post-event-hit operation.
//
// Send a hit event to register a user view.
//
// POST /event/hit
func (UnimplementedHandler) PostEventHit(ctx context.Context, req OptEventHit, params PostEventHitParams) (r PostEventHitRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PostUser implements post-user operation.
//
// Create a new user.
//
// POST /users
func (UnimplementedHandler) PostUser(ctx context.Context, req OptUserCreate) (r PostUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PostWebsites implements post-websites operation.
//
// Add a new website.
//
// POST /websites
func (UnimplementedHandler) PostWebsites(ctx context.Context, req OptWebsiteCreate) (r PostWebsitesRes, _ error) {
	return r, ht.ErrNotImplemented
}
