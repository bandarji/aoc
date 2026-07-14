package adventofcode

import "fmt"

func NewAOCDay2018(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y18D01{}, nil
	case 2:
		return &Y18D02{}, nil
	case 3:
		return &Y18D03{}, nil
	case 4:
		return &Y18D04{}, nil
	case 5:
		return &Y18D05{}, nil
	case 6:
		return &Y18D06{}, nil
	case 7:
		return &Y18D07{}, nil
	case 8:
		return &Y18D08{}, nil
	case 9:
		return &Y18D09{}, nil
	case 10:
		return &Y18D10{}, nil
	case 11:
		return &Y18D11{}, nil
	case 12:
		return &Y18D12{}, nil
	case 13:
		return &Y18D13{}, nil
	case 14:
		return &Y18D14{}, nil
	case 15:
		return &Y18D15{}, nil
	case 16:
		return &Y18D16{}, nil
	case 17:
		return &Y18D17{}, nil
	case 18:
		return &Y18D18{}, nil
	case 19:
		return &Y18D19{}, nil
	case 20:
		return &Y18D20{}, nil
	case 21:
		return &Y18D21{}, nil
	case 22:
		return &Y18D22{}, nil
	case 23:
		return &Y18D23{}, nil
	case 24:
		return &Y18D24{}, nil
	case 25:
		return &Y18D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2018, day %d", day)
	}
}
