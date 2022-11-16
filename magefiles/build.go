package main

import (
	"log"

	"github.com/magefile/mage/sh"
)

const (
	goldenDir        = "goldenmaster"
	generatedTestDir = "test"
)

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func gitClone(url, dir string) {
	must(sh.RunV("git", "clone", "--depth=1", url, dir))
}

func runDiff(dir1, dir2 string) {
	must(sh.RunV("git", "--no-pager", "diff", "--no-index", dir1, dir2))
}

func goInstall(pkg string) {
	must(sh.RunV("go", "install", pkg))
}

func genSol(sol string) {
	must(sh.RunV("apigear", "g", "sol", sol))
}

func rmDir(dir string) {
	must(sh.Rm(dir))
}

// Install installs the apigear cli and testbed-apis.
func Install() {
	goInstall("github.com/apigear-io/cli/cmd/apigear@latest")
}

// Diff runs the generator and compares the output with the golden master.
func Diff() {
	genSol("./apigear/goldenmaster.solution.yaml")
	runDiff(goldenDir, generatedTestDir)
}

// Clean removes all generated files.
func Clean() {
	rmDir(generatedTestDir)
}

func DiffOnly() {
	runDiff(goldenDir, generatedTestDir)
}

func GenTest() {
	rmDir(generatedTestDir)
	genSol("./apigear/goldenmaster.solution.yaml")
}
