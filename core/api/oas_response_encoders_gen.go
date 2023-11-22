// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func encodeDeleteUserResponse(response DeleteUserRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *DeleteUserOK:
		w.WriteHeader(200)

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ConflictError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(409)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeleteWebsitesIDResponse(response DeleteWebsitesIDRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *DeleteWebsitesIDOK:
		w.WriteHeader(200)

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetEventPingResponse(response GetEventPingRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *GetEventPingOKHeaders:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Cache-Control" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Cache-Control",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					return e.EncodeValue(conv.StringToString(response.CacheControl))
				}); err != nil {
					return errors.Wrap(err, "encode Cache-Control header")
				}
			}
			// Encode "Last-Modified" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Last-Modified",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					return e.EncodeValue(conv.StringToString(response.LastModified))
				}); err != nil {
					return errors.Wrap(err, "encode Last-Modified header")
				}
			}
		}
		w.WriteHeader(200)

		writer := w
		if _, err := io.Copy(writer, response.Response); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetUserResponse(response GetUserRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *UserGet:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetWebsiteIDSummaryResponse(response GetWebsiteIDSummaryRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *StatsSummary:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetWebsitesResponse(response GetWebsitesRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *GetWebsitesOKApplicationJSON:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetWebsitesIDResponse(response GetWebsitesIDRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *WebsiteGet:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetWebsitesIDActiveResponse(response GetWebsitesIDActiveRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *StatsActive:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePatchUserResponse(response PatchUserRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *UserGet:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ConflictError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(409)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePatchWebsitesIDResponse(response PatchWebsitesIDRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *WebsiteGet:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostAuthLoginResponse(response PostAuthLoginRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PostAuthLoginOK:
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Set-Cookie" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Set-Cookie",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					return e.EncodeValue(conv.StringToString(response.SetCookie))
				}); err != nil {
					return errors.Wrap(err, "encode Set-Cookie header")
				}
			}
		}
		w.WriteHeader(200)

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostEventHitResponse(response PostEventHitRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PostEventHitOK:
		w.WriteHeader(200)

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFoundError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostUserResponse(response PostUserRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *UserGetHeaders:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Set-Cookie" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Set-Cookie",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					return e.EncodeValue(conv.StringToString(response.SetCookie))
				}); err != nil {
					return errors.Wrap(err, "encode Set-Cookie header")
				}
			}
		}
		w.WriteHeader(201)

		e := new(jx.Encoder)
		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ForbiddenError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(403)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ConflictError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(409)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostWebsitesResponse(response PostWebsitesRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *WebsiteGet:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(201)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *BadRequestError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UnauthorisedError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ConflictError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(409)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *InternalServerError:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
