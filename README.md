# getdl

gawe golek url download

# Current Support

1. Oploverz
2. Samehadaku
3. Doronime
4. Lendrive

# install

```
$ make install
```

# uninstall

```
$ make uninstall
```

# example configuration

```
code: "265"
resolution: "720"
file_hosting: "zippy"
browser: "xdg-open"
open_in_browser: true
domain:
  doronime: "doronime.id"
  samehadaku: "65.108.132.145"
  lendrive: "lendrive.web.id"
  oploverz: "oploverz.co.in"
```

# todo

- [x] scrape data (judul, deskripsi, link download)
- [x] simple command
- [x] buat config format json
- [ ] bypass shortlink
- [ ] klok bisa async sih
- [x] tambah cobra biar enak management commandnya
- [ ] get direct file download
- [ ] terintegrasi dengan aria2c
- [ ] buat konfig per website
