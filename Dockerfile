# Use uma imagem base oficial do Golang para construir o binário
FROM golang:1.19-alpine AS builder

# Crie o diretório de trabalho
WORKDIR /app

# Copie os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixe as dependências
RUN go mod download

# Copie o código-fonte para o diretório de trabalho
COPY . .

# Compile o binário
RUN go build -o kube-carga .

# Use uma imagem base mínima para o contêiner final
FROM alpine:3.16

# Crie um diretório para a aplicação
WORKDIR /app

# Copie o binário do estágio de construção
COPY --from=builder /app/kube-carga .

# Copie o arquivo de configuração, se necessário
# COPY config.yaml /app/config.yaml

# Defina o comando de entrada
ENTRYPOINT ["./kube-carga"]

# Exemplo de como passar o arquivo de configuração como argumento
# CMD ["-kubeconfig", "/caminho/para/seu/kubeconfig"]