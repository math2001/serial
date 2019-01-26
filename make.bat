@echo off
prompt ^>$S
IF EXIST serial.exe DEL serial.exe
goimports -w .
rem go build -i
rem IF EXIST serial.exe serial.exe
go run serial.go COM7 9600
prompt