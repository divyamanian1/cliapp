package pnl

import "fmt"

func CalculatePL(c float32, pc float32) string {
	if c > pc {
		s := (c - pc) * 100
		return fmt.Sprintf("Profit made for 10 shares: %.2f.", s)
	}
	s := (pc - c) * 100
	return fmt.Sprintf("Loss made for 10 shares: %.2f.", s)
}
