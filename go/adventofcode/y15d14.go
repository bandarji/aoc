package adventofcode

import (
	"fmt"
	"strings"
)

const y15d14seconds = 2503

type y15d14ReindeerList []*y15d14Reindeer

type y15d14Reindeer struct {
	name                                string
	speed, flyingDuration, restDuration int
	distance                            int
	flyingTTL, restTTL                  int
	points                              int
}

func (r *y15d14Reindeer) cycle() {
	if r.flyingTTL > 0 {
		r.distance += r.speed
		r.flyingTTL--
	} else if r.restTTL > 0 {
		r.restTTL--
		if r.restTTL == 0 {
			r.flyingTTL = r.flyingDuration
		}
	}
}

type Y15D14 struct{}

func (d *Y15D14) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D14) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d14(d.GetInput(year, day), y15d14seconds, 1))
}

func (d *Y15D14) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d14(d.GetInput(year, day), y15d14seconds, 2))
}

func y15d14(input string, seconds, part int) (winningDistance int) {
	cycle := 0
	reindeers := y15d14ParseInput(input)
	for cycle < seconds {
		for _, r := range reindeers {
			if r.flyingTTL > 0 {
				r.distance += r.speed
				r.flyingTTL--
			} else if r.restTTL > 0 {
				r.restTTL--
			}
			if r.restTTL == 0 && r.flyingTTL == 0 {
				r.restTTL = r.restDuration
				r.flyingTTL = r.flyingDuration
			}
		}
		if cycle != 0 {
			maxDistance := y15d14FindMaxDistance(reindeers)
			for _, r := range reindeers {
				if r.distance == maxDistance {
					r.points++
				}
			}
		}
		cycle++
	}
	switch part {
	case 1:
		winningDistance = y15d14FindWinningDistance(reindeers)
	case 2:
		winningDistance = y15d14FindWinningPoints(reindeers)
	}
	return
}

func y15d14FindWinningPoints(reindeers y15d14ReindeerList) (winningPoints int) {
	for _, reindeer := range reindeers {
		if reindeer.points > winningPoints {
			winningPoints = reindeer.points
		}
	}
	return
}

func y15d14FindMaxDistance(reindeers y15d14ReindeerList) (maxDistance int) {
	for _, reindeer := range reindeers {
		if reindeer.distance > maxDistance {
			maxDistance = reindeer.distance
		}
	}
	return
}

func y15d14FindWinningDistance(reindeers y15d14ReindeerList) (winningDistance int) {
	for _, reindeer := range reindeers {
		if reindeer.distance > winningDistance {
			winningDistance = reindeer.distance
		}
	}
	return
}

func y15d14ParseInput(input string) y15d14ReindeerList {
	reindeers := y15d14ReindeerList{}
	for _, line := range strings.Split(input, "\n") {
		reindeers = append(reindeers, y15d14ParseLine(line))
	}
	return reindeers
}

func y15d14ParseLine(line string) *y15d14Reindeer {
	reindeer := &y15d14Reindeer{}
	fmt.Sscanf(
		line,
		"%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
		&reindeer.name, &reindeer.speed, &reindeer.flyingDuration, &reindeer.restDuration,
	)
	return reindeer
}
