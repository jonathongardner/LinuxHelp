#### Merge pdfs
```
pdftk file1.pdf file2.pdf cat ouput output-file.pdf
```

#### Simple math
```
echo $((1+1))
```

#### Mount an image
```
fdisk -l image.img
offsetVar=$(($1*512))
mount -o loop,offset=$offsetVar image.img /mnt/dir
```

#### Send notification over ssh
```
ssh user@ip 'DISPLAY=:0 notify-send "'$2'"'
```

#### Copy screenshot over ssh
```
ssh user@ip 'DISPLAY=:0 import -window root Desktop/screenshot.png'
scp user@ip:/home/$2/Desktop/screenshot.png .
ssh user@ip 'rm Desktop/screenshot.png'
```
