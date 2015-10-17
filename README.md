Go Twitter utility tools
========================

[![Build Status][travis-img]][travis-url]
[![Coverage Status][coveralls-img]][coveralls-url]

Twitter API を利用したりスクレイピングするコマンドラインツールです。

Usage
-----
```
NAME:
   got - Go Twitter utility tools

USAGE:
   got [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   fetch-images, fi
   help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version
```

### fetch-image
screen name を指定してユーザの画像をダウンロードできます。

```
NAME:
   got fetch-images - fetch user images

USAGE:
   got fetch-images [command options] <screen_name>

OPTIONS:
   --size "10"  size of fetch image
   --lastid "0" fetch start tweet ID
   --orig       set original image flag
   --download   auto download images
```

Download
--------
[Releases](https://github.com/ww24/got/releases)

* `all-in-one.tar.gz` と `all-in-one.zip` には全てのプラットフォーム向けの実行ファイルが含まれています。
* `got-` から始まるファイルは、各プラットフォーム向けの実行ファイルです。環境に合わせてダウンロードして下さい。

LICENSE
-------
[MIT](LICENSE)

[travis-url]: https://travis-ci.org/ww24/got
[travis-img]: https://api.travis-ci.org/ww24/got.svg
[coveralls-url]: https://coveralls.io/github/ww24/got?branch=master
[coveralls-img]: https://coveralls.io/repos/ww24/got/badge.svg?branch=master&service=github
