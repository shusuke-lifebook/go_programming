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
- **encoding/json**はJSONを扱う標準パッケージです。
- JSONのデータはbyteで作成する。
- **json.Unmarshal関数**に変数bとポインタ変数pを渡す。

  ```go
  package main

  import (
    "encoding/json"
    "fmt"
  )

  type Person struct {
    Name      string
    Age       int
    Nicknames []string
  }

  func main() {
    b := []byte(`{"name":"mike","age":20,"nicknames":["a","b","c"]}`)
    var p Person

    if err := json.Unmarshal(b, &p); err != nil {
      fmt.Println(err)
    }
    fmt.Println(p.Name, p.Age, p.Nicknames)
  }

  ```

### 7-2-2 json.Marshal関数で構造体をJSONに変換しよう
- 構造体のデータをJSONに変換するには、**json.Marshal**関数を使う。

  ```go
  package main

  import (
    "encoding/json"
    "fmt"
  )

  type Person struct {
    Name      string
    Age       int
    Nicknames []string
  }

  func main() {
    b := []byte(`{"name":"mike","age":20, "nicknames":["a","b","c"]}`)
    var p Person
    if err := json.Unmarshal(b, &p); err != nil {
      fmt.Println(err)
    }
    fmt.Println(p.Name, p.Age, p.Nicknames)

    v, _ := json.Marshal(p)
    fmt.Println(string(v))
  }

  ```

- **JSONに変換する際のキー名を指定しよう**
**構造体のフィールド名がJSONのキーに変換されるときの名前を指定することができる**

  ```go
  package main

  import (
    "encoding/json"
    "fmt"
  )

  type Person struct {
    Name      string   `json:"name"`
    Age       int      `json:"age"`
    Nicknames []string `json:"nicknames"`
  }

  func main() {
    b := []byte(`{"name":"micke","age":20,"nicknames":["a","b","c"]}`)
    var p Person

    if err := json.Unmarshal(b, &p); err != nil {
      fmt.Println(err)
    }
    fmt.Println(p.Name, p.Age, p.Nicknames)

    v, _ := json.Marshal(p)
    fmt.Println(string(v))
  }

  ```

- **JSONでの型を変更しよう**
  - 構造体とJSONとの間で変換を行うときに、違う型として扱うことができる
  - 「json:"age,string"」とした場合、json.Marshal関数でJSONに変換したときのageの値はstring型になる

    ```go
    package main

    import (
      "encoding/json"
      "fmt"
    )

    type Person struct {
      Name      string   `json:"name"`
      Age       int      `json:"age,string"`
      Nicknames []string `json:"nicknames"`
    }

    func main() {
      b := []byte(`{"name":"micke","age":"20","nicknames":["a","b","c"]}`)
      var p Person

      if err := json.Unmarshal(b, &p); err != nil {
        fmt.Println(err)
      }
      fmt.Println(p.Name, p.Age, p.Nicknames)

      v, _ := json.Marshal(p)
      fmt.Println(string(v))
    }

    ```
- **JSONに変換するときに非表示にしよう**
  - 「`json:"-"`」としていると、そのフィールドはjson.Unmarshal関数やjson.Marshal関数での変換時にJSON側では無視される。

    ```go
    package main

    import (
      "encoding/json"
      "fmt"
    )

    type Person struct {
      Name      string   `json:"-"`
      Age       int      `json:"age,string"`
      Nicknames []string `json:"nicknames"`
    }

    func main() {
      b := []byte(`{"name":"micke","age":"20","nicknames":["a","b","c"]}`)
      var p Person

      if err := json.Unmarshal(b, &p); err != nil {
        fmt.Println(err)
      }
      fmt.Println(p.Name, p.Age, p.Nicknames)

      v, _ := json.Marshal(p)
      fmt.Println(string(v))
    }

    ```

### 7-2-3 omitemptyでデフォルト値を省略しよう
- **omitempty**を指定することでJSONに反映させず省略することが可能
  ```go
  package main

  import (
    "encoding/json"
    "fmt"
  )

  type Person struct {
    Name      string   `json:"-"`
    Age       int      `json:"age,omitempty"`
    Nicknames []string `json:"nicknames"`
  }

  func main() {
    b := []byte(`{"name":"micke","age":0,"nicknames":["a","b","c"]}`)
    var p Person

    if err := json.Unmarshal(b, &p); err != nil {
      fmt.Println(err)
    }
    fmt.Println(p.Name, p.Age, p.Nicknames)

    v, _ := json.Marshal(p)
    fmt.Println(string(v))
  }
  ```
- **空の構造体にomitemptyを適用する**

  ```go
  package main

  import (
    "encoding/json"
    "fmt"
  )

  type T struct {
  }

  type Person struct {
    Name      string   `json:"name,omitempty"`
    Age       int      `json:"age,omitempty"`
    Nicknames []string `json:"nicknames,omitempty"`
    T         *T       `json:"T,omitempty"`
  }

  func main() {
    b := []byte(`{"name":"","age":20,"nicknames":[]}`)
    var p Person

    if err := json.Unmarshal(b, &p); err != nil {
      fmt.Println(err)
    }
    fmt.Println(p.Name, p.Age, p.Nicknames)

    v, _ := json.Marshal(p)
    fmt.Println(string(v))
  }

  ```

### 7-2-4 json.Marshal関数をカスタマイズしよう
- json.Marshalを独自処理に変更したい場合、MarshalJSONメソッドを実装する。

  ```go
  package main

  import (
    "encoding/json"
    "fmt"
  )

  type T struct {
  }

  type Person struct {
    Name      string   `json:"name,omitempty"`
    Age       int      `json:"age,omitempty"`
    Nicknames []string `json:"nicknames,omitempty"`
    T         *T       `json:"T,omitempty"`
  }

  func (p Person) MarshalJSON() ([]byte, error) {
    v, err := json.Marshal(&struct {
      Name string
    }{
      Name: "Mr." + p.Name,
    })
    return v, err
  }

  func main() {
    b := []byte(`{"name":"Mike","age":20,"nicknames":[]}`)
    var p Person

    if err := json.Unmarshal(b, &p); err != nil {
      fmt.Println(err)
    }
    fmt.Println(p.Name, p.Age, p.Nicknames)

    v, _ := json.Marshal(p)
    fmt.Println(string(v))
  }

  ```

### 7-2-5 json.Unmarshal関数をカスタマイズしよう
- json.Unmarshal関数でJSONを構造体に変換する処理を独自にカスタマイズする場合、**UnmarshalJSONメソッド**を実装する。

  ```go
  package main

  import (
    "encoding/json"
    "fmt"
  )

  type T struct {
  }

  type Person struct {
    Name      string   `json:"name,omitempty"`
    Age       int      `json:"age,omitempty"`
    Nicknames []string `json:"nicknames,omitempty"`
    T         *T       `json:"T,omitempty"`
  }

  func (p Person) MarshalJSON() ([]byte, error) {
    v, err := json.Marshal(&struct {
      Name string
    }{
      Name: "Mr." + p.Name,
    })
    return v, err
  }

  func (p *Person) UnmarshalJSON(b []byte) error {
    type Person2 struct {
      Name string
    }
    var p2 Person2
    err := json.Unmarshal(b, &p2)
    if err != nil {
      fmt.Println(err)
    }
    p.Name = p2.Name + "!"
    return err
  }

  func main() {
    b := []byte(`{"name":"Mike","age":20,"nicknames":[]}`)
    var p Person

    if err := json.Unmarshal(b, &p); err != nil {
      fmt.Println(err)
    }
    fmt.Println(p.Name, p.Age, p.Nicknames)

    v, _ := json.Marshal(p)
    fmt.Println(string(v))
  }

  ```

## 7-3 データベースを利用しよう
- SQLiteを例に、簡単なリレーショナルデータベースとSQLの使い方を記載する。

### 7-3-1 SQLiteを利用する準備をしよう
- [go-sqlite3](https://github.com/mattn/go-sqlite3)というパッケージを使います
  - go get github.com/mattn/go-sqlite3
  - SQLでデータベースを操作するには、**database/sql**パッケージを使用する
    - sql.Open関数でデータベースを開く
    - SQL文を作る
    - SQL文を実行する
    - **Closeメソッド**でデータベースを閉じる

    ```go
    package main

    import (
      "database/sql"
      "log"

      _ "github.com/mattn/go-sqlite3"
    )

    var DbConnection *sql.DB

    func main() {
      DbConnection, _ := sql.Open("sqlite3", "./example.sql")
      defer DbConnection.Close()
      cmd := `CREATE TABLE IF NOT EXISTS person (
        name STRING,
        age INT
      )`

      _, err := DbConnection.Exec(cmd)
      if err != nil {
        log.Fatalln(err)
      }
    }

    ```
- **INSERT文でレコードを挿入しよう**
  ```go
  package main

  import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
  )

  var DbConnection *sql.DB

  func main() {
    DbConnection, _ := sql.Open("sqlite3", "./example.sql")
    defer DbConnection.Close()
    cmd := `CREATE TABLE IF NOT EXISTS person (
      name STRING,
      age INT
    )`

    _, err := DbConnection.Exec(cmd)
    if err != nil {
      log.Fatalln(err)
    }

    cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
    _, err = DbConnection.Exec(cmd, "Nancy", 20)

    if err != nil {
      log.Fatalln(err)
    }
  }

  ```

- **UPDATE文でレコードを更新しよう**
  ```go
  package main

  import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
  )

  var DbConnection *sql.DB

  func main() {
    DbConnection, _ := sql.Open("sqlite3", "./example.sql")
    defer DbConnection.Close()
    cmd := `CREATE TABLE IF NOT EXISTS person (
      name STRING,
      age INT
    )`

    _, err := DbConnection.Exec(cmd)
    if err != nil {
      log.Fatalln(err)
    }

    cmd = "UPDATE person SET age = ? WHERE name = ?"
    _, err = DbConnection.Exec(cmd, 25, "Mike")

    if err != nil {
      log.Fatalln(err)
    }
  }

  ```
- **SELECT文で複数のレコードを取得しよう**

  ```go
  package main

  import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
  )

  var DbConnection *sql.DB

  type Person struct {
    Name string
    Age  int
  }

  func main() {
    DbConnection, _ := sql.Open("sqlite3", "./example.sql")
    defer DbConnection.Close()
    cmd := `CREATE TABLE IF NOT EXISTS person (
      name STRING,
      age INT
    )`

    _, err := DbConnection.Exec(cmd)
    if err != nil {
      log.Fatalln(err)
    }

    cmd = "SELECT * FROM person"
    rows, err := DbConnection.Query(cmd)
    defer rows.Close()
    var pp []Person
    for rows.Next() {
      var p Person
      err := rows.Scan(&p.Name, &p.Age)
      if err != nil {
        log.Println(err)
      }
      pp = append(pp, p)
    }

    for _, p := range pp {
      fmt.Println(p.Name, p.Age)
    }

  }

  ```

- **SELECT文で1件のレコードを取得しよう**
  - レコードを1件取得したい場合は、**QueryRowメソッド**を使用する

    ```go
    package main

    import (
      "database/sql"
      "fmt"
      "log"

      _ "github.com/mattn/go-sqlite3"
    )

    var DbConnection *sql.DB

    type Person struct {
      Name string
      Age  int
    }

    func main() {
      DbConnection, _ := sql.Open("sqlite3", "./example.sql")
      defer DbConnection.Close()
      cmd := `CREATE TABLE IF NOT EXISTS person (
        name STRING,
        age INT
      )`

      _, err := DbConnection.Exec(cmd)
      if err != nil {
        log.Fatalln(err)
      }

      cmd = "SELECT * FROM person where age = ?"
      row := DbConnection.QueryRow(cmd, 20)
      var p Person
      err = row.Scan(&p.Name, &p.Age)
      if err != nil {
        if err == sql.ErrNoRows {
          log.Println("No row")
        } else {
          log.Println(err)
        }
      }
      fmt.Println(p.Name, p.Age)
    }

    ```

- **DELETE文でレコードを削除しよう**
  ```go
  package main

  import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
  )

  var DbConnection *sql.DB

  type Person struct {
    Name string
    Age  int
  }

  func main() {
    DbConnection, _ := sql.Open("sqlite3", "./example.sql")
    defer DbConnection.Close()
    cmd := `CREATE TABLE IF NOT EXISTS person (
      name STRING,
      age INT
    )`

    _, err := DbConnection.Exec(cmd)
    if err != nil {
      log.Fatalln(err)
    }

    cmd = "DELETE FROM person where name = ?"
    _, err = DbConnection.Exec(cmd, "Nancy")
    if err != nil {
      log.Fatalln(err)
    }
  }

  ```

### 7-3-2 SQLインジェクションの例を確認しよう
- 以下のプログラムを実行すると何も表示されない。
- ターミナルからデータベースを確認すると、「Mr.X」のレコードが追加されている。
- 意図しないSQLが実行されることを**SQLインジェクション**という。

  ```go
  package main

  import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
  )

  var DbConnection *sql.DB

  type Person struct {
    Name string
    Age  int
  }

  func main() {
    DbConnection, _ := sql.Open("sqlite3", "./example.sql")
    defer DbConnection.Close()
    cmd := `CREATE TABLE IF NOT EXISTS person (
      name STRING,
      age INT
    )`

    _, err := DbConnection.Exec(cmd)
    if err != nil {
      log.Fatalln(err)
    }

    tableName := "person; INSERT INTO person (name, age) VALUES ('Mr.X', 100)"
    cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
    rows, _ := DbConnection.Query(cmd)
    defer rows.Close()
    var pp []Person
    for rows.Next() {
      var p Person
      err := rows.Scan(&p.Name, &p.Age)
      if err != nil {
        log.Println(err)
      }
      pp = append(pp, p)
    }
    err = rows.Err()
    if err != nil {
      log.Fatalln(err)
    }
    for _, p := range pp {
      fmt.Println(p.Name, p.Age)
    }
  }

  ```

## 7-4 Webアプリケーションを作成しよう
- パソコンやスマホのブラウザから利用するWebページで動作しているWebアプリケーションは、これまで記載してきた、HTTPやJSON、データベースなどの処理を組み合わせて構成されている
- シンプルにテキストデータを読み込み、GETリクエストとPOSTリクエストを処理するWebアプリケーションを例にして、基本的なWebアプリケーションのコードについて記載する

### 7-4-1 テキストを編集して表示するアプリケーションを作ろう
- GoでWebアプリケーションを作っていきます。
- Goの公式ページを紹介されている[アプリケーション](https://go.dev/doc/articles/wiki)を例に記載する。

### 7-4-2 OSパッケージでテキストファイルの読み込みをしよう
- 最初にGoの標準パッケージOSパッケージを使ってテキストファイルに対して読み書きを行う処理を作っていきましょう
- os.WriteFile関数でファイルにデータを書き込もう
  ```go
  package main

  import "os"

  type Page struct {
    Title string
    Body  []byte
  }

  func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
  }

  func main() {
    p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
    p1.save()
  }

  ```

- **os.ReadFile関数でファイルからデータを読み込む**

  ```go
  package main

  import (
    "fmt"
    "os"
  )

  type Page struct {
    Title string
    Body  []byte
  }

  func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
  }

  func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
      return nil, err
    }
    return &Page{Title: title, Body: body}, nil
  }

  func main() {
    p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
    p1.save()

    p2, _ := loadPage(p1.Title)
    fmt.Println(string(p2.Body))
  }

  ```

### 7-4-3 Webサーバーを立ち上げよう
- **http**パッケージを使い、Webサーバーの処理を作成する。
- **http.ListenAndServe**関数でWebサーバーを立ち上げる。

  ```go
  package main

  import (
    "fmt"
    "log"
    "net/http"
    "os"
  )

  type Page struct {
    Title string
    Body  []byte
  }

  func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
  }

  func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
      return nil, err
    }
    return &Page{Title: title, Body: body}, nil
  }

  func viewHandler(w http.ResponseWriter, r *http.Request) {
    // /view/test
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
  }

  func main() {
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
  }

  ```

### 7-4-4 HTMLテンプレートを利用しよう
- **text/template**パッケージを使って読み込んでいこう
  ```go
  func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
  }

  func viewHandler(w http.ResponseWriter, r *http.Request) {
    // /view/test
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
  }
  ```

### 7-4-5 Webページからの入力内容をファイルに保存しよう
  ```go
  func saveHandler(w http.ResponseWriter, r *http.Request) {
    // /save/test
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
  }

  ```

### 7-4-6 Webアプリケーションのコードを効率化しよう
- template.Must関数でtemplate.ParseFiles関数の引数に「edit.html」、「view.html」を指定する。
- ExecuteTemplateメソッドでページ内容をテンプレートに反映し、もしエラーがあれば、http.StatusInternalServerErrorを返す。

  ```go
  var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

  func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
  ```

- **Handlerで共通の処理を纏める**
  ```go
  package main

  import (
    "log"
    "net/http"
    "os"
    "regexp"
    "text/template"
  )

  type Page struct {
    Title string
    Body  []byte
  }

  func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
  }

  func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
      return nil, err
    }
    return &Page{Title: title, Body: body}, nil
  }

  var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

  func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }

  func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    // /view/test
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
  }

  func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    // /edit/test
    p, err := loadPage(title)
    if err != nil {
      p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
  }

  func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    // /save/test
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
  }

  var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

  func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      m := validPath.FindStringSubmatch(r.URL.Path)
      if m == nil {
        http.NotFound(w, r)
        return
      }
      fn(w, r, m[2])
    }
  }

  func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
    log.Fatal(http.ListenAndServe(":8080", nil))
  }

  ```