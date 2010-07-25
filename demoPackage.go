package demoPackage

import "os"

func Foo() {
	os.Stdout.WriteString("Foo\n");
}

func internalFoo() {
	os.Stdout.WriteString("internalFoo\n");
}

