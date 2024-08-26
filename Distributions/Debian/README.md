# DEBIAN

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
