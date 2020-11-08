# LinuxHelp
## Getting started
Generate SSH Key for git
```bash
ssh-keygen -t rsa -b 4096
```

Copy the public key to git
```bash
cat ~/.ssh/id_rsa.pub
```

Install `git` if needed
```bash
sudo apt-get install git
```

Clone Repository
```bash
mkdir ~/Projects && cd ~/Projects && git clone git@github.com:jonathongardner/LinuxHelp.git
```

Run start up script
```bash
cd ~/Project/LinuxHelp && ./SetUp.sh
```

Contains a bash_profile that displays helpful git information when in a directory tracked by git.
