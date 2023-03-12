# Estado de mi sistema
 
### Estado de la bateria

`acpi`
`acpi -i -b`
    
**upower para saber el estado de la bateria**

`upower -i /org/freedesktop/UPower/devices/battery_BAT0`

### Monitores de recursos 

`top`
`htop`
`free -m`
   
**Otros monitores de recursos**

1. Glances
2. Bpytop
`sudo apt-get install bpytop`
3. Bashtop
  ```bash
  git clone https://github.com/aristocratos/bashtop.git
  cd bashtop
  sudo make install
```
  run: `bashtop`
   
**matar procesos**

`killall firefox`

o para forzarlo:

`killall -9 firefox`

### Translate-Shell

**Install**
```bash
git clone https://github.com/soimort/translate-shell && cd translate-shell
make
sudo make install
```
or 

`sudo apt-get install translate-shell`

**How to use Translate-Shell**

`trans [worl]`

**Translate a file**

`trans : es file//home/user/namefile.txt`

**Use interactive mode**

`trans -shell en:es thanks`

### Mecanografia

`sudo apt install TTYPER`or
`tt` or `GNU typist`
