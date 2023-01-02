# API仕様
[go-swagger](https://goswagger.io/install.html)を使ったswagger.jsonの生成
```
$ pwd
任意のディレクトリ/recipe-api
$ docker run --rm -it --user $(id -u):$(id -g) -e XDG_CACHE_HOME=/tmp/.cache -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger generate spec -o ./swagger.json
```
# 参考文献
- [ginを最速でマスターしよう](https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a)
- [[Golang/Go言語]GinフレームワークでAPI開発1~Hello アプリ~](https://selfnote.work/20211115/programming/golang-gin-api-1/)
- [[Golang/Go言語]GinフレームワークでAPI開発2 CRUDの実装](https://selfnote.work/20211120/programming/golang-gin-api-2/)
- [[Golang/Go言語]GinフレームワークでAPI開発3 OpenAPIの導入](https://selfnote.work/20211125/programming/golang-gin-api-3/)