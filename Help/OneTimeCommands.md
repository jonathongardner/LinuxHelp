# One Time Commands...

#### Grep
```BASH
grep -rIi --color=always --exclude=*.{file1,...} || --include=*.{file1,...}
```
 - I -> ignore binary files
 - i -> case insensitive

## Git
#### Find all commits that relate to a file
```BASH
git log --full-history -- SOMEFILE
```
#### Find commit hashes
```BASH
git log --reverse --ancestry-path GITHASH^..master
```
## Linux
#### Hide files in in folder
```BASH
echo 'file-to-hide' >> .hidden
```

### Ubuntu
#### Move Window Buttons to left side
```BASH
gsettings set org.gnome.desktop.wm.preferences button-layout 'close,minimize,maximize:'
# Back
gsettings set org.gnome.desktop.wm.preferences button-layout ':close,maximize,minimize'
```
#### Move Keep monitors together when changing workspaces
```BASH
gsettings set org.gnome.mutter workspaces-only-on-primary false
```

# MAC
### mojave
[hidutil](https://developer.apple.com/library/archive/technotes/tn2450/_index.html)
#### Get key map for specific keyboard
```BASH
hidutil property -m '{"ProductID":0xB34C,"VendorID":0x046D}' --get "UserKeyMapping"
```
Info can be found in System Report
#### Remap key for specific keyboard (Right Alt to Right Ctrl)
```BASH
hidutil property -m '{"ProductID":0xB34C,"VendorID":0x046D}' --set '{"UserKeyMapping": [{"HIDKeyboardModifierMappingSrc":0x7000000e6,  "HIDKeyboardModifierMappingDst":0x7000000e4}] }'
```
