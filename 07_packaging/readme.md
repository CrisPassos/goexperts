# Packaging

Aula de Packaging no Go(lang):

- Introduzindo go mod

  - Como trabalhar com pacotes no GO de diferentes formas
  - O Go trabalha com Módulos
  - Primeira coisa que criamos ao iniciar um projeto é criar um modulo
  - Ao fazer isso damos opção de trabalhar fora do diretório do GOPATH
  - Para fazemos isso usamos o go mod
  - Após a criação do modulo ele vai entender que aquela pasta vai conter um projeto em GO
  - pasta CMD onde armazenamos os arquivos de execucao

- Acessando pacotes criados

  - Para eu importar o metodo de outra pasta eu tenho que importar através do nome no github + nome da pasta:
  - vide code

- Exportação de Objetos

  - O Go não é orientado a objeto, mas parece muito com
  - Vamos trabalhar com modificadores de acessos (comportamento similar)
  - Usamos letras maiusculas ou minisculas para dizer que um objeto é export ou não
  - usamos as regras para definir o que é exportado ou não

- Entendendo GO MOD

  - go mod tidy analiza todo o projeto e verifica os pacotes existentes, caso ainda sejam existentes e em uso ele vai no github e instal tudo
  - funciona como o package.json
  - go.sum é como o package-lock.json
  - outra forma de instalar é atraves do go get

- Trabalhando com GO MOD replace

  - Trabalhando com go mod diferentes
  - vide projeto
  - Quando preciso usar o go mod de um projeto que estou desenvolvendo
  - A primeira alternativa é rodar o comando -replace
  - O problema com o replace é pq quando for publicar eu vou ter problemas

- Usando Workspaces
  - Recurso para trabalhar com MONOREPO
  - Trabalhando com Workspaces, antes de publicar as libraries
  - Trabalhar com workspaces locais
  - criamos um ficheiro chamado go.work
  - Go Workspace é algo independente do sistema
  - Para executar o go mod tidy com o workspace, devemos usar o go get, publicar no github ou executar o go mod tidy -e

## Commands

- Criando e Inicializa um módulo

```
    go mod init github.com/CrisPassos/goexperts/tree/main/07_packaging/01_introduzindo_go_mod
```

- Instala as depedencias

```
    go mod tidy
```

- Adicionando uma library através do Go GET

```
    go get github.com/google/uuid
```

- Replace (substitui o caminho relativo pelo diretorio)

```
    go mod edit -replace github.com/CrisPassos/goexperts/tree/main/07_packaging/03/math=../math
```

- Permite criar um workspace na raiz do diretório

```
    go work init ./math ./system
```

- Ignora os pacotes que não foram achados, usados para o go workspace

```
    go mod tidy -e
```

## Links

https://gorm.io
