# Simple Bing Wallpaper Downloader in Go

Learning Go, so decided to make something that could download nice wallpaper for me.

Set it as a cron job, and immerse yourself with new backgrounds.

## Compiling this stuff
```
export GO111MODULE=on
go get github.com/PuerkitoBio/goquery
$ make
```

## Usage
```
# Saves current wallpaper to ~/Pictures
$ bing-wallpaper -output-dir ~/Pictures

# Saves wallpaper with a custom filename
$ bing-wallpaper -output-dir ~/Pictures -filename wallpaper1
```

## Copyright stuff

All the wallpaper downloaded by this is from Bing, not me.

Pictures are copyrights of the individual picture taker.

Please don't come after me D:
