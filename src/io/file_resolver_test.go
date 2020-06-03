package io

import (
	"fmt"
	"log"
	"testing"
)

func TestCalCuDir(t *testing.T) {
	tLine, cLine, err := CalCuDir("/Users/lihuichao/github/comment-calculator/src/core")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total line: %v\tcomment line: %v\n", tLine, cLine)
}
