package faker

import (
	"math/rand"
	"time"
)

type Fake struct {
	Locale string
}

type Faker interface {
	Name() NameInterface
	Address() AddressInterface
}

func New(locale string) *Fake {
	return &Fake{
		locale,
	}
}

func random(i int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(i)
}
