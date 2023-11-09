package taskqueue

import (
	"context"
	"testing"
	"time"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/require"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/eventloop"
	"go.k6.io/k6/js/modulestest"
)

func TestTaskQueue(t *testing.T) {
	// really basic test
	t.Parallel()
	rt := goja.New()
	vu := &modulestest.VU{
		RuntimeField: rt,
		InitEnvField: &common.InitEnvironment{},
		CtxField:     context.Background(),
		StateField:   nil,
	}
	loop := eventloop.New(vu)
	fq := New(loop.RegisterCallback)
	var i int
	require.NoError(t, rt.Set("a", func() {
		fq.Queue(func() error {
			fq.Queue(func() error {
				fq.Queue(func() error {
					i++
					fq.Close()
					return nil
				})
				i++
				return nil
			})
			i++
			return nil
		})
	}))

	err := loop.Start(func() error {
		_, err := vu.Runtime().RunString(`a()`)
		return err
	})
	require.NoError(t, err)
	require.Equal(t, i, 3)
}

func TestTwoTaskQueues(t *testing.T) {
	// try to find any kind of races through running multiple queues and having them race with each other
	t.Parallel()
	rt := goja.New()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	t.Cleanup(cancel)
	vu := &modulestest.VU{
		RuntimeField: rt,
		CtxField:     ctx,
	}
	loop := eventloop.New(vu)
	fq := New(loop.RegisterCallback)
	fq2 := New(loop.RegisterCallback)
	var i int
	incrimentI := func() { i++ }
	var j int
	incrimentJ := func() { j++ }
	var k int
	incrimentK := func() { k++ }

	require.NoError(t, rt.Set("a", func() {
		for s := 0; s < 5; s++ { // make multiple goroutines
			go func() {
				for p := 0; p < 1000000; p++ {
					fq.Queue(func() error { // queue a task to increment integers
						incrimentI()
						incrimentJ()
						return nil
					})
					time.Sleep(time.Millisecond) // this is here mostly to not get a goroutine that just loops
					select {
					case <-ctx.Done():
						return
					default:
					}
				}
			}()
			go func() { // same as above but with the other queue
				for p := 0; p < 1000000; p++ {
					fq2.Queue(func() error {
						incrimentI()
						incrimentK()
						return nil
					})
					time.Sleep(time.Millisecond)
					select {
					case <-ctx.Done():
						return
					default:
					}
				}
			}()
		}
	}))

	go func() {
		<-ctx.Done()
		fq.Close()
		fq2.Close()
	}()

	err := loop.Start(func() error {
		_, err := vu.Runtime().RunString(`a()`)
		return err
	})
	require.NoError(t, err)
	loop.WaitOnRegistered()
	require.Equal(t, i, k+j)
	require.Greater(t, k, 100)
	require.Greater(t, j, 100)
}
