# sikalabsx/master

This project contains the binaries of the `master` branch of the SikaLabs tools [slu](https://github.com/sikalabs/slu), [slr](https://github.com/sikalabs/slr), and [tergum](https://github.com/sikalabs/tergum). It's intended to be used for quick installation of the latest development version without need of the release.

## Install on Mac

```bash
brew install sikalabsx/tap/master
```

## Install on Linux

```bash
slu ib master
slu ib master_slu
slu ib master_slr
slu ib master_tergum
```

## Install as Original Binaries on Linux

```bash
slu ib master_slu && mv /usr/local/bin/master_slu /usr/local/bin/slu
slu ib master_slr && mv /usr/local/bin/master_slr /usr/local/bin/slr
slu ib master_tergum && mv /usr/local/bin/master_tergum /usr/local/bin/tergum
```

## Build Locally

```bash
make build-all
```

## Update Dependencies

```bash
make update-deps
```
