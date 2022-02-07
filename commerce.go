package faker

import (
	"encoding/json"
	"fmt"
	"github.com/wander4747/faker-go/locale"
)

type commerceStruct struct {
	Shop        []string `json:"shop"`
	ProductName []string `json:"product_name"`
	Category    []string `json:"category"`
}

type commerce struct {
	*Fake
	File []byte
	data *commerceStruct
}

type CommerceInterface interface {
	Shop() string
	ProductName() string
	Category() string
	Price(min, max int, symbol string) string
}

func (f *Fake) Commerce() CommerceInterface {
	loader := locale.Loader(f.Locale, locale.COMMERCE)

	data, err := json.Marshal(loader)
	if err != nil || loader == nil {
		panic("error converter struct")
	}

	commerces := &commerce{f, data, nil}
	commerces.data, _ = commerces.getData()
	return commerces
}

func (n *commerce) Shop() string {
	i := random(len(n.data.Shop))
	return n.data.Shop[i]
}

func (n *commerce) ProductName() string {
	i := random(len(n.data.ProductName))
	return n.data.ProductName[i]
}

func (n *commerce) Category() string {
	i := random(len(n.data.Category))
	return n.data.Category[i]
}

func (n *commerce) Price(min, max int, symbol string) string {
	if min < 0 || max < 0 {
		return fmt.Sprintf("%v%v", symbol, 0.0)
	}

	r := randFloat(min, max)
	return fmt.Sprintf("%v%.2f", symbol, r)
}

func (n *commerce) getData() (*commerceStruct, error) {
	var commerces commerceStruct

	err := json.Unmarshal(n.File, &commerces)
	if err != nil {
		return nil, err
	}

	return &commerces, nil
}
