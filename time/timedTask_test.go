package main

import "testing"

func BenchmarkTick1(b *testing.B){
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		byTimeTick1()
	}
}

func BenchmarkTick2(b *testing.B){
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		byTimeTick2()
	}
}

func BenchmarkSleep(b *testing.B){
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		byTimeSleep()
	}
}

func BenchmarkAfter(b *testing.B){
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		byTimeAfter()
	}
}