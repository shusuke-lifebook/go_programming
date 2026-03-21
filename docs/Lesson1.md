# Lesson 1 Goの基本
- Goの基本の書き方から、データや変数の扱い型、関数の作り方を見ていく。

## 1-1 基本的な処理の流れを学ぼう
- Goの基本的なコードの書き方と処理の流れについて説明
- Goはパッケージという単位でソースコードを管理する
- 多くの組み込み関数はパッケージのインポートが必要となる

### 1-1-1 main関数とinit関数の働きを知ろう
- **main関数**
  - 1行目にmainという**パッケージ**を宣言している。
  - 3行目では**fmt**というパッケージをインポートしている
  - 5～7行目は**main関数**の定義
    ```go
    package main

    import "fmt"

    func main() {
      fmt.Println("Hello world!")
    }

    ```
- **複数の文字を出力する**
  - fmt.Println関数には、複数の引数を渡すことができる。,(カンマ)で引数を区切って入れると、引数の値を連結して出力する
- **関数の呼び出し**
  - 新しいbazz関数を定義してmain関数から呼び出してみる。
  - 関数を定義するときは**func**の後に関数名(){}と続ける。
  - 実行したい処理を{}の中に記載する。
  - 定義した関数は、関数名()で呼び出すことができる。
    ```go
    package main

    import "fmt"

    func bazz() {
      fmt.Println("Bazz")
    }

    func main() {
      bazz()
      fmt.Println("Hello world!", "golang")
    }

    ```
- **init**関数
  - main関数以外にも、**init**という特別な関数がある。
  - init関数が定義されている場合、main関数より先に呼ばれる。そのため、次のような流れで処理が実行される。
    - init関数の呼び出し
      - fmt.Println("init!")を実行
    - main関数の呼び出し
    - bazz関数の呼び出し
      - fmt.Println("Bazz")を実行
    - main関数に戻る
    - fmt.Println("Hello world!", "golang")を実行
    ```go
    package main

    import "fmt"

    func init() {
      fmt.Println("Init!")
    }

    func bazz() {
      fmt.Println("Bazz")
    }

    func main() {
      bazz()
      fmt.Println("Hello, world!", "golang")
    }

    ``` 
- **コメントアウト**
  - コメントアウトの書き方は２つある
    - 行頭に//をつける方法
    - 複数の行を/* */で囲む方法

### 1-1-2 複数のパッケージをインポートしよう
- Goにはさまざまな標準パッケージが用意されており、複数のパッケージを利用する場面がある
- 複数のパッケージのインポートする場合は、importの後に()をつけ、()内に改行を区切ってパッケージ名を記述する
- パッケージや関数の情報は「go doc パッケージ名」または、「go doc パッケージ名 関数名」で確認することが可能
  ```go
  package main

  import (
    "fmt"
    "time"
  )

  func main() {
    fmt.Println("Hello, world!", time.Now())
  }

  ```

## 1-2 変数の作り方をマスターしよう
- 変数は、数値や文字列など何らかのデータを入れる箱のようなもの
- 変数を作ることを「宣言」、変数に値をいれることを「代入」という。
- また、データには種類があり、変数を宣言するときに代入したいデータ種類を指定する。

### 1-2-1 変数を宣言しよう
- 変数の宣言には**var**を使う。このvarは、変数という意味を持つvariableの略です。
- varの後の変数名とデータの種類を表す型を書く。
- 宣言と同時に初期化したい場合は、続けて=(イコール)と代入したい値を書く
  ```go
  package main

  import "fmt"

  func main() {
    var i int = 1
    fmt.Println(i)
  }

  ```
- **複数の変数を宣言する**
  - importと同様に、varの後に()をつけて、()内で複数の変数を宣言できる。
    ```go
    package main

    import "fmt"

    func main() {
      var (
        i    int     = 1
        f64  float64 = 1.2
        s    string  = "test"
        t, f bool    = true, false
      )
      fmt.Println(i, f64, s, t, f)
    }

    ```
- **短縮変数宣言(short variable declearation)**
  - 短縮変数宣言という記述方法にすると、varを省略して変数を宣言できる
  - 短縮変数宣言は、変数名と値を:=**(コロンとイコール)**でつなげるだけです。
  - 例えば、「xi := 1」とすると、int型の変数xiが定義され、xiに1が代入される。
    ```go
    package main

    import "fmt"

    func main() {
      xi := 1
      xf64 := 1.2
      xs := "test"
      xt, xf := true, false
      fmt.Println(xi, xf64, xs, xt, xf)
    }

    ```
  - 「:=」を使って変数を宣言すると、データ型が自動的に設定されるので注意
  - **fmt.Printf**関数で変数のデータ型を確認することできる。

- **宣言方法の使い分け**
  - varを使う方法と短縮変数宣言の違いは、関数の外で定義できるかどうか
    - var を使う場合、関数の外で定義することが可能
    - 短縮変数宣言は関数の中でしか使えない

## 1-3 データ型について学ぼう

## 1-4 データ構造のしくみを学ぼう

## 1-5 関数で処理をまとめよう