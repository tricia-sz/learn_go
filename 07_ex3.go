package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t %v\t %v", x, y, z)
	print(s)

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
	42       James Bond      true%
*/
