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

## 1-3 データ型について学ぼう

## 1-4 データ構造のしくみを学ぼう

## 1-5 関数で処理をまとめよう