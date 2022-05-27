export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

all:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bing-wallpaper ./bing-wallpaper.go


.PHONY: clean
clean:
	rm bing-wallpaper

install:
	cp -v bing-wallpaper /usr/local/bin/
	cp -v *{service,timer} /etc/systemd/system/

reinstall:
	rm -f /usr/local/bin/bing-wallpaper
	rm -f /etc/systemd/system/*{service,timer}
