package locale

import (
	"github.com/wander4747/faker-go/locale/pt_BR"
)

const (
	NAME    string = "name"
	ADDRESS string = "address"
)

var data = map[string]interface{}{}

func Loader(locale, typeFake string) interface{} {
	if typeFake == NAME {
		return names(locale)
	}
	if typeFake == ADDRESS {
		return addresses(locale)
	}

	return nil
}

func names(locale string) interface{} {
	data["pt_BR"] = pt_BR.Names

	return data[locale]
}

func addresses(locale string) interface{} {
	data["pt_BR"] = pt_BR.Addresses

	return data[locale]
}
