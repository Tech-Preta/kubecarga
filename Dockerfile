# Use uma imagem base oficial do Golang para construir o binário
FROM golang:1.19-alpine AS builder

# Crie o diretório de trabalho
WORKDIR /app

# Copie os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixe as dependências
RUN go mod download

# Copie apenas os arquivos de código-fonte necessários
COPY main.go ./
COPY src/ ./src/

# Compile o binário
RUN go build -o kube-carga .

# Use uma imagem base mínima para o contêiner final
FROM alpine:3.16

# Crie um diretório para a aplicação
WORKDIR /app

# Adicione um usuário não-root
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Copie o binário do estágio de construção
COPY --from=builder /app/kube-carga .

# Altere a propriedade dos arquivos para o usuário não-root
RUN chown -R appuser:appgroup /app

# Troque para o usuário não-root
USER appuser

# Defina o ponto de entrada do contêiner
ENTRYPOINT ["./kube-carga"]