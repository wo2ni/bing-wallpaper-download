# Golang 简单必应每日壁纸下载器.

## Arch Linux

### 从AUR安装.

```
yay -Syy bing-wallpaper-download        #国际用户.
yay -Syy bing-wallpaper-download-cn     #大陆用户.
```

* **手工构建.**

```
pacman -Syy git
git clone https://github.com/wo2ni/bing-wallpaper-download.git
makepkg -s
```

## 手动编译安装(适用与Debian,Fedora,RHEL系列发行版等).

* **大陆内会将www.bing.com解析为:https//s.cn.bing.net**

### 编译国际版本.

```
$ make app      #构建不同架构版本.
$ make          #仅构建Linux版本.
$ sudo make install
```

### 编译大陆版本.
```
sed -i 's/return fmt.Sprintf("%s%s", bingURL, link), nil/return fmt.Sprintf("%s", link), nil/g' bing-wallpaper.go
$ make app      #构建不同架构版本.
$ make          #仅构建Linux版本.
$ sudo make install
```


## 用法
```
# 查看使用方法.
$ bing-wallpaper -h

# 保存图片到 ~/Pictures
$ bing-wallpaper -o ~/Pictures/

# 使用自定义文件名保存壁纸.
$ bing-wallpaper -o ~/Pictures/ -filename wallpaper1
```

## 版权相关.

此下载的所有壁纸都来自必应,而不是我.
图片是个人图片拍摄者的版权.
请勿用于商业用途.
