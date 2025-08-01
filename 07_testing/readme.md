# Testing

Aula de Testing no Go(lang):

- Iniciando com testes automatizados

  - Testes Automatizados são muito importantes para o Go(lang)
  - O GO possui uma suite de testes já incluida na linguagem
  - sempre trabalhamos com o ficheiro seja \_test.go

- Testando em batch

  - testar varías informações ao mesmo tempo
  - tipo testar várias entradas diferentes (casos extremos)
  - usamos o t.Errorf para exbir o erro no console
  - esse é o jeito nativo ou simples de fazer, porém tem bibliotecas para isso

- Verificando Cobertura de Código

  - Recursos para testar usando GO
  - Quando vamos testar além dos testes verificamos alguma cobertura
  - Aqui vamos trabalhar com coverage
  - Ferramenta GOTOOL

- Trabalhando com benchmarking

  - O Go(lang) tem uns recursos a mais para rodar testes de benchmarking
  - Quando precisamos verificar performance das rotinas ou funções
  - eu consigo rodar um número inifinito de requisições

- Fuzzing

  - Testes de mutação, gera resultados experados aleatoriamente para testar os mais diferentes casos
  - Tenta quebrar os padrões colocando os valores máximos e minimos nos testes

- Iniciando com Testify

  - Library do Go (lang) para testes
  - Pacote bem famoso no processo de testes
  - aqui usamos os assets

- Trabalhando com Mocks
  - usando a library do testify
  -

## Commands

- Commandos para executar os testes

```
    go test .
```

```
    go test -v
```

- Comando para rodar coverage

```
    go test -coverprofile=coverage.out
```

- Comando para rodar go tool

```
    go tool cover -html=coverage.out
```

- Comando para rodar o benchmarking

```
    go test -bench=.
```

```
    go test -bench= -run=ˆ#
```

- Todas as opções de parametros que podemos observar

```
    go help test
```

- Rodar o FUZZ

```
    go test -fuzz=.
```

- Rodar somente o FUZZ sem rodar os outros testes

```
    go test -fuzz=. -run=ˆ#
```

## Links

- github.com/stretchr/testify
