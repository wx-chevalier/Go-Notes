package firebase

import "testing"

func Test_firebaseClient_Get(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := Authenticate()
			if err != nil {
				t.Error(err)
			}
			_, err = f.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("firebaseClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_firebaseClient_Set(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"test", "case"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := Authenticate()
			if err != nil {
				t.Error(err)
			}
			if err := f.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("firebaseClient.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
