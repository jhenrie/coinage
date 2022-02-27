package coinage

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	COPPER   int64 = 1
	SILVER         = 10
	ELECTRUM       = 50
	GOLD           = 100
	PLATNUM        = 1000
)

func ParseCoinage(swap string) (*CoinPurse, error) {
	purse := NewEmptyPurse("")
	var scount string
	var coin string
	var remain string
	remain = swap

	if ok, err := validate(remain); !ok {
		return nil, err
	}

	for remain != "" {
		scount, coin, remain = nextSegment(remain)
		if scount == "" || coin == "" {
			return nil, fmt.Errorf("invalid coinage string")
		}

		count, err := strconv.Atoi(scount)
		if err != nil {
			return nil, fmt.Errorf("non-integer count for coin in string")
		}

		switch strings.ToLower(coin) {
		case "c", "cp":
			purse.AddCopper(int64(count))
		case "s", "sp":
			purse.AddSilver(int64(count))
		case "e", "ep":
			purse.AddElectrum(int64(count))
		case "g", "gp":
			purse.AddGold(int64(count))
		case "p", "pp":
			purse.AddPlatnum(int64(count))
		}
	}

	return purse, nil
}

func nextSegment(swap string) (string, string, string) {
	for i := 0; i < len(swap); i++ {
		switch c := int(swap[i]); c {
		case 'C', 'c', 'S', 's', 'E', 'e', 'G', 'g', 'P', 'p':
			if len(swap) > i+1 {
				if swap[i+1] == 'p' || swap[i+1] == 'P' {
					return swap[0:i], swap[i : i+2], swap[i+2:]
				} else {
					return swap[0:i], swap[i : i+1], swap[i+1:]
				}
			} else {
				return swap[0:i], swap[i : i+1], ""
			}
		}
	}
	return "", "", ""
}

func validate(swap string) (bool, error) {
	if match, _ := regexp.MatchString("[1-9]*[P,p,E,e,G,g,S,s,C,c]*", swap); !match {
		return false, fmt.Errorf("invalid coinage string")
	}
	return true, nil
}
