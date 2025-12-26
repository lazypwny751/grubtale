# Grubtale

<img width="794" height="599" alt="Screenshot from 2025-12-26 00-43-23" src="https://github.com/user-attachments/assets/e7e32b28-a728-4475-a44e-3078d1edc5a6" />

An Undertale-inspired GRUB theme that fills your boot process with **DETERMINATION**.

# Installation

## Simple steps to get Grubtale up and running on your system.
```sh
git clone "https://github.com/lazypwny751/grubtale.git" && cd "grubtale"
make install
sudo ./grubtale
```

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
