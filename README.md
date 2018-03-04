# apig-sonar-share
APIGのソナー結果をシェアするサーバー  
[DEMO](http://35.185.8.180/v1/apig/coins)  
[Swagger](http://35.185.8.180/swagger/)  

# 使い方

beegoとbeeを取得しておく必要があります。

````
go get github.com/astaxie/beego
go get github.com/astaxie/bee
````

あとは、このリポジトリを「$GOPATH/src/github.com/scrpgil/apig-sonar-share」にcloneし  
bee runコマンドを打てば起動します。  

````
git clone https://github.com/scrpgil/apig-sonar-share.git
~~cloneしたフォルダに移動~~
bee run
````

conf/app.confにAPI_KEYを記述する箇所があるのでそこに自分のAPI_KEYを入力してください。

````
[api]
key=ここに書く
````

デフォルトはポート80で立ち上がります。
