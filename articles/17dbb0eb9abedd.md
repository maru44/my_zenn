---
title: "genericsのjsonを埋め込みたかった"
emoji: "🦖"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: ["go"]
published: true
---

## generics の無名埋め込みができない

1.18 になり go に待望の generics がやってきました。
個人的に凄い便利でいいなーと思って使っています。

しかし現状`struct`を json に `marshal` した際にその階層に展開できないなーということに気がつきました。

**BAD これはできない**

```go:bad.go
type umekomi[T any] struct {
    Name string
    T
}
```

**これはできるが...**

```go
type gene[T any] struct {
    Name string
    Extra T
}
```

例えば以下のようなものを json にすると

```go:sample.go
type ext struct {
    Age int
    BornAt time.Time
}

var bornThisWay = gene[ext]{
    Name: "Gaga",
    Extra: ext{
        Age: 0,
        BornAt: time.Now(),
    },
}
```

こんな感じになってしまう

```json
{
  "Name": "Gaga",
  "Extra": {
    "Age": 0,
    "BornAt": "****"
  }
}
```

同じ階層に展開したいなーと思いました
別に手前で構造体を定義するときはこれに則った上で定義すれば何も問題ないが、他所の API とかを引っ張ってくる際には generics 版無名埋め込みが欲しいなーと

## proposal 出してみた

go にプロポーザルを出してみた

https://github.com/golang/go/issues/52138

英語テキトーなのは申し訳ない。衝動に駆られ 2 分で書いたため聳え立つ糞みたいな英語だ。

緊張の初 `proposal` & 初 `PR`

その結果は...

`dup` と判断され一瞬にしてクローズに 😭

そもそもこちらの意図がちゃんと伝わったかどうか不明だがまあいいだろう。

作法に則っていなかったらしく怒られてしまったが、 **英語わかりまシェーン、ダンケシェーン** という気持ちでやっている。

**余談:**

test はちゃめちゃやないかい...
なんでこんなんになっとるんや 😰

## 諦めないゾイ

こんなの作った
https://github.com/maru44/gson

使い方

```go
package main

import (
    "fmt"

    "github.com/maru44/gson"
)

type good[T any] struct {
    Name string
    Free T `json:"..."`
}

type more struct {
    Age int
    Country string
}

func main() {
    got := good[more]{}
    in := `{"Name": "Foo", "Age": 20, "Country": "U.S.A"}`

    _ = gson.Unmarshal([]byte(in), got)
}

```

とやると

```
good[more]{
    Name: "Foo",
    more: more{
        Age: 20,
        Country: "U.S.A",
    },
}
```

こんな感じのが帰ってくる
逆(marshal)も然り

:::message
コードを見てもらうとわかりますが、完全に `encoding/json`のパクリだが、そこは引用・オマージュ・再構築 (古塔 ◯ み風)と捉えてもらえるとありがたいです。
:::
