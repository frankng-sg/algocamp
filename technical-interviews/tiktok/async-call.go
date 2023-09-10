/*
Title
	Weak Dependencies with timeout

Question description
	Implement a GetBalance(uid) API

	You have 2 downstream calls:
	- wallet.GetBalance(uid) - strong
	- voucher.GetVoucherInfo(uid) - weak
	- Overall latency must be within 50ms max
	for weak dependencies, should default to some reasonable value if not completed within deadline

Response Structure:
{
    balance: 10,
    voucher_info: {
        voucher_id: 1000
    }
}
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type WalletService struct{}

func (w *WalletService) GetBalance(uid int, timeout time.Duration) (float64, error) {
	time.Sleep(timeout)
	return 10.0, nil
}

type VoucherInfo struct {
	VoucherID int `json:"voucher_id"`
}

type VoucherService struct{}

func (v *VoucherService) GetVoucherInfo(uid int, timeout time.Duration) (int, error) {
	time.Sleep(timeout)
	return 1001, nil
}

type WalletInfo struct {
	Balance     float64     `json:"balance"`
	VoucherInfo VoucherInfo `json:"voucher_info"`
}

func GetCustomerWallet(uid int) (WalletInfo, error) {
	timeoutChan := time.Tick(50 * time.Millisecond)
	walletInfo := WalletInfo{}
	balanceChan := make(chan float64)
	voucherChan := make(chan int)

	go func() {
		balanceInfo, _ := walletService.GetBalance(uid, time.Duration(10*time.Millisecond))
		balanceChan <- balanceInfo
	}()

	go func() {
		voucherInfo, _ := voucherService.GetVoucherInfo(uid, time.Duration(45*time.Millisecond))
		voucherChan <- voucherInfo
	}()

	for {
		select {
		case balance := <-balanceChan:
			walletInfo.Balance = balance
		case voucher := <-voucherChan:
			walletInfo.VoucherInfo = VoucherInfo{VoucherID: voucher}
		case <-timeoutChan:
			return walletInfo, nil
		}
	}
}

var walletService WalletService
var voucherService VoucherService

func clearTerminal() {
	// Use different clear commands based on the operating system.
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	walletService = WalletService{}
	voucherService = VoucherService{}
	startTime := time.Now()
	clearTerminal()
	fmt.Println(GetCustomerWallet(1))
	fmt.Println(time.Since(startTime).Milliseconds())
}
