# cgd
This is a small command-line utility that creates a hierarchy of directories meant to store game installation files. The organization is as follows:

``` text
- sorting name (eg. doom.2016)
-- release name (eg. doom-P2P)
--- dlc
--- extras
--- instructions
--- setup
--- updates
```

## Building
``` shell
git clone https://github.com/sodamouse/cgd.git
cd src
go build .
```

## Usage
``` shell
Usage of cgd:
  -a string
        [Mandatory] Specifies the root directory name
  -b string
        [Mandatory] Specifies the release directory name
  -r string
        [Optional] Specifies the base path (default "/home/soda/repo/cgd")
  -v    [Optional] Display program version information
```
