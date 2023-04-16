/*
The MIT License (MIT)

Copyright (c) 2022 sodamouse

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import "os"
import "flag"
import "log"
import "fmt"

func createDirectoryStructure(root string, sortingName string, releaseName string) {
	e := os.Chdir(root)
	checkFatal(e)

	_, e = os.Stat(sortingName)
	if os.IsNotExist(e) {
		e = os.Mkdir(sortingName, 0777)
		checkFatal(e)
	}

	e = os.Chdir(sortingName)
	checkFatal(e)

	e = os.Mkdir(releaseName, 0777)
	checkFatal(e)

	e = os.Chdir(releaseName)
	checkFatal(e)

	subDirs := [...]string{"dlc", "essentials", "extras", "instructions", "setup", "updates"}

	for _, dir := range subDirs {
		os.Mkdir(dir, 0777)
	}
}

func main() {
	rootDirPath, e := os.Getwd()
	checkFatal(e)

	sortingName := flag.String("a", "", "[Mandatory] Specifies the root directory name")
	releaseName := flag.String("b", "", "[Mandatory] Specifies the release directory name")
	showVersion := flag.Bool("v", false, "[Optional] Display program version information")
	flag.StringVar(&rootDirPath, "r", rootDirPath, "[Optional] Specifies the base path")
	flag.Parse()

	if *showVersion {
		fmt.Println("cgd version 1.0 (Go)")
		return
	}

	if *sortingName == "" || *releaseName == "" {
		flag.Usage()
		return
	}

	createDirectoryStructure(rootDirPath, *sortingName, *releaseName)
}

func checkFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
