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
