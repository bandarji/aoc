package adventofcode

import (
	"fmt"
	"time"
)

type Y15D01 struct{}

func (d *Y15D01) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D01) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d01p1(d.GetInput(year, day)), time.Since(start))
}

func (d *Y15D01) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d01p2(d.GetInput(year, day)), time.Since(start))
}

type Y15D02 struct{}

func (d *Y15D02) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D02) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d02p1(d.GetInput(year, day)), time.Since(start))
}

func (d *Y15D02) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d02p2(d.GetInput(year, day)), time.Since(start))
}

type Y15D03 struct{}

func (d *Y15D03) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D03) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d03p1(d.GetInput(year, day)), time.Since(start))
}

func (d *Y15D03) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d03p2(d.GetInput(year, day)), time.Since(start))
}

type Y15D04 struct{}

func (d *Y15D04) GetInput(year, day int) string {
	return y15d04input
}

func (d *Y15D04) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d04(d.GetInput(year, day), 5), time.Since(start))
}

func (d *Y15D04) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d04(d.GetInput(year, day), 6), time.Since(start))
}

type Y15D05 struct{}

func (d *Y15D05) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D05) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d05(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D05) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d05(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D06 struct{}

func (d *Y15D06) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D06) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d06(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D06) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d06(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D07 struct{}

func (d *Y15D07) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D07) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d07(d.GetInput(year, day), "a", 1), time.Since(start))
}

func (d *Y15D07) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d07(d.GetInput(year, day), "a", 2), time.Since(start))
}

type Y15D08 struct{}

func (d *Y15D08) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D08) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d08(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D08) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d08(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D09 struct{}

func (d *Y15D09) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D09) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d09(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D09) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d09(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D10 struct{}

func (d *Y15D10) GetInput(year, day int) string {
	return y15d10input
}

// Reduced cycles from 40 and 50 to 4 and 5
// Use 'make test' to validate locally
// Edit this file with 40 and 50 but prepare to wait a few minutes for the run
func (d *Y15D10) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d10(d.GetInput(year, day), 4), time.Since(start))
}

func (d *Y15D10) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d10(d.GetInput(year, day), 5), time.Since(start))
}

type Y15D11 struct{}

func (d *Y15D11) GetInput(year, day int) string {
	return y15d11input
}

func (d *Y15D11) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, y15d11(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D11) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, y15d11(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D12 struct{}

func (d *Y15D12) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D12) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d12(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D12) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d12(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D13 struct{}

func (d *Y15D13) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D13) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d13(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D13) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d13(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D14 struct{}

func (d *Y15D14) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D14) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d14(d.GetInput(year, day), y15d14seconds, 1), time.Since(start))
}

func (d *Y15D14) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d14(d.GetInput(year, day), y15d14seconds, 2), time.Since(start))
}

type Y15D15 struct{}

func (d *Y15D15) GetInput(year, day int) string {
	return y15d15Input
}

func (d *Y15D15) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d15(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D15) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d15(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D16 struct{}

func (d *Y15D16) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D16) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d16(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D16) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d16(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D17 struct{}

func (d *Y15D17) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D17) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d17(d.GetInput(year, day), 150, 1), time.Since(start))
}

func (d *Y15D17) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d17(d.GetInput(year, day), 150, 2), time.Since(start))
}

type Y15D18 struct{}

func (d *Y15D18) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D18) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d18(d.GetInput(year, day), 100, 1), time.Since(start))
}

func (d *Y15D18) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d18(d.GetInput(year, day), 100, 2), time.Since(start))
}

type Y15D19 struct{}

func (d *Y15D19) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D19) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d19(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D19) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d19(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D20 struct{}

func (d *Y15D20) GetInput(year, day int) string {
	return y15d20Input
}

func (d *Y15D20) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d20(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D20) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d20(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D21 struct{}

func (d *Y15D21) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D21) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d21(d.GetInput(year, day), y15d21StartingHP, 1), time.Since(start))
}

func (d *Y15D21) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d21(d.GetInput(year, day), y15d21StartingHP, 2), time.Since(start))
}

type Y15D22 struct{}

func (d *Y15D22) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D22) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d22(d.GetInput(year, day), 1, 50, 500), time.Since(start))
}

func (d *Y15D22) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d22(d.GetInput(year, day), 2, 50, 500), time.Since(start))
}

type Y15D23 struct{}

func (d *Y15D23) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D23) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d23(d.GetInput(year, day), "b", 1), time.Since(start))
}

func (d *Y15D23) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d23(d.GetInput(year, day), "b", 2), time.Since(start))
}

type Y15D24 struct{}

func (d *Y15D24) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D24) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d24(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y15D24) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y15d24(d.GetInput(year, day), 2), time.Since(start))
}

type Y15D25 struct{}

func (d *Y15D25) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D25) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y15d25(d.GetInput(year, day), 20151125), time.Since(start))
}

func (d *Y15D25) Part2(year, day int) string {
	return "2015 Complete!"
}

type Y16D01 struct{}

func (d *Y16D01) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D01) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d01(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D01) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d01(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D02 struct{}

func (d *Y16D02) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D02) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, y16d02(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D02) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, y16d02(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D03 struct{}

func (d *Y16D03) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D03) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d03(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D03) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d03(d.GetInput(year, day), 2), time.Since(start))
}
