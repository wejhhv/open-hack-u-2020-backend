# open-hack-u-2020-backend
## 本番環境
heroku: https://open-hack-u-2020-backend.herokuapp.com/  
frontend: yet...
## 資料
Gorm: http://gorm.io/ja_JP/docs/  
echo:https://echo.labstack.com/guide  
schedule: https://docs.google.com/document/d/1ciZnMsEzIa827yUnZu9w0Gn9ODGrfo-bImdd5iO8lE0/edit  
project management: https://github.com/KazuwoKiwame12/open-hack-u-2020-backend/projects/1
## 他のREADME.md
- [APIの設計](https://github.com/KazuwoKiwame12/OpenHackU2020_backend/tree/master/Controller#api%E3%81%AE%E8%A8%AD%E8%A8%88)
- [DBの設計](https://github.com/KazuwoKiwame12/OpenHackU2020_backend/tree/master/DB#db%E3%81%AE%E8%A8%AD%E8%A8%88)
- [Serviceフォルダについて](https://github.com/KazuwoKiwame12/OpenHackU2020_backend/tree/master/Service#api%E3%81%AE%E4%B8%BB%E8%A6%81%E6%A9%9F%E8%83%BD%E3%81%AE%E5%AE%9F%E8%A3%85)
## 初期設定
- 1: .evnを.env.exampleより作成
```bash
$ cp .env.example .env
```
- 2: go.modにあるライブラリを自動インストール
```bash
$ go build
```
- 3: .envの項目を埋めDBとの接続
  
## 開発
### ファイル構成
- DB: Database関連の処理
  - Model: 各テーブルのCRUD機能
- Controller: API_URLと紐付ける機能
  - EmotionListOfPrefectures.go
  - CommentListInPrefecture.go
  - CommentDetail.go
  - UserEdit.go
  - RegisterData.go
  - DeleteData.go
- Service: Controllerで用いる主要機能の提供
### Git・Githubについて
- Issue活用
  - Issueは様々な問題や疑問、課題を共有するための機能　
  - ラベル付けや、担当者へのassignなども行なう　
  - **終了**したissueは閉じること
- プロジェクト活用
  - 自身の進めるタスクなどのスケジュールや予定などを、まとめ共有するための機能
  - 一週間ごとにプロジェクトをまとめると進捗管理がしやすい
  - Issueとの関連付けが可能である→ [参考文献](https://help.github.com/ja/github/managing-your-work-on-github/adding-issues-and-pull-requests-to-a-project-board)
- Wiki活用
  - 自分が良いと思ったことや、開発の多末になるであろう知識を共有する機能
  - 参考記事があれば、リンクを貼っておきましょう
- **ブランチ**について
  - **masterに直接Pushは禁止**です（できないようにします）
  - **新たな作業(issue)を行う時は毎回masterをfetch, pullすること**
  - ブランチは作業ごとに切り替えること(最新のmasterの状態で作業を始めるため)
  - マージされたブランチ(リモート、ローカル)は削除すること
  - commitのメッセージは内容がわかるように
  - PullRequestには**Close #番号**とコメントに含めること
  - 作業が未完成の状態のPullRequestにはコメントに**WIP**と文字を入力すること

### 注意点
1. export(外部で使用する)する関数は関数名の最初の文字を大文字にしなければならない  
2. go.modのmoduleの名前 = "project-name"がモジュールのパスとして扱われる。 結果、binの下に作成されるのは"project-name"であるため、開発時は自分の名前に変えること。**必ずその変更はcommitしないこと**
3. 自分でPullRequestをmergeしてはいけない
4. 使用しないファイルを作成していると、herokuのデプロイ時にエラーが発生する
