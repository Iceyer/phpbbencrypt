package phpbbencrypt

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	t.Log(Encrypt("123456"))
	t.Log(Verfiy("123456", "$H$9wXDkkAvY1XanMapy6OgnEVNS0ifOu1"))
	data := "TestString 测试字符"
	t.Log(Verfiy(data, Encrypt(data)))
}

func BenchmarkEncrypti16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encrypt("1234567890123456")
	}
}
func BenchmarkVerfiy16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !Verfiy("1234567890123456", "$H$9lSllFvHIqO1vTyVkTV/LIY8JD0Ut01") {
			b.Error("Verfiy Failed")
		}
	}
}

func BenchmarkEncrypti6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encrypt("123456")
	}
}

func BenchmarkVerfiy6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !Verfiy("123456", "$H$9wXDkkAvY1XanMapy6OgnEVNS0ifOu1") {
			b.Error("Verfiy Failed")
		}
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
