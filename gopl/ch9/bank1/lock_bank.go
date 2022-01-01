package bank1

// 通过阻塞方式，来保证并发安全

var deposits = make(chan int)
var balances = make(chan int)

func SafeDeposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func caller() {

	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go caller()
}
