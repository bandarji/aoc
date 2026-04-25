package adventofcode

import "fmt"

func NewAOCDay2016(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y16D01{}, nil
	case 2:
		return &Y16D02{}, nil
	case 3:
		return &Y16D03{}, nil
	case 4:
		return &Y16D04{}, nil
	case 5:
		return &Y16D05{}, nil
	case 6:
		return &Y16D06{}, nil
	case 7:
		return &Y16D07{}, nil
	case 8:
		return &Y16D08{}, nil
	case 9:
		return &Y16D09{}, nil
	case 10:
		return &Y16D10{}, nil
	case 11:
		return &Y16D11{}, nil
	case 12:
		return &Y16D12{}, nil
	case 13:
		return &Y16D13{}, nil
	case 14:
		return &Y16D14{}, nil
	case 15:
		return &Y16D15{}, nil
	case 16:
		return &Y16D16{}, nil
	case 17:
		return &Y16D17{}, nil
	case 18:
		return &Y16D18{}, nil
	case 19:
		return &Y16D19{}, nil
	case 20:
		return &Y16D20{}, nil
	case 21:
		return &Y16D21{}, nil
	case 22:
		return &Y16D22{}, nil
	case 23:
		return &Y16D23{}, nil
	case 24:
		return &Y16D24{}, nil
	case 25:
		return &Y16D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2016, day %d", day)
	}
}
