package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"image/color"
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

type HouseData struct {
	CRIM, ZN, INDUS, CHAS, NOX, RM, AGE, DIS, RAD, TAX, PTRATIO, B, LSTAT, MEDV float64
}

type Houses []HouseData

var iterations int

func main() {
	flag.IntVar(&iterations, "n", 1000, "number of iterations")
	flag.Parse()

	data, err := readData("data.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}

	var xys plotter.XYs

	for _, v := range data {
		xys = append(xys, struct{ X, Y float64 }{v.DIS, v.AGE})
	}

	err = plotData("out.png", xys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func readData(path string) (Houses, error) {
	// Open the file
	csvfile, err := os.Open("../data/boston_house_prices.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
		return nil, err
	}
	defer csvfile.Close()

	var data Houses
	s := bufio.NewScanner(csvfile)

	for s.Scan() {
		var CRIM, ZN, INDUS, CHAS, NOX, RM, AGE, DIS, RAD, TAX, PTRATIO, B, LSTAT, MEDV float64
		_, err := fmt.Sscanf(s.Text(), `%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f`, &CRIM, &ZN, &INDUS, &CHAS, &NOX, &RM, &AGE, &DIS, &RAD, &TAX, &PTRATIO, &B, &LSTAT, &MEDV)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		data = append(data, HouseData{CRIM, ZN, INDUS, CHAS, NOX, RM, AGE, DIS, RAD, TAX, PTRATIO, B, LSTAT, MEDV})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return data, nil
}

func plotData(path string, xys plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}

	// create scatter with all data points
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	m, c := linearRegression(xys, 0.01)

	// create fake linear regression result
	l, err := plotter.NewLine(plotter.XYs{
		{1, 1*m + c}, {12, 12*m + c},
	})

	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}

func linearRegression(xys plotter.XYs, alpha float64) (m, c float64) {
	// const (
	// 	min   = -100.0
	// 	max   = 100.0
	// 	delta = 0.1
	// )

	// minCost := math.MaxFloat64
	// for im := min; im < max; im += delta {
	// 	for ic := min; ic < max; ic += delta {
	// 		cost := computeCost(xys, im, ic)
	// 		if cost < minCost {
	// 			minCost = cost
	// 			m, c = im, ic
	for i := 0; i < iterations; i++ {
		dm, dc := computeGradient(xys, m, c)
		m += -dm * alpha
		c += -dc * alpha
		// fmt.Printf("grad(%.2f, %.2f) = (%.2f, %.2f)\n", m, c, dm, dc)
		fmt.Printf("cost(%.2f,%.2f) = %.2f\n", m, c, computeCost(xys, m, c))
	}

	return m, c
}

func computeCost(xys plotter.XYs, m, c float64) float64 {
	// cost 1/N * sum(y-(m*x+c))^2

	s := 0.0
	for _, xy := range xys {
		d := xy.Y - (xy.X*m + c)
		s += d * d
	}
	return s / float64(len(xys))
}

func computeGradient(xys plotter.XYs, m, c float64) (dm, dc float64) {
	for _, xy := range xys {
		d := xy.Y - (xy.X*m + c)
		dm += -xy.X * d
		dc += -d
	}
	n := float64(len(xys))
	return 2 / n * dm, 2 / n * dc
}
