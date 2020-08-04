package timer

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetCalculateTime(t *testing.T) {
	type args struct {
		currentTimer time.Time
		d            string
	}

	current := GetNowTime()
	duration, _ := time.ParseDuration("2h")
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				currentTimer: current,
				d:            "2h",
			},
			want:    current.Add(duration),
			wantErr: false,
		},
		{
			name: "test error",
			args: args{
				currentTimer: current,
				d:            "2",
			},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCalculateTime(tt.args.currentTimer, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCalculateTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCalculateTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNowTime1(t *testing.T) {
	fmt.Println(GetNowTime().Format("2006-01-02 15:04:05"))
}
