# Sorveteria Três Estrelas - Frontend

Frontend da aplicação para controle de vendas e fluxo de caixa da <https://www.instagram.com/sorveteria3estrelas_/>

## Tabela de conteúdos

- [Tabela de conteúdos](#tabela-de-conteúdos)
- [Pré-requisitos](#pré-requisitos)
- [Rodando a aplicação](#rodando-a-aplicação-)
- [Contribuindo com o projeto](#contribuindo-com-o-projeto)

## Pré-requisitos

O projeto é escrito em [https://nodejs.org/en/](nodejs), então você precisa ter a linguagem instalada na sua máquina.

O [Makefile](https://www.gnu.org/software/make/) contem alguns scripts para iniciar a aplicação, suas dependencias, executar os testes e etc, você pode instalar com o linux com `sudo apt install make`, no Mac `brew install make` e no windows pode usar o [chocollatey](https://chocolatey.org/) com `choco install make`

Para saber quais sãos os comandos disponíveis, execute `make help` no terminal, você verá algo como:

```bash
❯ make help

Usage:
  make [target]

Targets:
help        Display this help
lint        Run golangci-lint
format      Format code
run         Run project
docker      Build a docker image
```

## Rodando a aplicação 🎲

```bash
# Clone este repositório
$ git clone <https://github.com/uesleicarvalhoo/sorveteria-tres-estrelas>

# Acesse a pasta do projeto no terminal
$ cd sorveteria-tres-estrelas

# Você pode facilmente iniciar as dependencias de desenvolvimento com o comando
$ docker compose up -d postgres redis zipkin kong-gateway backend

# Isso vai iniciar alguns containers:
# PostgreSQL    localhost:5432  -> Banco de dados da aplicação
# Redis         localhost:6379  -> Cache
# Zipkin        localhost:9411  -> Exporter das métricas
# Kong          localhost:8001  -> Kong para fazer o proxy e metrificar a aplicação
# backend       localhost:5000  -> backend da aplicação

# Depois de executar as dependencias de desenvolvimento, acesse a pasta onde está o frontend
$ cd frontend

# E para iniciar a aplicação em si, é só executar:
$ make run
# O servidor inciará na porta:8080 acesse <http://localhost:8080/>
```

## Contribuindo com o projeto

```bash
# Clone este repositório
$ git clone <https://github.com/uesleicarvalhoo/sorveteria-tres-estrelas>

# Acesse a pasta do projeto no terminal/cmd
$ cd sorveteria-tres-estrelas/frontend

# Faça suas alterações

# Formate o código
$ make format

# Abra uma pull request e ela será analisada
```