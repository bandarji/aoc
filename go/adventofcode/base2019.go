package adventofcode

import "fmt"

func NewAOCDay2019(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y19D01{}, nil
	case 2:
		return &Y19D02{}, nil
	case 3:
		return &Y19D03{}, nil
	case 4:
		return &Y19D04{}, nil
	case 5:
		return &Y19D05{}, nil
	case 6:
		return &Y19D06{}, nil
	case 7:
		return &Y19D07{}, nil
	case 8:
		return &Y19D08{}, nil
	case 9:
		return &Y19D09{}, nil
	case 10:
		return &Y19D10{}, nil
	case 11:
		return &Y19D11{}, nil
	case 12:
		return &Y19D12{}, nil
	case 13:
		return &Y19D13{}, nil
	case 14:
		return &Y19D14{}, nil
	case 15:
		return &Y19D15{}, nil
	case 16:
		return &Y19D16{}, nil
	case 17:
		return &Y19D17{}, nil
	case 18:
		return &Y19D18{}, nil
	case 19:
		return &Y19D19{}, nil
	case 20:
		return &Y19D20{}, nil
	case 21:
		return &Y19D21{}, nil
	case 22:
		return &Y19D22{}, nil
	case 23:
		return &Y19D23{}, nil
	case 24:
		return &Y19D24{}, nil
	case 25:
		return &Y19D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2019, day %d", day)
	}
}
