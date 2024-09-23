# OpenManage-API

**OpenManage-API** é uma API simples e poderosa construída com a linguagem **Go** (Golang) usando o framework **Gin Gonic**. Este projeto tem como objetivo servir como uma introdução prática ao desenvolvimento de APIs em Go, abordando conceitos essenciais como manipulação de dados, rotas RESTful e integração com banco de dados.

A API oferece funcionalidades básicas de gerenciamento de vagas, denominadas **openings**, e permite criar, listar, atualizar e remover registros. Com uma estrutura clara e um código bem documentado, o projeto é ideal para iniciantes que desejam entender a construção de uma API CRUD (Create, Read, Update, Delete) em Go.

## Funcionalidades

A **OpenManage-API** oferece as seguintes funcionalidades para gerenciamento de vagas:

- **Criar uma vaga (opening):** Adicione novas oportunidades ao sistema.
- **Listar vagas (openings):** Visualize todas as vagas cadastradas.
- **Atualizar uma vaga (opening):** Altere detalhes de uma vaga existente.
- **Remover uma vaga (opening):** Exclua uma vaga do sistema.

## Tecnologias Utilizadas

- **Linguagem:** Go (Golang)
- **Framework:** Gin Gonic - para criação e gerenciamento de rotas.
- **SQLite** Integração com um banco de dados
- **ORM: GORM**
- **Go Modules**

## Estrutura do Projeto

- `main.go`: Ponto de entrada da aplicação.
- `controllers/`: Contém a lógica de controle das rotas e endpoints.
- `models/`: Define a estrutura das entidades e a integração com o banco de dados.
- `routes/`: Configuração das rotas da API.
- `config/`: Configurações da aplicação, como banco de dados e variáveis de ambiente.

## Endpoints da API

Aqui estão os principais endpoints da API e suas funcionalidades:

### 1. Criar uma nova vaga

- **Método:** `POST`
- **Endpoint:** `/openings`
- **Descrição:** Cria uma nova vaga no sistema.
- **Corpo da Requisição:**
  ```json
  {
	"role": "Programador Back end",
	"company":"any company",
	"location":"São miguel",
	"remote":true,
	"salary":520000,
	"link":"company.jobs.com.br"
  }
### 2. Listar todas as vagas

- **Método:** `GET`
- **Endpoint:** `/openings`
- **Descrição:** Retorna uma lista de todas as vagas cadastradas.

### 3. Atualizar uma vaga existente

- **Método:** `PUT`
- **Endpoint:** `/openings/{id}`
- **Descrição:** Atualiza os detalhes de uma vaga específica.
- **Parâmetro de Rota:** `id` - ID da vaga a ser atualizada.
- **Corpo da Requisição:**
  ```json
  {
	"role": "Programador Back end",
	"company":"any company",
	"location":"São miguel",
	"remote":true,
	"salary":520000,
	"link":"company.jobs.com.br"
  }

## 4. Remover uma Vaga

**Método:** `DELETE`  
**Endpoint:** `/openings/{id}`  
**Descrição:** Remove uma vaga do sistema.

**Parâmetro de Rota:**  
- `id` - ID da vaga a ser removida.

## Como Executar o Projeto

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/LUC4SNUN3S/OpenManage-API.git
   
2. **Navegue até o diretório do projeto:**
    ```bash
    cd OpenManage-API

3. **Instale as dependencias**
     ```bash
     go mod tidy
4. Configure o banco de dados (se necessário):
   Configure as variaveis de ambiente para conexção com o banco de dados no arquivo `.env` (se aplicavel).

5. Execute a aplicação:
   ```bash
   go run main.go
6. Acesse a API:
   Acesse `http://localhost:8080`para começar a interagir com a API.
  

         
    
   
  

  
