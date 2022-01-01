package bank2

// 使用缓冲量为1的channel作为锁🔒
var (
	semaphore = make(chan struct{}, 1) //二元信号量
	balance   = 0
)

func Deposit(amount int) {
	semaphore <- struct{}{} //获取token
	balance += amount
	<-semaphore //释放token
}

func Balance() int {
	semaphore <- struct{}{} //获取token
	b := balance
	<-semaphore //释放token
	return b
}
