> EN esta guía estaremos describiendo de manera detallada, el proceso de instalación de la distribución GNU/Linux de liberación continua, Arch Linux.
> Iniciamos l LiveCD en modo ROOT (Administrador) para empezar a descargar los programas necesarios que vamos a usar y realizar configuraciones básicas.

## Configuración de teclado
**Establecer distribuciones del teclado temporal**
Ponemos nuestro teclado para Español
`root@archiso ~ # loadkeys es`
Ponemos nuestro teclado para Latin-American
`root@archiso ~ # loadkeys la-latin1`
Ponemos nuestro teclado para Americano
`root@archiso ~ # loadkeys us`

## Comprobando si su sistema es UEFI:
> Compruebe el tipo de arranque, ejecutando este comando para ver una lista con `ls`

`root@archiso ~ # ls /sys/firmware/efi/efivars`
Si hay archivos en **efivars** significa que tu sistema es **UEFI/EFI**

## Conexión a Internet por Wifi
> IS desea conectarse por wifi, es conveniente primero verificar si se ha cargado correctamente el controlador de la tarjeta wifi. Compruebe que en la salida del siguiente comando exista un fichero que inicie con la letra `wl...`
- wlan0
- wlp2s0
- Otros...

`root@archiso ~ # ls /sys/class/net`
Si no existe ningún fichero que inicie con la letra `wl`, entonces no podrá realizar la instalación via wifi.
Si por el contrario el fichero existe entonces sera posible conectarse a Internet via wifi.
Estaremos usando dos programas que vienen ya activados en el LIVE CD:
**IWD**
`root@archiso ~ # iwctl --passphrase 'password' station wlan0 connect 'NameRed'`
**NMCLI**
`root@archiso ~ # nmcli dev wifi -a connect 'NameRed' password 'password'`

## Cambiar el Idioma en LiveCD
Vamos a configurar el idioma temporal de la herramientas disponibles a nuestro idioma español, sobre todo las de particionado.
`root@archiso ~ # echo "es_ES.UTF-8 UTF-8" > /etc/locale.gen`
Ahora vamos aplicar la configuración.
`root@archiso ~ # locale-gen`
Exportamos la variable LANG para finalizar la configuración regional temporal.
`root@archiso ~ # export LANG=es_ES.UTF-8`
Comprobamos que este en español con pacman
`root@archiso ~ # pacman -Sy`

## Actualizar reloj del sistema
`root@archiso ~ # timedatectl set-ntp true`

## Tabla de partición para UEFI
Discos Duros en GPT permite mas de 128 particiones primarias

**Para usar esta tecnología depende de tu placa madre si soporta GPT y UEFI ya que es para sistemas modernos y actuales.**

Para consultar la tabla de particiones que tiene su disco duro donde va a instalar el Sistema Operativo use el siguiente comando:
`root@archiso ~ # fdisk -l`
> - Es importante saber cal es la ruta de nuestro disco duro
> - Nuestro caso es **/dev/sda: 500Gib**
> - Los resultados que terminan en [rom, loop o airoot] pueden ignorarse
> - En este caso /dev/loop0 es la imagen ISO de ArchLinux