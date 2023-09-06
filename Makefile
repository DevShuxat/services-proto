PWD=$(shell pwd)
SERVICE=eater-svc
MIGRATION_PATH=${PWD}/src/infrastructure/migrations
migration-file:
docker run -v ${MIGRATION_PATH}:/migrations migrate/migrate create -ext sql -dir /migrations -seq $(NAME)
clear:
rm -rf ${PWD}/bin/${SERVICE}
bin: clear
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ${PWD}/bin/${SERVICE} ${PWD}/main.go
chmod +x ${PWD}/bin/${SERVICE}
bin-linux: clear
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${PWD}/bin/${SERVICE} ${PWD}/main.go
chmod +x ${PWD}/bin/${SERVICE}
bin-windows: clear
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o ${PWD}/bin/${SERVICE} ${PWD}/main.go
chmod +x ${PWD}/bin/${SERVICE}