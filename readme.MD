# CLI Stress Test

Um CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema gera um relatório no terminal com informações após a execução dos testes.

### Parametros de entrada:

- **--url** - URL do serviço a ser testado.
- **--requests** - Número total de requests.
- **--concurrency** - Número de chamadas simultâneas.

### Relatório: 
O relatório é apresentado ao final do teste no terminal e apresenta as seguintes informações:
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).


### Como executar: 

No terminal digite para clonar o repositório: 

```
git clone https://github.com/psilva1982/stress_test_challange.git
```

Crie o build da imagem:

```
docker build -t loadtest .
```

Execute a imagem: 

```
docker run loadtest --url=https://www.google.com --requests=1000 --concurrency=10
```
