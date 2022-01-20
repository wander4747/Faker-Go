package faker

import (
	"encoding/json"
	"fmt"
	"github.com/wander4747/faker-go/locale"
	"strconv"
	"strings"
)

type addressStruct struct {
	State            []string `json:"state"`
	StateAbbr        []string `json:"state_abbr"`
	City             []string `json:"city"`
	Country          []string `json:"country"`
	Region           []string `json:"region"`
	RegionAbbr       []string `json:"region_abbr"`
	SecondaryAddress []string `json:"secondary_address"`
	StreetName       []string `json:"street_name"`
	StreetPrefix     []string `json:"street_prefix"`
	Neighborhood     []string `json:"neighborhood"`
}

type address struct {
	*Fake
	File []byte
	data *addressStruct
}

type AddressInterface interface {
	Name() string
	ZipCode() string
	State() string
	StateAbbr() string
	City() string
	Neighborhood() string
	Country() string
	Region() string
	RegionAbbr() string
	SecondaryAddress() string
	StreetName() string
	StreetPrefix() string
}

func (f *Fake) Address() AddressInterface {
	loader := locale.Loader(f.Locale, locale.ADDRESS)

	data, err := json.Marshal(loader)
	if err != nil || loader == nil {
		panic("error converter struct")
	}

	address := &address{f, data, nil}
	address.data, err = address.getData()
	return address
}

func (s *address) Name() string {
	number := strconv.Itoa(random(1000))
	numberSecondary := strconv.Itoa(random(10))
	return s.StreetPrefix() + " " + s.StreetName() + "," + number +
		" " + strings.Replace(s.SecondaryAddress(), "#", numberSecondary, 1) +
		" - " + s.Neighborhood() + " - " + s.City() +
		" - " + s.StateAbbr() + ", " + s.ZipCode()
}

func (s *address) ZipCode() string {
	return fmt.Sprintf("%05d", random(99999)) + "-" + fmt.Sprintf("%03d", random(999))
}

func (s *address) State() string {
	i := random(len(s.data.State))
	return s.data.State[i]
}

func (s *address) StateAbbr() string {
	i := random(len(s.data.StateAbbr))
	return s.data.StateAbbr[i]
}

func (s *address) City() string {
	i := random(len(s.data.City))
	return s.data.City[i]
}

func (s *address) Neighborhood() string {
	i := random(len(s.data.Neighborhood))
	return s.data.Neighborhood[i]
}

func (s *address) Country() string {
	i := random(len(s.data.Country))
	return s.data.Country[i]
}

func (s *address) Region() string {
	i := random(len(s.data.Region))
	return s.data.Region[i]
}

func (s *address) RegionAbbr() string {
	i := random(len(s.data.RegionAbbr))
	return s.data.RegionAbbr[i]
}

func (s *address) SecondaryAddress() string {
	numberSecondary := strconv.Itoa(random(10))
	i := random(len(s.data.SecondaryAddress))
	return strings.Replace(s.data.SecondaryAddress[i], "#", numberSecondary, 1)
}

func (s *address) StreetName() string {
	i := random(len(s.data.StreetName))
	return s.data.StreetName[i]
}

func (s *address) StreetPrefix() string {
	i := random(len(s.data.StreetPrefix))
	return s.data.StreetPrefix[i]
}

func (s *address) getData() (*addressStruct, error) {
	var addresses addressStruct

	err := json.Unmarshal(s.File, &addresses)
	if err != nil {
		return nil, err
	}

	return &addresses, nil
}
