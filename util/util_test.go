package util

import "testing"

// go test -v . 执行测试
// go test -v -count=10 . 执行10次测试
func TestGenShortId(t *testing.T) {
	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		t.Errorf("genShortId failed, %v", err)
	}
}

// 性能测试 go test -test.bench=".*"
// 性能分析 go test -test.bench=".*" -cpuprofile=cpu.profile .
// 使用pprof分析 go tool pprof util.test cpu.profile
// 输入 top 得到top N的数据
// 输入web 在 web页面打开 svg图  需要安装 graphviz
func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

func BenchmarkGenShortIdConsuming(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer() // 重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}
