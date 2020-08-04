package word

import "testing"

func TestCamelCaseToUnderscore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试数据 UserName",
			args: args{s: "UserName"},
			want: "user_name",
		},
		{
			name: "测试数据 user_name",
			args: args{s: "user_name"},
			want: "user_name",
		},
		{
			name: "测试数据 userName",
			args: args{s: "userName"},
			want: "user_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCaseToUnderscore(tt.args.s); got != tt.want {
				t.Errorf("CamelCaseToUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试 UserName",
			args: args{s: "UserName"},
			want: "username",
		},
		{
			name: "测试 userName",
			args: args{s: "userName"},
			want: "username",
		},
		{
			name: "测试 username",
			args: args{s: "username"},
			want: "username",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLower(tt.args.s); got != tt.want {
				t.Errorf("ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试 username",
			args: args{s: "username"},
			want: "USERNAME",
		},
		{
			name: "测试 userName",
			args: args{s: "userName"},
			want: "USERNAME",
		},
		{
			name: "测试 US",
			args: args{s: "US"},
			want: "US",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.args.s); got != tt.want {
				t.Errorf("ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnderscoreToLowerCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test user_name",
			args: args{s: "user_name"},
			want: "userName",
		},
		{
			name: "test user name",
			args: args{s: "user name"},
			want: "userName",
		},
		{
			name: "test userName",
			args: args{s: "userName"},
			want: "userName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnderscoreToLowerCamelCase(tt.args.s); got != tt.want {
				t.Errorf("UnderscoreToLowerCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test user name",
			args: args{s: "user name"},
			want: "UserName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnderscoreToUpperCamelCase(tt.args.s); got != tt.want {
				t.Errorf("UnderscoreToUpperCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
