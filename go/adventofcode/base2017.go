package adventofcode

import "fmt"

func NewAOCDay2017(day int) (DayRunner, error) {
	switch day {
	case 1:
		return &Y17D01{}, nil
	// case 2:
	// 	return &Y17D02{}, nil
	// case 3:
	// 	return &Y17D03{}, nil
	// case 4:
	// 	return &Y17D04{}, nil
	// case 5:
	// 	return &Y17D05{}, nil
	// case 6:
	// 	return &Y17D06{}, nil
	// case 7:
	// 	return &Y17D07{}, nil
	// case 8:
	// 	return &Y17D08{}, nil
	// case 9:
	// 	return &Y17D09{}, nil
	// case 10:
	// 	return &Y17D10{}, nil
	// case 11:
	// 	return &Y17D11{}, nil
	// case 12:
	// 	return &Y17D12{}, nil
	// case 13:
	// 	return &Y17D13{}, nil
	// case 14:
	// 	return &Y17D14{}, nil
	// case 15:
	// 	return &Y17D15{}, nil
	// case 16:
	// 	return &Y17D16{}, nil
	// case 17:
	// 	return &Y17D17{}, nil
	// case 18:
	// 	return &Y17D18{}, nil
	// case 19:
	// 	return &Y17D19{}, nil
	// case 20:
	// 	return &Y17D20{}, nil
	// case 21:
	// 	return &Y17D21{}, nil
	// case 22:
	// 	return &Y17D22{}, nil
	// case 23:
	// 	return &Y17D23{}, nil
	// case 24:
	// 	return &Y17D24{}, nil
	// case 25:
	// 	return &Y17D25{}, nil
	default:
		return nil, fmt.Errorf("no day runner for year 2017, day %d", day)
	}
}
