Maintainer="baris-inandi"
pkgname=barley
pkgver=r32.5e864b2
pkgrel=1
epoch=
pkgdesc="AUR helper with a familiar subcommand system"
arch=(x86_64)
url="https://github.com/baris-inandi/fe"
license=('GPLv3')
groups=()
depends=(bash sudo fzf)
makedepends=(git go)
checkdepends=()
optdepends=()
provides=(barley)
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=("git+$url")
noextract=()
md5sums=('SKIP')
validpgpkeys=()

pkgver() {
	cd $pkgname
	printf "${_pkgver}r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
	cd $pkgname
	make install
}

package() {
	cd $pkgname
	make install
}
