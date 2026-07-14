package adventofcode

import "fmt"

func NewAOCDay2023(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y23D01{}, nil
	case 2:
		return &Y23D02{}, nil
	case 3:
		return &Y23D03{}, nil
	case 4:
		return &Y23D04{}, nil
	case 5:
		return &Y23D05{}, nil
	case 6:
		return &Y23D06{}, nil
	case 7:
		return &Y23D07{}, nil
	case 8:
		return &Y23D08{}, nil
	case 9:
		return &Y23D09{}, nil
	case 10:
		return &Y23D10{}, nil
	case 11:
		return &Y23D11{}, nil
	case 12:
		return &Y23D12{}, nil
	case 13:
		return &Y23D13{}, nil
	case 14:
		return &Y23D14{}, nil
	case 15:
		return &Y23D15{}, nil
	case 16:
		return &Y23D16{}, nil
	case 17:
		return &Y23D17{}, nil
	case 18:
		return &Y23D18{}, nil
	case 19:
		return &Y23D19{}, nil
	case 20:
		return &Y23D20{}, nil
	case 21:
		return &Y23D21{}, nil
	case 22:
		return &Y23D22{}, nil
	case 23:
		return &Y23D23{}, nil
	case 24:
		return &Y23D24{}, nil
	case 25:
		return &Y23D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2023, day %d", day)
	}
}
