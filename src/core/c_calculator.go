package core

import (
	"bufio"
	"log"
	"strings"
)

type CCalculator struct {
	// the path of directory
}

func (c *CCalculator) GetLinesNumber(fileReader *bufio.Reader) (uint32, uint32, error) {

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
			typical, blockStart := strings.Index(line, CTypical), strings.Index(line, CBlockStart)
			if typical != -1 && blockStart != -1 {
				if typical < blockStart {
					commentLine++
				} else {
					inBlock = true
					blockLine++
				}
			} else if typical == -1 && blockStart != -1 {
				if strings.Contains(line, CBlockEnd) {
					// deal with inline block
					commentLine++
				} else {
					inBlock = true
					blockLine++
				}
			} else if blockStart == -1 && typical != -1 {
				commentLine++
			} else {
				// typical = blockStart == -1, so we will do nothing.
			}
		} else if inBlock {
			// deal with inline block
			blockLine++
			end := strings.Contains(line, CBlockEnd)
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
