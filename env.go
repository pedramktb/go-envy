package envy

import (
	"encoding"
	"errors"
	"os"
	"strconv"
	"strings"
)

var ErrUnsupportedType = errors.New("types is unsupported as it does not implement encoding.TextUnmarshaler")

func Env() map[string]string {
	sl := os.Environ()
	m := map[string]string{}
	for i := range sl {
		p := strings.SplitN(sl[i], "=", 2)
		if len(p) == 2 {
			m[p[0]] = p[1]
		}
	}
	return m
}

type envOptions[E any] struct {
	fallback *E
}

type Option[E any] func(o *envOptions[E])

// WithFallback sets the fallback value for the environment variable if it is not set
func WithFallback[E any](fallback E) Option[E] {
	return func(o *envOptions[E]) {
		o.fallback = &fallback
	}
}

// Env returns the value of the environment variable, if it exists and an error if the parsing fails
func Get[E any](env string, options ...Option[E]) (E, bool, error) {
	opts := envOptions[E]{}
	for i := range options {
		options[i](&opts)
	}

	str, ok := os.LookupEnv(env)
	if !ok {
		if opts.fallback != nil {
			return *opts.fallback, false, nil
		}
		return *new(E), false, nil
	}

	var val any
	var err error
	var e E
	var i64 int64
	var u64 uint64
	var f64 float64
	var c128 complex128

	switch v := any(&e).(type) {
	case encoding.TextUnmarshaler:
		err = v.UnmarshalText([]byte(str))
		val = e
	case *struct{}:
		val = struct{}{}
	case *string:
		val = str
	case *bool:
		val, err = strconv.ParseBool(str)
	case *uintptr:
		u64, err = strconv.ParseUint(str, 10, 0)
		val = uintptr(u64)
	case *int:
		i64, err = strconv.ParseInt(str, 10, 0)
		val = int(i64)
	case *int64:
		val, err = strconv.ParseInt(str, 10, 64)
	case *int32:
		i64, err = strconv.ParseInt(str, 10, 32)
		val = int32(i64)
	case *int16:
		i64, err = strconv.ParseInt(str, 10, 16)
		val = int16(i64)
	case *int8:
		i64, err = strconv.ParseInt(str, 10, 8)
		val = int8(i64)
	case *uint:
		u64, err = strconv.ParseUint(str, 10, 0)
		val = uint(u64)
	case *uint64:
		val, err = strconv.ParseUint(str, 10, 64)
	case *uint32:
		u64, err = strconv.ParseUint(str, 10, 32)
		val = uint32(u64)
	case *uint16:
		u64, err = strconv.ParseUint(str, 10, 16)
		val = uint16(u64)
	case *uint8:
		u64, err = strconv.ParseUint(str, 10, 8)
		val = uint8(u64)
	case *float64:
		val, err = strconv.ParseFloat(str, 64)
	case *float32:
		f64, err = strconv.ParseFloat(str, 32)
		val = float32(f64)
	case *complex128:
		val, err = strconv.ParseComplex(str, 128)
	case *complex64:
		c128, err = strconv.ParseComplex(str, 64)
		val = complex64(c128)
	default:
		return *new(E), true, ErrUnsupportedType
	}

	return val.(E), true, err
}
