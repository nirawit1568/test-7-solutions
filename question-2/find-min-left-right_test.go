package question2

import "testing"

func TestFindMinLeftRight(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test-case-1",
			args: args{input: "LLRR="},
			want: "210122",
		},
		{
			name: "test-case-2",
			args: args{input: "==RLL"},
			want: "000210",
		},
		{
			name: "test-case-3",
			args: args{input: "=LLRR"},
			want: "221012",
		},
		{
			name: "test-case-4",
			args: args{input: "RRL=R"},
			want: "012001",
		},
		{
			name: "test-case-5",
			args: args{input: "LLLLL"},
			want: "543210",
		},
		{
			name: "test-case-6",
			args: args{input: "=L=L=L"},
			want: "3322110",
		},
		{
			name: "test-case-7",
			args: args{input: "L=L=L="},
			want: "3221100",
		},
		{
			name: "test-case-8",
			args: args{input: "R=R=R="},
			want: "0112233",
		},
		{
			name: "test-case-9",
			args: args{input: "=R=R=R"},
			want: "0011223",
		},
		{
			name: "test-case-10",
			args: args{input: "==LLL"},
			want: "333210",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinLeftRight(tt.args.input); got != tt.want {
				t.Errorf("FindMinLeftRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
