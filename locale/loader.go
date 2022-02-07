package locale

import (
	"github.com/wander4747/faker-go/locale/pt_BR"
)

const (
	PERSON   string = "person"
	ADDRESS  string = "address"
	COMMERCE string = "commerce"
)

var data = map[string]interface{}{}

func Loader(locale, typeFake string) interface{} {
	if typeFake == PERSON {
		return names(locale)
	}
	if typeFake == ADDRESS {
		return addresses(locale)
	}
	if typeFake == COMMERCE {
		return commerces(locale)
	}

	return nil
}

func names(locale string) interface{} {
	data["pt_BR"] = pt_BR.Persons

	return data[locale]
}

func addresses(locale string) interface{} {
	data["pt_BR"] = pt_BR.Addresses

	return data[locale]
}

func commerces(locale string) interface{} {
	data["pt_BR"] = pt_BR.Commerces

	return data[locale]
}
