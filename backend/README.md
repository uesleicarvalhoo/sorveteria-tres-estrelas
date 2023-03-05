# sorveteria-tres-estrelas

API para controle de vendas e fluxo de caixa da <https://www.instagram.com/sorveteria3estrelas_/>

## Tabela de conteúdos

- [Tabela de conteúdos](#tabela-de-conteúdos)
- [Pré-requisitos](#pré-requisitos)
- [Rodando a aplicação](#rodando-a-aplicação-🎲)
- [Contribuindo com o projeto](#contribuindo-com-o-projeto)
- [Testes](#testes)
  - [Testes unitários](#testes-unitários)
  - [Testes de integração](#testes-de-integração)
  - [Testes end to end](#testes-end-to-end)
- [Deploy](#deploy)

## Pré-requisitos

O projeto é escrito em [https://go.dev/dl/](golang), então você precisa ter a linguagem instalada na sua máquina.

O [Makefile](https://www.gnu.org/software/make/) contem alguns scripts para iniciar a aplicação, suas dependencias, executar os testes e etc, você pode instalar com o linux com `sudo apt install make`, no Mac `brew install make` e no windows pode usar o [chocollatey](https://chocolatey.org/) com `choco install make`

Além disso o [Docker](https://www.docker.com/) é usado para executar as dependencias de desenvolvimento

Para saber quais sãos os comandos disponíveis, execute `make help` no terminal, você verá algo como:

```bash
Usage:
  make [target]

Targets:
help        Display this help
install-tools  Instal mockery, gofumpt, swago and golangci-lint
lint        Run golangci-lint
format      Format code
swagger     Generate swagger docs
run         Run app
compose     Init containers with dev dependencies
generate-mocks  Generate mock files
clean-mocks  Clean mock files
test        Run tests all tests
test/unit   Run unit tests
test/integration  Run integration tests
coverage    Run tests, make coverage report and display it into browser
clean       Remove Cache files
```

## Rodando a aplicação 🎲

```bash
# Clone este repositório
$ git clone <https://github.com/uesleicarvalhoo/sorveteria-tres-estrelas>

# Acesse a pasta do projeto no terminal
$ cd sorveteria-tres-estrelas

# Você pode facilmente iniciar as dependencias de desenvolvimento com o comando
$ make compose

# Isso vai iniciar alguns containers:
# PostgreSQL    localhost:5432  -> Banco de dados da aplicação
# Redis         localhost:6379

# E para iniciar a aplicação em si, é só executar:
$ make run
# O servidor inciará na porta:8000 acesse <http://localhost:8000/>
# O servidor já conta com documentação integrada, disponível no path /docs/swagger/index.html
```

## Contribuindo com o projeto

```bash
# Clone este repositório
$ git clone <https://github.com/uesleicarvalhoo/sorveteria-tres-estrelas>

# Acesse a pasta do projeto no terminal/cmd
$ cd sorveteria-tres-estrelas

# Faça suas alterações

# Formate o código
$ make format

# Garanta que os testes estão passando
$ make test

# Abra uma pull request e ela será analisada
```

## Testes

### Testes unitários

Os testes unitários podem ser executados com o comando

```bash
make test/unit
```

### Testes de integração

Os proprios testes de integração criam containers com as suas dependencias, então você pode executa-los com o comando

```bash
make test/integration
```

### Testes end to end

TODO

## Deploy

TODO
