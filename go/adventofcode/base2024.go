package adventofcode

import "fmt"

func NewAOCDay2024(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y24D01{}, nil
	case 2:
		return &Y24D02{}, nil
	case 3:
		return &Y24D03{}, nil
	case 4:
		return &Y24D04{}, nil
	case 5:
		return &Y24D05{}, nil
	case 6:
		return &Y24D06{}, nil
	case 7:
		return &Y24D07{}, nil
	case 8:
		return &Y24D08{}, nil
	case 9:
		return &Y24D09{}, nil
	case 10:
		return &Y24D10{}, nil
	case 11:
		return &Y24D11{}, nil
	case 12:
		return &Y24D12{}, nil
	case 13:
		return &Y24D13{}, nil
	case 14:
		return &Y24D14{}, nil
	case 15:
		return &Y24D15{}, nil
	case 16:
		return &Y24D16{}, nil
	case 17:
		return &Y24D17{}, nil
	case 18:
		return &Y24D18{}, nil
	case 19:
		return &Y24D19{}, nil
	case 20:
		return &Y24D20{}, nil
	case 21:
		return &Y24D21{}, nil
	case 22:
		return &Y24D22{}, nil
	case 23:
		return &Y24D23{}, nil
	case 24:
		return &Y24D24{}, nil
	case 25:
		return &Y24D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2024, day %d", day)
	}
}
