package logs

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	// test default logger with outer functions
	ctx := context.TODO()
	Debug(ctx, "debug: %d", 1)
	Info(ctx, "info: %d", 2)
	Warn(ctx, "warn: %d", 3)
	Error(ctx, "error: %d", 4)
}

func TestPanicLog(t *testing.T) {
	ctx := context.TODO()
	defer func() {
		a := recover()
		if a == nil {
			t.Fatal("program do not panic")
		}
		t.Log(a)
	}()
	Panic(ctx, "panic: %d", 5)
}

//func TestFatalLog(t *testing.T) {
//	ctx := context.TODO()
//	Fatal(ctx, "fatal: %d", 5)
//}
