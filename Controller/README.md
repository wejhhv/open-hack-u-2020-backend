## APIの設計
### 感情が地図に載る(感情の共有)
**各県の色16進数を返す**  
- APIのURL: `/emotion/prefectures`  
- Method: GET
- フロントエンドから受け取るデータ  
```
None
```
- フロントエンドへと渡すデータ
```
[
  {
    Prefecture: string,
    Color: string
  },
  ....
]
```

**コメントの一覧表示 or（感情の場所を指すピン）を各県ごとに分けて渡す(選択された県のみ渡す方が良い)**
- APIのURL: `/emotion/{prefecture}/comments`
- Method: GET
- フロントエンドから受け取るデータ
```
prefecture: string  
```
- フロントエンドへと渡すデータ
```
//成功
[
    {
      ID: int,
      EmotionID: int,
      Latitude: double,
      Longtitude: double,
      UserName: string,
      DateTime: dateTime,
    },
    .....
]

//失敗...prefectureの部分が英語大文字の県名になっていない
[]
```
**ピンを押して、コメント内容確認(紐ついた返信も)**
- APIのURL: `/emotion/comments/{comment_id}` 
- Method: GET   
- フロントエンドから受け取るデータ  
```
prefecture: string  
comment_id: int
```
- フロントエンドへと渡すデータ
```
//成功
{
  CommentContent: {
    Comment: string,
    EmotionID: int,
    UserName: string
    DateTime: dateTime,
  }
  Response: [
      {
        ID: int,
        UserName: string,
        Comment: string,
        DateTime: dateTime,
      },
      ......
  ]
}

//失敗...存在しないID or 県名を指定したとき
{
    "CommentContent": {
        "Comment": "",
        "EmotionID": 0,
        "UserName": "",//これでエラー判別
        "DateTime": "0001-01-01T00:00:00Z"
    },
    "Response": []
}
```
**自分のコメント取得(paginate)**
- APIのURL: `/emotion/mycomments/:user_id/:page_num/:page_size`
- Method: GET
- フロントエンドから受け取るデータ
```
user_id: string  
page_num: string
page_size: string
```
- フロントエンドへと渡すデータ
```
//成功
[
    {
      ID: int,
      EmotionID: int,
      Latitude: double,
      Longtitude: double,
      UserName: string,
      DateTime: dateTime,
    },
    .....
]

//失敗...prefectureの部分が英語大文字の県名になっていない
[]
```
### ユーザ情報登録
**ユーザ情報登録(名前が重複しないようにする)**  
- APIのURL: `/user/register`
- Method: POST
- フロントエンドから受け取るデータ    
```
name: string  
```
- フロントエンドへと渡すデータ
```
//成功
{
  ID: int,
  Name: string,
  Point: int
}

//失敗...名前が重複したとき
{
    "ID": 0, ***これでフロントはエラー判定ができる****
    "Name": string,
    "Point": 0
}
```
**ユーザ情報編集** 
- APIのURL: `/user/edit`
- Method: PATCH
- フロントエンドから受け取るデータ  
```
id: int
newName: string
```
- フロントエンドへと渡すデータ
```
//成功
{
  ID: int,
  Name: string,
  Point: int
}

//失敗...名前が重複したとき
{
    "ID": 0, ***これでフロントはエラー判定ができる****
    "Name": "",
    "Point": 0
}
```

### 感情の登録など
**コメントの登録**
- APIのURL: `/comment/register`
- Method: POST
- フロントエンドから受け取るデータ
```
[    
  {      
    user_id: int,
    emotion_id: int,   
    comment: string,   
    latitude: double,      
    longtitude: double,      
    prefecture: string,      
  }
]
```
- フロントエンドへと渡すデータ
```
{
  hasSuccess: boolean
}
```

**コメントに対する返信の保存**
- APIのURL: `/comment/response/register`
- Method: POST
- フロントエンドから受け取るデータ
```
[    
  {      
    user_id: int,
    comment_id: int,   
    comment: string,        
  }
]
```
- フロントエンドへと渡すデータ
```
{
  hasSuccess: boolean
}
```
**コメント削除**
- APIのURL: `/comment/delete`
- Method: DELETE
- フロントエンドから受け取るデータ
```
comment_id: int
```
- フロントエンドへと渡すデータ
```
{
  hasSuccess: boolean
}
```
**返信削除**
- APIのURL: `/comment/response/delete`
- Method: DELETE
- フロントエンドから受け取るデータ
```
response_id: int
```
- フロントエンドへと渡すデータ
```
{
  hasSuccess: boolean
}
```
