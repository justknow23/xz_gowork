package tools

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestGetSpecialEndDate(t *testing.T) {
	type args struct {
		now          time.Time
		formatString string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			"GetSpecialEndDate",
			args{
				time.Date(2021, 02, 27, 11, 59, 59, 0, time.UTC),
				"1m",
			},
			time.Date(2021, 02, 28, 23, 59, 59, 0, time.UTC),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSpecialEndDate(tt.args.now, tt.args.formatString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpecialEndDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSpecialEndDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLastDateOfMonth(t *testing.T) {

	A := GetZeroTime(time.Now()).String()
	B := GetFirstDateOfMonth(time.Now()).String()
	assert.NotNil(t, A)
	assert.NotNil(t, B)
}
