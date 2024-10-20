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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinLeftRight(tt.args.input); got != tt.want {
				t.Errorf("FindMinLeftRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
