# bcon
文字列のバイト数カウントツール

* utf-8をはじめ、sjisなどの多くのエンコーディングに対応してます（ライブラリのおかげ） 

## Used library
* [cobra](https://github.com/spf13/cobra)
* [charsetutil](https://github.com/yuin/charsetutil)

## Install
so easy. love golang :)

```
go get -u github.com/yasukotelin/bcon
```

## How to use

```
$ bcon --help
bcon is the tool to count up byte number of string

Usage:
  bcon [flags]

Flags:
  -c, --cset string   charset flag (default "utf8")
  -h, --help          help for bcon
      --version       version for bcon
```

文字列がUTF-8のときに何Byteになるか知りたいとき。デフォルトで `charset` はUTF-8になっているのでそのまま文字列を引数に指定します。

```
$ bcon あいうえお
15
```

文字列がUTF-8以外で任意のエンコーディングのときに何Byteになるか知りたい。 `--cset` または `-c` フラグでエンコーディングを指定します。

```
$ bcon -c sjis あいうえお
10
```

## Supported character sets

https://encoding.spec.whatwg.org/#names-and-labels