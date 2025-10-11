# Help
## Fix yubikey Fedora
```
gpg-agent --daemon --enable-ssh-support
systemctl restart pcscd
```