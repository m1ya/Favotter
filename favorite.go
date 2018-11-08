package main

import (
  . "./keys"
  . "fmt"
  "net/url"
)

func main(){
	api := GetApi()
	
  // 検索のオプション
  v := url.Values{}
  v.Set("count", "10")
   // 検索結果取得
  searchResponse, err := api.GetSearch("KEY_WORD", v)
   if err != nil {
    panic(err)
  }
  // Likeする
  for _, tweet := range searchResponse.Statuses {
    tweet, err := api.Favorite(tweet.Id)
    if err == nil {
      Println(tweet.Text)
    }
  }
}