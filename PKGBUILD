# Maintainer: BadBoy <luckmelove2@gmail.com>

pkgname='bing-wallpaper-download'
_pkgname='bing-wallpaper'
pkgver=0.1
pkgrel=1
pkgdesc="Bing wallpaper Download."
arch=('x86_64' 'i686' 'arm' 'armv6h' 'armv7h' 'aarch64')
license=('Apache')
url='https://github.com/wo2ni/bing-wallpaper-download'
depends=('glibc' 'feh')
makedepends=('go')
install="${pkgname}.install"

source=("git+${url}.git")

sha256sums=('SKIP')

build() {
    cd "${srcdir}/${pkgname}"
    make -j $(nproc)
}

package() {
    msg "Install ${pkgname}"
    install -dm755 "${pkgdir}/usr/local/bin/"
    install -Dm744 "${srcdir}/${pkgname}/${_pkgname}" "${pkgdir}/usr/local/bin"

    msg "Install ${pkgname} systemd timer and service"
    install -dm755 "${pkgdir}/etc/systemd/system/"
    for i in "${srcdir}/${pkgname}"/*{timer,service}; do
        install -Dm644 "${i}" "${pkgdir}/etc/systemd/system/"
    done
}
