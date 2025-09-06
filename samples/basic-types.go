package samples

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
)

var (
	ToBe    bool       = true
	Str     string     = "Hello Go!"
	MaxInt  uint64     = 1<<64 - 1
	Sqrt128 complex128 = cmplx.Sqrt(-5 + 12i)

	IntMax     int     = int(math.MaxInt64)
	IntMin     int     = int(math.MinInt64)
	Int8Max    int8    = int8(math.MaxInt8)
	Int8Min    int8    = int8(math.MinInt8)
	Int16Max   int16   = int16(math.MaxInt16)
	Int16Min   int16   = int16(math.MinInt16)
	Int32Max   int32   = int32(math.MaxInt32)
	Int32Min   int32   = int32(math.MinInt32)
	Int64Max   int64   = int64(math.MaxInt64)
	Int64Min   int64   = int64(math.MinInt64)
	UintMax    uint    = ^uint(0)
	UintMin    uint    = uint(0)
	Uint8Max   uint8   = uint8(math.MaxUint8)
	Uint8Min   uint8   = uint8(0)
	Uint16Max  uint16  = uint16(math.MaxUint16)
	Uint16Min  uint16  = uint16(0)
	Uint32Max  uint32  = uint32(math.MaxUint32)
	Uint32Min  uint32  = uint32(0)
	Uint64Max  uint64  = uint64(math.MaxUint64)
	Uint64Min  uint64  = uint64(0)
	UintptrMax uintptr = ^uintptr(0)
	UintptrMin uintptr = uintptr(0)

	ByteVal byte = 'a'
	RuneVal rune = '好'

	Float32Max float32 = float32(math.MaxFloat32)
	Float32Min float32 = float32(-math.MaxFloat32)
	Float64Max float64 = float64(math.MaxFloat64)
	Float64Min float64 = float64(-math.MaxFloat64)

	Complex32Max complex64  = complex(math.MaxFloat32, math.MaxFloat32)
	Complex32Min complex64  = complex(-math.MaxFloat32, -math.MaxFloat32)
	Complex64Max complex128 = complex(math.MaxFloat64, math.MaxFloat64)
	Complex64Min complex128 = complex(-math.MaxFloat64, -math.MaxFloat64)
)

func BasicTypesSamples() {
	// Create a slice of structs to hold the variables and each values
	varTypes := []struct {
		Name  string
		Value interface{}
	}{
		{"ToBe", ToBe},
		{"Str", Str},
		{"MaxInt", MaxInt},
		{"Sqrt128", Sqrt128},
		{"IntMax", IntMax},
		{"IntMin", IntMin},
		{"Int8Max", Int8Max},
		{"Int8Min", Int8Min},
		{"Int16Max", Int16Max},
		{"Int16Min", Int16Min},
		{"Int32Max", Int32Max},
		{"Int32Min", Int32Min},
		{"Int64Max", Int64Max},
		{"Int64Min", Int64Min},
		{"UintMax", UintMax},
		{"UintMin", UintMin},
		{"Uint8Max", Uint8Max},
		{"Uint8Min", Uint8Min},
		{"Uint16Max", Uint16Max},
		{"Uint16Min", Uint16Min},
		{"Uint32Max", Uint32Max},
		{"Uint32Min", Uint32Min},
		{"Uint64Max", Uint64Max},
		{"Uint64Min", Uint64Min},
		{"UintptrMax", UintptrMax},
		{"UintptrMin", UintptrMin},
		{"ByteVal", ByteVal},
		{"RuneVal", RuneVal},
		{"Float32Max", Float32Max},
		{"Float32Min", Float32Min},
		{"Float64Max", Float64Max},
		{"Float64Min", Float64Min},
		{"Complex32Max", Complex32Max},
		{"Complex32Min", Complex32Min},
		{"Complex64Max", Complex64Max},
		{"Complex64Min", Complex64Min},
	}

	for _, v := range varTypes {
		fmt.Printf("Name: %s, Type: %s,  Value: %v \n", v.Name, reflect.TypeOf(v.Value), v.Value)
	}
}
