/*
	@From：Go语言中文网
    @Author：站长polaris
*/

/*
   @File : timedChannelOperation
   @Author: Aeishen
   @Date: 2019/12/17 21:18
   @Description: 定时Channel操作
*/

package main

import (
	"context"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type anyType int


/*
定时Channel操作：
	有时，你想要为你的 Channel 操作定时：持续尝试做一些事情，如果不能在一段时间内完成就放弃继续尝试。
	要做到这一点，你可以使用 context 或者 time，两者都很好。context可能更惯用，而 time 则更高效，但它们几乎是完全相同的,如下:
	由于并不真正关心性能（毕竟我们是在等待），我发现唯一的区别是使用 context 的解决方案会执行更多的分配（也因为使用 Timer
    的那种可以进一步优化以回收Timer）。请注意，重复使用timer是非常复杂的。因此请记住，如果仅仅为了节省10 allocs/op 的
    资源损耗而去复用timer，很可能并不值得。如果你感兴趣，这里有关于如何使用 timer 的文章。
*/
func ToChanTimedContext(ctx context.Context, d time.Duration, message anyType, c chan<- anyType) (written bool) {
    ctx, cancel := context.WithTimeout(ctx, d)
    defer cancel()
    select {
    case c <- message:
        return true
    case <-ctx.Done():
       return false
    }
}

func ToChanTimedTimer(d time.Duration, message anyType, c chan<- anyType) (written bool) {
    t := time.NewTimer(d)
    defer t.Stop()
    select {
    case c <- message:
        return true
    case <-t.C:
        return false
    }
}


/*
先来先服务：
	有时你希望将相同的消息写入多个 Channel，先写入任何可用的 Channel，但绝不要在同一 Channel 上两次写入相同的消息。
	要做到这一点，有两种方法：你可以使用局部变量屏蔽 Channel，并相应地禁用 select的case子句，或者使用 goroutine/wait 方案。如下:
	请注意，在这种情况下，性能可能很重要。而且在编写生成Goroutine的解决方案时，所花费的时间几乎是使用 select 的解决方案的4 倍。
	如果在编译期不知道 Channel 的数量，则第一个解决方案将变得更为复杂，但仍然有可能实现，而第二个解决方案则基本保持不变。
	注意：如果你的程序有许多未知大小的活动部件，则有必要进行重新设计，因为这很可能简化它。
*/
func FirstComeFirstServedSelect(message anyType, a, b chan<- anyType) {
    for i := 0; i < 2; i++ {
        select {
        case a <- message:
            a = nil
        case b <- message:
            b = nil
        }
	}
}

func FirstComeFirstServedGoroutines(message anyType, a, b chan<- anyType) {
    var wg sync.WaitGroup
	wg.Add(2)
	go func() { a <- message; wg.Done() }()
    go func() { b <- message; wg.Done() }()
    wg.Wait()
}

/*
	如果你的代码在你检查后仍然有未绑定的活动部分，这里有两个解决方案来提供支持 :
	不用说：使用反射的解决方案比使用 Goroutine 的解决方案慢几个数量级，所以请不要使用它。
*/
func FirstComeFirstServedGoroutinesVariadic(message anyType, chs ...chan<- anyType) {
    var wg sync.WaitGroup
    wg.Add(len(chs))
    for _, c := range chs {
        c := c
        go func() { c <- message; wg.Done() }()
	}
    wg.Wait()
}

func FirstComeFirstServedSelectVariadic(message anyType, chs ...chan<- anyType) {
    cases := make([]reflect.SelectCase, len(chs))
    for i, ch := range chs {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(message),
		}
	}
	for i := 0; i < len(chs); i++ {
		chosen, _, _ := reflect.Select(cases)
        cases[chosen].Chan = reflect.ValueOf(nil)
    }
}

/*
整合在一起：
	如果你想在一段时间内尝试几次发送并且如果它在这里花费了太多时间就中止尝试，这里有两种解决方案：一种是time+select，
	另一种是 context+go。如果在编译期知道Channel 的数量，则第一种更好，否则，就应该使用另一个方案。
*/
func ToChansTimedTimerSelect(d time.Duration, message anyType, a, b chan anyType) (written int) {
    t := time.NewTimer(d)
    for i := 0; i < 2; i++ {
        select {
        case a <- message:
            a = nil
        case b <- message:
            b = nil
        case <-t.C:
            return i
        }
    }
    t.Stop()
    return 2
}

func ToChansTimedContextGoroutines(ctx context.Context, d time.Duration, message anyType, ch ...chan anyType) (written int) {
    ctx, cancel := context.WithTimeout(ctx, d)
    defer cancel()
    var (
         wr int32
         wg sync.WaitGroup
    )
    wg.Add(len(ch))
    for _, c := range ch {
        c := c
        go func() {
            defer wg.Done()
            select {
            case c <- message:
               atomic.AddInt32(&wr, 1)
            case <-ctx.Done():
            }
        }()
    }
    wg.Wait()
    return int(wr)
}

func main() {

}
