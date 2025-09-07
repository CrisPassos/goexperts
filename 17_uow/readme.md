# GraphQL

Aula UOW no Go(lang):

- Apresentando caso
  - Criação de um projeto simples como o do SQLC
- Entendo inconsistência
  - No final vamos fazer uma inserção de categorias e cursos na base de dados
- Entendendo Unit of Wor
  - UOW: Uint of Work
  - Cria um espaço entre as trasanções BEGIN -> Transações -> COMMIT/ROLLBACK
    -> repositorios -> uow -> getRepository -> Repository(TX)
    register -> getRepository -> unregister -> DO(call) recebe toda a transação e tudo que for acontecer
- Criando interface no UOW
- Registrando repositórios
- Implementando principais métodos
- Implementando GetRepository
- Criando novo case de Uso
- Testando implementação com uow

## commands

- command

```
  command
```

## Links
