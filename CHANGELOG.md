# CHANGELOG

Todas as mudanças notáveis ​​neste projeto serão documentadas neste arquivo.

O formato é baseado em [Mantenha um Changelog](https://keepachangelog.com/pt-BR/1.1.0/)
e este projeto adere a [Versionamento Semântico](https://semver.org/lang/pt-BR/).

<!--
## [Unreleased] - yyyy-mm-dd

Here we write upgrading notes for brands. It's a team effort to make them as
straightforward as possible.

### Added

### Changed

### Fixed

### Breaking Changes
-->

## [1.1.0] - 2025-01-22

### Added
- Adicionado suporte ao build multiplataforma e Docker via GoReleaser.
- Adicionado arquivo `.dockerignore`.

### Changed
- Refatorado o workflow de release para utilizar GoReleaser e remoção de arquivos de configuração depreciados.
- Atualização dos nomes dos workflows e modificação do gatilho de release para a branch main.
- Atualização das versões do Go nos workflows.

### Fixed
- Correções no workflow de release para permitir gatilhos apenas em pull requests e tags.
- Ajuste no formato da imagem Docker e manipulação de conteúdo na configuração do GoReleaser.
- Limpeza de artefatos antes de rodar o GoReleaser.
- Diversas correções em pipelines com erros.

### Removed
- Remoção da configuração do CodeQL.

### Build & Dependencies
- Atualização de dependências: `golang.org/x/net`, `k8s.io/client-go`, `k8s.io/api`, `k8s.io/apimachinery`, `google.golang.org/protobuf`.

**Changelog completo:** [v1.1.0](https://github.com/Tech-Preta/kubecarga/releases/tag/v1.1.0)

---

## [1.0.0] - 2024-02-27

### Added
- Primeira versão do Kubecarga, uma ferramenta de teste de carga em Go para clusters Kubernetes.
- Teste de carga contínuo: permite a criação contínua de deployments em intervalos configuráveis.
- Configuração flexível dos parâmetros de carga (número de deployments, intervalos, recursos).
- Integração total com ambientes Kubernetes, sem necessidade de ferramentas externas.

### Contributors
- @nataliagranato fez sua primeira contribuição.

**Changelog completo:** [v1.0.0](https://github.com/Tech-Preta/kubecarga/releases/tag/v1.0.0)

---

## [0.0.0] - 2025-03-25

### Added
- Esse é apenas um exemplo de descrição.
