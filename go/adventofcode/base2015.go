package adventofcode

import "fmt"

func NewAOCDay2015(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y15D01{}, nil
	case 2:
		return &Y15D02{}, nil
	case 3:
		return &Y15D03{}, nil
	case 4:
		return &Y15D04{}, nil
	case 5:
		return &Y15D05{}, nil
	case 6:
		return &Y15D06{}, nil
	case 7:
		return &Y15D07{}, nil
	case 8:
		return &Y15D08{}, nil
	case 9:
		return &Y15D09{}, nil
	case 10:
		return &Y15D10{}, nil
	case 11:
		return &Y15D11{}, nil
	case 12:
		return &Y15D12{}, nil
	case 13:
		return &Y15D13{}, nil
	case 14:
		return &Y15D14{}, nil
	case 15:
		return &Y15D15{}, nil
	case 16:
		return &Y15D16{}, nil
	case 17:
		return &Y15D17{}, nil
	case 18:
		return &Y15D18{}, nil
	case 19:
		return &Y15D19{}, nil
	case 20:
		return &Y15D20{}, nil
	case 21:
		return &Y15D21{}, nil
	case 22:
		return &Y15D22{}, nil
	case 23:
		return &Y15D23{}, nil
	case 24:
		return &Y15D24{}, nil
	case 25:
		return &Y15D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2015, day %d", day)
	}
}
