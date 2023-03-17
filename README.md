# Sorveteria três estrelas

Aplicação completa para gerenciar as vendas da [sorveteria três estrelas](https://www.instagram.com/sorveteria3estrelas_/)

## Tabela de conteúdos

- [Tabela de conteúdos](#tabela-de-conteúdos)
- [Rodando a aplicação](#rodando-a-aplicação-)
- [Contribuindo com o projeto](#contribuindo-com-o-projeto)
- [Contato](#contato)
- [TODO's](#todos)

## Rodando a aplicação 🎲

```bash
# Clone este repositório
$ git clone <https://github.com/uesleicarvalhoo/sorveteria-tres-estrelas>

# Acesse a pasta do projeto no terminal
$ cd sorveteria-tres-estrelas

# Você pode facilmente iniciar as dependencias de desenvolvimento com o comando
$ docker compose build && docker compose up -d

# Isso vai iniciar alguns containers:
# PostgreSQL    localhost:5432  -> Banco de dados da aplicação
# Zipkin        localhost:9411  -> Exporter das métricas
# Kong          localhost:8001  -> Kong para fazer o proxy e metrificar a aplicação
# backend       localhost:5000  -> backend da aplicação
# frontend      localhost:8080  -> frontend da aplicação

# Você pode acessar o site através da url <http://localhost:8000/>
# E caso queira acessar a API use a url <http://localhost:8000/api>
```

## Contribuindo com o projeto

```bash
# Clone este repositório
$ git clone <https://github.com/uesleicarvalhoo/sorveteria-tres-estrelas>

# Acesse a pasta do projeto no terminal/cmd
$ cd sorveteria-tres-estrelas/frontend

# Faça suas alterações e verifique o se há necessidade de algum passo adicional no README.md dos projetos que você realizou alguma alteração

# Abra uma pull request e ela será analisada
```

## Deploy

Para realizar o deploy de uma nova versão, basta gerar uma nova release e a [action de deploy](https://github.com/uesleicarvalhoo/sorveteria-tres-estrelas/actions/workflows/release.yml) vai ser executada e fazer o deploy da nova tag no servidor remoto

## Contato

Olá, sou Ueslei Carvalho 👋🏻 criador e mantenedor deste projeto. Caso queira entrar em contato comigo, fique a vontade para utilizar qualquer um dos canais abaixo! :)

<https://www.linkedin.com/in/uesleicarvalhoo/>

📧 uesleicdoliveira@gmail.com

📷 <https://www.instagram.com/uesleicarvalhoo/>

## TODO's

- [ ] Desenhar a arquitetura
- [ ] Adicionar uma lógica para a criação do primeiro usuário no Backend
- [ ] Adicionar prints mostrando o sistema
- [ ] Adicionar prints do trace do back pro front
- [ ] Adicionar workflows de lint para o frontend
- [ ] Adicionar logica para rodar o lint/testes do front e back só quando o repositório do projeto for alterado
