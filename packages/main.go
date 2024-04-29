package main

import (
	"firstmodule/subpackage"
	"fmt"

	"github.com/badoux/checkmail"
)

// Working with packages demands working with modules
// a module is a set of packages that compose a project
// packages are divided by folders
//
// To create a new module, type
// go mod init <modulename> in your
// commands terminal, (I'm using windows powershell
// as you can see in module.ps1)
// after running the command you will
// find the file go.mod with its name
// and go version

func main() {
	fmt.Println("from main")
	// Note that we are using upper case letter
	subpackage.Write()

	// Uncomment the line below to see it fail
	// subpackage.write2()
	subpackage.Write3()
	// to youse an external package you ust reference
	// the name after the last slash
	error := checkmail.ValidateFormat("em     ailx@domain.com")
	fmt.Println(error)
}
