# 🖥️ Bambah Desktop

> Aplicação desktop multiplataforma para controle financeiro pessoal

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Fyne](https://img.shields.io/badge/GUI-Fyne-blue)](https://fyne.io/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## 📋 Sobre

**Bambah Desktop** é uma aplicação gráfica multiplataforma desenvolvida em Go usando o framework Fyne. Oferece uma interface intuitiva para o gerenciamento de finanças pessoais, permitindo o registro de receitas e despesas de forma rápida e organizada.

## ✨ Funcionalidades

- 💰 **Interface gráfica intuitiva** para registro de receitas e despesas
- 🏦 **Seleção dinâmica de contas** carregadas do backend
- 🏷️ **Categorização automática** com listas dinâmicas
- 📅 **Controle de datas** com seletor integrado
- ✅ **Validação em tempo real** de campos obrigatórios
- 🔄 **Sincronização automática** com banco de dados via SDK
- 🖱️ **Interface responsiva** com formulários limpos

## 🚀 Instalação

### Pré-requisitos

- Go 1.24 ou superior
- Conta no [Appwrite](https://appwrite.io)
- Dependências do sistema para Fyne (varie por SO)

#### Dependências do Sistema

**Linux (Ubuntu/Debian):**
```bash
sudo apt-get install gcc pkg-config libgl1-mesa-dev xorg-dev
```

**Linux (Fedora/CentOS):**
```bash
sudo dnf install gcc pkg-config mesa-libGL-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel libXxf86vm-devel
```

**macOS:**
```bash
# Xcode command line tools
xcode-select --install
```

**Windows:**
```bash
# Requer TDM-GCC ou similares
# Ou use Go com CGO_ENABLED=1
```

### Configuração

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/marcelocg/bambah-desktop.git
   cd bambah-desktop
   ```

2. **Instale as dependências:**
   ```bash
   go mod tidy
   ```

3. **Configure as variáveis de ambiente:**
   ```bash
   export APPWRITE_ENDPOINT=https://cloud.appwrite.io/v1
   export APPWRITE_PROJECT_ID=seu-project-id
   export APPWRITE_DATABASE_ID=seu-database-id
   export APPWRITE_API_KEY=sua-api-key
   export APPWRITE_COLLECTION_ID=entries  # Opcional, padrão: "entries"
   ```

4. **Compile e execute:**
   ```bash
   go run cmd/main.go
   ```

### Build para Distribuição

```bash
# Build local
go build -o bambah-desktop cmd/main.go

# Build cross-platform
GOOS=windows GOARCH=amd64 go build -o bambah-desktop.exe cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o bambah-desktop-mac cmd/main.go
GOOS=linux GOARCH=amd64 go build -o bambah-desktop-linux cmd/main.go
```

## 💡 Como Usar

### Interface Principal

A aplicação possui uma interface simples e intuitiva:

1. **Tipo de Lançamento**: Escolha entre "Receita" ou "Despesa"
2. **Conta**: Selecione a conta onde será registrada a transação
3. **Categoria**: Escolha a categoria apropriada
4. **Valor**: Digite o valor da transação (sempre positivo)
5. **Data**: Selecione a data (padrão: hoje)
6. **Descrição**: Adicione uma descrição opcional

### Funcionalidades da Interface

- **Validação em Tempo Real**: Campos obrigatórios são validados automaticamente
- **Carregamento Dinâmico**: Contas e categorias são carregadas do backend
- **Ajuste Automático de Sinal**: Despesas ficam negativas automaticamente
- **Limpeza Rápida**: Botão para limpar todos os campos
- **Feedback Visual**: Mensagens de sucesso e erro

## 🏗️ Arquitetura

### Estrutura do Projeto

```
bambah-desktop/
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── internal/
│   ├── ui/
│   │   ├── app.go             # Janela principal e configuração
│   │   └── forms/
│   │       └── entry_form.go   # Formulário de lançamentos
│   └── config/                 # Configurações (futuro)
├── go.mod                      # Dependências do módulo
├── go.sum                      # Checksums das dependências
└── README.md                   # Documentação do projeto
```

### Integração com SDK

A aplicação utiliza o [Bambah SDK](https://github.com/marcelocg/bambah-sdk) para:

- Abstração completa do backend
- Operações CRUD de lançamentos financeiros
- Carregamento dinâmico de contas e categorias
- Configuração automática via variáveis de ambiente

```go
// Criação do serviço (automática via SDK)
service, err := bambah.NewFinancialServiceFromEnv()

// Uso das funcionalidades
accounts, _ := service.ListAccounts()
categories, _ := service.ListCategories()
service.CreateFinancialEntry(entry)
```

## 🛠️ Tecnologias Utilizadas

- **[Go](https://golang.org)** - Linguagem de programação
- **[Fyne](https://fyne.io/)** - Framework GUI multiplataforma
- **[Bambah SDK](https://github.com/marcelocg/bambah-sdk)** - Abstração de backend
- **[Appwrite](https://appwrite.io)** - Backend-as-a-Service

## 🔧 Desenvolvimento

### Comandos de Desenvolvimento

```bash
# Executar aplicação
go run cmd/main.go

# Executar testes
go test ./...

# Formatar código
go fmt ./...

# Análise estática
go vet ./...

# Build de desenvolvimento
go build -o bambah-desktop cmd/main.go
```

### Estrutura de Desenvolvimento

Para contribuir com o projeto:

1. Fork o repositório
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Implemente suas mudanças
4. Execute testes e validações
5. Commit suas mudanças (`git commit -m 'Adiciona nova feature'`)
6. Push para a branch (`git push origin feature/nova-feature`)
7. Abra um Pull Request

## 🤝 Projetos Relacionados

- [**bambah-sdk**](https://github.com/marcelocg/bambah-sdk) - SDK comum para o ecossistema Bambah
- [**bambah-cli**](https://github.com/marcelocg/bambah-cli) - Interface de linha de comando

## 📝 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 📞 Suporte

Se você encontrar algum problema ou tiver sugestões:

- Abra uma [issue](https://github.com/marcelocg/bambah-desktop/issues)
- Entre em contato através do email

---

<div align="center">
  <sub>Feito com ❤️ em Go e Fyne</sub>
</div>