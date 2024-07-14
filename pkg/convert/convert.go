package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(s.String(), 10, 32)
	return uint32(v), err
}

func (s StrTo) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(s.String(), 10, 64)
	return v, err
}

func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32()
	return v
}

func (s StrTo) MustUint64() uint64 {
	v, _ := s.Uint64()
	return v
}
