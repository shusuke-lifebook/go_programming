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

### 1-2-2 constを使って定数を宣言しよう
- 不変変数(定数)は**const**を使って宣言する。
- 定数は関数内でも定義できるが、基本的には関数外で定義します。
- 定数Piは名前の頭文字が大文字ですが、他のファイルから呼び出される場合には大文字にする
- varと同じく、const()で複数の定数を定義することができる。

  ```go
  package main

  import "fmt"

  const Pi = 3.14

  const (
    Username = "test_user"
    Password = "test_pass"
  )

  func main() {
    fmt.Println(Pi, Username, Password)
  }

  ```
- **constとオーバーフロー**
  - constの上限値を超える値を扱うときの説明
  - constにはオーバーフローが発生するようなコードを書いても問題ない
    ```go
    package main

    import "fmt"

    // var big int = 9223372036854775807
    // var big int = 9223372036854775807 + 1
    const big = 9223372036854775807 + 1

    func main() {
      fmt.Println(big - 1)
    }
    ```

## 1-3 データ型について学ぼう
- データ型とは、コードで扱うデータの種類のこと。
- 数値はint型、浮動小数点はfloat型、文字列はstring型、真偽値はbool型
- Goには他にもさまざまなデータ型が用意されている

### 1-3-1 数値型の基本を知ろう
- **数値型**
  ```go
  package main

  import "fmt"

  func main() {
    var (
      u8  uint8     = 255
      i8  int8      = 127
      f32 float32   = 0.2
      c64 complex64 = -5 + 12i
    )
    fmt.Println(u8, i8, f32, c64)
  }

  ```
- **演算子を使った数値の操作**
  - 数値は、+,-などの**算術演算子**を使った式で計算を行える。
    ```go
    package main

    import "fmt"

    func main() {
      fmt.Println("1 + 1 =", 1+1)
      fmt.Println("10 - 1 =", 10-1)
      fmt.Println("10 / 2 =", 10/2)
      fmt.Println("10 / 3 =", 10/3)
      fmt.Println("10.0 / 3 =", 10.0/3)
      fmt.Println("10 / 3.0=", 10/3.0)
      fmt.Println("10 % 2 =", 10%2)
      fmt.Println("10 % 3 =", 10%3)
    }
    ```

  - 次に「++」を使った**インクリメント**と、「--」を使った**デクリメント**
    ```go
    package main

    import "fmt"

    func main() {
      x := 0
      fmt.Println(x)
      x++
      fmt.Println(x)
      x--
      fmt.Println(x)
    }

    ```
  - **シフト演算子**
    - 2進数で表した値を左右にシフト移動して行う計算のこと「<<」「>>」を使う。
    ```go
    package main

    import "fmt"

    func main() {
      fmt.Println(1 << 0)
      fmt.Println(1 << 1)
      fmt.Println(1 << 2)
      fmt.Println(1 << 3)
    }

    ```

### 1-3-2 文字列型の基本をしろう
- **"(ダブルクォート)** や **`(バッククォート)**で囲んだ部分は文字列として扱われる
- 文字列を+でつなげると、文字列の連結を行う

- **文字列から指定した文字を取得する**
  - 文字列は1文字目から順番に**インデックス**という順番を表す番号が0から振られる。
  - "Hello World"[0]でインデックスが0の「H」を取得できる
  - Goの場合はASCIIコードで出力される。文字として出力する場合、string()を使った型変換が必要
    ```go
    package main

    import "fmt"

    func main() {
      fmt.Println("Hello World"[0]) // ASCIIコードで出力
      fmt.Println(string("Hello World"[0]))
    }

    ```

- **文字列型の変数宣言**
  - 文字列型は**string型**です。string型の変数は次のように宣言する
    ```go
    package main

    import "fmt"

    func main() {
      var s string = "Hello World"
      fmt.Println(s)
    }

    ```
- **文字列の置き換え**
  - 文字列操作のための便利な処理(関数)が纏められたstringsパッケージをインポートして、**strings.Replace**関数を使う

    ```go
    package main

    import (
      "fmt"
      "strings"
    )

    func main() {
      var s string = "Hello World"
      fmt.Println(s)
      s = strings.Replace(s, "H", "X", 1)
      fmt.Println(s)
    }

    ```
    
  - stringsパッケージには、他にも指定した文字列を探す**strings.Contain**関数がある。
    ```go
    package main

    import (
      "fmt"
      "strings"
    )

    func main() {
      var s string = "Hello World"
      fmt.Println(s)
      fmt.Println(strings.Contains(s, "World"))
    }

    ```
- 文字列の改行
  - 文字列の途中で改行を入れる場合は「\n」を入れる
    ```go
    package main

    import "fmt"

    func main() {
      fmt.Println("Hello\nWorld")
    }

    ```

### 1-3-3 論理値型の基本を知ろう
- 論理型の**bool**です。短縮変数宣言を使って、変数のtとfを作り、それぞれにtrue/falseを代入する
- fmt.Printf関数でデータ型と値を出力する
  ```go
  package main

  import "fmt"

  func main() {
    t, f := true, false
    fmt.Printf("%T %v\n", t, t)
    fmt.Printf("%T %v\n", f, f)
  }

  ```
- **論理演算子**
  - **&&演算子(論理積)**は左右の値がどちらもtrueであればtrueを出力し、左右のどちらか、もしくは両方がfalseであれば、falseを出力する
  - **||演算子(論理和)**
    - 左右のどちらかがtrueであればtrueを出力、どちらもfalseであればfalseを出力する
  - **!演算子(否定)**は、値を反転した結果を出力する

### 1-3-4 データ型を変換してみよう
- **数値のcast**
  - データ型を変換する**cast**ともいう。
  - int型の変数を定義し、**float64**に変換する例
    ```go
    package main

    import "fmt"

    func main() {
      var x int = 1
      xx := float64(x)
      fmt.Printf("%T %v %f\n", xx, xx, xx)
    }

    ```
- **文字列のcast**
  - string型をint型に変換した場合、**strconv.Atoi関数を使う**
    ```go
    package main

    import (
      "fmt"
      "strconv"
    )

    func main() {
      var s string = "14"
      i, _ := strconv.Atoi(s)
      fmt.Printf("%T %v", i, i)
    }

    ```

## 1-4 データ構造のしくみを学ぼう
- Goには配列(Array)とスライス(Slice)というデータ型がある
  - 配列は固定長
  - スライスは可変長
- マップ(Map)というデータ構造もある

#### 1-4-1 配列(Array)の基本を学ぼう
- 配列もvarを利用して宣言する。「配列名 [要素数]データ型」という形で配列の要素数とデータ型を指定する
  ```go
  package main

  import "fmt"

  func main() {
    var a [2]int
    a[0] = 100
    a[1] = 200
    fmt.Println(a)
  }

  ```
- 配列の宣言ともに値を代入することも可能。宣言時に値を入れる場合は、{}に,(カンマ)で値を区切って入れる。
  ```go
  package main

  import "fmt"

  func main() {
    a := [2]int{100, 200}
    fmt.Println(a)
  }

  ```

#### 1-4-2 スライス(Slice)の基本を学ぼう
- スライスを宣言するには、配列と同じであるが、[]に要素数を指定しないで空とする。
  ```go
  package main

  import "fmt"

  func main() {
    n := []int{1, 2, 3, 4, 5}
    fmt.Println(n)
  }

  ```
- **要素を出力する**
  - **配列[インデックス]**と言った形で要素を取得することが可能
    ```go
    package main

    import "fmt"

    func main() {
      n := []int{1, 2, 3, 4, 5}
      fmt.Println(n[2])
    }

    ```
  - **配列[開始値:終了値]**という書き方で、取得するは要素の範囲を指定することができる。
    ```go
    package main

    import "fmt"

    func main() {
      n := []int{1, 2, 3, 4, 5}
      fmt.Println(n[2:4])
      fmt.Println(n[:2])
      fmt.Println(n[2:])
      fmt.Println(n[:])
    }

    ```
  - **要素の追加**
    - **append関数**を使うことでスライスの要素数をあとから増やせる。
    ```go
    package main

    import "fmt"

    func main() {
      n := []int{1, 2, 3, 4, 5}
      fmt.Println(n)
      n = append(n, 100)
      fmt.Println(n)
      n = append(n, 200, 300, 400)
      fmt.Println(n)
    }

    ```
  - **多次元スライス**
    - スライスの中にスライスを入れることができる。[][]データ型{...}と宣言時に[][]2つ書きます。
    - スライスの中からスライスを取り出すには、スライス[インデックス][インデックス]という形で2つのスライスに対してインデックスを指定する
    ```go
    package main

    import "fmt"

    func main() {
      var board = [][]int{
        []int{0, 1, 2},
        []int{3, 4, 5},
        []int{6, 7, 8},
      }
      fmt.Println(board)
      fmt.Println(board[1])
      fmt.Println(board[1][2])
    }

    ```

#### 1-4-3 make関数でスライスを作ろう
- **make**関数を使うと、要素の値が0で初期化されたスライスを作成することが可能
- 「n := make([]int, 3, 5)」でint型で長さ(length)3、容量(capacity)が5のスライス定義し変数nに代入する
- スライスの長さはlen関数、容量はcap関数で確認できる。
  ```go
  package main

  import "fmt"

  func main() {
    n := make([]int, 3, 5)
    fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)
  }

  ```
- **長さ0のスライス**
  - 長さ0のスライスはmake関数を使う方法と使わない方法の2つのやり方で作成できる
    ```go
    package main

    import "fmt"

    func main() {
      b := make([]int, 0)
      var c []int
      fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
      fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)
    }

    ```

#### 1-4-4 バイト配列を知ろう
- 要素がbyte型のスライスは、「[]byte{値1, 値2,...}」でbyte型の要素を持つスライスを定義できる
- 要素がbyte型のスライスはや配列は、**バイト配列**とも呼べれる。
  - 要素の値はASCIIコードとして扱える
  - string()関数でcastすると文字列が得られる。
    ```go
    package main

    import "fmt"

    func main() {
      b := []byte{72, 73}
      fmt.Println(b)
      fmt.Println(string(b))
    }

    ```

#### 1-4-5 マップ(Map)の基本を学ぼう
- マップも複数の値を1つにすることができるデータ構造で、**キー(Key)と値(Value)**の組み合わせ管理する。
- 「map[キーのデータ型]値のデータ型{キー1:値1, キー2:値2, ....}」という形で宣言する
  ```go
  package main

  import "fmt"

  func main() {
    m := map[string]int{"apple": 100, "banana": 200}
    fmt.Println(m)
    fmt.Println(m["apple"])
  }

  ```
- **値の有無の確認**
  - 要素の存在するかを確認するときには、「変数1」,「変数2」 := マップ[キー]という書き方をつかう。
  - 変数1に要素の値、変数2に真偽値で要素の有無を返す
    ```go
    package main

    import "fmt"

    func main() {
      m := map[string]int{"apple": 100, "banana": 200}
      fmt.Println(m)
      v, ok := m["apple"]
      fmt.Println(v, ok)
      v2, ok2 := m["orange"]
      fmt.Println(v2, ok2)
    }

    ```
- **マップとmake関数**
  - make関数で空のマップを作ってから、空のマップに対して値を入れていくことができる。
    ```go
    package main

    import "fmt"

    func main() {
      m := make(map[string]int)
      m["pc"] = 5000
      fmt.Println(m)
    }

    ```

## 1-5 関数で処理をまとめよう