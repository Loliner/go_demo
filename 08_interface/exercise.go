package main

import "fmt"

// ============================================================
// 练习：用 interface 实现一个简单的支付系统
// ============================================================
//
// 目标：
// 1. 定义 PaymentMethod interface，包含两个方法：
//    - Pay(amount float64) error
//    - Name() string
//
// 2. 实现两种支付方式：
//    - CreditCard：有信用额度 limit，余额不足时返回 error
//    - WeChatPay：有账户余额 balance，余额不足时返回 error
//
// 3. 实现 Checkout 函数，接受 PaymentMethod 和金额，完成支付并打印结果
//
// 4. 用 Type Switch 实现 paymentInfo(p PaymentMethod) 函数，
//    打印不同支付方式的详细信息（如信用卡额度、微信余额）

// TODO: 在这里写你的代码
type PaymentMethod interface {
	Pay(amount float64) error
	Name() string
}

type CreditCard struct {
	Limit float64
	Usage float64
}

func (cc *CreditCard) Pay(amount float64) error {
	if cc.Usage+amount > cc.Limit {
		return fmt.Errorf("payment failed: credit limit exceeded")
	}
	cc.Usage += amount
	fmt.Printf("%s paid %.2f\n", cc.Name(), amount)
	return nil
}

func (cc *CreditCard) Name() string {
	return "CreditCard"
}

type WeChatPay struct {
	Balance float64
}

func (wx *WeChatPay) Pay(amount float64) error {
	if amount > wx.Balance {
		return fmt.Errorf("payment failed: insufficient balance")
	}
	wx.Balance = wx.Balance - amount
	fmt.Printf("%s paid %.2f\n", wx.Name(), amount)
	return nil
}

func (wx *WeChatPay) Name() string {
	return "WeChatPay"
}

func Checkout(method PaymentMethod, amount float64) {
	if err := method.Pay(amount); err != nil {
		fmt.Printf("Checkout error: %s\n", err)
	}
}

func paymentInfo(p PaymentMethod) {
	switch m := p.(type) {
	case *CreditCard:
		fmt.Printf("%s, limit: %.2f\n", m.Name(), m.Limit)
	case *WeChatPay:
		fmt.Printf("%s, balance: %.2f\n", m.Name(), m.Balance)
	default:
		fmt.Printf("Unknown payment method: %s\n", p.Name())
	}
}

func exercise() {
	fmt.Println("=== Exercise: interface ===")

	// 测试用例（完成代码后取消注释）:

	cc := &CreditCard{Limit: 1000}
	wx := &WeChatPay{Balance: 500}

	Checkout(cc, 200) // => CreditCard paid 200.00
	Checkout(wx, 300) // => WeChatPay paid 300.00
	Checkout(wx, 300) // => payment failed: insufficient balance

	paymentInfo(cc) // => CreditCard, limit: 1000.00
	paymentInfo(wx) // => WeChatPay, balance: 200.00
}
