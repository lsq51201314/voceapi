@echo off
cd %cd%
echo ��������ͼ�� ������
echo IDI_ICON1 ICON "main.ico" > main.rc
windres -o main.syso main.rc
echo ���ڱ���Windowsƽ̨ ������
go build -ldflags "-s -w" -o voceapi.exe
echo ���ڱ���Linuxƽ̨ ������
SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=linux
go build -ldflags "-s -w" -o voceapi
echo ������� ������
pause