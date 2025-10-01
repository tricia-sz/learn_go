### 01_ Hello World
- Estrutura básica:
  - package main.
  - func main:
  - import.
- Packages:
  - Pacotes são coleções de funções pré prontas.
  - Notação: pacote.Identificador Exemplo.Println()
  - Documentação: Println
- Variaveis: 
- Variaveis nao utilizadas? Nao pode: _ nelas.
- ... funções variáticas.
- Lição principal: package main, func, main, pacoteIdentificador.

### 02_ Operador Curto de Declaração
- := (Formato de gopher)
- Uso:
 - Tipagem automática
 - Só pode repetir se houverem variáveis novas
 - != do assignment operator (operador de atribuilção)
 - Só funciona dentro de codeblock, ou seja, dentro de escopo.
- Termologia:
 - keywords (palavra-chave) sao termos reservados. Ex: package, func, main,import etc.
 - operadores, operandos
 - statement (decaração, afirmação) -> uma linha...
 - scoop (abrangencia)
 - package-level scope
- Lição Principal:
- := utilizado para criar novas variaveis dentro de codeblocks
- = usado para atribuir valores a variaveis ja existentes

### 03_ A palavra-chave var
- Variavel declarada em um code block é undefined em outro
- Para variaveis com abrangencia maior, package level scope, utilizamos `var`
- funciona em escopo global
- Prestar atenção nas chaves, colchetes e parenteses.

### 04_ Explorando Tipos
- Tipos em Go são estáticos.
- O tipo pode ser deduzido pelo compilador:
 - x := 10
 - var y = "Tia do batman"
- Ou pode ser declarado especificamente:
 - var w string "isso é uma string"
 - var z init = 15
- Tipos de dados primitivos: Disponiveis na linguagem nativamente com os blocos de contrucao.
 - Ex.  int, string, bool.

- Tipos de dados compostos: Sao tipos compostos de tipos primitivos, e criado pelo usuario. 
 - Ex. Array, slice, struct, map.

### 05_ Valor zero
- Declaração vs. Inicialização vs. Atribuicao de valor vaviaveis: caixas postais,
- O que é valor zero.
- Os zeros:
 - ints: 0
 - floats: 0.0
 - booleans: false
 - strings: ""
 - pointers, funcs, interfaces, channels, maps, nil.
 - Use := sempre que possivel
 - Usar var para package-level scope.

 