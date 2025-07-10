# Fundação

Aula de Fundação do Go(lang):

- Entendendo a primeira linha

  - todos os packages devem seguir o nome do diretório
  - tudo começa com package, todo package começa com o nome do diretório, caso contrário começa com main

- Declarando e atribuição

  - se uma varíavel nasceu com um determinado tipo ela vai morrer com esse tipo, não é possível mudar
  - se uma varíavel for declarada, mas nunca usada o compilador vai reclamar

- Criação de tipos

  - criação dos próprios tipos, além dos nativos
  - utilizamos a palavra reservada `type``
  - podemos usar uma lib externa para criar os tipos

- Importando fmt e tipagem

  - library fmt é padrão do GO
  - %d - digito
  - %v - string
  - \n - pula linha

- Percorrendo Arrays

  - array, trabalha como varíavel com tamanho fixo que deve ser percorrido
  - len() método nativo que conta o tamanho do array
  - normalmente não trabalhamos muito com array e sim Slices

- Slices

  - "arrays infinitos""
  - dividos em ponteiro, tamanho e capacidade
  - por de trás dos panos trabalha com arrays

- Maps

  - HashTables
  - Key/Value, sem ordenação necessariamente

- Funções
- Funções variádicas

  - trabalha com parametros infinitos, quando não sabemos quantos parametros vamos passar
  - usamos o ... para fazer isso

- Closures

  - Funçoes anonimas
  - ter funcoes dentro de outra

- Iniciando com Structs

  - O GO não é orientado a objetos
  - Existe o GO way de programar
  - Structs é o que mais se aproxima das classes que conhecemos
  - Struct é um tipo de classe, mas não é uma classe
  - Criar uma estrutura para armazenar as informcoes tipo cliente
  - Um Struct é um tipo de dado composto que agrupa variáveis de diferentes tipos sob um mesmo nome.
  - é possivel mudar o valor de um campo de um struct, mas não é possível mudar o tipo do campo

- Composição de Structs

  - Composição é a forma de criar um novo tipo de struct a partir de outros tipos de struct
  - É possível criar structs dentro de structs, o que permite uma estrutura mais complexa e organizada
  - Pode ser encadeada como a Herança na orientação a objetos, mas não é exatamente o mesmo conceito

- Metódos em Structs

  - É possível criar métodos para structs, o que permite encapsular comportamentos relacionados a esses tipos de dados
  - Métodos são funções associadas a um tipo específico, permitindo que você execute ações relacionadas a esse tipo
  - usamos () para definir o tipo que o método pertence a struct

- Interfaces

## Commands

```
  comand
```

## Links
