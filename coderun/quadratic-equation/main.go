// https://coderun.yandex.ru/selections/quickstart/problems/quadratic-equation
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var a, b, c int
	fmt.Fscanf(reader, "%d %d %d", &a, &b, &c)

	d := b*b - 4*a*c

	if d < 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	A := float64(a)
	B := float64(b)
	D := float64(d)

	x1 := (-B + math.Sqrt(D)) / (2.0 * A)
	x2 := (-B - math.Sqrt(D)) / (2.0 * A)

	if d == 0 {
		fmt.Fprintln(writer, 1)
		fmt.Fprintln(writer, x1)
		return
	}

	fmt.Fprintln(writer, 2)
	if x1 < x2 {
		fmt.Fprintf(writer, "%f\n%f\n", x1, x2)
	} else {
		fmt.Fprintf(writer, "%f\n%f\n", x2, x1)
	}
}
