## GraphQL Client tool

## 実行例

[graphqurl](https://github.com/hasura/graphqurl)という curl likeなツールをインストール
$ npm install -g graphqurl

$ gq http://localhost:8080/graphql -q '{user(id: "1"){userName}}'

curl -i -X POST -d '{query: User {user(id: "1"){userName}}}' http://localhost:8080/graphql

```bash
# キー指定
curl -i -X POST -d '{query: user(id: "1"){userName}}' http://localhost:8080/graphql

# 全件取得
curl -i -X POST -d '{query: userList{userName}}' http://localhost:8080/graphql
```