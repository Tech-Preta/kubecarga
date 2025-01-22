# Nome do binário
BINARY = kube-carga

# Diretório de origem
SRC_DIR = .

# Variáveis do Go
GO_FILES = $(shell find $(SRC_DIR) -name '*.go')
GO_MOD = go.mod

# Comandos
.PHONY: all build clean run

all: build

build: $(GO_FILES) $(GO_MOD)
    @echo ">> Building the binary..."
    @go build -o $(BINARY) $(SRC_DIR)

clean:
    @echo ">> Cleaning the binary..."
    @rm -f $(BINARY)

run: build
    @echo ">> Running the application..."
    @./$(BINARY) -kubeconfig ~/.kube/config