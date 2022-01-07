package faker

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestFake_Name(t *testing.T) {
	f := New("pt_BR")
	tests := []struct {
		name   string
		locale string
		want   NameInterface
	}{
		{"Success", "pt_BR", f.Name()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := f.Name()
			require.Equal(t, tt.want, got)

		})
	}
}

func Test_name(t *testing.T) {
	fake, file, names := getTestName()

	type fields struct {
		Fake *Fake
		File []byte
		data *nameStruct
	}
	tests := []struct {
		name     string
		fields   fields
		typeName string
		want     string
	}{
		{"FirstName", fields{fake, file, &names}, "firstName", "Wander"},
		{"LastName", fields{fake, file, &names}, "lastName", "Douglas"},
		{"FullName", fields{fake, file, &names}, "fullName", "Wander Douglas"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &name{
				Fake: tt.fields.Fake,
				File: tt.fields.File,
				data: tt.fields.data,
			}
			var got string
			if tt.typeName == "firstName" {
				got = n.FirstName()
			}
			if tt.typeName == "lastName" {
				got = n.LastName()
			}
			if tt.typeName == "fullName" {
				got = n.FullName()
			}

			require.Equal(t, tt.want, got)
		})
	}
}

func Test_name_getData(t *testing.T) {
	fake, file, names := getTestName()

	type fields struct {
		Fake *Fake
		File []byte
		data *nameStruct
	}
	tests := []struct {
		name    string
		fields  fields
		want    *nameStruct
		wantErr bool
	}{
		{"Success", fields{fake, file, &names}, &names, false},
		{"Error", fields{fake, []byte(""), &names}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &name{
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

func getTestName() (*Fake, []byte, nameStruct) {
	fake := New("pt_BR")

	json := `{
		"first_name": [
			"Wander"
		],
		"last_name": [
			"Douglas"
		]
	}`

	names := nameStruct{
		FirstName: []string{"Wander"},
		LastName:  []string{"Douglas"},
	}

	return fake, []byte(json), names
}
