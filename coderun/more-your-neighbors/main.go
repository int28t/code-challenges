// https://coderun.yandex.ru/selections/quickstart/problems/more-your-neighbors/solutions
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

	var a, b, c, cnt, ans int
	for {
		_, err := fmt.Fscan(reader, &c)
		if err != nil {
			break
		}
		cnt++
		if cnt >= 3 && b > a && b > c {
			ans++
		}
		a = b
		b = c
	}

	fmt.Fprintln(writer, ans)
}
