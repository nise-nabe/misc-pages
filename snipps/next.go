package main

import (
        "index/suffixarray"
        "log"
        "regexp"
)

func main() {
        str := `1234 2 3333 4 5  aaaaa
s
aaa
`
        next := func(str string) func() string {
                reg, _ := regexp.Compile("\\S+")
                is := suffixarray.New([]byte(str)).FindAllIndex(reg, -1)
                return func() (result string) {
                        if len(is) < 1 {
                                return ""
                        }
                        result = str[is[0][0]:is[0][1]]
                        is = is[1:]
                        return
                }
        }(str)
        for x := next(); x != ""; x = next() {
                log.Println(x)
        }
}
