package core

import (
	"bufio"
	"log"
	"strings"
)

// for python style comment
type PCalculator struct {
}

func (c *PCalculator) GetLinesNumber(fileReader *bufio.Reader) (uint32, uint32, error) {

	var (
		totalLine   uint32 = 0
		commentLine uint32 = 0
		inBlock     bool   = false
		blockLine   uint32 = 0
	)
	for {
		line, err := fileReader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatal(err)
			}
			// reach the end of file
			return totalLine, commentLine, nil
		}
		if line == "\n" {
			continue
		}
		if !inBlock {
			typical, blockStart := strings.Index(line, PTypical), strings.Index(line, PBlockStart)
			if typical != -1 && blockStart != -1 {
				if typical < blockStart {
					commentLine++
				} else {
					inBlock = true
					blockLine++
				}
			} else if typical == -1 && blockStart != -1 {
				inBlock = true
				blockLine++
			} else if blockStart == -1 && typical != -1 {
				commentLine++
			} else {
				// typical = blockStart == -1, so we will do nothing.
			}
		} else if inBlock {
			// doesn't support inline doc string (e.g, '''hello''') for simplicity
			// the reason is that `PBlockEnd` and `PBlockStart` have the same format, we need an parser to get the
			// concrete meaning
			blockLine++
			end := strings.Contains(line, PBlockEnd)
			if end {
				inBlock = false
				commentLine += blockLine
				blockLine = 0
			}
		}
		totalLine++
	}
	return totalLine, commentLine, nil
}
