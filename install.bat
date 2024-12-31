@echo off
SET BINARY_NAME=urlactions.exe
SET DESTINATION=C:\Windows\System32

REM Ensure the binary exists in the current directory
IF NOT EXIST "%BINARY_NAME%" (
    ECHO Error: %BINARY_NAME% not found in the current directory.
    EXIT /B 1
)

REM Copy the binary to the destination
ECHO Installing %BINARY_NAME% to %DESTINATION%...
COPY %BINARY_NAME% %DESTINATION%

IF %ERRORLEVEL% EQU 0 (
    ECHO Installation successful!
    ECHO You can now run the tool using: urlactions
) ELSE (
    ECHO Installation failed.
    EXIT /B 1
)
