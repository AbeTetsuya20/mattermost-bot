package test

import (
	"fmt"
	"rat/domain"
	"testing"
)

func Test_user_test(t *testing.T) {
	type args struct {
		user *domain.User
	}
	type want struct {
		user *domain.User
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "1ユーザ当選",
			args: args{
				user: domain.NewUser(
					"tetsuya.abe@email.com",
					"abe tetsuya",
					1234,
				),
			},
			want: want{
				user: domain.NewUser(
					"tetsuya.abe@email.com",
					"abe tetsuya",
					1234,
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q", tt.name), func(t *testing.T) {
			if got := tt.args.user; got != tt.want.user {
				t.Errorf("User() = %v, want %v", got, tt.want.user)
			}
		})
	}
}
