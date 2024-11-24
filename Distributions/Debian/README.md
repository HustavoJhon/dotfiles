# DEBIAN

![Debian](https://img.shields.io/badge/Debian-D70A53?style=for-the-badge&logo=debian&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
![Shell Script](https://img.shields.io/badge/shell_script-%23121011.svg?style=for-the-badge&logo=gnu-bash&logoColor=white)
![Neovim](https://img.shields.io/badge/NeoVim-%2357A143.svg?&style=for-the-badge&logo=neovim&logoColor=white)

```bash
# directory user
sudo apt update
sudo apt install xdg-user-dirs
xdg-user-dirs-update
```

```bash
# wifi
ls /sys/class/net

sudo nano /etc/wpa_supplicant/wpa_supplicant.conf
```

```conf
network={
	ssid="name_wifi"
	psk="password"
}
```

```bash
sudo wpa_supplicant -B -i wlp1s0 -c /etc/wpa_supplicant/wpa_supplicant.conf

sudo dhclient wlp1s0
````

```zsh
sudo apt update
sudo apt install python3
sudo apt install python3.11-venv
```

**Warp**

> eliminar mensaje de introduccion
```zsh
sudo rm /etc/motd
```

**ASCII**
[shadown](https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=Type%20Something%20)

**Programs**

```bash
sudo apt update
sudo apt install snapd

sudo snap install obsidian --classic
```

```bash
sudo apt install cmus

mkdir -p ~/.cmus
wget https://raw.githubusercontent.com/egel/cmus-gruvbox/master/gruvbox-dark.theme
# Then open Cmus and set new theme :colorscheme gruvbox-dark.
```

**RSA SSH keys**

```bash
ssh-keygen -t rsa -b 2048 -C "user@email"
```

configure SSH to point a different directory

config file
```bash
# Github.com
Host github.com
    HostName github.com
    IdentifyFile ~/.ssh/id_rsa_github

# Azure.com
Host ssh.dev.azure.com
    HostName ssh.dev.azure.com
    IdentifyFile ~/.ssh/id_rsa_azure
```
