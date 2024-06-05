go test . -v
go test -run=TestDoCheckClose -v

# 通道关闭原则
    一个常用的使用Go通道的原则是不要在数据接收方或者在有多个发送者的情况下关闭通道。换句话说，我们只应该让一个通道唯一的发送者关闭此通道。

1. 粗鲁的方式（非常不推荐）

如果 channel 已经被关闭，再次关闭会产生 panic，这时通过 recover 使程序恢复正常。
```
func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// 一个函数的返回结果可以在defer调用中修改。
			justClosed = false
		}
	}()

	// 假设ch != nil。
	close(ch)   // 如果 ch 已关闭，将 panic
	return true // <=> justClosed = true; return
}
```

2. 礼貌的方式

使用 sync.Once 或互斥锁(sync.Mutex)确保 channel 只被关闭一次。

```
type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}
```

3. 优雅的方式

- 情形一：M个接收者和一个发送者，发送者通过关闭用来传输数据的通道来传递发送结束信号。
- 情形二：一个接收者和N个发送者，此唯一接收者通过关闭一个额外的信号通道来通知发送者不要再发送数据了。
- 情形三：M个接收者和N个发送者，它们中的任何协程都可以让一个中间调解协程帮忙发出停止数据传送的信号。