package io

import (
	"bufio"
	"github.com/Sunnycheey/comment-calculator/src/core"
	"log"
	"os"
	"path"
)

func CalCuDir(dir string) (uint32, uint32, error) {

	// return total line and comment line
	// open dir
	// please use the full path

	directory, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer directory.Close()

	info, err := directory.Readdir(0)

	if err != nil {
		log.Fatal(err)
	}

	var commentLine uint32 = 0
	var totalLine uint32 = 0

	for _, file := range info {
		if !file.IsDir() {
			filePath := path.Join(dir, file.Name())
			extension := path.Ext(filePath)
			// open the file
			srcFile, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}
			fileReader := bufio.NewReader(srcFile)
			t, c, err := CalcuFile(fileReader, extension)
			if err != nil {
				log.Fatal(err)
			}
			totalLine += t
			commentLine += c
			srcFile.Close()
		} else {
			innerDir := path.Join(dir, file.Name())
			t, c, err := CalCuDir(innerDir)
			if err != nil {
				log.Fatal(err)
			}
			totalLine += t
			commentLine += c
		}
	}
	return totalLine, commentLine, nil
}

func CalcuFile(fileReader *bufio.Reader, extension string) (uint32, uint32, error) {

	var commentLine uint32 = 0
	var totalLine uint32 = 0
	var calculator core.Calculator

	if extension == core.C || extension == core.Go || extension == core.Java || extension == core.CPP {
		calculator = &core.CCalculator{}
	} else if extension == core.Python {
		calculator = &core.PCalculator{}
	}
	t, c, err := calculator.GetLinesNumber(fileReader)
	if err != nil {
		log.Fatal(err)
	}
	totalLine += t
	commentLine += c
	return totalLine, commentLine, nil
}
