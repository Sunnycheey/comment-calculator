package main

import (
	"flag"
	"github.com/Sunnycheey/comment-calculator/src/io"
	"log"
)

func main() {
	_ = `Compute #comment line, #total line as well as the percentage of comment line in a project.
	Currently support language: Java(.java), Go(.go), C++(.cpp), C(.c), Python(.py)

Usage:
	calculator --dir <dir>

Options:
	-h --help	show this message
	--dir	set the source code directory
`
	dir := flag.String("dir", "src", "the path of source code directory")
	flag.Parse()
	t, c, err := io.CalCuDir(*dir)
	if err != nil {
		log.Fatal(err)
	}
	percent := float64(c) / float64(t) * 100
	log.Printf("total line: %d\tcomment line: %d\tpercentage: %f%", t, c, percent)
}
