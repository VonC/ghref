@echo off
setlocal
set GOPATH=%~dp0/deps
set GOBIN=%~dp0/bin
cd %~dp0
rem cd
rem set GO
%GOROOT%\bin\go.exe install
set PATH=.;%PATH%
set a=%1
if NOT "%a%" == "" (
	if NOT "%a%" == "bump" (
		git tag -m "%a% release" %a%
		set a=bump
	)
)
if "%a%" == "bump" (
  git updver
)
endlocal
doskey ghref=%~dp0\bin\ghref.exe $*
