package adventofcode

import "fmt"

func NewAOCDay2021(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y21D01{}, nil
	case 2:
		return &Y21D02{}, nil
	case 3:
		return &Y21D03{}, nil
	case 4:
		return &Y21D04{}, nil
	case 5:
		return &Y21D05{}, nil
	case 6:
		return &Y21D06{}, nil
	case 7:
		return &Y21D07{}, nil
	case 8:
		return &Y21D08{}, nil
	case 9:
		return &Y21D09{}, nil
	case 10:
		return &Y21D10{}, nil
	case 11:
		return &Y21D11{}, nil
	case 12:
		return &Y21D12{}, nil
	case 13:
		return &Y21D13{}, nil
	case 14:
		return &Y21D14{}, nil
	case 15:
		return &Y21D15{}, nil
	case 16:
		return &Y21D16{}, nil
	case 17:
		return &Y21D17{}, nil
	case 18:
		return &Y21D18{}, nil
	case 19:
		return &Y21D19{}, nil
	case 20:
		return &Y21D20{}, nil
	case 21:
		return &Y21D21{}, nil
	case 22:
		return &Y21D22{}, nil
	case 23:
		return &Y21D23{}, nil
	case 24:
		return &Y21D24{}, nil
	case 25:
		return &Y21D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2021, day %d", day)
	}
}
