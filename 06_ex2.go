package main

import "fmt"

var x int
var y string
var z bool

func main() {

	fmt.Printf("%v\n %v\n %v\n", x, y, z)

}

/*
 - Use var para declarar 3 variaveis.  Elas devem ter package-level scope. Nao atruibua  valores a essas variaveis. Utilize os seguintes identificadores e tipos:
 1. X devera ser int
 2. y devera ser string
 3. z devera  ter tipo bool
- Na funcao main:
 1. Demostre valores de cada identificador
 2. O compilador atribuiu valores para essas variaveis. Como esses valores se chamam?
*/

/*
	output:
	0  false
*/
