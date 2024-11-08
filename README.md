# Golang-Stress-Test

Um sistema de teste de carga simples em Go para realizar testes de desempenho em serviços web.

## Comu usar

### Pré-requisitos

- Go (>= 1.23) instalado
- Docker (opcional)

### Executando localmente

1. Clone o repositório:

    ```bash
    git clone git@github.com:felipecardosodeoliveira/Golang-Stress-Test.git
    cd Golang-Stress-Test

2. Compile e execute o programa:

    go run cmd/main.go --url<URL> --requests=<NÚMERO_DE_requests> --concurrency=<NÚMERO_DE_CONCORRÊNCIA>

### Executando com Docker

Você pode construir uma imagem docker e executar a aplicação.

1. Build da imagem:

    docker build -t golang_stress_test .

2. Execute a imagem:

    docker run golang_stress_test --url=https://fullcycle.com.br/ --requests=1000 --concurrency=10

    Resultado do teste
    ```bash
    docker run golang_stress_test --url=https://fullcycle.com.br/ --requests=1000 --concurrency=10
    Tempo total: 2m36.401545297s
    Total de requests realizados: 1000
    Requests com status 200: 0
    Distribuição de outros códigos de status:

    docker run golang_stress_test --url=https://google.com.br/ --requests=1000 --concurrency=10
    Tempo total: 4m33.443621686s
    Total de requests realizados: 1000
    Requests com status 200: 0
    Distribuição de outros códigos de status:
   
