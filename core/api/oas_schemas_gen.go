// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
)

type BadRequestError struct {
	Error BadRequestErrorError `json:"error"`
}

// GetError returns the value of Error.
func (s *BadRequestError) GetError() BadRequestErrorError {
	return s.Error
}

// SetError sets the value of Error.
func (s *BadRequestError) SetError(val BadRequestErrorError) {
	s.Error = val
}

func (*BadRequestError) getWebsiteIDSummaryRes() {}
func (*BadRequestError) getWebsitesIDActiveRes() {}
func (*BadRequestError) getWebsitesIDRes()       {}
func (*BadRequestError) getWebsitesRes()         {}
func (*BadRequestError) patchUsersUserIdRes()    {}
func (*BadRequestError) patchWebsitesIDRes()     {}
func (*BadRequestError) postAuthLoginRes()       {}
func (*BadRequestError) postAuthRefreshRes()     {}
func (*BadRequestError) postUserRes()            {}
func (*BadRequestError) postWebsitesRes()        {}

type BadRequestErrorError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *BadRequestErrorError) GetCode() int32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *BadRequestErrorError) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *BadRequestErrorError) SetCode(val int32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *BadRequestErrorError) SetMessage(val string) {
	s.Message = val
}

type ConflictError struct {
	Error ConflictErrorError `json:"error"`
}

// GetError returns the value of Error.
func (s *ConflictError) GetError() ConflictErrorError {
	return s.Error
}

// SetError sets the value of Error.
func (s *ConflictError) SetError(val ConflictErrorError) {
	s.Error = val
}

func (*ConflictError) patchUsersUserIdRes() {}
func (*ConflictError) postUserRes()         {}
func (*ConflictError) postWebsitesRes()     {}

type ConflictErrorError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *ConflictErrorError) GetCode() int32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *ConflictErrorError) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *ConflictErrorError) SetCode(val int32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *ConflictErrorError) SetMessage(val string) {
	s.Message = val
}

// DeleteWebsitesIDOK is response for DeleteWebsitesID operation.
type DeleteWebsitesIDOK struct{}

func (*DeleteWebsitesIDOK) deleteWebsitesIDRes() {}

// Website hit event.
// Ref: #/components/schemas/EventHit
type EventHit struct {
	// UUID generated for each user to link multiple events on the same page together.
	B string `json:"b"`
	// Page URL including query parameters.
	U string `json:"u"`
	// Referrer URL.
	R string `json:"r"`
	// Event type consisting of either 'pagehide', 'unload', 'load', 'hidden' or 'visible'.
	E string `json:"e"`
	// Title of page.
	T OptInt `json:"t"`
	// Timezone of the user.
	D OptString `json:"d"`
	// Screen width.
	W OptInt `json:"w"`
	// Screen height.
	H OptInt `json:"h"`
	// Time spent on page. Only sent on unload.
	M OptInt `json:"m"`
}

// GetB returns the value of B.
func (s *EventHit) GetB() string {
	return s.B
}

// GetU returns the value of U.
func (s *EventHit) GetU() string {
	return s.U
}

// GetR returns the value of R.
func (s *EventHit) GetR() string {
	return s.R
}

// GetE returns the value of E.
func (s *EventHit) GetE() string {
	return s.E
}

// GetT returns the value of T.
func (s *EventHit) GetT() OptInt {
	return s.T
}

// GetD returns the value of D.
func (s *EventHit) GetD() OptString {
	return s.D
}

// GetW returns the value of W.
func (s *EventHit) GetW() OptInt {
	return s.W
}

// GetH returns the value of H.
func (s *EventHit) GetH() OptInt {
	return s.H
}

// GetM returns the value of M.
func (s *EventHit) GetM() OptInt {
	return s.M
}

// SetB sets the value of B.
func (s *EventHit) SetB(val string) {
	s.B = val
}

// SetU sets the value of U.
func (s *EventHit) SetU(val string) {
	s.U = val
}

// SetR sets the value of R.
func (s *EventHit) SetR(val string) {
	s.R = val
}

// SetE sets the value of E.
func (s *EventHit) SetE(val string) {
	s.E = val
}

// SetT sets the value of T.
func (s *EventHit) SetT(val OptInt) {
	s.T = val
}

// SetD sets the value of D.
func (s *EventHit) SetD(val OptString) {
	s.D = val
}

// SetW sets the value of W.
func (s *EventHit) SetW(val OptInt) {
	s.W = val
}

// SetH sets the value of H.
func (s *EventHit) SetH(val OptInt) {
	s.H = val
}

// SetM sets the value of M.
func (s *EventHit) SetM(val OptInt) {
	s.M = val
}

// GetEventPingOK is response for GetEventPing operation.
type GetEventPingOK struct {
	LastModified OptString
}

// GetLastModified returns the value of LastModified.
func (s *GetEventPingOK) GetLastModified() OptString {
	return s.LastModified
}

// SetLastModified sets the value of LastModified.
func (s *GetEventPingOK) SetLastModified(val OptString) {
	s.LastModified = val
}

func (*GetEventPingOK) getEventPingRes() {}

type GetWebsitesOKApplicationJSON []WebsiteGet

func (*GetWebsitesOKApplicationJSON) getWebsitesRes() {}

type InternalServerError struct {
	Error InternalServerErrorError `json:"error"`
}

// GetError returns the value of Error.
func (s *InternalServerError) GetError() InternalServerErrorError {
	return s.Error
}

// SetError sets the value of Error.
func (s *InternalServerError) SetError(val InternalServerErrorError) {
	s.Error = val
}

func (*InternalServerError) deleteWebsitesIDRes()    {}
func (*InternalServerError) getEventPingRes()        {}
func (*InternalServerError) getUsersUserIdRes()      {}
func (*InternalServerError) getWebsiteIDSummaryRes() {}
func (*InternalServerError) getWebsitesIDActiveRes() {}
func (*InternalServerError) getWebsitesIDRes()       {}
func (*InternalServerError) getWebsitesRes()         {}
func (*InternalServerError) patchUsersUserIdRes()    {}
func (*InternalServerError) patchWebsitesIDRes()     {}
func (*InternalServerError) postAuthLoginRes()       {}
func (*InternalServerError) postAuthRefreshRes()     {}
func (*InternalServerError) postEventHitRes()        {}
func (*InternalServerError) postUserRes()            {}
func (*InternalServerError) postWebsitesRes()        {}

type InternalServerErrorError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *InternalServerErrorError) GetCode() int32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *InternalServerErrorError) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *InternalServerErrorError) SetCode(val int32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *InternalServerErrorError) SetMessage(val string) {
	s.Message = val
}

type NotFoundError struct {
	Error NotFoundErrorError `json:"error"`
}

// GetError returns the value of Error.
func (s *NotFoundError) GetError() NotFoundErrorError {
	return s.Error
}

// SetError sets the value of Error.
func (s *NotFoundError) SetError(val NotFoundErrorError) {
	s.Error = val
}

func (*NotFoundError) deleteWebsitesIDRes()    {}
func (*NotFoundError) getUsersUserIdRes()      {}
func (*NotFoundError) getWebsiteIDSummaryRes() {}
func (*NotFoundError) getWebsitesIDActiveRes() {}
func (*NotFoundError) getWebsitesIDRes()       {}
func (*NotFoundError) getWebsitesRes()         {}
func (*NotFoundError) patchUsersUserIdRes()    {}

type NotFoundErrorError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *NotFoundErrorError) GetCode() int32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *NotFoundErrorError) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *NotFoundErrorError) SetCode(val int32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *NotFoundErrorError) SetMessage(val string) {
	s.Message = val
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptEventHit returns new OptEventHit with value set to v.
func NewOptEventHit(v EventHit) OptEventHit {
	return OptEventHit{
		Value: v,
		Set:   true,
	}
}

// OptEventHit is optional EventHit.
type OptEventHit struct {
	Value EventHit
	Set   bool
}

// IsSet returns true if OptEventHit was set.
func (o OptEventHit) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptEventHit) Reset() {
	var v EventHit
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptEventHit) SetTo(v EventHit) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptEventHit) Get() (v EventHit, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptEventHit) Or(d EventHit) EventHit {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat32 returns new OptFloat32 with value set to v.
func NewOptFloat32(v float32) OptFloat32 {
	return OptFloat32{
		Value: v,
		Set:   true,
	}
}

// OptFloat32 is optional float32.
type OptFloat32 struct {
	Value float32
	Set   bool
}

// IsSet returns true if OptFloat32 was set.
func (o OptFloat32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat32) Reset() {
	var v float32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat32) SetTo(v float32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat32) Get() (v float32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat32) Or(d float32) float32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPostAuthLoginReq returns new OptPostAuthLoginReq with value set to v.
func NewOptPostAuthLoginReq(v PostAuthLoginReq) OptPostAuthLoginReq {
	return OptPostAuthLoginReq{
		Value: v,
		Set:   true,
	}
}

// OptPostAuthLoginReq is optional PostAuthLoginReq.
type OptPostAuthLoginReq struct {
	Value PostAuthLoginReq
	Set   bool
}

// IsSet returns true if OptPostAuthLoginReq was set.
func (o OptPostAuthLoginReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPostAuthLoginReq) Reset() {
	var v PostAuthLoginReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPostAuthLoginReq) SetTo(v PostAuthLoginReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPostAuthLoginReq) Get() (v PostAuthLoginReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPostAuthLoginReq) Or(d PostAuthLoginReq) PostAuthLoginReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPostAuthRefreshReq returns new OptPostAuthRefreshReq with value set to v.
func NewOptPostAuthRefreshReq(v PostAuthRefreshReq) OptPostAuthRefreshReq {
	return OptPostAuthRefreshReq{
		Value: v,
		Set:   true,
	}
}

// OptPostAuthRefreshReq is optional PostAuthRefreshReq.
type OptPostAuthRefreshReq struct {
	Value PostAuthRefreshReq
	Set   bool
}

// IsSet returns true if OptPostAuthRefreshReq was set.
func (o OptPostAuthRefreshReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPostAuthRefreshReq) Reset() {
	var v PostAuthRefreshReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPostAuthRefreshReq) SetTo(v PostAuthRefreshReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPostAuthRefreshReq) Get() (v PostAuthRefreshReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPostAuthRefreshReq) Or(d PostAuthRefreshReq) PostAuthRefreshReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUserCreate returns new OptUserCreate with value set to v.
func NewOptUserCreate(v UserCreate) OptUserCreate {
	return OptUserCreate{
		Value: v,
		Set:   true,
	}
}

// OptUserCreate is optional UserCreate.
type OptUserCreate struct {
	Value UserCreate
	Set   bool
}

// IsSet returns true if OptUserCreate was set.
func (o OptUserCreate) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUserCreate) Reset() {
	var v UserCreate
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUserCreate) SetTo(v UserCreate) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUserCreate) Get() (v UserCreate, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUserCreate) Or(d UserCreate) UserCreate {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUserCreateLanguage returns new OptUserCreateLanguage with value set to v.
func NewOptUserCreateLanguage(v UserCreateLanguage) OptUserCreateLanguage {
	return OptUserCreateLanguage{
		Value: v,
		Set:   true,
	}
}

// OptUserCreateLanguage is optional UserCreateLanguage.
type OptUserCreateLanguage struct {
	Value UserCreateLanguage
	Set   bool
}

// IsSet returns true if OptUserCreateLanguage was set.
func (o OptUserCreateLanguage) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUserCreateLanguage) Reset() {
	var v UserCreateLanguage
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUserCreateLanguage) SetTo(v UserCreateLanguage) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUserCreateLanguage) Get() (v UserCreateLanguage, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUserCreateLanguage) Or(d UserCreateLanguage) UserCreateLanguage {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUserPatch returns new OptUserPatch with value set to v.
func NewOptUserPatch(v UserPatch) OptUserPatch {
	return OptUserPatch{
		Value: v,
		Set:   true,
	}
}

// OptUserPatch is optional UserPatch.
type OptUserPatch struct {
	Value UserPatch
	Set   bool
}

// IsSet returns true if OptUserPatch was set.
func (o OptUserPatch) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUserPatch) Reset() {
	var v UserPatch
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUserPatch) SetTo(v UserPatch) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUserPatch) Get() (v UserPatch, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUserPatch) Or(d UserPatch) UserPatch {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptWebsiteCreate returns new OptWebsiteCreate with value set to v.
func NewOptWebsiteCreate(v WebsiteCreate) OptWebsiteCreate {
	return OptWebsiteCreate{
		Value: v,
		Set:   true,
	}
}

// OptWebsiteCreate is optional WebsiteCreate.
type OptWebsiteCreate struct {
	Value WebsiteCreate
	Set   bool
}

// IsSet returns true if OptWebsiteCreate was set.
func (o OptWebsiteCreate) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptWebsiteCreate) Reset() {
	var v WebsiteCreate
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptWebsiteCreate) SetTo(v WebsiteCreate) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptWebsiteCreate) Get() (v WebsiteCreate, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptWebsiteCreate) Or(d WebsiteCreate) WebsiteCreate {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptWebsitePatch returns new OptWebsitePatch with value set to v.
func NewOptWebsitePatch(v WebsitePatch) OptWebsitePatch {
	return OptWebsitePatch{
		Value: v,
		Set:   true,
	}
}

// OptWebsitePatch is optional WebsitePatch.
type OptWebsitePatch struct {
	Value WebsitePatch
	Set   bool
}

// IsSet returns true if OptWebsitePatch was set.
func (o OptWebsitePatch) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptWebsitePatch) Reset() {
	var v WebsitePatch
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptWebsitePatch) SetTo(v WebsitePatch) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptWebsitePatch) Get() (v WebsitePatch, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptWebsitePatch) Or(d WebsitePatch) WebsitePatch {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// PostAuthLoginOK is response for PostAuthLogin operation.
type PostAuthLoginOK struct {
	SetCookie OptString
}

// GetSetCookie returns the value of SetCookie.
func (s *PostAuthLoginOK) GetSetCookie() OptString {
	return s.SetCookie
}

// SetSetCookie sets the value of SetCookie.
func (s *PostAuthLoginOK) SetSetCookie(val OptString) {
	s.SetCookie = val
}

func (*PostAuthLoginOK) postAuthLoginRes() {}

type PostAuthLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetEmail returns the value of Email.
func (s *PostAuthLoginReq) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *PostAuthLoginReq) GetPassword() string {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *PostAuthLoginReq) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *PostAuthLoginReq) SetPassword(val string) {
	s.Password = val
}

// PostAuthRefreshOK is response for PostAuthRefresh operation.
type PostAuthRefreshOK struct {
	SetCookie OptString
}

// GetSetCookie returns the value of SetCookie.
func (s *PostAuthRefreshOK) GetSetCookie() OptString {
	return s.SetCookie
}

// SetSetCookie sets the value of SetCookie.
func (s *PostAuthRefreshOK) SetSetCookie(val OptString) {
	s.SetCookie = val
}

func (*PostAuthRefreshOK) postAuthRefreshRes() {}

type PostAuthRefreshReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetEmail returns the value of Email.
func (s *PostAuthRefreshReq) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *PostAuthRefreshReq) GetPassword() string {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *PostAuthRefreshReq) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *PostAuthRefreshReq) SetPassword(val string) {
	s.Password = val
}

// PostEventHitNotFound is response for PostEventHit operation.
type PostEventHitNotFound struct{}

func (*PostEventHitNotFound) postEventHitRes() {}

// PostEventHitOK is response for PostEventHit operation.
type PostEventHitOK struct{}

func (*PostEventHitOK) postEventHitRes() {}

// Return the number of active realtime users.
// Ref: #/components/schemas/StatsActive
type StatsActive struct {
	Visitors int `json:"visitors"`
}

// GetVisitors returns the value of Visitors.
func (s *StatsActive) GetVisitors() int {
	return s.Visitors
}

// SetVisitors sets the value of Visitors.
func (s *StatsActive) SetVisitors(val int) {
	s.Visitors = val
}

func (*StatsActive) getWebsitesIDActiveRes() {}

// Ref: #/components/schemas/StatsSummary
type StatsSummary struct {
	Uniques   OptInt     `json:"uniques"`
	Pageviews OptInt     `json:"pageviews"`
	Bounces   OptFloat32 `json:"bounces"`
	Duration  OptInt     `json:"duration"`
}

// GetUniques returns the value of Uniques.
func (s *StatsSummary) GetUniques() OptInt {
	return s.Uniques
}

// GetPageviews returns the value of Pageviews.
func (s *StatsSummary) GetPageviews() OptInt {
	return s.Pageviews
}

// GetBounces returns the value of Bounces.
func (s *StatsSummary) GetBounces() OptFloat32 {
	return s.Bounces
}

// GetDuration returns the value of Duration.
func (s *StatsSummary) GetDuration() OptInt {
	return s.Duration
}

// SetUniques sets the value of Uniques.
func (s *StatsSummary) SetUniques(val OptInt) {
	s.Uniques = val
}

// SetPageviews sets the value of Pageviews.
func (s *StatsSummary) SetPageviews(val OptInt) {
	s.Pageviews = val
}

// SetBounces sets the value of Bounces.
func (s *StatsSummary) SetBounces(val OptFloat32) {
	s.Bounces = val
}

// SetDuration sets the value of Duration.
func (s *StatsSummary) SetDuration(val OptInt) {
	s.Duration = val
}

func (*StatsSummary) getWebsiteIDSummaryRes() {}

type UnauthorisedError struct {
	Error UnauthorisedErrorError `json:"error"`
}

// GetError returns the value of Error.
func (s *UnauthorisedError) GetError() UnauthorisedErrorError {
	return s.Error
}

// SetError sets the value of Error.
func (s *UnauthorisedError) SetError(val UnauthorisedErrorError) {
	s.Error = val
}

func (*UnauthorisedError) deleteWebsitesIDRes()    {}
func (*UnauthorisedError) getWebsiteIDSummaryRes() {}
func (*UnauthorisedError) getWebsitesIDActiveRes() {}
func (*UnauthorisedError) getWebsitesIDRes()       {}
func (*UnauthorisedError) getWebsitesRes()         {}
func (*UnauthorisedError) patchUsersUserIdRes()    {}
func (*UnauthorisedError) patchWebsitesIDRes()     {}
func (*UnauthorisedError) postAuthLoginRes()       {}
func (*UnauthorisedError) postAuthRefreshRes()     {}
func (*UnauthorisedError) postUserRes()            {}
func (*UnauthorisedError) postWebsitesRes()        {}

type UnauthorisedErrorError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *UnauthorisedErrorError) GetCode() int32 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *UnauthorisedErrorError) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *UnauthorisedErrorError) SetCode(val int32) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *UnauthorisedErrorError) SetMessage(val string) {
	s.Message = val
}

// User model for admin.
// Ref: #/components/schemas/UserCreate
type UserCreate struct {
	Email    string                `json:"email"`
	Password string                `json:"password"`
	Language OptUserCreateLanguage `json:"language"`
}

// GetEmail returns the value of Email.
func (s *UserCreate) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *UserCreate) GetPassword() string {
	return s.Password
}

// GetLanguage returns the value of Language.
func (s *UserCreate) GetLanguage() OptUserCreateLanguage {
	return s.Language
}

// SetEmail sets the value of Email.
func (s *UserCreate) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *UserCreate) SetPassword(val string) {
	s.Password = val
}

// SetLanguage sets the value of Language.
func (s *UserCreate) SetLanguage(val OptUserCreateLanguage) {
	s.Language = val
}

type UserCreateLanguage string

const (
	UserCreateLanguageEn UserCreateLanguage = "en"
)

// AllValues returns all UserCreateLanguage values.
func (UserCreateLanguage) AllValues() []UserCreateLanguage {
	return []UserCreateLanguage{
		UserCreateLanguageEn,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s UserCreateLanguage) MarshalText() ([]byte, error) {
	switch s {
	case UserCreateLanguageEn:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *UserCreateLanguage) UnmarshalText(data []byte) error {
	switch UserCreateLanguage(data) {
	case UserCreateLanguageEn:
		*s = UserCreateLanguageEn
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// User model for admin.
// Ref: #/components/schemas/UserGet
type UserGet struct {
	// Unique identifier for the given user.
	ID          string `json:"id"`
	Email       string `json:"email"`
	Language    string `json:"language"`
	DateCreated int64  `json:"dateCreated"`
	DateUpdated int64  `json:"dateUpdated"`
}

// GetID returns the value of ID.
func (s *UserGet) GetID() string {
	return s.ID
}

// GetEmail returns the value of Email.
func (s *UserGet) GetEmail() string {
	return s.Email
}

// GetLanguage returns the value of Language.
func (s *UserGet) GetLanguage() string {
	return s.Language
}

// GetDateCreated returns the value of DateCreated.
func (s *UserGet) GetDateCreated() int64 {
	return s.DateCreated
}

// GetDateUpdated returns the value of DateUpdated.
func (s *UserGet) GetDateUpdated() int64 {
	return s.DateUpdated
}

// SetID sets the value of ID.
func (s *UserGet) SetID(val string) {
	s.ID = val
}

// SetEmail sets the value of Email.
func (s *UserGet) SetEmail(val string) {
	s.Email = val
}

// SetLanguage sets the value of Language.
func (s *UserGet) SetLanguage(val string) {
	s.Language = val
}

// SetDateCreated sets the value of DateCreated.
func (s *UserGet) SetDateCreated(val int64) {
	s.DateCreated = val
}

// SetDateUpdated sets the value of DateUpdated.
func (s *UserGet) SetDateUpdated(val int64) {
	s.DateUpdated = val
}

func (*UserGet) getUsersUserIdRes()   {}
func (*UserGet) patchUsersUserIdRes() {}
func (*UserGet) postUserRes()         {}

// User model for admin.
// Ref: #/components/schemas/UserPatch
type UserPatch struct {
	Email    OptString `json:"email"`
	Password OptString `json:"password"`
	Language OptString `json:"language"`
}

// GetEmail returns the value of Email.
func (s *UserPatch) GetEmail() OptString {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *UserPatch) GetPassword() OptString {
	return s.Password
}

// GetLanguage returns the value of Language.
func (s *UserPatch) GetLanguage() OptString {
	return s.Language
}

// SetEmail sets the value of Email.
func (s *UserPatch) SetEmail(val OptString) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *UserPatch) SetPassword(val OptString) {
	s.Password = val
}

// SetLanguage sets the value of Language.
func (s *UserPatch) SetLanguage(val OptString) {
	s.Language = val
}

// Website details.
// Ref: #/components/schemas/WebsiteCreate
type WebsiteCreate struct {
	Name     string  `json:"name"`
	Hostname string  `json:"hostname"`
	IsActive OptBool `json:"isActive"`
}

// GetName returns the value of Name.
func (s *WebsiteCreate) GetName() string {
	return s.Name
}

// GetHostname returns the value of Hostname.
func (s *WebsiteCreate) GetHostname() string {
	return s.Hostname
}

// GetIsActive returns the value of IsActive.
func (s *WebsiteCreate) GetIsActive() OptBool {
	return s.IsActive
}

// SetName sets the value of Name.
func (s *WebsiteCreate) SetName(val string) {
	s.Name = val
}

// SetHostname sets the value of Hostname.
func (s *WebsiteCreate) SetHostname(val string) {
	s.Hostname = val
}

// SetIsActive sets the value of IsActive.
func (s *WebsiteCreate) SetIsActive(val OptBool) {
	s.IsActive = val
}

// Website details.
// Ref: #/components/schemas/WebsiteGet
type WebsiteGet struct {
	// Unique identifier for the website.
	ID       string `json:"id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	IsActive bool   `json:"isActive"`
}

// GetID returns the value of ID.
func (s *WebsiteGet) GetID() string {
	return s.ID
}

// GetName returns the value of Name.
func (s *WebsiteGet) GetName() string {
	return s.Name
}

// GetHostname returns the value of Hostname.
func (s *WebsiteGet) GetHostname() string {
	return s.Hostname
}

// GetIsActive returns the value of IsActive.
func (s *WebsiteGet) GetIsActive() bool {
	return s.IsActive
}

// SetID sets the value of ID.
func (s *WebsiteGet) SetID(val string) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *WebsiteGet) SetName(val string) {
	s.Name = val
}

// SetHostname sets the value of Hostname.
func (s *WebsiteGet) SetHostname(val string) {
	s.Hostname = val
}

// SetIsActive sets the value of IsActive.
func (s *WebsiteGet) SetIsActive(val bool) {
	s.IsActive = val
}

func (*WebsiteGet) getWebsitesIDRes()   {}
func (*WebsiteGet) patchWebsitesIDRes() {}
func (*WebsiteGet) postWebsitesRes()    {}

// Website details.
// Ref: #/components/schemas/WebsitePatch
type WebsitePatch struct {
	Name     OptString `json:"name"`
	Hostname OptString `json:"hostname"`
	IsActive OptBool   `json:"isActive"`
}

// GetName returns the value of Name.
func (s *WebsitePatch) GetName() OptString {
	return s.Name
}

// GetHostname returns the value of Hostname.
func (s *WebsitePatch) GetHostname() OptString {
	return s.Hostname
}

// GetIsActive returns the value of IsActive.
func (s *WebsitePatch) GetIsActive() OptBool {
	return s.IsActive
}

// SetName sets the value of Name.
func (s *WebsitePatch) SetName(val OptString) {
	s.Name = val
}

// SetHostname sets the value of Hostname.
func (s *WebsitePatch) SetHostname(val OptString) {
	s.Hostname = val
}

// SetIsActive sets the value of IsActive.
func (s *WebsitePatch) SetIsActive(val OptBool) {
	s.IsActive = val
}
