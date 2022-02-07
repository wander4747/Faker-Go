package faker

import (
	"encoding/json"
	"github.com/wander4747/faker-go/locale"
)

type personStruct struct {
	FirstName []string `json:"first_name"`
	LastName  []string `json:"last_name"`
}

type person struct {
	*Fake
	File []byte
	data *personStruct
}

type PersonInterface interface {
	FirstName() string
	LastName() string
	FullName() string
	Age() int
}

func (f *Fake) Person() PersonInterface {
	loader := locale.Loader(f.Locale, locale.PERSON)

	data, err := json.Marshal(loader)
	if err != nil || loader == nil {
		panic("error converter struct")
	}

	person := &person{f, data, nil}
	person.data, err = person.getData()
	return person
}

func (n *person) FirstName() string {
	i := random(len(n.data.FirstName))
	return n.data.FirstName[i]
}

func (n *person) LastName() string {
	i := random(len(n.data.LastName))
	return n.data.LastName[i]
}

func (n *person) FullName() string {
	return n.FirstName() + " " + n.LastName()
}

func (n *person) Age() int {
	return random(100)
}

func (n *person) getData() (*personStruct, error) {
	var persons personStruct

	err := json.Unmarshal(n.File, &persons)
	if err != nil {
		return nil, err
	}

	return &persons, nil
}
