# go-backend

Para executar a migração para atualizar a base de dados, basta seguir esse passo a passo:
1 - Criar .env em .devcontainer e definir os valores:
DB_URL=
PGHOST=
PGPORT=
PGUSER=
PGPASSWORD=
PGDATABASE=

2 - Executar o comando:
`migrate -path migrations -database $DB_URL up`
