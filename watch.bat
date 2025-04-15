@echo off
:loop
cls
go build -o go-smartcerti.exe .
if %ERRORLEVEL% NEQ 0 (
    echo Build failed.
) else (
    echo Running app...
    go-smartcerti.exe
)
timeout /t 2 >nul
goto loop
