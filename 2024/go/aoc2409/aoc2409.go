package aoc2409

import (
	"aoc24/aoc2400"
	"fmt"
	"strings"
	"time"
)

const AOC2409_TEST string = `2333133121414131402`

type Disk struct {
	Sectors []int
}

func (d *Disk) Defrag(gap, finalSector int) {
	id := d.Sectors[finalSector]
	d.Sectors[finalSector] = -1
	d.Sectors[gap] = id
}

func (d *Disk) Checksum() (checksum int) {
	for index, fileID := range d.Sectors {
		if fileID > 0 {
			checksum += index * fileID
		}
	}
	return
}

func ReadDiskMap(input string) (files, gaps []int) {
	files = []int{}
	gaps = []int{}
	size := 0
	for i, c := range strings.Split(input, "") {
		size = aoc2400.StrToInt(c)
		if i%2 == 0 {
			files = append(files, size)
		} else {
			gaps = append(gaps, size)
		}
	}
	return
}

func MakeDisk(files, gaps []int) Disk {
	sectors := []int{}
	id := 0
	for i, count := range files {
		for j := range count {
			_ = j
			sectors = append(sectors, id)
		}
		id++
		if i < len(gaps) {
			for j := range gaps[i] {
				_ = j
				sectors = append(sectors, -1)
			}
		}
	}
	return Disk{Sectors: sectors}
}

// func MakeIDMap(diskmap string) (idmap string) {
// 	// 12345 -> 0..111....22222
// 	// 2333133121414131402 -> 00...111...2...333.44.5555.6666.777.888899
// 	s := strings.Builder{}
// 	id := 0
// 	readingID := true
// 	for _, c := range diskmap {
// 		v := aoc2400.StrToInt(string(c))
// 		if readingID {
// 			for i := 0; i < v; i++ {
// 				s.WriteString(fmt.Sprintf("%d", id))
// 			}
// 			id++
// 			readingID = false
// 		} else {
// 			for i := 0; i < v; i++ {
// 				s.WriteString(".")
// 			}
// 			readingID = true
// 		}
// 	}
// 	idmap = s.String()
// 	return
// }

// func CompactDisk(idMap string) (disk []string) {
// 	viz := 100
// 	if viz > len(idMap) {
// 		viz = len(idMap) - 4
// 	}
// 	disk = []string{}
// 	disk = append(disk, strings.Split(idMap, "")...)
// 	count := 0
// 	for !Uncompacted(disk) {
// 		if count < 50 {
// 			// CountDigitsAndDots(disk)
// 			fmt.Printf("Disk [%s |||| %s]\n", strings.Join(disk[:viz], ""), strings.Join(disk[len(disk)-viz:], ""))
// 		}
// 		// fmt.Printf("Uncompacted - %s\n", strings.Join(disk, ""))
// 		Compact(disk)
// 		count++
// 	}
// 	return
// }

// func CountDigitsAndDots(disk []string) {
// 	digits := 0
// 	dots := 0
// 	for i := 0; i < len(disk); i++ {
// 		if disk[i] == "." {
// 			digits = i + 1
// 			break
// 		}
// 	}
// 	for i := len(disk) - 1; i >= 0; i-- {
// 		if disk[i] >= "0" && disk[i] <= "9" {
// 			dots = i + 1
// 			break
// 		}
// 	}
// 	fmt.Printf("Disk len = %d, digits = %d, dots = %d\n", len(disk), digits, dots)
// }

// func Compact(disk []string) {
// 	firstDot := 0
// 	LastDigit := 0
// 	s := ""
// 	for i := 0; i < len(disk); i++ {
// 		if disk[i] == "." {
// 			firstDot = i
// 			break
// 		}
// 	}
// 	for i := len(disk) - 1; i >= 0; i-- {
// 		if disk[i] >= "0" && disk[i] <= "9" {
// 			LastDigit = i
// 			s = disk[i]
// 			break
// 		}
// 	}
// 	if firstDot < LastDigit {
// 		disk[firstDot] = s
// 		disk[LastDigit] = "."
// 	}
// }

// func Uncompacted(disk []string) (uncompacted bool) {
// 	uncompacted = true
// 	foundDot := false
// 	for _, c := range disk {
// 		if c >= "0" && c <= "9" {
// 			if foundDot {
// 				return false
// 			}
// 		}
// 		if c == "." {
// 			foundDot = true
// 		}
// 	}
// 	return
// }

// func Checksum(diskmapCompact []string) (checksum int) {
// 	for i, c := range diskmapCompact {
// 		// if c == "." {
// 		// 	break
// 		// }
// 		if c >= "0" && c <= "9" {
// 			checksum += i * aoc2400.StrToInt(c)
// 		}
// 	}
// 	return
// }

func Aoc240901(input string) (checksum int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	disk := MakeDisk(ReadDiskMap(input))
	gap, finalSector := 0, 0
	for {
		gap = disk.FirstGapIndex()
		if gap == -1 {
			break
		}
		finalSector = disk.FindFinalSector()
		if finalSector == -1 || finalSector <= gap {
			break
		}
		disk.Defrag(gap, finalSector)
	}
	fmt.Println("Yea, got here!")
	checksum = disk.Checksum()
	// 89425419840 - ugh!
	return
}

func (d Disk) FirstGapIndex() (gap int) {
	for i, sector := range d.Sectors {
		if sector == -1 {
			return i
		}
	}
	return -1
}

func (d Disk) FindFinalSector() (finalSector int) {
	for i := len(d.Sectors) - 1; i >= 0; i-- {
		if d.Sectors[i] != -1 {
			return i
		}
	}
	return -1
}

func (d Disk) FileStartPoint(filePosition int) (startPoint int) {
	if filePosition < 0 {
		return -1
	}
	if filePosition >= len(d.Sectors) {
		return -1
	}
	if d.Sectors[filePosition] == -1 {
		return -1
	}
	startPoint = filePosition
	fileID := d.Sectors[filePosition]
	for startPoint >= 0 && d.Sectors[startPoint] == fileID {
		startPoint--
	}
	startPoint++
	return
}

func (d Disk) FileSize(filePosition int) (fileSize int) {
	if filePosition < 0 {
		return 0
	}
	if filePosition >= len(d.Sectors) {
		return 0
	}
	if d.Sectors[filePosition] == -1 {
		return 0
	}
	fileID := d.Sectors[filePosition]
	start := filePosition
	for start >= 0 && d.Sectors[start] == fileID {
		start--
	}
	start++
	end := filePosition
	for end < len(d.Sectors) && d.Sectors[end] == fileID {
		end++
	}
	fileSize = end - start
	return
}

func (d Disk) GapSize(gapPosition int) (gapSize int) {
	for i := gapPosition; i < len(d.Sectors) && d.Sectors[i] == -1; i++ {
		gapSize++
	}
	return
}

func (d *Disk) ShiftFile(gapPos, filePos, fileSize int) {
	fileID := d.Sectors[filePos]
	for i := 0; i < fileSize; i++ {
		d.Sectors[filePos+i] = -1
	}
	for i := 0; i < fileSize; i++ {
		d.Sectors[gapPos+i] = fileID
	}
}

func Aoc240902(input string) (checksum int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	filePosition := -1
	gapCandidate := -1
	disk := MakeDisk(ReadDiskMap(input))
	maxFileID := -1
	for _, sector := range disk.Sectors {
		if sector > maxFileID {
			maxFileID = sector
		}
	}
	for fileID := maxFileID; fileID >= 0; fileID-- {
		filePosition = -1
		for i := len(disk.Sectors) - 1; i >= 0; i-- {
			if disk.Sectors[i] == fileID {
				filePosition = disk.FileStartPoint(i)
				break
			}
		}
		if filePosition == -1 {
			continue
		}
		fileSize := disk.FileSize(filePosition)
		gapCandidate = -1
		for i := 0; i < filePosition; i++ {
			if disk.Sectors[i] == -1 {
				gapSize := disk.GapSize(i)
				if gapSize >= fileSize {
					gapCandidate = i
					break
				}
				i += gapSize - 1
			}
		}
		if gapCandidate != -1 && gapCandidate < filePosition {
			disk.ShiftFile(gapCandidate, filePosition, fileSize)
		}
	}
	checksum = disk.Checksum()
	// 89425419840 - ugh!
	return
}
