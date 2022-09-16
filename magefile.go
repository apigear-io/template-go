//go:build mage

package main

import (
	"log"

	"github.com/magefile/mage/sh"
)

const (
	goldenDir = "goldenmaster"
	tmpDir    = "tmp"
	apisDir   = "testbed-apis"
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
func genX(out string) {
	must(sh.RunV("cli", "g", "x", "-t", ".", "-i", "testbed-apis/apigear", "-o", out, "-f", "all", "--force"))
}

func rmDir(dir string) {
	must(sh.Rm(dir))
}

// Install installs the apigear cli and testbed-apis.
func Install() {
	rmDir("testbed-apis")
	goInstall("github.com/apigear-io/cli@latest")
	gitClone("https://github.com/apigear-io/go-testbed-apis.git", "testbed-apis")
}

// Master generates the golden master.
func Master() {
	rmDir(goldenDir)
	genX(goldenDir)
}

// Diff runs the generator and compares the output with the golden master.
func Diff() {
	rmDir(tmpDir)
	genX(tmpDir)
	runDiff(goldenDir, tmpDir)
}

// Clean removes all generated files.
func Clean() {
	rmDir(tmpDir)
	rmDir(apisDir)
}
