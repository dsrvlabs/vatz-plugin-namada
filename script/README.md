# Script for vatz & vatz-plugin-namada
## 1. install_vatz-plugin-namada.sh
- Run this script to install vatz and vatz-plugin-namada. 
```
bash install_vatz-plugin-namada.sh
```
## 2. default_example.yaml
- Replace default.yaml with this file after `vatz init`.
- You must enter hostname and webhook.
- Change the port if necessary.
```
cp default_example.yaml /root/vatz/default.yaml
```
## 3. vatz_start.sh
- Running this script will run vatz and vatz-plugin-namada. 
- The log is stored in `/var/log/vatz`.
  - You can change the log path if necessary.
- Enter valoper_address and voter_address.
- Modify namada home to suit your needs.
```
bash vatz_start.sh
```
## 4. vatz_stop.sh
- Running this script will stop both vatz and vatz-plugin-namada as they are currently running.
```
bash vatz_stop.sh
```
