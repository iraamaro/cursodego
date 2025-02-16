# go build

`go build`

Mudar nome do executável

`go build -o meuapp`

Executável para Windows

`GOOS=windows GOARCH=386 go build -o meuapp.exe`

Executável para Linux ARM com verbose

`GOOS=linux GOARCH=arm go build -o meuapp-rasp -v`