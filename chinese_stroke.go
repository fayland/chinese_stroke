package chinese_stroke

import (
	"bufio"
	// "fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var StrokeMap map[string]int

func init() {
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	panic(err)
	// }
	// filepath := path.Join(dir, "/dat/Stroke.dat")

	// FIXME
	gopath := os.Getenv("GOPATH")
	filepath = path.Join(gopath, "/chinese_stroke/dat/Stroke.dat")

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	StrokeMap = make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		StrokeMap[fields[0]], _ = strconv.Atoi(fields[1])
	}
}

func GetStroke(char string) int {
	// from "\u597d" to 597D
	qr := strings.ToUpper(strings.Replace(strings.Replace(strconv.QuoteToASCII(char), "\\u", "", 1), "\"", "", 2))
	stroke, _ := StrokeMap[qr]

	return stroke
}
