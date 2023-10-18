
実行
go run ./cmd/


### gomockの生成
gomock
https://github.com/uber-go/mock
go install go.uber.org/mock/mockgen@latest
mockの作成はそれぞれのフォルダで実行
package名とフォルダ名をあわせるために、隣のフォルダに設定
mockgen -source=./repository/task.go -destination=./repository_mock/task.mock.go
