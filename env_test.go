package envy_test

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/pedramktb/go-envy"
	"github.com/stretchr/testify/assert"
)

func Test_Env(t *testing.T) {
	err := os.Setenv("TEST_KEY", "test_value")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv("TEST_KEY2", "")
	if err != nil {
		t.Fatal(err)
	}

	envs := envy.Env()
	assert.Equal(t, "test_value", envs["TEST_KEY"])
	assert.Equal(t, "", envs["TEST_KEY2"])
}

func Test_Get(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		value   string
		wantSet bool
		want    any
		wantErr error
	}{
		{
			name:    "set",
			key:     "SET",
			value:   "",
			wantSet: true,
			want:    struct{}{},
		},
		{
			name:    "rune",
			key:     "RUNE",
			value:   "1",
			wantSet: true,
			want:    rune(1),
		},
		{
			name:    "byte",
			key:     "BYTE",
			value:   "1",
			wantSet: true,
			want:    byte(1),
		},
		{
			name:    "string",
			key:     "STRING",
			value:   "string",
			wantSet: true,
			want:    string("string"),
		},
		{
			name:    "bool",
			key:     "BOOL",
			value:   "true",
			wantSet: true,
			want:    bool(true),
		},
		{
			name:    "uintptr",
			key:     "UINTPTR",
			value:   "1",
			wantSet: true,
			want:    uintptr(1),
		},
		{
			name:    "int",
			key:     "INT",
			value:   "-1",
			wantSet: true,
			want:    int(-1),
		},
		{
			name:    "int64",
			key:     "INT64",
			value:   "-1",
			wantSet: true,
			want:    int64(-1),
		},
		{
			name:    "int32",
			key:     "INT32",
			value:   "-1",
			wantSet: true,
			want:    int32(-1),
		},
		{
			name:    "int16",
			key:     "INT16",
			value:   "-1",
			wantSet: true,
			want:    int16(-1),
		},
		{
			name:    "int8",
			key:     "INT8",
			value:   "-1",
			wantSet: true,
			want:    int8(-1),
		},
		{
			name:    "uint",
			key:     "UINT",
			value:   "1",
			wantSet: true,
			want:    uint(1),
		},
		{
			name:    "uint64",
			key:     "UINT64",
			value:   "1",
			wantSet: true,
			want:    uint64(1),
		},
		{
			name:    "uint32",
			key:     "UINT32",
			value:   "1",
			wantSet: true,
			want:    uint32(1),
		},
		{
			name:    "uint16",
			key:     "UINT16",
			value:   "1",
			wantSet: true,
			want:    uint16(1),
		},
		{
			name:    "uint8",
			key:     "UINT8",
			value:   "1",
			wantSet: true,
			want:    uint8(1),
		},
		{
			name:    "float64",
			key:     "FLOAT64",
			value:   "1.1",
			wantSet: true,
			want:    float64(1.1),
		},
		{
			name:    "float32",
			key:     "FLOAT32",
			value:   "1.1",
			wantSet: true,
			want:    float32(1.1),
		},
		{
			name:    "complex128",
			key:     "COMPLEX128",
			value:   "1.1+2.2i",
			wantSet: true,
			want:    complex128(complex(float64(1.1), float64(2.2))),
		},
		{
			name:    "complex64",
			key:     "COMPLEX64",
			value:   "1.1+2.2i",
			wantSet: true,
			want:    complex64(complex(float32(1.1), float32(2.2))),
		},
		{
			name:    "with unmarshaler",
			key:     "WITH_UNMARSHALER",
			value:   "a4541d10-5a6d-4f86-b157-362a6c227a9e",
			wantSet: true,
			want: uuid.UUID{
				0xa4, 0x54, 0x1d, 0x10, 0x5a, 0x6d, 0x4f, 0x86, 0xb1, 0x57, 0x36, 0x2a, 0x6c, 0x22, 0x7a, 0x9e,
			},
		},
		{
			name:    "unsupported type",
			key:     "UNSUPPORTED_TYPE",
			value:   "'X':1",
			wantSet: true,
			want:    struct{ X int }{},
			wantErr: envy.ErrUnsupportedType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.Setenv(tt.key, tt.value)
			if err != nil {
				t.Fatal(err)
			}
			var got any
			var set bool
			switch tt.want.(type) {
			case struct{}:
				got, set, err = envy.Get[struct{}](tt.key)
			case string:
				got, set, err = envy.Get[string](tt.key)
			case bool:
				got, set, err = envy.Get[bool](tt.key)
			case uintptr:
				got, set, err = envy.Get[uintptr](tt.key)
			case int:
				got, set, err = envy.Get[int](tt.key)
			case int64:
				got, set, err = envy.Get[int64](tt.key)
			case int32:
				got, set, err = envy.Get[int32](tt.key)
			case int16:
				got, set, err = envy.Get[int16](tt.key)
			case int8:
				got, set, err = envy.Get[int8](tt.key)
			case uint:
				got, set, err = envy.Get[uint](tt.key)
			case uint64:
				got, set, err = envy.Get[uint64](tt.key)
			case uint32:
				got, set, err = envy.Get[uint32](tt.key)
			case uint16:
				got, set, err = envy.Get[uint16](tt.key)
			case uint8:
				got, set, err = envy.Get[uint8](tt.key)
			case float64:
				got, set, err = envy.Get[float64](tt.key)
			case float32:
				got, set, err = envy.Get[float32](tt.key)
			case complex128:
				got, set, err = envy.Get[complex128](tt.key)
			case complex64:
				got, set, err = envy.Get[complex64](tt.key)
			case uuid.UUID:
				got, set, err = envy.Get[uuid.UUID](tt.key)
			case struct{ X int }:
				got, set, err = envy.Get[struct{ X int }](tt.key)
			default:
				t.Fatalf("unhandled test case type: %T", tt.want)
			}
			assert.Equal(t, tt.wantSet, set)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
			err = os.Unsetenv(tt.key)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
