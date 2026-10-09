// https://coderun.yandex.ru/selections/quickstart/problems/season-tasks
package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int64
	fmt.Fscanf(os.Stdin, "%d %d", &a, &b)
	fmt.Fprintf(os.Stdout, "%d", a+b)
}
