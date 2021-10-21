package tests

import (
	"testing"
)

/*func TestSum(t *testing.T) {
	sum := Sum(100, 200)
	if sum == 300 {
		t.Log("the result is ok")
	} else {
		t.Errorf("the result is wrong")
	}

	sum = Sum(200, 200)
	if sum == 400 {
		t.Log("the result is ok")
	} else {
		t.Errorf("the result is wrong")
	}

	sum = Sum(-200, 200)
	if sum == 0 {
		t.Log("the result is ok")
	} else {
		t.Errorf("the result is wrong")
	}
}*/

func TestSum1(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{a: -100, b: -200},
			want: -300,
		},
		{
			name: "case2",
			args: args{a: 100, b: -300},
			want: -200,
		},
		{
			name: "case3",
			args: args{a: 200, b: 200},
			want: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}


func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Sum(100,200)
	}
}