# Golang 简单必应每日壁纸下载器.

## 手动编译.
```
$ export GO111MODULE=on
$ go get github.com/PuerkitoBio/goquery
$ make
$ sudo make install
```

## 用法
```
# 保存图片到 ~/Pictures
$ bing-wallpaper -o ~/Pictures

# Saves wallpaper with a custom filename
# 使用自定义文件名保存壁纸.
$ bing-wallpaper -o ~/Pictures -filename wallpaper1
```

## 版权相关.

此下载的所有壁纸都来自必应,而不是我.
图片是个人图片拍摄者的版权.
请勿用于商业用途.
