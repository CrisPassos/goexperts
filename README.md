- comandos docker

docker compose up -d
docker compose exec mysql bash
mysql -u root -p
USE goexpert

drop table produtos

- comandos go
  go mod init

  go test .
  go test -v
  go test -coverprofile=coverage.out
  go tool cover -html=coverage.out
  go test -bench=.
  go test -fuzz=.
