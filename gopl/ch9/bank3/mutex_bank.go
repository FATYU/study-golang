package bank3

import "sync"

//使用sync中的mutex进行加锁
var (
	mutex   sync.Mutex //
	balance int        //共享变量
)

func Deposit(amount int) {

	mutex.Lock()
	balance += amount //临界区代码
	mutex.Unlock()
}

func Balance() int {
	mutex.Lock()
	b := balance
	mutex.Lock()
	return b

}

//使用defer方式编写Balance

func BalanceWithDefer() int {
	mutex.Lock()
	defer mutex.Unlock() //最终会被调用，释放锁
	return balance
}
