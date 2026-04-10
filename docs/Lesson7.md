# Lesson 7 Webアプリケーションの作成
- Webアプリケーションは以下のような様々な要素で構成される
  - HTTPリクエスト
  - JSON
  - データベース
- 各要素についてそれぞれ概要を説明しつつ、最後には簡単なWebアプリケーションを作成して理解を深めていきましょう！

## 7-1 HTTPリクエストを送信しよう
- Webページなどでデータをやり取りする際には、HTTP(Hypertext Transfer Protocol)という取り決めに従う必要がある。
- HTTPに従ってデータをアクセスする方法には、
  - Webページからデータを取得するGETリクエスト
  - データを登録するPOSTリクエスト
- Goの標準パッケージを利用して、基本的なHTTPリクエストを送信する方法について記載する。

### 7-1-1 GETリクエストを送信しよう
- **net/http**は、WebサーバーやWebクライアントを作るためのパッケージです。
- Webでデータをやりとりする場合、主にHTTPというプロトコル(取り決め)に従って通信を行う
- **http.Get関数**では指定したURLにGETリクエストを送信できます。

  ```go
  package main

  import (
    "fmt"
    "io"
    "net/http"
  )

  func main() {
    resp, _ := http.Get("http://example.com")
    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
  }

  ```
- **net/url**パッケージでURLを解析しよう
  - URLが正しい形式かどうかを解析するには、**net/url**パッケージをインポートして、url.Parse関数を利用する。
    ```go
    package main

    import (
      "fmt"
      "net/url"
    )

    func main() {
      base, _ := url.Parse("http://example.com")
      fmt.Println(base)
    }

    ```
- **クエリパラメータを追加しよう**
  - HTTPリクエストの際、URLの後ろに追加した渡す値のことを**クエリパラメータ**という。
  - 「http://example.com/test?a=1&b=2」というURLの「a=1&b=2」のように?の後ろに部分がクエリパラメータ
  - **ResolveReference**メソッドで使用するクエリパラメータを追加することができる。

    ```go
    import (
      "fmt"
      "net/url"
    )

    func main() {
      base, _ := url.Parse("http://example.com")
      reference, _ := url.Parse("/test?a=1&b=2")
      endpoint := base.ResolveReference(reference).String()
      fmt.Println(endpoint)
    }

    ```

### 7-1-2 http.NewRequest関数でリクエストを作成しよう
- 今度は**http.NewRequest関数**でリクエストを作成し、GETリクエストを送信する方法を実行してみよう。
  ```go
  package main

  import (
    "fmt"
    "net/http"
    "net/url"
  )

  func main() {
    base, _ := url.Parse("http://example.com")
    reference, _ := url.Parse("/text?a=1&b=2")
    endpoint := base.ResolveReference(reference).String()
    fmt.Println(endpoint)
    req, _ := http.NewRequest("GET", endpoint, nil)
  }

  ```
- **クエリパラメータを確認しよう**
  - **Queryメソッド(req.URL.Query())**で取り出すことが可能

    ```go
    package main

    import (
      "fmt"
      "net/http"
      "net/url"
    )

    func main() {
      base, _ := url.Parse("http://example.com")
      reference, _ := url.Parse("/text?a=1&b=2")
      endpoint := base.ResolveReference(reference).String()
      fmt.Println(endpoint)
      req, _ := http.NewRequest("GET", endpoint, nil)
      q := req.URL.Query()
      fmt.Print(q)
    }

    ```
- **Addメソッドでクエリパラメータを追加する**
  - http.NewRequest関数でリクエストを作成する場合、**Addメソッド**でクエリパラメータを追加することができる。
  - **Encodeメソッド**を利用してエンコードする必要がある。

    ```go
    package main

    import (
      "fmt"
      "net/http"
      "net/url"
    )

    func main() {
      base, _ := url.Parse("http://example.com")
      reference, _ := url.Parse("/text?a=1&b=2")
      endpoint := base.ResolveReference(reference).String()
      fmt.Println(endpoint)
      req, _ := http.NewRequest("GET", endpoint, nil)
      q := req.URL.Query()
      q.Add("c", "3&%")
      fmt.Println(q)
      fmt.Print(q.Encode())
      req.URL.RawQuery = q.Encode()
    }

    ```
- **http.ClientでHTTPリクエストを送信しよう**
  - 作成したHTTPリクエストを使ってGETリクエストを行う際には、**http.Client**という型でクライアントを作成する必要がある。

    ```go
    package main

    import (
      "fmt"
      "io"
      "net/http"
      "net/url"
    )

    func main() {
      base, _ := url.Parse("http://example.com")
      reference, _ := url.Parse("/text?a=1&b=2")
      endpoint := base.ResolveReference(reference).String()
      fmt.Println(endpoint)
      req, _ := http.NewRequest("GET", endpoint, nil)
      q := req.URL.Query()
      q.Add("c", "3&%")
      fmt.Println(q)
      fmt.Print(q.Encode())
      req.URL.RawQuery = q.Encode()

      var client *http.Client = &http.Client{}
      resp, _ := client.Do(req)
      body, _ := io.ReadAll(resp.Body)
      fmt.Println(string(body))
    }

    ```

### 7-1-3 POSTリクエストを送信しよう
- POSTリクエストは、http.NewRequestでリクエストを作成する際に、メソッドを"POST"に指定する。
- POSTの場合はリクエストボディにデータを入れて渡す。

  ```go
  package main

  import (
    "bytes"
    "fmt"
    "io"
    "net/http"
    "net/url"
  )

  func main() {
    base, _ := url.Parse("http://example.com")
    reference, _ := url.Parse("/test?a=1&b=2")
    endpoint := base.ResolveReference(reference).String()
    fmt.Println(endpoint)
    req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))

    var client *http.Client = &http.Client{}
    resp, _ := client.Do(req)
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
  }

  ```

## 7-2 JSONと構造体を相互に変換しよう
- JSONとは、「JavaScript Object Notation」の略で、Web上でデータをやり取りする際によく用いられるデータの記述形式です。
- encoding/jsonという標準パッケージを用いることで、JSONのデータをGoの構造体として扱い処理することが可能。

### 7-2-1 json.Unmarshal関数でJSONを構造体に変換しよう

## 7-3 データベースを利用しよう

## 7-4 Webアプリケーションを作成しよう