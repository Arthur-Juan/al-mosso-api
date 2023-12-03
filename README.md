# al-mosso-api
API do projeto Al-Mosso. Aplicações para a internet

## Descrição do Projeto

Este projeto consiste em uma API desenvolvida em Golang para facilitar o processo de marcação de reservas em um restaurante. A aplicação utiliza PostgreSQL como banco de dados, Docker e docker-compose para facilitar a implantação e escalabilidade, testes automatizados para garantir a qualidade do código e envio de e-mails para confirmação das reservas.

## Tecnologias Utilizadas

* Linguagem Golang: Desenvolvido principalmente em Golang, uma linguagem eficiente e de fácil manutenção.

* Banco de Dados PostgreSQL: Utiliza o PostgreSQL como sistema de gerenciamento de banco de dados.

* Docker e docker-compose: Facilita a implantação e o gerenciamento da aplicação em diferentes ambientes.

* Testes Automatizados: Utiliza frameworks e bibliotecas de teste em Golang para garantir a qualidade do código.

* Envio de E-mail: Integração com um serviço de envio de e-mails para notificação.


## Arquitetura 

O projeto adota uma arquitetura modularizada para garantir a escalabilidade, a manutenibilidade e a clareza do código. A estrutura do diretório é organizada da seguinte maneira:

* cmd/api/main.go -> entrypoint do projeto
* config/config.go -> reúne as configurações do projeto
* config/pg.go -> configurações referentes ao banco de dados Postgres
* internal/* -> funcionalidades do projeto
* internal/handlers/* -> contém os handlers HTTP que recebem os inputs do usuário e retornam o output
* internal/services/* -> contém a regra de negócio de cada funcionalidade
* internal/entity/* -> contém as entidades do sistema e suas lógicas
* internal/routes/ * -> contém as configurações das rotas e middlaweres, seguindo o pradrão rest
* internal/interfaces/* -> contém abstrações de implementações que podem ser mudadas futuramente
* pkg/* -> contém implementações de dependências do projeto
* pkg/token/* -> implementa a funcionalidade de token de autenticação, no caso, JWT
* pkg/database/* -> implementa a comunicação com o banco de dados, migrations, seeders, e schemas para se comunicar com a base
* pkg/fileHandler/* -> implementa gerenciamento de arquivos, implementando validações de segurança utilizando assinatura das images
* pkg/cryptography/* -> implementa algumas questões de criptografia como validação e geração de hashes
* pkg/emailPkg/* -> implementa integração com serviço de email 




