echo off
set PWD=%~dp0
%1 start "" mshta vbscript:CreateObject("Shell.Application").ShellExecute("cmd.exe","/c %~s0 ::","","runas",1)(window.close)&&exit
%PWD%\k8sproxy.exe -u %BASE_URL%
pause
echo on
