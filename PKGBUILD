# Maintainer: Raafat Turki <raafat.turki@pm.me>

pkgname=shortsig
pkgver=0.1
pkgrel=1
pkgdesc="remote routine execution through TCP"
arch=('x86_64')
license=('GPL')
source=("git+https://github.com/RaafatTurki/shortsig.git")
md5sums=('SKIP')

build() {
  cd "$srcdir/$pkgname"
  go build
}

package() {
  cd "$srcdir/$pkgname"

  mkdir -p $pkgdir/usr/bin
  cp $pkgname $pkgdir/usr/bin

  mkdir -p $pkgdir/usr/local/man/man1
  cp docs/${pkgname}.1 $pkgdir/usr/local/man/man1
}
