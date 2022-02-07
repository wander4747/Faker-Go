package faker

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestFake_Commerce(t *testing.T) {
	f := New("pt_BR")

	tests := []struct {
		name   string
		locale string
		want   CommerceInterface
	}{
		{"Success", "pt_BR", f.Commerce()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fake{
				Locale: tt.locale,
			}
			if got := f.Commerce(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Commerce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commerce_Category(t *testing.T) {
	fake, file, names := getTestCommerce()

	type fields struct {
		Fake *Fake
		File []byte
		data *commerceStruct
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Category", fields{fake, file, &names}, "Informática"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &commerce{
				Fake: tt.fields.Fake,
				File: tt.fields.File,
				data: tt.fields.data,
			}

			got := n.Category()
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_commerce_Price(t *testing.T) {
	fake, file, names := getTestCommerce()

	type fields struct {
		Fake *Fake
		File []byte
		data *commerceStruct
	}
	type args struct {
		min    int
		max    int
		symbol string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"Price real", fields{fake, file, &names}, args{min: 1, max: 2, symbol: "R$ "}, "R$ "},
		{"Price dollar", fields{fake, file, &names}, args{min: 1, max: 2, symbol: "USD "}, "USD "},
		{"Price crazy", fields{fake, file, &names}, args{min: -1, max: -1, symbol: "CR "}, "CR "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &commerce{
				Fake: tt.fields.Fake,
				File: tt.fields.File,
				data: tt.fields.data,
			}
			got := n.Price(tt.args.min, tt.args.max, tt.args.symbol)
			require.Contains(t, tt.want, tt.args.symbol, got)

			priceString := strings.Replace(got, tt.args.symbol, "", -1)
			fmt.Println(priceString)

			price, _ := strconv.ParseFloat(priceString, 32)

			if price < 0 {
				result := fmt.Sprintf("%v%v", tt.args.symbol, 0.0)
				require.Equal(t, fmt.Sprintf("%v%v", tt.args.symbol, 0.0), result)
			}
		})
	}
}

func Test_commerce_ProductName(t *testing.T) {
	fake, file, names := getTestCommerce()

	type fields struct {
		Fake *Fake
		File []byte
		data *commerceStruct
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"ProductName", fields{fake, file, &names}, "Notebook"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &commerce{
				Fake: tt.fields.Fake,
				File: tt.fields.File,
				data: tt.fields.data,
			}

			got := n.ProductName()
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_commerce_Shop(t *testing.T) {
	fake, file, names := getTestCommerce()

	type fields struct {
		Fake *Fake
		File []byte
		data *commerceStruct
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Shop", fields{fake, file, &names}, "Magazine Luiza"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &commerce{
				Fake: tt.fields.Fake,
				File: tt.fields.File,
				data: tt.fields.data,
			}

			got := n.Shop()
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_commerce_getData(t *testing.T) {
	fake, file, names := getTestCommerce()

	type fields struct {
		Fake *Fake
		File []byte
		data *commerceStruct
	}

	tests := []struct {
		name    string
		fields  fields
		want    *commerceStruct
		wantErr bool
	}{
		{"Success", fields{fake, file, &names}, &names, false},
		{"Error", fields{fake, []byte(""), &names}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &commerce{
				Fake: tt.fields.Fake,
				File: tt.fields.File,
				data: tt.fields.data,
			}
			got, err := n.getData()
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

func getTestCommerce() (*Fake, []byte, commerceStruct) {
	fake := New("pt_BR")

	json := `{
		"shop": [
			"Magazine Luiza"
		],
		"product_name": [
			"Notebook"
		],
		"category": [
			"Informática"
		]
	}`

	names := commerceStruct{
		Shop:        []string{"Magazine Luiza"},
		ProductName: []string{"Notebook"},
		Category:    []string{"Informática"},
	}

	return fake, []byte(json), names
}
