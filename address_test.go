package faker

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestFake_Address(t *testing.T) {
	f := New("pt_BR")

	tests := []struct {
		name   string
		locale string
		want   AddressInterface
	}{
		{"Success", "pt_BR", f.Address()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fake{
				Locale: tt.locale,
			}
			if got := f.Address(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_address(t *testing.T) {
	fake, file, addresses := getTestAddress()

	type fields struct {
		Fake *Fake
		File []byte
		data *addressStruct
	}
	tests := []struct {
		name     string
		fields   fields
		typeName string
		want     string
	}{
		{"Name", fields{fake, file, &addresses}, "name",
			"Avenida Santo Antonio,476 Bloco 4 - Buritis - Belo Horizonte - MG, 06838-834"},
		{"ZipCode", fields{fake, file, &addresses}, "zipcode", "30880-460"},
		{"State", fields{fake, file, &addresses}, "state", "Minas Gerais"},
		{"State Abbr", fields{fake, file, &addresses}, "state_abbr", "MG"},
		{"City", fields{fake, file, &addresses}, "city", "Belo Horizonte"},
		{"Neighborhood", fields{fake, file, &addresses}, "neighborhood", "Buritis"},
		{"Country", fields{fake, file, &addresses}, "country", "Brazil"},
		{"Region", fields{fake, file, &addresses}, "region", "Sudeste"},
		{"Region Abbr", fields{fake, file, &addresses}, "region_abbr", "SE"},
		{"Secondary Address", fields{fake, file, &addresses}, "secondary_address", "Bloco 1"},
		{"Street Name", fields{fake, file, &addresses}, "street_name", "Santo Antonio"},
		{"Street Prefix", fields{fake, file, &addresses}, "street_prefix", "Avenida"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &address{
				Fake: tt.fields.Fake,
				File: tt.fields.File,
				data: tt.fields.data,
			}

			var got string

			if tt.typeName == "name" {
				got = s.Name()
				require.Contains(t, tt.want, "-", got)
				require.Contains(t, tt.want, "Bloco", got)
				require.Contains(t, tt.want, "MG", got)
				require.NotEmpty(t, tt.want, got)
				return
			}

			if tt.typeName == "zipcode" {
				got = s.ZipCode()
				require.Contains(t, tt.want, "-", got)
				return
			}

			if tt.typeName == "state" {
				got = s.State()
			}

			if tt.typeName == "state_abbr" {
				got = s.StateAbbr()
			}

			if tt.typeName == "city" {
				got = s.City()
			}

			if tt.typeName == "neighborhood" {
				got = s.Neighborhood()
			}

			if tt.typeName == "country" {
				got = s.Country()
			}

			if tt.typeName == "region" {
				got = s.Region()
			}

			if tt.typeName == "region_abbr" {
				got = s.RegionAbbr()
			}

			if tt.typeName == "secondary_address" {
				got = s.SecondaryAddress()
				require.Contains(t, tt.want, "Bloco", got)
				return
			}

			if tt.typeName == "street_name" {
				got = s.StreetName()
			}
			if tt.typeName == "street_prefix" {
				got = s.StreetPrefix()
			}

			require.Equal(t, tt.want, got)
		})
	}
}

func Test_address_getData(t *testing.T) {
	fake, file, addresses := getTestAddress()

	type fields struct {
		Fake *Fake
		File []byte
		data *addressStruct
	}
	tests := []struct {
		name    string
		fields  fields
		want    *addressStruct
		wantErr bool
	}{
		{"Success", fields{fake, file, &addresses}, &addresses, false},
		{"Error", fields{fake, []byte(""), &addresses}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &address{
				Fake: tt.fields.Fake,
				File: tt.fields.File,
				data: tt.fields.data,
			}
			got, err := s.getData()
			if (err != nil) != tt.wantErr {
				t.Errorf("getData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTestAddress() (*Fake, []byte, addressStruct) {
	fake := New("pt_BR")

	json := `{
		  "state":[
			"Minas Gerais"
		  ],
		  "state_abbr": [
			"MG"
		  ],
		  "city": [
			"Belo Horizonte"
		  ],
		  "country": [
			"Brazil"
		  ],
		  "region": [
			"Sudeste"
		  ],
		  "region_abbr": [
			"SE"
		  ],
		  "secondary_address": [
			"Bloco #"
		  ],
		  "street_name": [
			"Santo Antonio"
		  ],
		  "street_prefix": [
			"Avenida"
		  ],
		  "neighborhood" : [
			"Buritis"
		  ]
		}`

	address := addressStruct{
		State:            []string{"Minas Gerais"},
		StateAbbr:        []string{"MG"},
		City:             []string{"Belo Horizonte"},
		Country:          []string{"Brazil"},
		Region:           []string{"Sudeste"},
		RegionAbbr:       []string{"SE"},
		SecondaryAddress: []string{"Bloco #"},
		StreetName:       []string{"Santo Antonio"},
		StreetPrefix:     []string{"Avenida"},
		Neighborhood:     []string{"Buritis"},
	}

	return fake, []byte(json), address
}
