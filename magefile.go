//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/sh"
)

const (
	golden_dir = "goldenmaster"
	test_dir   = "test"
)

func runDiff(dir1, dir2 string) error {
	_ = sh.RunV("git", "--no-pager", "diff", "--no-index", dir1, dir2)
	return nil
}

func rmDir(dir string) error {
	return os.RemoveAll(dir)
}

func runGenerator(output string) error {
	return sh.Run("cli", "g", "x", "-t", ".", "-i", "testbed.api", "-o", output, "-f", "all", "--force")
}

// Install installs the apigear cli.
func Install() error {
	return sh.Run("go", "install", "github.com/apigear-io/cli@latest")
}

// Gen generates the golden master.
func Gen() error {
	output := golden_dir
	if err := rmDir(output); err != nil {
		return err
	}
	if err := runGenerator(output); err != nil {
		return err
	}
	return nil
}

// Test runs the generator and compares the output with the golden master.
func Test() error {
	// mg.Deps(Clean)
	// mg.Deps(Install)
	// mg.Deps(Gen)
	output := test_dir
	if err := runGenerator(output); err != nil {
		return err
	}
	if err := runDiff(golden_dir, output); err != nil {
		return err
	}
	return nil
}

// Clean removes all generated files.
func Clean() error {
	return rmDir("test")
}
