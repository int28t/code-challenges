// https://coderun.yandex.ru/selections/quickstart/problems/triangle
package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c int
	fmt.Fscanf(os.Stdin, "%d\n%d\n%d", &a, &b, &c)
	if a < b+c && b < a+c && c < a+b {
		fmt.Fprintf(os.Stdout, "YES\n")
	} else {
		fmt.Fprintf(os.Stdout, "NO\n")
	}
}
