package adventofcode

import "fmt"

func NewAOCDay2025(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y25D01{}, nil
	case 2:
		return &Y25D02{}, nil
	case 3:
		return &Y25D03{}, nil
	case 4:
		return &Y25D04{}, nil
	case 5:
		return &Y25D05{}, nil
	case 6:
		return &Y25D06{}, nil
	case 7:
		return &Y25D07{}, nil
	case 8:
		return &Y25D08{}, nil
	case 9:
		return &Y25D09{}, nil
	case 10:
		return &Y25D10{}, nil
	case 11:
		return &Y25D11{}, nil
	case 12:
		return &Y25D12{}, nil
	case 13:
		return &Y25D13{}, nil
	case 14:
		return &Y25D14{}, nil
	case 15:
		return &Y25D15{}, nil
	case 16:
		return &Y25D16{}, nil
	case 17:
		return &Y25D17{}, nil
	case 18:
		return &Y25D18{}, nil
	case 19:
		return &Y25D19{}, nil
	case 20:
		return &Y25D20{}, nil
	case 21:
		return &Y25D21{}, nil
	case 22:
		return &Y25D22{}, nil
	case 23:
		return &Y25D23{}, nil
	case 24:
		return &Y25D24{}, nil
	case 25:
		return &Y25D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2025, day %d", day)
	}
}
