# init-order
golangのinit()の挙動のテスト。結論は次の通り。
- init()はpackageが初めてimportされるときに実行される
- packageが他のpackageをimportしている場合、それらすべてのpackageのimport完了後にinit()が実行される

## 実行結果
```
$ go run main.go
a.init
c.init
b.init
d.init (1)
d.init (2)
d.init (3)
main.init
main.main
```
