Go Twitter utility tools
========================
Twitter API を利用したりスクレイピングするツールです。

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
