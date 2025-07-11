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

  - Interfaces são um conjunto de métodos que um tipo deve implementar
  - Elas permitem definir comportamentos sem especificar a implementação exata
  - É possível criar interfaces para agrupar diferentes tipos que compartilham comportamentos comuns
  - A interface é como se fosse um implements
  - Qualquer CLiente que tenha o método desativar vai implmentar a interface Pessoa
  - Possiblidade de usar diversos tipos de uma forma mais simples
  - Poder de abstração, encapsulamento e polimorfismo
  - A interface do Go so permite assinatura de métodos, não permite atributos

- Ponteiros

  - Ponteiros são variáveis que armazenam o endereço de memória de outra variável
  - Eles permitem manipular diretamente a memória, o que pode ser útil para otimizar o desempenho
  - É possível criar ponteiros para tipos primitivos e structs
  - O operador `&` é usado para obter o endereço de uma variável, enquanto `*` é usado para desreferenciar um ponteiro
  - É bastante simples usar ponteiros no GO
  - Pondeito = Memória -> Endereço -> Valor
  - toda vez que usamos `*` estamos apontando para o endereço de memória
  - usamos ponteiros para não duplicar a memória, economizando recursos, funciona como um cache, escopo global
  - ponteiros são passados por referência, o que significa que as alterações feitas em uma variável apontada por um ponteiro afetam a variável original
  - diferente de outras linguagens, eu não vou buscar o valor da variavel, eu vou buscar o uso do endereço de memória

- Quando usar ponteiros

  - Quando você precisa modificar o valor de uma variável dentro de uma função
  - Quando você deseja economizar memória, evitando a cópia de grandes estruturas de dados
  - Quando você precisa compartilhar dados entre diferentes partes do código sem duplicação
  - Quando eu vou fazer apenas uma cópia dos dados eu não preciso de ponteiros
  - Se eu quero valores mutaveis eu trabalho com ponteiros

- Ponteiros e Structs

  - Ponteiros são especialmente úteis ao trabalhar com structs, pois permitem modificar os campos de uma struct sem criar cópias desnecessárias
  - Ao passar um ponteiro para uma função, você pode alterar os valores dos campos da struct diretamente
  - Isso é útil para evitar a duplicação de grandes estruturas de dados e melhorar o desempenho do programa

- Interfaces Vazias

  - Interfaces vazias são um recurso poderoso do Go que permite trabalhar com qualquer tipo de dado
  - solução usada antes do generics
  - parece como se fosse o "any"

- Type assertation

  - Type assertion é uma maneira de verificar o tipo de uma variável em tempo de execução
  - É útil quando você trabalha com interfaces vazias e precisa garantir que um valor tenha um tipo específico
  - A sintaxe é `value.(Type)`, onde `value` é a variável e `Type` é o tipo esperado
  - Foca um tipo específico, se não for o tipo esperado vai gerar um panic

- Generics

  - Generics são uma maneira de escrever código que pode trabalhar com diferentes tipos de dados sem duplicação
  - Eles permitem criar funções e tipos que podem operar em qualquer tipo, mantendo a segurança de tipo
  - A sintaxe é `func Name[T any](param T)`, onde `T` é um tipo genérico
  - É possível usar generics para criar funções e estruturas de dados reutilizáveis

- Pacotes e módulos

## Commands

```
  comand
```

## Links
