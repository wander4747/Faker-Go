package faker

import (
	"math/rand"
	"time"
)

type Fake struct {
	Locale string
}

type Faker interface {
	Person() PersonInterface
	Address() AddressInterface
	Commerce() CommerceInterface
}

func New(locale string) *Fake {
	return &Fake{
		locale,
	}
}

func random(i int) int {
	r := newRandom()
	return r.Intn(i)
}

func randFloat(min, max int) float64 {
	r := newRandom()
	return float64(r.Intn(max-min)) + float64(min) + r.Float64()
}

func newRandom() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s)
}
