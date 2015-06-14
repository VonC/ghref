@echo off
setlocal
set GOPATH=%~dp0/deps
set GOBIN=%~dp0/bin
cd %~dp0
rem cd
rem set GO
%GOROOT%\bin\go.exe install
set PATH=.;%PATH%
git updver
endlocal
doskey ghref=%~dp0\bin\ghref.exe $*
