# set PowerShell to UTF-8
[console]::InputEncoding = [console]::OutputEncoding = New-Object System.Text.UTF8Encoding


# PROMOT
Import-Module posh-git
Import-Module oh-my-posh
$omp_config = Join-Path $PSScriptRoot "./hustavo.omp.json"
oh-my-posh init pwsh --config $omp_config | Invoke-Expression
# Set-PoshPrompt name_theme


# Icons
Import-Module -Name Terminal-Icons


# PSReadLine
Set-PsReadLineOption -EditMode Emacs
Set-PsReadLineOption -BellStyle None
Set-PsReadLineKeyHandler -Chord 'Ctrl+d' -Function DeleteChar
Set-PsReadLineOption -PredictionSource History


# FzF
Import-Module PSFzf
Set-PSFzfOption -PSReadLineChordProvider 'Ctrl+f' -PSReadLineChordReverseHistory 'Ctrl+r'

# Env
$env:GIT_SSH = "C:\Windows\system32\OpenSSH\ssh.exe"

# Set-Alias 
Set-Alias touch ni
Set-Alias home sl
Set-Alias ll ls
Set-Alias l ls
Set-Alias x clear
Set-Alias gg "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Google Chrome.lnk"
Set-Alias brave "C:ProgramData\Microsoft\Wondows\Start Menu\Programs\Brave.lnk"


#Utilities
function which ($command)
{
    Get-Command -Name $command -ErrorAction SilentlyContinue |
        Select-Object -ExpandProperty Path -ErrorAction SilentlyContinue
}
