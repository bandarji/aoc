package adventofcode

import "fmt"

func NewAOCDay2022(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y22D01{}, nil
	case 2:
		return &Y22D02{}, nil
	case 3:
		return &Y22D03{}, nil
	case 4:
		return &Y22D04{}, nil
	case 5:
		return &Y22D05{}, nil
	case 6:
		return &Y22D06{}, nil
	case 7:
		return &Y22D07{}, nil
	case 8:
		return &Y22D08{}, nil
	case 9:
		return &Y22D09{}, nil
	case 10:
		return &Y22D10{}, nil
	case 11:
		return &Y22D11{}, nil
	case 12:
		return &Y22D12{}, nil
	case 13:
		return &Y22D13{}, nil
	case 14:
		return &Y22D14{}, nil
	case 15:
		return &Y22D15{}, nil
	case 16:
		return &Y22D16{}, nil
	case 17:
		return &Y22D17{}, nil
	case 18:
		return &Y22D18{}, nil
	case 19:
		return &Y22D19{}, nil
	case 20:
		return &Y22D20{}, nil
	case 21:
		return &Y22D21{}, nil
	case 22:
		return &Y22D22{}, nil
	case 23:
		return &Y22D23{}, nil
	case 24:
		return &Y22D24{}, nil
	case 25:
		return &Y22D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2022, day %d", day)
	}
}
