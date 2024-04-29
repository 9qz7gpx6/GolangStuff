// Of note
// the folder name or file name doesn't matter.
// the compiler will use what is declared in
// the files. by converntion i use to create
// the the packages with the same name of folders

package subpackage

import "fmt"

// it is important to pay attention to the
// letter case

// it defines the member visibility

// Functions stating with upper case
// are visible to other packages.
func Write() {
	fmt.Println("from auxiliar")
}

// Functions starting with lower case
// are visible only inside their own package
func write2() {
	fmt.Println("write2 from auxiliar")
}

// Just a curiosity
// Exported functions must have a comment
// exported functions without comments generate
// warnings
func Write3() {
	write2()
}
