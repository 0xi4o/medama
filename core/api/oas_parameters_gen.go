// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteWebsitesIDParams is parameters of delete-websites-id operation.
type DeleteWebsitesIDParams struct {
	// Session token for authentication.
	MeSess string
	// Hostname for the website.
	Hostname string
}

func unpackDeleteWebsitesIDParams(packed middleware.Parameters) (params DeleteWebsitesIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "_me_sess",
			In:   "cookie",
		}
		params.MeSess = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "hostname",
			In:   "path",
		}
		params.Hostname = packed[key].(string)
	}
	return params
}

func decodeDeleteWebsitesIDParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteWebsitesIDParams, _ error) {
	c := uri.NewCookieDecoder(r)
	// Decode cookie: _me_sess.
	if err := func() error {
		cfg := uri.CookieParameterDecodingConfig{
			Name:    "_me_sess",
			Explode: true,
		}
		if err := c.HasParam(cfg); err == nil {
			if err := c.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.MeSess = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "_me_sess",
			In:   "cookie",
			Err:  err,
		}
	}
	// Decode path: hostname.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "hostname",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Hostname = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "hostname",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetEventPingParams is parameters of get-event-ping operation.
type GetEventPingParams struct {
	// If this exists, then user exists in cache and is not a unique user.
	LastModified OptString
}

func unpackGetEventPingParams(packed middleware.Parameters) (params GetEventPingParams) {
	{
		key := middleware.ParameterKey{
			Name: "Last-Modified",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.LastModified = v.(OptString)
		}
	}
	return params
}

func decodeGetEventPingParams(args [0]string, argsEscaped bool, r *http.Request) (params GetEventPingParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: Last-Modified.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "Last-Modified",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLastModifiedVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotLastModifiedVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.LastModified.SetTo(paramsDotLastModifiedVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "Last-Modified",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}

// GetUsersUserIdParams is parameters of get-users-userId operation.
type GetUsersUserIdParams struct {
	// Id of an existing user.
	UserId string
}

func unpackGetUsersUserIdParams(packed middleware.Parameters) (params GetUsersUserIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(string)
	}
	return params
}

func decodeGetUsersUserIdParams(args [1]string, argsEscaped bool, r *http.Request) (params GetUsersUserIdParams, _ error) {
	// Decode path: userId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "userId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.UserId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetWebsiteIDSummaryParams is parameters of get-website-id-summary operation.
type GetWebsiteIDSummaryParams struct {
	// Session token for authentication.
	MeSess string
	// Start time (seconds) in Unix epoch format.
	Start OptString
	// End time (seconds) in Unix epoch format.
	End OptString
	// Hostname for the website.
	Hostname string
}

func unpackGetWebsiteIDSummaryParams(packed middleware.Parameters) (params GetWebsiteIDSummaryParams) {
	{
		key := middleware.ParameterKey{
			Name: "_me_sess",
			In:   "cookie",
		}
		params.MeSess = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "start",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Start = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "end",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.End = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "hostname",
			In:   "path",
		}
		params.Hostname = packed[key].(string)
	}
	return params
}

func decodeGetWebsiteIDSummaryParams(args [1]string, argsEscaped bool, r *http.Request) (params GetWebsiteIDSummaryParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	c := uri.NewCookieDecoder(r)
	// Decode cookie: _me_sess.
	if err := func() error {
		cfg := uri.CookieParameterDecodingConfig{
			Name:    "_me_sess",
			Explode: true,
		}
		if err := c.HasParam(cfg); err == nil {
			if err := c.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.MeSess = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "_me_sess",
			In:   "cookie",
			Err:  err,
		}
	}
	// Decode query: start.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStartVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Start.SetTo(paramsDotStartVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "start",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: end.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "end",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEndVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotEndVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.End.SetTo(paramsDotEndVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "end",
			In:   "query",
			Err:  err,
		}
	}
	// Decode path: hostname.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "hostname",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Hostname = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "hostname",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetWebsitesParams is parameters of get-websites operation.
type GetWebsitesParams struct {
	// Session token for authentication.
	MeSess string
}

func unpackGetWebsitesParams(packed middleware.Parameters) (params GetWebsitesParams) {
	{
		key := middleware.ParameterKey{
			Name: "_me_sess",
			In:   "cookie",
		}
		params.MeSess = packed[key].(string)
	}
	return params
}

func decodeGetWebsitesParams(args [0]string, argsEscaped bool, r *http.Request) (params GetWebsitesParams, _ error) {
	c := uri.NewCookieDecoder(r)
	// Decode cookie: _me_sess.
	if err := func() error {
		cfg := uri.CookieParameterDecodingConfig{
			Name:    "_me_sess",
			Explode: true,
		}
		if err := c.HasParam(cfg); err == nil {
			if err := c.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.MeSess = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "_me_sess",
			In:   "cookie",
			Err:  err,
		}
	}
	return params, nil
}

// GetWebsitesIDParams is parameters of get-websites-id operation.
type GetWebsitesIDParams struct {
	// Session token for authentication.
	MeSess string
	// Hostname for the website.
	Hostname string
}

func unpackGetWebsitesIDParams(packed middleware.Parameters) (params GetWebsitesIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "_me_sess",
			In:   "cookie",
		}
		params.MeSess = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "hostname",
			In:   "path",
		}
		params.Hostname = packed[key].(string)
	}
	return params
}

func decodeGetWebsitesIDParams(args [1]string, argsEscaped bool, r *http.Request) (params GetWebsitesIDParams, _ error) {
	c := uri.NewCookieDecoder(r)
	// Decode cookie: _me_sess.
	if err := func() error {
		cfg := uri.CookieParameterDecodingConfig{
			Name:    "_me_sess",
			Explode: true,
		}
		if err := c.HasParam(cfg); err == nil {
			if err := c.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.MeSess = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "_me_sess",
			In:   "cookie",
			Err:  err,
		}
	}
	// Decode path: hostname.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "hostname",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Hostname = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "hostname",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetWebsitesIDActiveParams is parameters of get-websites-id-active operation.
type GetWebsitesIDActiveParams struct {
	// Session token for authentication.
	MeSess string
	// Hostname for the website.
	Hostname string
}

func unpackGetWebsitesIDActiveParams(packed middleware.Parameters) (params GetWebsitesIDActiveParams) {
	{
		key := middleware.ParameterKey{
			Name: "_me_sess",
			In:   "cookie",
		}
		params.MeSess = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "hostname",
			In:   "path",
		}
		params.Hostname = packed[key].(string)
	}
	return params
}

func decodeGetWebsitesIDActiveParams(args [1]string, argsEscaped bool, r *http.Request) (params GetWebsitesIDActiveParams, _ error) {
	c := uri.NewCookieDecoder(r)
	// Decode cookie: _me_sess.
	if err := func() error {
		cfg := uri.CookieParameterDecodingConfig{
			Name:    "_me_sess",
			Explode: true,
		}
		if err := c.HasParam(cfg); err == nil {
			if err := c.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.MeSess = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "_me_sess",
			In:   "cookie",
			Err:  err,
		}
	}
	// Decode path: hostname.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "hostname",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Hostname = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "hostname",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PatchUsersUserIdParams is parameters of patch-users-userId operation.
type PatchUsersUserIdParams struct {
	// ID of an existing user.
	UserId string
}

func unpackPatchUsersUserIdParams(packed middleware.Parameters) (params PatchUsersUserIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "userId",
			In:   "path",
		}
		params.UserId = packed[key].(string)
	}
	return params
}

func decodePatchUsersUserIdParams(args [1]string, argsEscaped bool, r *http.Request) (params PatchUsersUserIdParams, _ error) {
	// Decode path: userId.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "userId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.UserId = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "userId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PatchWebsitesIDParams is parameters of patch-websites-id operation.
type PatchWebsitesIDParams struct {
	// Session token for authentication.
	MeSess string
	// Hostname for the website.
	Hostname string
}

func unpackPatchWebsitesIDParams(packed middleware.Parameters) (params PatchWebsitesIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "_me_sess",
			In:   "cookie",
		}
		params.MeSess = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "hostname",
			In:   "path",
		}
		params.Hostname = packed[key].(string)
	}
	return params
}

func decodePatchWebsitesIDParams(args [1]string, argsEscaped bool, r *http.Request) (params PatchWebsitesIDParams, _ error) {
	c := uri.NewCookieDecoder(r)
	// Decode cookie: _me_sess.
	if err := func() error {
		cfg := uri.CookieParameterDecodingConfig{
			Name:    "_me_sess",
			Explode: true,
		}
		if err := c.HasParam(cfg); err == nil {
			if err := c.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.MeSess = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "_me_sess",
			In:   "cookie",
			Err:  err,
		}
	}
	// Decode path: hostname.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "hostname",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Hostname = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "hostname",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PostEventHitParams is parameters of post-event-hit operation.
type PostEventHitParams struct {
	UserAgent OptString
	// Used to infer user language.
	AcceptLanguage  OptString
	SecChUa         OptString
	SecChUaMobile   OptString
	SecChUaPlatform OptString
}

func unpackPostEventHitParams(packed middleware.Parameters) (params PostEventHitParams) {
	{
		key := middleware.ParameterKey{
			Name: "User-Agent",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.UserAgent = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "Accept-Language",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.AcceptLanguage = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "Sec-Ch-Ua",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.SecChUa = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "Sec-Ch-Ua-Mobile",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.SecChUaMobile = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "Sec-Ch-Ua-Platform",
			In:   "header",
		}
		if v, ok := packed[key]; ok {
			params.SecChUaPlatform = v.(OptString)
		}
	}
	return params
}

func decodePostEventHitParams(args [0]string, argsEscaped bool, r *http.Request) (params PostEventHitParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: User-Agent.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "User-Agent",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUserAgentVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUserAgentVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UserAgent.SetTo(paramsDotUserAgentVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "User-Agent",
			In:   "header",
			Err:  err,
		}
	}
	// Decode header: Accept-Language.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "Accept-Language",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotAcceptLanguageVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotAcceptLanguageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.AcceptLanguage.SetTo(paramsDotAcceptLanguageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "Accept-Language",
			In:   "header",
			Err:  err,
		}
	}
	// Decode header: Sec-Ch-Ua.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "Sec-Ch-Ua",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSecChUaVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSecChUaVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.SecChUa.SetTo(paramsDotSecChUaVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "Sec-Ch-Ua",
			In:   "header",
			Err:  err,
		}
	}
	// Decode header: Sec-Ch-Ua-Mobile.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "Sec-Ch-Ua-Mobile",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSecChUaMobileVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSecChUaMobileVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.SecChUaMobile.SetTo(paramsDotSecChUaMobileVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "Sec-Ch-Ua-Mobile",
			In:   "header",
			Err:  err,
		}
	}
	// Decode header: Sec-Ch-Ua-Platform.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "Sec-Ch-Ua-Platform",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSecChUaPlatformVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSecChUaPlatformVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.SecChUaPlatform.SetTo(paramsDotSecChUaPlatformVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "Sec-Ch-Ua-Platform",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}
