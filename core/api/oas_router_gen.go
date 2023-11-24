// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "auth/login"
				if l := len("auth/login"); len(elem) >= l && elem[0:l] == "auth/login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "POST":
						s.handlePostAuthLoginRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
			case 'e': // Prefix: "event/"
				if l := len("event/"); len(elem) >= l && elem[0:l] == "event/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'h': // Prefix: "hit"
					if l := len("hit"); len(elem) >= l && elem[0:l] == "hit" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handlePostEventHitRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}
				case 'p': // Prefix: "ping"
					if l := len("ping"); len(elem) >= l && elem[0:l] == "ping" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetEventPingRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'u': // Prefix: "user"
				if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "DELETE":
						s.handleDeleteUserRequest([0]string{}, elemIsEscaped, w, r)
					case "GET":
						s.handleGetUserRequest([0]string{}, elemIsEscaped, w, r)
					case "PATCH":
						s.handlePatchUserRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handlePostUserRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "DELETE,GET,PATCH,POST")
					}

					return
				}
			case 'w': // Prefix: "website"
				if l := len("website"); len(elem) >= l && elem[0:l] == "website" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "hostname"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'p': // Prefix: "pages"
							if l := len("pages"); len(elem) >= l && elem[0:l] == "pages" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetWebsiteIDPagesRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 's': // Prefix: "summary"
							if l := len("summary"); len(elem) >= l && elem[0:l] == "summary" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetWebsiteIDSummaryRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				case 's': // Prefix: "s"
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleGetWebsitesRequest([0]string{}, elemIsEscaped, w, r)
						case "POST":
							s.handlePostWebsitesRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET,POST")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "hostname"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "DELETE":
								s.handleDeleteWebsitesIDRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							case "GET":
								s.handleGetWebsitesIDRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							case "PATCH":
								s.handlePatchWebsitesIDRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "DELETE,GET,PATCH")
							}

							return
						}
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "auth/login"
				if l := len("auth/login"); len(elem) >= l && elem[0:l] == "auth/login" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						// Leaf: PostAuthLogin
						r.name = "PostAuthLogin"
						r.summary = "Session Token Authentication."
						r.operationID = "post-auth-login"
						r.pathPattern = "/auth/login"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'e': // Prefix: "event/"
				if l := len("event/"); len(elem) >= l && elem[0:l] == "event/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'h': // Prefix: "hit"
					if l := len("hit"); len(elem) >= l && elem[0:l] == "hit" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "POST":
							// Leaf: PostEventHit
							r.name = "PostEventHit"
							r.summary = "Send Hit Event."
							r.operationID = "post-event-hit"
							r.pathPattern = "/event/hit"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'p': // Prefix: "ping"
					if l := len("ping"); len(elem) >= l && elem[0:l] == "ping" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: GetEventPing
							r.name = "GetEventPing"
							r.summary = "Unique User Check."
							r.operationID = "get-event-ping"
							r.pathPattern = "/event/ping"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'u': // Prefix: "user"
				if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "DELETE":
						// Leaf: DeleteUser
						r.name = "DeleteUser"
						r.summary = "Delete User."
						r.operationID = "delete-user"
						r.pathPattern = "/user"
						r.args = args
						r.count = 0
						return r, true
					case "GET":
						// Leaf: GetUser
						r.name = "GetUser"
						r.summary = "Get User Info."
						r.operationID = "get-user"
						r.pathPattern = "/user"
						r.args = args
						r.count = 0
						return r, true
					case "PATCH":
						// Leaf: PatchUser
						r.name = "PatchUser"
						r.summary = "Update User Info."
						r.operationID = "patch-user"
						r.pathPattern = "/user"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						// Leaf: PostUser
						r.name = "PostUser"
						r.summary = "Create New User."
						r.operationID = "post-user"
						r.pathPattern = "/user"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'w': // Prefix: "website"
				if l := len("website"); len(elem) >= l && elem[0:l] == "website" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "hostname"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'p': // Prefix: "pages"
							if l := len("pages"); len(elem) >= l && elem[0:l] == "pages" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: GetWebsiteIDPages
									r.name = "GetWebsiteIDPages"
									r.summary = "Get Page Stats."
									r.operationID = "get-website-id-pages"
									r.pathPattern = "/website/{hostname}/pages"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						case 's': // Prefix: "summary"
							if l := len("summary"); len(elem) >= l && elem[0:l] == "summary" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: GetWebsiteIDSummary
									r.name = "GetWebsiteIDSummary"
									r.summary = "Get Stat Summary."
									r.operationID = "get-website-id-summary"
									r.pathPattern = "/website/{hostname}/summary"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				case 's': // Prefix: "s"
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "GetWebsites"
							r.summary = "List Websites."
							r.operationID = "get-websites"
							r.pathPattern = "/websites"
							r.args = args
							r.count = 0
							return r, true
						case "POST":
							r.name = "PostWebsites"
							r.summary = "Add Website."
							r.operationID = "post-websites"
							r.pathPattern = "/websites"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "hostname"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							switch method {
							case "DELETE":
								// Leaf: DeleteWebsitesID
								r.name = "DeleteWebsitesID"
								r.summary = "Delete Website."
								r.operationID = "delete-websites-id"
								r.pathPattern = "/websites/{hostname}"
								r.args = args
								r.count = 1
								return r, true
							case "GET":
								// Leaf: GetWebsitesID
								r.name = "GetWebsitesID"
								r.summary = "Get Website."
								r.operationID = "get-websites-id"
								r.pathPattern = "/websites/{hostname}"
								r.args = args
								r.count = 1
								return r, true
							case "PATCH":
								// Leaf: PatchWebsitesID
								r.name = "PatchWebsitesID"
								r.summary = "Update Website."
								r.operationID = "patch-websites-id"
								r.pathPattern = "/websites/{hostname}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
