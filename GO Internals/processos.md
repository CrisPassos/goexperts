# Go Internals - Processo

Aulas mais teoricas sobre o Go(lang):

- Processos:
  - Instância de um programa em execucão
  - Componentes para o processo funcionar é preciso de componentes organizados
    - Endereçamento (região da memória dedicada a um processo)
    - Contextos
      - Conjunto de dados que o SO "salva"para gerenciar um processo (para um processo rodar em determinado aspecto eu preciso executar esse Contexto)
        - EX.: Programa Counter (PC) ou Instruction Pointer (IP)
        - Possui o enderco da próxima instrução que o precessador irá executar
        - Auxília no Contexto Switch
    - Registro do Processador
      - Áreas que temporariamente armazenam no CPU dados e endereços para realizar a execução
      - Dados
        - Ex: Realiza operações aritméticas e lógicas
      - Registro de Endereços
        - Armazenamento em memória, incluindo stacks pointers
        - Ex: Ao acessar uma variável o CPU possui um registro na memória para guardar seu valor
