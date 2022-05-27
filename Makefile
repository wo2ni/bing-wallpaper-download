export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

all:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bing-wallpaper ./bing-wallpaper.go

install:
	cp -v bing-wallpaper /usr/local/bin/
	cp -v *{service,timer} /etc/systemd/system/

.PHONY: clean
clean:
	rm bing-wallpaper
