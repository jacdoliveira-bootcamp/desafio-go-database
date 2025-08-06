# --------------------------------------------------------
# Etapa de desenvolvimento - Utiliza imagem oficial do Go
# --------------------------------------------------------
FROM golang:1.24.3 AS dev

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia arquivos de dependências e makefile para melhor cache das camadas
COPY go.mod go.sum Makefile ./

# Faz o download apenas das dependências Go
RUN go mod download

# Instala a ferramenta Air (hot reload para Go)
RUN go install github.com/air-verse/air@latest

# Copia todo o restante do código para dentro do contêiner
COPY . .

# Expõe a porta da aplicação
EXPOSE 8080

CMD ["air"]