# Teste de Carga em Clusters Kubernetes

Este é um exemplo de teste de carga em Go para clusters Kubernetes. O teste cria deployments de forma contínua em intervalos regulares para simular carga no cluster.

## Pré-requisitos

- Go instalado na máquina local
- Acesso a um cluster Kubernetes (local ou remoto)
- Configuração do Kubernetes (`kubectl`) ou arquivo `kubeconfig` para autenticação e autorização

## O que o código faz

O código cria deployments no cluster Kubernetes em intervalos regulares. No exemplo fornecido, ele cria deployments usando a imagem `nginx`, mas pode ser adaptado para outras imagens e tipos de recursos do Kubernetes.

## Como executar

1. Clone este repositório:

```bash
git clone https://github.com/Tech-Preta/k8s-test.git
cd k8s-test
