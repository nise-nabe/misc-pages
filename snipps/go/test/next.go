package main

import (
        "index/suffixarray"
        "io/ioutil"
        "log"
        "regexp"
        "time"
)

func main() {
        str, _ := ioutil.ReadFile("large.in")
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
        }(string(str))
        t := time.Now().UnixNano()
        for i := 0; i < 100000000; i++ {
                next()
        }
        log.Println(time.Now().UnixNano() - t)
}
