package contact

import (
	"reflect"
	"testing"

	"github.com/dhruvbehl/address-book/api"
)

func TestGetAllContact(t *testing.T) {
	tests := []struct {
		name    string
		want    *[]api.Contact
		wantErr bool
	}{
		{
			"success get all",
			&[]api.Contact{
				{
					Id: "100",
					FirstName: "MockFirstName",
					LastName: "MockLastName",
					Phone: "xxxxxxxxxx",
					Address: "MockAddress",
				},
			}, false,
		},
	}
	for i, tt := range tests {
		t.Logf("%v) Initiating test case: [%v]",i, tt.name)
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllContact()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllContact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetContactById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *api.Contact
		wantErr bool
	}{
		{
			"success get by id",
			args{
				id: "100",
			},
			&api.Contact{
				Id: "100",
				FirstName: "MockFirstName",
				LastName: "MockLastName",
				Phone: "xxxxxxxxxx",
				Address: "MockAddress",
			},
			false,
		},
		{
			"failure get by id, id not found",
			args{
				id: "201",
			},
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Logf("%v) Initiating test case: [%v]",i, tt.name)
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetContactById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContactById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContactById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteContactById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"success delete by id",
			args{
				id: "100",
			},
			false,
		},
		{
			"failure delete by id, ",
			args{
				id: "301",
			},
			true,
		},
	}
	for i, tt := range tests {
		t.Logf("%v) Initiating test case: [%v]",i, tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteContactById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteContactById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpsertContact(t *testing.T) {
	type args struct {
		newContact api.Contact
		id         string
	}
	tests := []struct {
		name    string
		args    args
		want    *api.Contact
		wantErr bool
	}{
		{
			"success create",
			args{
				id: "",
				newContact: api.Contact{
					Id: "400",
					FirstName: "MockFirstName-test-create",
					LastName: "MockLastName",
					Phone: "xxxxxxxxxx",
					Address: "MockAddress",
				},
			},
			&api.Contact{
				Id: "400",
				FirstName: "MockFirstName-test-create",
				LastName: "MockLastName",
				Phone: "xxxxxxxxxx",
				Address: "MockAddress",
			},
			false,
		},
		{
			"success update",
			args{
				id: "400",
				newContact: api.Contact{
					Id: "400",
					FirstName: "MockFirstName-test-update",
					LastName: "",
					Phone: "",
					Address: "",
				},
			},
			&api.Contact{
				Id: "400",
				FirstName: "MockFirstName-test-update",
				LastName: "MockLastName",
				Phone: "xxxxxxxxxx",
				Address: "MockAddress",
			},
			false,
		},
		{
			"failure update id mismatch",
			args{
				id: "401",
				newContact: api.Contact{
					Id: "100",
					FirstName: "MockFirstName-update",
					LastName: "",
					Phone: "",
					Address: "",
				},
			},
			nil,
			true,
		},
		{
			"failure update id not found",
			args{
				id: "401",
				newContact: api.Contact{
					Id: "400",
					FirstName: "MockFirstName-update",
					LastName: "",
					Phone: "",
					Address: "",
				},
			},
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Logf("%v) Initiating test case: [%v]",i, tt.name)
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpsertContact(tt.args.newContact, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpsertContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpsertContact() = %v, want %v", got, tt.want)
			}
		})
	}
}
