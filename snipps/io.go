package main

import (
        "bufio"
        "fmt"
        "io"
        "io/ioutil"
        "os"
        "strconv"
)

func run() {
}

func main() {
        s = NewInOut(os.Stdin, os.Stdout)
        run()
        s.Flush()
}

var (
        s *InOut
)

// io start

type InOut struct {
        in      []byte
        *bufio.Writer
}

func NewInOut(r io.Reader, w io.Writer) *InOut {
        in, _ := ioutil.ReadAll(r)
        return &InOut{in, bufio.NewWriter(w)}
}

func (s *InOut) next() (r int) {
        buf := s.in
        p := 1
        for (buf[0] < '0' || '9' < buf[0]) && buf[0] != '-' {
                buf = buf[1:]
        }
        if buf[0] == '-' {
                p = -1
                buf = buf[1:]

        }
        for '0' <= buf[0] && buf[0] <= '9' {
                r = 10*r + int(buf[0]-'0')
                buf = buf[1:]
        }
        r *= p
        s.in = buf
        return
}

func (s *InOut) nextStr() (r string) {
        buf := s.in
        for buf[0] == '\n' || buf[0] == ' ' {
                buf = buf[1:]
        }
        p := 0
        for buf[p] != '\n' && buf[p] != ' ' {
                p++
        }
        r = string(buf[0:p])
        s.in = buf[p:]
        return
}

func (s *InOut) print(os ...interface{}) {
        for _, o := range os {
                switch o.(type) {
                case byte:
                        s.WriteByte(o.(byte))
                case string:
                        s.WriteString(o.(string))
                case int:
                        s.WriteString(strconv.Itoa(o.(int)))
//                case int64: // for gc-2010-07-14 ie. ideone.com
//                        s.WriteString(strconv.Itoa64(o.(int64)))
                default:
                        s.WriteString(fmt.Sprint(o))
                }
        }
}

func (s *InOut) println(os ...interface{}) {
        for _, o := range os {
                s.print(o)
        }
        s.print("\n")
}

func (s *InOut) printlnNow(o interface{}) {
        fmt.Println(o)
}

// io end
