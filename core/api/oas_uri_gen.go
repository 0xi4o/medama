// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

// EncodeURI encodes FilterFixed as URI form.
func (s *FilterFixed) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("eq", func(e uri.Encoder) error {
		if val, ok := s.Eq.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"eq\"")
	}
	if err := e.EncodeField("neq", func(e uri.Encoder) error {
		if val, ok := s.Neq.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"neq\"")
	}
	if err := e.EncodeField("in", func(e uri.Encoder) error {
		if val, ok := s.In.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"in\"")
	}
	if err := e.EncodeField("not_in", func(e uri.Encoder) error {
		if val, ok := s.NotIn.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"not_in\"")
	}
	return nil
}

var uriFieldsNameOfFilterFixed = [4]string{
	0: "eq",
	1: "neq",
	2: "in",
	3: "not_in",
}

// DecodeURI decodes FilterFixed from URI form.
func (s *FilterFixed) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FilterFixed to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "eq":
			if err := func() error {
				var sDotEqVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotEqVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Eq.SetTo(sDotEqVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"eq\"")
			}
		case "neq":
			if err := func() error {
				var sDotNeqVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotNeqVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Neq.SetTo(sDotNeqVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"neq\"")
			}
		case "in":
			if err := func() error {
				var sDotInVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotInVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.In.SetTo(sDotInVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"in\"")
			}
		case "not_in":
			if err := func() error {
				var sDotNotInVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotNotInVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.NotIn.SetTo(sDotNotInVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"not_in\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FilterFixed")
	}

	return nil
}

// EncodeURI encodes FilterString as URI form.
func (s *FilterString) EncodeURI(e uri.Encoder) error {
	if err := e.EncodeField("eq", func(e uri.Encoder) error {
		if val, ok := s.Eq.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"eq\"")
	}
	if err := e.EncodeField("neq", func(e uri.Encoder) error {
		if val, ok := s.Neq.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"neq\"")
	}
	if err := e.EncodeField("contains", func(e uri.Encoder) error {
		if val, ok := s.Contains.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"contains\"")
	}
	if err := e.EncodeField("not_contains", func(e uri.Encoder) error {
		if val, ok := s.NotContains.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"not_contains\"")
	}
	if err := e.EncodeField("starts_with", func(e uri.Encoder) error {
		if val, ok := s.StartsWith.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"starts_with\"")
	}
	if err := e.EncodeField("not_starts_with", func(e uri.Encoder) error {
		if val, ok := s.NotStartsWith.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"not_starts_with\"")
	}
	if err := e.EncodeField("ends_with", func(e uri.Encoder) error {
		if val, ok := s.EndsWith.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"ends_with\"")
	}
	if err := e.EncodeField("not_ends_with", func(e uri.Encoder) error {
		if val, ok := s.NotEndsWith.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"not_ends_with\"")
	}
	if err := e.EncodeField("in", func(e uri.Encoder) error {
		if val, ok := s.In.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"in\"")
	}
	if err := e.EncodeField("not_in", func(e uri.Encoder) error {
		if val, ok := s.NotIn.Get(); ok {
			return e.EncodeValue(conv.StringToString(val))
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "encode field \"not_in\"")
	}
	return nil
}

var uriFieldsNameOfFilterString = [10]string{
	0: "eq",
	1: "neq",
	2: "contains",
	3: "not_contains",
	4: "starts_with",
	5: "not_starts_with",
	6: "ends_with",
	7: "not_ends_with",
	8: "in",
	9: "not_in",
}

// DecodeURI decodes FilterString from URI form.
func (s *FilterString) DecodeURI(d uri.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode FilterString to nil")
	}

	if err := d.DecodeFields(func(k string, d uri.Decoder) error {
		switch k {
		case "eq":
			if err := func() error {
				var sDotEqVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotEqVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Eq.SetTo(sDotEqVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"eq\"")
			}
		case "neq":
			if err := func() error {
				var sDotNeqVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotNeqVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Neq.SetTo(sDotNeqVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"neq\"")
			}
		case "contains":
			if err := func() error {
				var sDotContainsVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotContainsVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.Contains.SetTo(sDotContainsVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"contains\"")
			}
		case "not_contains":
			if err := func() error {
				var sDotNotContainsVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotNotContainsVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.NotContains.SetTo(sDotNotContainsVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"not_contains\"")
			}
		case "starts_with":
			if err := func() error {
				var sDotStartsWithVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotStartsWithVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.StartsWith.SetTo(sDotStartsWithVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"starts_with\"")
			}
		case "not_starts_with":
			if err := func() error {
				var sDotNotStartsWithVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotNotStartsWithVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.NotStartsWith.SetTo(sDotNotStartsWithVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"not_starts_with\"")
			}
		case "ends_with":
			if err := func() error {
				var sDotEndsWithVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotEndsWithVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.EndsWith.SetTo(sDotEndsWithVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"ends_with\"")
			}
		case "not_ends_with":
			if err := func() error {
				var sDotNotEndsWithVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotNotEndsWithVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.NotEndsWith.SetTo(sDotNotEndsWithVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"not_ends_with\"")
			}
		case "in":
			if err := func() error {
				var sDotInVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotInVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.In.SetTo(sDotInVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"in\"")
			}
		case "not_in":
			if err := func() error {
				var sDotNotInVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					sDotNotInVal = c
					return nil
				}(); err != nil {
					return err
				}
				s.NotIn.SetTo(sDotNotInVal)
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"not_in\"")
			}
		default:
			return nil
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode FilterString")
	}

	return nil
}
