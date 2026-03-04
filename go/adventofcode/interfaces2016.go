package adventofcode

import (
	"fmt"
	"time"
)

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

type Y16D04 struct{}

func (d *Y16D04) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D04) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d04(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D04) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d04(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D05 struct{}

func (d *Y16D05) GetInput(year, day int) string {
	return y16d05Input
}

func (d *Y16D05) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, y16d05(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D05) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, y16d05(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D06 struct{}

func (d *Y16D06) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D06) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, y16d06(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D06) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, y16d06(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D07 struct{}

func (d *Y16D07) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D07) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d07(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D07) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d07(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D08 struct{}

func (d *Y16D08) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D08) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, y16d08(d.GetInput(year, day), 1, y16d08Tall, y16d08Wide), time.Since(start))
}

func (d *Y16D08) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, y16d08(d.GetInput(year, day), 2, y16d08Tall, y16d08Wide), time.Since(start))
}

type Y16D09 struct{}

func (d *Y16D09) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D09) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d09(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D09) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d09(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D10 struct{}

func (d *Y16D10) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D10) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d10(d.GetInput(year, day), 1, 17, 61), time.Since(start))
}

func (d *Y16D10) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d10(d.GetInput(year, day), 2, 17, 61), time.Since(start))
}

type Y16D11 struct{}

func (d *Y16D11) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D11) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d11(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D11) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d11(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D12 struct{}

func (d *Y16D12) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D12) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d12(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D12) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d12(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D13 struct{}

func (d *Y16D13) GetInput(year, day int) string {
	return "1352"
}

func (d *Y16D13) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d13(d.GetInput(year, day), y16d13DestinationX, y16d13DestinationY, 1), time.Since(start))
}

func (d *Y16D13) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d13(d.GetInput(year, day), y16d13DestinationX, y16d13DestinationY, 2), time.Since(start))
}

type Y16D14 struct{}

func (d *Y16D14) GetInput(year, day int) string {
	return y16d14Salt
}

func (d *Y16D14) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d14(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D14) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d14(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D15 struct{}

func (d *Y16D15) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D15) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d15(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D15) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d15(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D16 struct{}

func (d *Y16D16) GetInput(year, day int) string {
	return y16d16Input
}

func (d *Y16D16) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, y16d16(d.GetInput(year, day), y16d16DiskLengthPart1), time.Since(start))
}

func (d *Y16D16) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, y16d16(d.GetInput(year, day), y16d16DiskLengthPart2), time.Since(start))
}

type Y16D17 struct{}

func (d *Y16D17) GetInput(year, day int) string {
	return y16d17Input
}

func (d *Y16D17) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, y16d17(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D17) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, y16d17(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D18 struct{}

func (d *Y16D18) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D18) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d18(d.GetInput(year, day), 40), time.Since(start))
}

func (d *Y16D18) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d18(d.GetInput(year, day), 400_000), time.Since(start))
}

type Y16D19 struct{}

func (d *Y16D19) GetInput(year, day int) string {
	return "3004953"
}

func (d *Y16D19) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d19(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y16D19) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d19(d.GetInput(year, day), 2), time.Since(start))
}

type Y16D20 struct{}

func (d *Y16D20) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y16D20) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y16d20(d.GetInput(year, day), 4_294_967_295, 1), time.Since(start))
}

func (d *Y16D20) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y16d20(d.GetInput(year, day), 4_294_967_295, 2), time.Since(start))
}
