# Módulos Privados

Aula de Módulos Privados Eventos no Go(lang):

- Introdução
  - Como baixar módulos privados, por exemplo, módulos internos da empresa
- Configuranto GOPrivate
- Autenticando no Bitbucket
  - no machine usamos api.bitbucket
- Go Proxy vc Vendor
  - Em vez de baixar diretamente do github, o GO tem um proxy que armazena os packages antes de ir acessar no github

## commands

- Infomamos ao Go que temos modulos privados

```
     go env | grep Private
```

- Mudando a propriedade do GOPrivate

```
     export GOPRIVATE=github.com/devfullcycle/fcutils-secret
```

- Baixar os packages quando não há certezas se os modulos não vai mudar

```
     go mod vendor
```

## Links

- https://github.com/devfullcycle/fcutils/tree/main/pkg/events
