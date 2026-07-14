package adventofcode

import "fmt"

func NewAOCDay2020(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y20D01{}, nil
	case 2:
		return &Y20D02{}, nil
	case 3:
		return &Y20D03{}, nil
	case 4:
		return &Y20D04{}, nil
	case 5:
		return &Y20D05{}, nil
	case 6:
		return &Y20D06{}, nil
	case 7:
		return &Y20D07{}, nil
	case 8:
		return &Y20D08{}, nil
	case 9:
		return &Y20D09{}, nil
	case 10:
		return &Y20D10{}, nil
	case 11:
		return &Y20D11{}, nil
	case 12:
		return &Y20D12{}, nil
	case 13:
		return &Y20D13{}, nil
	case 14:
		return &Y20D14{}, nil
	case 15:
		return &Y20D15{}, nil
	case 16:
		return &Y20D16{}, nil
	case 17:
		return &Y20D17{}, nil
	case 18:
		return &Y20D18{}, nil
	case 19:
		return &Y20D19{}, nil
	case 20:
		return &Y20D20{}, nil
	case 21:
		return &Y20D21{}, nil
	case 22:
		return &Y20D22{}, nil
	case 23:
		return &Y20D23{}, nil
	case 24:
		return &Y20D24{}, nil
	case 25:
		return &Y20D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2020, day %d", day)
	}
}
