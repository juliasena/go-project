# Go Project

## Introdução

Este projeto foi desenvolvido em Go e tem como finalidade a manipulação de dados a partir do arquivo `base_teste.tsx`. O projeto busca oferecer uma solução clara e organizada para o desafio proposto, focando na inserção eficiente e estruturada de dados no banco de dados PostgreSQL.

## Funcionalidades

- Conexão com o banco de dados PostgreSQL.
- Leitura e processamento dos dados contidos em `base_teste.tsx`.
- Inserção dinâmica de dados nas tabelas do banco.
- Validação de dados para garantir integridade.

## Instalação

Para começar a usar este projeto, siga as etapas abaixo para configurar o ambiente Go e instalar as dependências necessárias:

1. **Instalar Go**
   - Certifique-se de ter o Go instalado em sua máquina. Você pode baixar a versão mais recente do Go [aqui](https://golang.org/dl/). Após a instalação, verifique se o Go está instalado corretamente executando o seguinte comando no terminal:

   ```bash
   go version
   
2. **Instalar PostgreSQL**
   - Instale o PostgreSQL em sua máquina para poder conectar-se ao banco de dados. Você encontra a versão de instalação [aqui](https://www.postgresql.org/download/). Caso tenha dúvidas na hora da configuração, recomendo este vídeo no YouTube, que é de fácil entendimento: [Vídeo de Instalação do PostgreSQL](https://www.youtube.com/watch?v=UbX-2Xud1JA) (versão para Windows).

4. **Uso de IDE Gráfica**
   - Existem diversas opções no mercado, mas recomendo o Visual Studio Code (VSCode), uma plataforma que suporta várias linguagens. Para instalar, você pode clicar [aqui](https://code.visualstudio.com/download).


Após isso, sua máquina estará pronta para receber o código deste projeto. Para isso, você só precisa dar um <git clone> ou clicar no botão "Download" deste repositório.

## Executando o Projeto

Depois de clonar o repositório, siga as instruções abaixo para executar o projeto:

1. **Navegue até o diretório do projeto:**
cd nome-do-repositorio

2. **Compile e execute o código:**
go run main.go


## Banco de dados

Após a instalação, é essencial criar um banco de dados específico para o projeto e, dentro dele, criar uma tabela que corresponda às colunas do arquivo base. Dessa forma, poderemos alimentar a tabela com os dados formatados adequadamente.

Como mostra o print abaixo:
![image](https://github.com/user-attachments/assets/6459c007-9a41-41de-a273-20a1e2e2e47d)

![image](https://github.com/user-attachments/assets/6058ff0f-6ab5-48ab-bf5e-7c04e1aa5bf4)

E desta maneira conseguiremos subir os dados rodado a partir do script.


## Contribuição
Sinta-se à vontade para contribuir para este projeto. Você pode fazer isso através de pull requests ou reportando issues.
