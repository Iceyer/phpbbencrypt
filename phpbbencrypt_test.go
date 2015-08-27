package phpbbencrypt

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	t.Log(Encrypt("123456"))
	t.Log(Verify("123456", "$H$9wXDkkAvY1XanMapy6OgnEVNS0ifOu1"))
	data := "TestString 测试字符"
	t.Log(Verify(data, Encrypt(data)))
}

func BenchmarkEncrypti16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encrypt("1234567890123456")
	}
}
func BenchmarkVerify16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !Verify("1234567890123456", "$H$9lSllFvHIqO1vTyVkTV/LIY8JD0Ut01") {
			b.Error("Verify Failed")
		}
	}
}

func BenchmarkEncrypti6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encrypt("123456")
	}
}

func BenchmarkVerify6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !Verify("123456", "$H$9wXDkkAvY1XanMapy6OgnEVNS0ifOu1") {
			b.Error("Verify Failed")
		}
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
