export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

os-archs=darwin:amd64 darwin:arm64 freebsd:386 freebsd:amd64 linux:386 linux:amd64 linux:arm linux:arm64 windows:386 windows:amd64 linux:mips64 linux:mips64le linux:mips:softfloat linux:mipsle:softfloat

all:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bing-wallpaper ./bing-wallpaper.go

app:
	@$(foreach n, $(os-archs),\
		os=$(shell echo "$(n)" | cut -d : -f 1);\
		arch=$(shell echo "$(n)" | cut -d : -f 2);\
		gomips=$(shell echo "$(n)" | cut -d : -f 3);\
		target_suffix=$${os}_$${arch};\
		echo "Build $${os}-$${arch}...";\
		env CGO_ENABLED=0 GOOS=$${os} GOARCH=$${arch} GOMIPS=$${gomips} go build -trimpath -ldflags "$(LDFLAGS)" -o ./release/bing-wallpaper_$${target_suffix} ./bing-wallpaper.go;\
		echo "Build $${os}-$${arch} done";\
	)
	@mv ./release/bing-wallpaper_windows_386 ./release/bing-wallpaper_windows_386.exe
	@mv ./release/bing-wallpaper_windows_amd64 ./release/bing-wallpaper_windows_amd64.exe

.PHONY: clean
clean:
	rm bing-wallpaper

appclean: 
	rm -rf ./release/

install:
	cp -v bing-wallpaper /usr/local/bin/
	cp -v *{service,timer} /etc/systemd/system/

reinstall:
	rm -f /usr/local/bin/bing-wallpaper
	rm -f /etc/systemd/system/*{service,timer}
