package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},

		//Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbbbb", 1},
		{"abcabcabcd", 4},
		{"fijklmn", 7},
		//Chinese Support
		{"我爱go语言！", 7},
		{"我要学会go语言！", 9},
		{"一二三二一", 3},
	}

	for _, tt := range tests {
		actual := LengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; "+
				"expected %d",
				actual, tt.s, tt.ans)
		}
	}

}

func BenchmarkSubstr(b *testing.B) {
	s := "我要学会go语言！"
	//我们来做个超长数据，那么clear memory的比重就不大了。
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Log("len(s) = %d", len(s))
	ans := 9
	//将准备输入数据的时间剔除
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//这里要decode，所以变成大头
		actual := LengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d",
				actual, s, ans)
		}
	}

}
