package models

import "testing"

func TestNewDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDB(); got == nil {
				t.Errorf("NewDB() = %v, want %v", got, "not nil")
			}
		})
	}
}

func Test_db_GetScore(t *testing.T) {
	type fields struct {
		score int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{"base-case", fields{1}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				score: tt.fields.score,
			}
			got, err := d.GetScore()
			if (err != nil) != tt.wantErr {
				t.Errorf("db.GetScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("db.GetScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_db_SetScore(t *testing.T) {
	type fields struct {
		score int64
	}
	type args struct {
		score int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"base-case", fields{1}, args{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{
				score: tt.fields.score,
			}
			if err := d.SetScore(tt.args.score); (err != nil) != tt.wantErr {
				t.Errorf("db.SetScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
