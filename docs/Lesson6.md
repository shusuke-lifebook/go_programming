# Lesson 6 パッケージ
- Goのコードはパッケージという単位で構成される。
- 自分で作成した別のパッケージのコードを利用したり、Goの標準パッケージやサードパーティ製のパッケージを利用することで、よりさまざまなコードを書くことができる

## 6-1 パッケージでコードを管理しよう
- 今までは、mainパッケージに属する単独のファイルを使ってきました。
- ただし、コードの量が増えていくとコードを他のファイルに書いて読みたくなると思う。
- ここではパッケージごとにファイルを分け別のパッケージから関数や型を呼び出す方法について記載する

### 6-1-1 パッケージ単位でコードを分けよう
- Goのコードは、1つまたは複数のファイルで構成される**パッケージ**という単位で管理する
- 「mylib」という名前でフォルダを作成する。
- 新しいパッケージを作成するさいには新しいフォルダを作成する必要がある、パッケージ名とフォルダ名を対応させる必要がある。
- 「mylib/math.go」
  ```go
  // Package mylib provides math utilities.
  package mylib

  func Average(s []int) int {
    total := 0
    for _, i := range s {
      total += i
    }
    return total / len(s)
  }

  ```
- main.go
  ```go
  package main

  import (
    "awesomeProject/mylib"
    "fmt"
  )

  func main() {
    s := []int{1, 2, 3, 4, 5}
    fmt.Println(mylib.Average(s))
  }

  ```
#### 6-1-2 関数や型をエクスポートしよう
- エクスポート機能：
  - 関数や型の名前を大文字ではじめると、他のパッケージが呼び出せるようになる
  - 逆に小文字からはじめると、他のパッケージから関数や型を呼び出せずエラーとなる。
  - 同じパッケージから呼び出す場合は小文字で書いてもアクセスできる

## 6-2 テスト実行しよう
- 作成したコードが正しく動作するかを確認するには、テスト用のコードを作成する
- テストを作成することで、コードに変更があった場合でも、テストが成功すればコードの動作が問題ないことが確認できる。
- Goの標準パッケージであるtestingを使ってテストを作成し実行する方法を記載する

### 6-2-1 testingパッケージでテストを作ろう
- **testing**パッケージを使って、テストを書く
  - [testingのドキュメント](https://pkg.go.dev/testing)
- **math.go**のテストを作成する
  - テストを作る際は、テストする対象と同じフォルダに、テスト用のコードを記述するファイルを入れる
  - testingパッケージは、ファイル名の末尾が「_test.go」のファイルを読み込んでテストを行う。
  - 「go test ./...」と入力してコマンド実行する。現在のディレクトリとその下にあるすべての「_test.go」がついたファイルを実行できる

    ```go
    package mylib

    import "testing"

    func TestAverage(t *testing.T) {
      v := Average([]int{1, 2, 3, 4, 5})
      if v != 3 {
        t.Error("Expected 3, got", v)
      }
    }

    ```
- **testingのメソッド**
  - **t.Fail** テストが失敗し、かつテストの実行を続けたいときに使う。
  - **t.Skip** テストが必要ないときに使う

## 6-3 コードの形式を整えよう
- コードを読みやすくするために、形式を整えることは重要。
- ただし、手動でコードを修正していくのは非常に大変かつ修正漏れが発生する恐れがある。
- ツールを使ってある程度の修正することが一般的です。
- ここでは、gofmtというコードを修正するためのツールについて記載する。

### 6-3-1 gofmtでコードの形式を整えよう
- **gofmt**は、Goのコードを書き方を修正するツールです。
  - 「gofmt math.go」(gofmt 修正するファイル名)で実行すると整形したコードが表示される。
  - 「gofmt -w math.go」のように-wをつけると修正したコードで上書きしてくれる。

## 6-4 サードパーティのパッケージを利用しよう
- Goにはさまざまな機能を持つ標準パッケージがあります。
- 開発したい機能によっては、サードパーティのパッケージを活用する場合もあります。
- ここでは、サードパーティのパッケージをインストールして使用する方法について説明する。

### 6-4-1 サードパーティのパッケージのインストール
- [第三者が公開しているパッケージはこちらから検索できます](https://pkg.go.dev/)
- 「talib」という株価を分析するサードパーティのパッケージをインストールする
- 「quote」という株価などの情報をダウンロードするパッケージをインストールする
- Goでは、**go get**というコマンドでサードパーティのパッケージをインストールする
  - go get github.com/markcheno/go-talib
  - go get github.com/markcheno/go-quote@latest
    ```go
    package main

    import (
      "fmt"

      "github.com/markcheno/go-quote"
      "github.com/markcheno/go-talib"
    )

    func main() {
      // Coinbase からデータ取得（BTC-USD の例）
      spy, err := quote.NewQuoteFromCoinbase("BTC-USD", "2021-04-01", "2021-04-04", quote.Daily)
      if err != nil {
        panic(err)
      }

      fmt.Print(spy.CSV())

      // RSI 計算
      rsi2 := talib.Rsi(spy.Close, 2)
      fmt.Println(rsi2)
    }

    ```

## 6-5 ドキュメントを作成しよう
- コードを作成したら、ドキュメントを整備することも大切です。
- Goには、コードをドキュメント化するためのツールが用意されている。
- ドキュメントを書く際にはいくつかのルールがあります。
- ここでは、ドキュメントの書き方と、その内容を確認する方法について記述する。

### 6-5-1 go docでコードの説明を確認しよう
- **go doc**は、Goのドキュメント確認するツールです。
- ターミナルで「go doc fmt.Println」とコマンドを入力すると、Printlnの説明文がでる。
- 自分で作成したmath.goのAverage関数にも説明を書いていきます。
- 説明文にはコメントアウトして書きます。先頭に//を書く方法と、/**/で囲む方法があるがどちらで記載しても問題ない。
- また、文の最初を関数名と同じ(ここでは「Average」とする必要がある。)
  ```go
  // Average returns the average of a series of numbers
  func Average(s []int) int {
    total := 0
    for _, i := range s {
      total += i
    }
    return total / len(s)
  }

  ```
- 「go doc mylib.Average」とターミナルでコマンドを実行すると、コードに書いた説明が表示される。

### 6-5-2 godocでブラウザ上のドキュメントを確認しよう
- go docと似たコマンドに、**godoc**がある。godocは説明文を書いたコードをローカルのWebページで実行することができる
- godocは1.12まで利用可能であったが、それ以降のバージョンでは非推奨。
- 以下のツールを利用することにする
  - [Arrow](https://github.com/navid-m/arrow)
  - [doc2go](https://abhinav.github.io/doc2go/)
  - インストール
    - Arrow 
      - コマンド: go install github.com/navid-m/arrow@latest
      - コマンド：arrow -v
    - doc2go 
      - コマンド: go install go.abhg.dev/doc2go@latest
      - コマンド：doc2go -version
  - 使い方
    - Arrow
      ```
      arrow .
      open docs/index.html   # macOS
      xdg-open docs/index.html  # Linux
      start docs/index.html     # Windows
      ```
    - doc2go
      ```
      doc2go -out www/ ./...
      ```
      → www/ を GitHub Pages に置くだけで公開できる。




## 6-6 便利な標準パッケージを活用しよう