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
2. Bashtop
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
