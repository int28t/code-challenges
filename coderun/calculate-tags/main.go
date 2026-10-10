// https://coderun.yandex.ru/selections/quickstart/problems/calculate-tags
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	if n <= 2 {
		fmt.Fprintln(writer, n)
		return
	}

	var t [2]int
	t[0] = 1
	t[1] = 1

	ans := 2
	for i := 2; i < n; i++ {
		cur := t[0] + t[1]
		ans += cur
		t[0] = t[1]
		t[1] = cur
	}

	fmt.Fprintln(writer, ans)
}
