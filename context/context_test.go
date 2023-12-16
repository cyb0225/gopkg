package context

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func Test_WithoutCancel(t *testing.T) {
	t.Run("test without cancel success and test get values", func(t *testing.T) {
		key, value := "hello", "world"
		ctx := context.WithValue(context.TODO(), key, value)

		ctxWithCancel, cancel := context.WithCancel(ctx)

		ctxWithoutCancel := WithoutCancel(ctxWithCancel)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			cancel()
			select {
			case <-ctxWithoutCancel.Done():
				t.Error("newCtx still canceled")
				return
			case <-time.Tick(time.Millisecond * 10):
				t.Log("timeout")
				got := ctxWithoutCancel.Value(key).(string)
				assert.Equal(t, value, got)
			}
		}()

		wg.Wait()
	})
}
