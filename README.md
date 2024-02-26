# Teste de Carga em Clusters Kubernetes

Este é um exemplo de teste de carga em Go para clusters Kubernetes. O teste cria deployments de forma contínua em intervalos regulares para simular carga no cluster.

## Pré-requisitos

- Go instalado na máquina local
- Acesso a um cluster Kubernetes (local ou remoto)
- Configuração do Kubernetes (`kubectl`) ou arquivo `kubeconfig` para autenticação e autorização

## O que o código faz

O código cria deployments no cluster Kubernetes em intervalos regulares. No exemplo fornecido, ele cria deployments usando a imagem `nginx`, mas pode ser adaptado para outras imagens e tipos de recursos do Kubernetes.

## Pontos de atenção

- Caso não seja setado no momento da execução o caminho do KUBECONFIG ele herdará das variáveis de ambiente.

- O teste é para criação de deployments, caso deseje mudar o Namespace, Name, Imagem e etc realize a alteração entre as linhas 67-89.

### Clone este repositório

```bash
git clone https://github.com/Tech-Preta/k8s-test.git
cd k8s-test
```

### Inicialize o módulo Go

Execute o seguinte comando para inicializar um novo módulo Go:

```bash
go mod init
```

### Instalação das dependências

Para instalar as dependências necessárias, execute os seguintes comandos:

```bash
go get k8s.io/client-go/kubernetes
go get k8s.io/client-go/rest
go get k8s.io/client-go/tools/clientcmd
```

Esses comandos irão baixar e instalar os pacotes e suas dependências, e também atualizarão os arquivos go.mod e go.sum.

### Construindo o projeto

Agora que temos todas as dependências instaladas, podemos criar o executável para isso utilize o comando abaixo:

```bash
go build -o kube-carga .
```

Isso irá compilar o código e criar o executável no diretório atual.

### Como executar

Finalmente podemos executar a nossa aplicação com o comando abaixo:

```bash
sudo chmod +x kube-carga
./kube-carga -kubeconfig /caminho/para/seu/kubeconfig


## Próximas Funcionalidades

Aqui estão algumas funcionalidades que podem ser adicionadas no futuro:

### Limpeza dos Deployments

Adicionar uma funcionalidade para limpar os deployments criados pelo `kube-carga` no cluster após a execução do teste de carga.

#### Tarefas:

- [ ] Implementar a lógica para deletar os deployments após a conclusão do teste.
- [ ] Adicionar suporte para exclusão de outros recursos criados durante o teste, como Services e Pods.

### Verificação do Tempo de Implantação

Adicionar uma funcionalidade para verificar quanto tempo cada deployment demora para ser implantado no cluster.

#### Tarefas:

- [ ] Registrar o tempo de início da criação de cada deployment.
- [ ] Registrar o tempo de conclusão da criação de cada deployment.
- [ ] Calcular e registrar a diferença de tempo entre o início e a conclusão da implantação de cada deployment.
