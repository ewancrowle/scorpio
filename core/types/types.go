package types

import "fmt"

type (
	// a Alpha
	Amount int64

	Address [AddressLength]byte
)

const (
	// 100a -> 1 SCX
	CoinAmount Amount = 100

	// 20 x [8] -> 160 bytes
	AddressLength = 20
)

func (a Amount) ToCoin() float64 {
	return float64(a) / float64(CoinAmount)
}

func (a Amount) ToCoinStr() string {
	return fmt.Sprintf("%f SCX", a.ToCoin())
}
