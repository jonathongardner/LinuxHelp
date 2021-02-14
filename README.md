# LinuxHelp
Easily setup a Linux how I like it.

## Getting started
Install `git` if needed
```bash
sudo apt-get install git
```

Clone Repository
```bash
mkdir ~/Projects && cd ~/Projects && git clone https://github.com/jonathongardner/LinuxHelp.git
```

Generate SSH Key for git
```bash
ssh-keygen -t rsa -b 4096
```

Copy the public key to git
```bash
cd ~/Project/LinuxHelp/script && ./push-git-public-key
```

Change to use ssh in LinuxHelp repository
```
git remote set-url origin git@github.com:jonathongardner/LinuxHelp.git
```

Run start up script to install/export bash_profile, vimrc, and gitconfig.
```bash
cd ~/Project/LinuxHelp && ./SetUp.sh
```
