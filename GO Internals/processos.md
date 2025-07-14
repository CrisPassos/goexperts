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
    - HEAP
      - Utilizado para alocação de memória dinamica. Cresce e encolhe em tempo de execução conforme a necessidade de mais ou mennos espaços
      - conforme eu vou lendo mais os dados vão indo para a memória heap, como ela é dinamica, ela vai crescendo e encolhendo conforme a necessidade
      - Ex: Alocação de memória para estruturas de dados como listas, árvores, etc
      - Tipo um garbage collector, que vai limpando a memória quando não é mais utilizada
      - O custo da HEAP é maior que o custo da stack, pois a HEAP precisa de mais gerenciamento
    - Stack
      - Utilizada para armazenar variáveis locais e chamadas de função
      - Armazena informaçõees de controle de chamadas de função, como endereços de retorno, e parametros de função. Segue uma estrtutura LIFO (Last In, First Out - Ultimo a entrar, primeiro a sair)
      - Trabalhar com stack é mais eficiente, quando eu paro de usar a função ela é automaticamente liberada
      - Stack Overflow
        - Ocorre quando a pilha de chamadas de função excede o limite de memória alocada para ela
        - Ex: Recursão infinita ou uso excessivo de memória em chamadas de função
      - Ex: Variáveis locais, parâmetros de função e endereços de retorno
