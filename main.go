// MIT License
// Copyright (c) 2020 kurosawa
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package main

import (
	"fmt"
	"regexp"
	"github.com/mitchellh/go-ps"
	"log"
	"flag"
	"syscall"
)

func main() {
	var (
		l = flag.Bool("l", false, "show process list")
		k = flag.String("k", "KillProcessName", "regexp process name for kill")
	)
	flag.Parse()
	killname := *k
	if *l {
		processList()
	} else {
		processKill(killname)
	}

}

func processList() {
	processes, err := ps.Processes()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range processes {
		fmt.Printf("%d: %s\n", p.Pid(), p.Executable())
	}
}

func processKill(killname string) {
	processes, err := ps.Processes()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range processes {
		process_name := p.Executable()
		matched, err := regexp.Match(killname, []byte(process_name))
		if err != nil {
			log.Fatal(err)
		}
		if matched {
			syscall.Kill(p.Pid(), syscall.SIGKILL)
			fmt.Println("KILL pid:", p.Pid(), " pname:", process_name)
		}
	}
}
