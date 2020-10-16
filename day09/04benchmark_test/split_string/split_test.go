package split_string

import "testing"

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("127.0.0.1", ".")
	}
}
func BenchmarkSplitParallel(b *testing.B) {
	//b.SetParallelism(1) // 设置使用CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}
