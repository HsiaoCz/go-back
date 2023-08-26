package fb

import "testing"

// 性能比较函数
func benchmarkFeb(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		FeB(n)
	}
}

func BenchmarkFeb1(b *testing.B) {
	benchmarkFeb(b, 1)
}

func BenchmarkFeb2(b *testing.B) {
	benchmarkFeb(b, 5)
}

func BenchmarkFeb3(b *testing.B) {
	benchmarkFeb(b, 10)
}

func BenchmarkFeb4(b *testing.B) {
	benchmarkFeb(b, 20)
}

// benchmark的命令
// go test -bench=.
// 或者指定具体的要测试的函数
// b.ResetTimer重置测试时间

// 并行测试
func BenchmarkHello(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = FeB(10)
		}
	})
}
