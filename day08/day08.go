package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coordinates struct {
	x int
	y int
}

func findAntennas(antennaMap []string) map[string][]Coordinates {
	var antennas map[string][]Coordinates
	antennas = make(map[string][]Coordinates)
	for i := 0; i < len(antennaMap); i++ {
		for j := 0; j < len(antennaMap[i]); j++ {
			if string(antennaMap[i][j]) != "." {
				antenna, found := antennas[string(antennaMap[i][j])]
				if !found {
					antenna = []Coordinates{}
				}
				antenna = append(antenna, Coordinates{i, j})
				antennas[string(antennaMap[i][j])] = antenna
			}
		}
	}
	return antennas
}

func calculateDistances(antenna1 Coordinates, antenna2 Coordinates) int {
	p1 := math.Pow(float64(antenna1.x-antenna2.x), 2)
	p2 := math.Pow(float64(antenna1.y-antenna2.y), 2)
	return int(math.Sqrt(p1 + p2))
}

func calculateVector(antenna1 Coordinates, antenna2 Coordinates) (int, int) {
	vectorX := math.Abs(float64(antenna1.x) - float64(antenna2.x))
	vectorY := math.Abs(float64(antenna1.y) - float64(antenna2.y))
	return int(vectorX), int(vectorY)
}

func drawChar(antennaMap []string, x int, y int, c rune) {
	out := []rune(antennaMap[x])
	out[y] = c
	antennaMap[x] = string(out)
}

func printMap(antennaMap []string) {
	for i := 0; i < len(antennaMap); i++ {
		fmt.Println(antennaMap[i])
	}
}

func partOne(antennaMap []string) int {
	antennasMap := findAntennas(antennaMap)
	count := 0
	for _, antennas := range antennasMap {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				vectorX, vectorY := calculateVector(antennas[i], antennas[j])

				var antiNode1, antiNode2 Coordinates

				if antennas[i].x < antennas[j].x {
					if antennas[i].y < antennas[j].y {
						antiNode1.x = antennas[i].x - vectorX
						antiNode1.y = antennas[i].y - vectorY
						antiNode2.x = antennas[j].x + vectorX
						antiNode2.y = antennas[j].y + vectorY
					} else {
						antiNode1.x = antennas[i].x - vectorX
						antiNode1.y = antennas[i].y + vectorY
						antiNode2.x = antennas[j].x + vectorX
						antiNode2.y = antennas[j].y - vectorY
					}
				} else if antennas[i].x > antennas[j].x {
					if antennas[i].y < antennas[j].y {
						antiNode1.x = antennas[i].x + vectorX
						antiNode1.y = antennas[i].y + vectorY
						antiNode2.x = antennas[j].x - vectorX
						antiNode2.y = antennas[j].y - vectorY
					} else {
						antiNode1.x = antennas[i].x - vectorX
						antiNode1.y = antennas[i].y - vectorY
						antiNode2.x = antennas[j].x + vectorX
						antiNode2.y = antennas[j].y + vectorY
					}
				}

				if antiNode1.x > -1 && antiNode1.y > -1 && antiNode1.x < len(antennaMap) && antiNode1.y < len(antennaMap[antiNode1.x]) && antennaMap[antiNode1.x][antiNode1.y] != '#' {
					drawChar(antennaMap, antiNode1.x, antiNode1.y, '#')
					count++
				}

				if antiNode2.x > -1 && antiNode2.y > -1 && antiNode2.x < len(antennaMap) && antiNode2.y < len(antennaMap[antiNode2.x]) && antennaMap[antiNode2.x][antiNode2.y] != '#' {
					drawChar(antennaMap, antiNode2.x, antiNode2.y, '#')
					count++
				}

			}
		}
	}
	return count
}

func partTwo(antennaMap []string) int {
	antennasMap := findAntennas(antennaMap)
	count := 0
	for _, antennas := range antennasMap {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				vectorX, vectorY := calculateVector(antennas[i], antennas[j])

				var antiNode1, antiNode2 Coordinates
				antiNode1 = antennas[i]
				antiNode2 = antennas[j]

				for (antiNode1.x > -1 && antiNode1.y > -1 && antiNode1.x < len(antennaMap) && antiNode1.y < len(antennaMap[antiNode1.x])) || (antiNode2.x > -1 && antiNode2.y > -1 && antiNode2.x < len(antennaMap) && antiNode2.y < len(antennaMap[antiNode2.x])) {
					if antennas[i].x < antennas[j].x {
						if antennas[i].y < antennas[j].y {
							antiNode1.x = antiNode1.x - vectorX
							antiNode1.y = antiNode1.y - vectorY
							antiNode2.x = antiNode2.x + vectorX
							antiNode2.y = antiNode2.y + vectorY
						} else {
							antiNode1.x = antiNode1.x - vectorX
							antiNode1.y = antiNode1.y + vectorY
							antiNode2.x = antiNode2.x + vectorX
							antiNode2.y = antiNode2.y - vectorY
						}
					} else if antennas[i].x > antennas[j].x {
						if antennas[i].y < antennas[j].y {
							antiNode1.x = antiNode1.x + vectorX
							antiNode1.y = antiNode1.y + vectorY
							antiNode2.x = antiNode2.x - vectorX
							antiNode2.y = antiNode2.y - vectorY
						} else {
							antiNode1.x = antiNode1.x - vectorX
							antiNode1.y = antiNode1.y - vectorY
							antiNode2.x = antiNode2.x + vectorX
							antiNode2.y = antiNode2.y + vectorY
						}
					}

					fmt.Println(antiNode1, antiNode2)

					if antiNode1.x > -1 && antiNode1.y > -1 && antiNode1.x < len(antennaMap) && antiNode1.y < len(antennaMap[antiNode1.x]) && antennaMap[antiNode1.x][antiNode1.y] == '.' {
						drawChar(antennaMap, antiNode1.x, antiNode1.y, '#')
						count++
					}

					if antiNode2.x > -1 && antiNode2.y > -1 && antiNode2.x < len(antennaMap) && antiNode2.y < len(antennaMap[antiNode2.x]) && antennaMap[antiNode2.x][antiNode2.y] == '.' {
						drawChar(antennaMap, antiNode2.x, antiNode2.y, '#')
						count++
					}
				}
			}
		}
		count += len(antennas)
	}
	return count
}

func main() {
	inputFile, _ := os.Open("day08/input.txt")
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	antennaMap := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		antennaMap = append(antennaMap, scanner.Text())
	}

	antennaMapPartOne := append([]string(nil), antennaMap...)
	resultPartOne := partOne(antennaMapPartOne)
	printMap(antennaMapPartOne)
	fmt.Println("Part one:", resultPartOne)

	antennaMapPartTwo := append([]string(nil), antennaMap...)
	resultPartTwo := partTwo(antennaMapPartTwo)
	printMap(antennaMapPartTwo)
	fmt.Println("Part two:", resultPartTwo)
}
