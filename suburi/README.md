## GraphQL Client tool

## 実行例

```bash
# 登録(ユーザ)
curl -i -X POST -d 'mutation{createUser(userName:"mano", description: "des", photoURL: "photo", email: "email"){userId, userName, description, photoURL, email}}' http://localhost:8080/graphql

# 登録(イベント)
curl -i -X POST -d 'mutation{createEvent(userId:"1", eventName:"event", description:"des", location:"loc", startTime: "start", endTime: "end"){eventId, eventName}}' http://localhost:8080/graphql

# キー指定
curl -i -X POST -d '{query: user(id: "1"){userName}}' http://localhost:8080/graphql

# 全件取得
curl -i -X POST -d '{query: userList{userName,email}}' http://localhost:8080/graphql
```
