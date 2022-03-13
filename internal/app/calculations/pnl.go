package pnl

import "fmt"

const sharesCount int = 10

func CalculatePL(c float32, pc float32) string {
	if c > pc {
		s := (c - pc) * float32(sharesCount)
		return fmt.Sprintf("Profit made for %d shares: %.2f.", sharesCount, s)
	} else if c < pc {
		s := (pc - c) * float32(sharesCount)
		return fmt.Sprintf("Loss made for %d shares: %.2f.", sharesCount, s)
	} else {
		return fmt.Sprintf("No profit or loss")
	}
}
