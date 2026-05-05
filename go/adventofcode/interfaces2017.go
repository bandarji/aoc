package adventofcode

import (
	"fmt"
	"time"
)

type Y17D01 struct{}

func (d *Y17D01) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y17D01) Part1(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, y17d01(d.GetInput(year, day), 1), time.Since(start))
}

func (d *Y17D01) Part2(year, day int) string {
	start := time.Now()
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, y17d01(d.GetInput(year, day), 2), time.Since(start))
}

// type Y17D02 struct{}

// func (d *Y17D02) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D02) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, Y17d02(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D02) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, Y17d02(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D03 struct{}

// func (d *Y17D03) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D03) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d03(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D03) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d03(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D04 struct{}

// func (d *Y17D04) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D04) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d04(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D04) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d04(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D05 struct{}

// func (d *Y17D05) GetInput(year, day int) string {
// 	return Y17d05Input
// }

// func (d *Y17D05) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, Y17d05(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D05) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, Y17d05(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D06 struct{}

// func (d *Y17D06) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D06) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, Y17d06(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D06) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, Y17d06(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D07 struct{}

// func (d *Y17D07) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D07) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d07(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D07) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d07(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D08 struct{}

// func (d *Y17D08) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D08) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, Y17d08(d.GetInput(year, day), 1, Y17d08Tall, Y17d08Wide), time.Since(start))
// }

// func (d *Y17D08) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, Y17d08(d.GetInput(year, day), 2, Y17d08Tall, Y17d08Wide), time.Since(start))
// }

// type Y17D09 struct{}

// func (d *Y17D09) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D09) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d09(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D09) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d09(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D10 struct{}

// func (d *Y17D10) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D10) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d10(d.GetInput(year, day), 1, 17, 61), time.Since(start))
// }

// func (d *Y17D10) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d10(d.GetInput(year, day), 2, 17, 61), time.Since(start))
// }

// type Y17D11 struct{}

// func (d *Y17D11) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D11) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d11(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D11) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d11(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D12 struct{}

// func (d *Y17D12) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D12) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d12(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D12) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d12(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D13 struct{}

// func (d *Y17D13) GetInput(year, day int) string {
// 	return "1352"
// }

// func (d *Y17D13) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d13(d.GetInput(year, day), Y17d13DestinationX, Y17d13DestinationY, 1), time.Since(start))
// }

// func (d *Y17D13) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d13(d.GetInput(year, day), Y17d13DestinationX, Y17d13DestinationY, 2), time.Since(start))
// }

// type Y17D14 struct{}

// func (d *Y17D14) GetInput(year, day int) string {
// 	return Y17d14Salt
// }

// func (d *Y17D14) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d14(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D14) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d14(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D15 struct{}

// func (d *Y17D15) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D15) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d15(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D15) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d15(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D16 struct{}

// func (d *Y17D16) GetInput(year, day int) string {
// 	return Y17d16Input
// }

// func (d *Y17D16) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, Y17d16(d.GetInput(year, day), Y17d16DiskLengthPart1), time.Since(start))
// }

// func (d *Y17D16) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, Y17d16(d.GetInput(year, day), Y17d16DiskLengthPart2), time.Since(start))
// }

// type Y17D17 struct{}

// func (d *Y17D17) GetInput(year, day int) string {
// 	return Y17d17Input
// }

// func (d *Y17D17) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, Y17d17(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D17) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, Y17d17(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D18 struct{}

// func (d *Y17D18) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D18) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d18(d.GetInput(year, day), 40), time.Since(start))
// }

// func (d *Y17D18) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d18(d.GetInput(year, day), 400_000), time.Since(start))
// }

// type Y17D19 struct{}

// func (d *Y17D19) GetInput(year, day int) string {
// 	return "3004953"
// }

// func (d *Y17D19) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d19(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D19) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d19(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D20 struct{}

// func (d *Y17D20) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D20) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d20(d.GetInput(year, day), 4_294_967_295, 1), time.Since(start))
// }

// func (d *Y17D20) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d20(d.GetInput(year, day), 4_294_967_295, 2), time.Since(start))
// }

// type Y17D21 struct{}

// func (d *Y17D21) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D21) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %s (%v)", year, day, Y17d21(Y17d21Initial, d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D21) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %s (%v)", year, day, Y17d21(Y17d21Initial, d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D22 struct{}

// func (d *Y17D22) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D22) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d22(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D22) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d22(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D23 struct{}

// func (d *Y17D23) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D23) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d23(d.GetInput(year, day), Y17d23RegistersPart1), time.Since(start))
// }

// func (d *Y17D23) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d23(d.GetInput(year, day), Y17d23RegistersPart2), time.Since(start))
// }

// type Y17D24 struct{}

// func (d *Y17D24) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D24) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d24(d.GetInput(year, day), 1), time.Since(start))
// }

// func (d *Y17D24) Part2(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d (%v)", year, day, Y17d24(d.GetInput(year, day), 2), time.Since(start))
// }

// type Y17D25 struct{}

// func (d *Y17D25) GetInput(year, day int) string {
// 	return readContent(formatFilename(year, day))
// }

// func (d *Y17D25) Part1(year, day int) string {
// 	start := time.Now()
// 	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d (%v)", year, day, Y17d25(d.GetInput(year, day)), time.Since(start))
// }

// func (d *Y17D25) Part2(year, day int) string {
// 	return "2016 Complete!"
// }
