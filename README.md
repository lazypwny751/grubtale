# Grubtale

An Undertale-inspired GRUB theme that fills your boot process with **DETERMINATION**.

# Installation

## Compile the engine.
```
git clone "https://github.com/lazypwny751/grubtale.git" cd "grubtale"
make all
sudo ./grubtale
```

After that you can modify your "grub" file which at "/etc/default/grub", find(if doesn't exists you can create) and update the variable GRUB_THEME= like

```sh
GRUB_THEME="/boot/grub/themes/Grubtale/theme.txt"
# if you're using grub2
GRUB_THEME="/boot/grub2/themes/Grubtale/theme.txt"
``` 

save and exit, now you can type `sudo update-grub` but if you use arch linux derivates type (`grub-mkconfig -o /boot/grub/grub.cfg` or `grub2-mkconfig`)

## Requirements

- Go

# Contact

- **Discord**: lazypwny751
- **Twitter**: Ahmetta02120401s

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---


*"Despite everything, it's still GNU/Linux."*

**Stay determined!**
