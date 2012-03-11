package main

import (
        "bufio"
        "os"
        "strconv"
)

func main() {
        s := bufio.NewWriter(os.Stdout)
        for i := 0; i < 100000000; i++ {
                s.WriteString(strconv.Itoa(i))
                s.WriteString("\n")
        }
        s.Flush()
}
