package removeduplicatesfromsortedarray

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2},
			},
			want:     2,
			wantNums: []int{1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			t.Log("cek", tt.wantNums[0:tt.want])
		})
	}
}
