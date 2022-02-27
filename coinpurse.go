package coinage

import "fmt"

type CoinPurse struct {
	Owner string
	Count int64
}

func NewCoinPurse(owner string, count int64) *CoinPurse {
	return &CoinPurse{
		Owner: owner,
		Count: count,
	}
}

func NewEmptyPurse(owner string) *CoinPurse {
	return NewCoinPurse(owner, 0)
}

func (me *CoinPurse) Add(purse CoinPurse) {
	me.Count = me.Count + purse.Count
}

func (me *CoinPurse) AddCopper(coins int64) {
	me.Add(*NewCoinPurse("", coins*COPPER))
}

func (me *CoinPurse) AddSilver(coins int64) {
	me.Add(*NewCoinPurse("", coins*SILVER))
}

func (me *CoinPurse) AddElectrum(coins int64) {
	me.Add(*NewCoinPurse("", coins*ELECTRUM))
}

func (me *CoinPurse) AddGold(coins int64) {
	me.Add(*NewCoinPurse("", coins*GOLD))
}

func (me *CoinPurse) AddPlatnum(coins int64) {
	me.Add(*NewCoinPurse("", coins*PLATNUM))
}

func (me *CoinPurse) Remove(purse CoinPurse) error {
	if me.Count < purse.Count {
		return fmt.Errorf("insufficient funds")
	}
	me.Count = me.Count - purse.Count

	return nil
}

func (me *CoinPurse) RemoveCopper(coins int64) error {
	return me.Remove(*NewCoinPurse("", coins*COPPER))
}

func (me *CoinPurse) RemoveSilver(coins int64) error {
	return me.Remove(*NewCoinPurse("", coins*SILVER))
}

func (me *CoinPurse) RemoveElectrum(coins int64) error {
	return me.Remove(*NewCoinPurse("", coins*ELECTRUM))
}

func (me *CoinPurse) RemoveGold(coins int64) error {
	return me.Remove(*NewCoinPurse("", coins*GOLD))
}

func (me *CoinPurse) RemovePlatnum(coins int64) error {
	return me.Remove(*NewCoinPurse("", coins*PLATNUM))
}
