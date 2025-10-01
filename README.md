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

### 06_ Pacote fmt
- Setup: strings, ints, bools.
- Strings: interpreted string literals vs. raw string literals.
 - Rune literals.
 - Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte. 
- Format printing: documentação.
 - Grupo #1: Print → standard out
   - func Print(a ...interface{}) (n int, err error)
   - func Println(a ...interface{}) (n int, err error)
   - func Printf(format string, a ...interface{}) (n int, err error)
   - Format verbs. (%v %T)
 - Grupo #2: Print → string, pode ser usado como variável
    - func Sprint(a ...interface{}) string
    - func Sprintf(format string, a ...interface{}) string
    - func Sprintln(a ...interface{}) string
  - Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
    - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
    - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
    - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
  
### 07_ Criando próprio tipo
 - Revisando: tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
 - Tem uma história que Bill Kennedy dizia que se um dia fizesse uma tattoo, ela diria "type is life."
 - Grande parte dos aspectos mais avançados de Go dependem quase que exclusivamente de tipos.
 - Como fundação para estas ferramentas, vamos aprender a declarar nossos próprios tipos.
 - Revisando: tipos são fixos. Uma vez declarada uma variável como de um certo tipo, isso é imutável.
type hotdog int → var b hotdog (main hotdog)
 - Uma variável de tipo hotdog não pode ser atribuida com o valor de uma variável tipo int, mesmo que este seja o tipo subjacente de hotdog.