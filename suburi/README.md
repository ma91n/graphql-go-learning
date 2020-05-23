## GraphQL Client tool

## 実行例

```bash
# キー指定
curl -i -X POST -d '{query: user(id: "195360fb-c3aa-4b68-93dd-a8185f55fd6b"){userName}}' http://localhost:8080/graphql

# 全件取得
curl -i -X POST -d '{query: userList{userName,email}}' http://localhost:8080/graphql

# 登録
curl -i -X POST -d 'mutation{createUser(userName:"mano", description: "des", photoURL: "photo", email: "email"){userId, userName, description, photoURL, email}}' http://localhost:8080/graphql
```
