@echo off
go build -ldflags "-w -s" -o release/http/http.exe main.go

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-w -s" -o release/http/http32.exe main.go


@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/http/http_amd64 main.go


@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -ldflags "-w -s" -o release/http/http_arm64 main.go

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/down/down.exe download.go

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-w -s" -o release/down/down32.exe download.go


@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/down/down_amd64 download.go


@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -ldflags "-w -s" -o release/down/down_arm64 download.go

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/screen/Screen.exe Screenshot.go

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-w -s" -o release/screen/Screen32.exe Screenshot.go

@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/screen/Savejpg_amd64 Savejpg.go