# Lesson4 Struct オリエンテッド
- Goには、クラスや継承といったオブジェクト指向プログラムのための機能はありません。
- データ型にメソッドを作成したり、構造体の中に構造体を持たせて擬似的な継承をしたりすることができる

## 4-1 メソッドを作成しよう
- Goでは構造体を含むデータ型に紐づけた関数のことをメソッドと呼びます。
- ポインタからメソッドを呼び出すときとは注意すべきことがある。
- メソッドの基本的な使い方と注意点、そして他のプログラミング言語におけるコンストラクタのような初期化処理の実装方法について記載する。

### 4-1-1 型に紐づくメソッドを作成しよう
- 構造体に**メソッド**を作成する
  - int型のX,Yを持つ構造体Vertexを宣言する
  - 引数が構造体Vertexの変数v、返り値がint型のArea関数を作る
    ```go
    type Vertex struct {
      X, Y int
    }

    func Area(v Vertex) int {
      return v.X * v.Y
    }

    func main() {
      v := Vertex{3, 4}
      fmt.Println(Area(v))
    }

    ```
  - Area関数を構造体Vertexに結びつけたメソッドとして定義してみよう
  - メソッド作成時は「func (v Vertex) Area() int」のように、funcの後に()をつけ、その中に**レシーバー**と呼ばれる引数の名前と型を指定する。
  - メソッドを呼び出すには、「v.Area()」のように、**メソッドを結びつけた構造体の変数とメソッド名を.(ドット)でつないで実行します。**

    ```go
    package main

    import "fmt"

    type Vertex struct {
      X, Y int
    }

    func (v Vertex) Area() int {
      return v.X * v.Y
    }

    func Area(v Vertex) int {
      return v.X * v.Y
    }

    func main() {
      v := Vertex{3, 4}
      fmt.Println(Area(v))
      fmt.Println(v.Area())
    }

    ```

### 4-1-2 ポインタレシーバーと値レシーバー
- メソッドで紐づけれた構造体の値を書き換えたい場合は、メソッド作成時に**レシーバーに\*をつける**
  ```go
  package main

  import "fmt"

  type Vertex struct {
    X, Y int
  }

  func (v Vertex) Area() int {
    return v.X * v.Y
  }

  func (v *Vertex) Scale(i int) {
    v.X = v.X * i
    v.Y = v.Y * i
  }

  func Area(v Vertex) int {
    return v.X * v.Y
  }

  func main() {
    v := Vertex{3, 4}
    v.Scale(10)
    fmt.Println(v.Area())
  }

  ```

### 4-1-3 Newで初期化時の処理を実行しよう
- 初期化時に実行される処理(コンストラクタ)はGoでは**New**という関数を作成して行う
- 実際にNew関数を作成しみよう
  - VertexのX,Yを小文字にする
    - 他のパッケージから操作することはできず、このパッケージ内からのみ書き換えできるようになる。
  - New関数を作成する
    - 引数にint型のxとy、戻り値に\*Vertex(Vertexのポインタ)を設定する

      ```go
      package main

      import "fmt"

      type Vertex struct {
        x, y int
      }

      func (v Vertex) Area() int {
        return v.x * v.y
      }

      func (v *Vertex) Scale(i int) {
        v.x = v.x * i
        v.y = v.y * i
      }

      func Area(v Vertex) int {
        return v.x * v.y
      }

      func New(x, y int) *Vertex {
        return &Vertex{x, y}
      }

      func main() {
        v := New(3, 4)
        v.Scale(10)
        fmt.Println(v.Area())
      }

      ```

### 4-1-4 構造体以外の型のメソッド
- typeを使うと、組み込み型に新しい名前をつけた独自の型を作成することが可能
- その独自の型にメソッドを持たすことができる
  - 例) int型にMyIntという別の名前を付けて新しい型を作ってみる。
  
    ```go
    package main

    import "fmt"

    type MyInt int

    func (i MyInt) Double() int {
      return int(i * 2)
    }

    func main() {
      myInt := MyInt(10)
      fmt.Println(myInt.Double())
    }

    ```
  - **構造体以外の型のメソッドの注意点**
    - 1つ注意点として、「return int(i * 2)」と返り値をint型にcastしていますが、castしないとエラーがでます。

      ```go
      package main

      import "fmt"

      type MyInt int

      func (i MyInt) Double() int {
        fmt.Printf("%T %v\n", i, i)
        fmt.Printf("%T %v\n", 1, 1)
        return int(i * 2)
      }

      func main() {
        myInt := MyInt(10)
        fmt.Println(myInt.Double())
      }

      ```

## 4-2 構造体の埋め込みをしよう
- Goでは、構造体の中に構造体を持たせることで、オブジェクト指向プログラミングにおける継承のようなことができる
- ここでは、構造体を埋め込む方法を他のプログラミング言語の例と比較しつつ説明しておく。

### 4-2-1 構造体の中に構造体を埋め込もう
- Goの**埋め込み(Embedded)**という仕組みについて説明する。
- これは他のプログラミング言語では、**継承**などと呼ばれる処理にあたる。
- **Goで構造体を埋め込む**
  - 構造体の中に構造体の中に埋め込む

    ```go
    package main

    import "fmt"

    type Vertex struct {
      x, y int
    }

    func (v Vertex) Area() int {
      return v.x * v.y
    }

    func (v *Vertex) Scale(i int) {
      v.x = v.x * i
      v.y = v.y * i
    }

    type Vertex3D struct {
      Vertex
      z int
    }

    func (v Vertex3D) Area3D() int {
      return v.x * v.y * v.z
    }

    func (v *Vertex3D) Scale3D(i int) {
      v.x = v.x * i
      v.y = v.y * i
      v.z = v.z * i
    }

    func New(x, y, z int) *Vertex3D {
      return &Vertex3D{Vertex{x, y}, z}
    }

    func main() {
      v := New(3, 4, 5)
      v.Scale(10)
      fmt.Println(v.Area())
      fmt.Println(v.Area3D())
    }

    ```

## 4-3 インターフェースを使ったプログラムをつくろ
- Goのインターフェースは、メソッドの名前のみを宣言したもので、そのメソッドを持つ型はインターフェースを実装していると判定される。
- Goのインターフェースの使い方について説明しつつ、ダックタイピングについても解説していく。

### 4-3-1 インターフェースを作成しよう
- Goにおける**インターフェース**を説明していく。Humanというインターフェースを作成し、{}の中に「Say()」メソッドを書く
- インターフェースでは、**メソッド名のみを宣言し、処理のコードを書かない**
- 構造体を作成してインターフェースに当てはめる
  - string型のNameというフィールドを持つPersonという構造体を作成する。
  - Personに紐づくメソッドとしてSayを作成する
- Human型(インターフェース型)の変数mikeを宣言し、構造体Personに{"Mike"}を代入する。
- 「mike.Say()」と実行すると、「Mike」と表示される。

  ```go
  package main

  import "fmt"

  type Human interface {
    Say()
  }

  type Person struct {
    Name string
  }

  func (p Person) Say() {
    fmt.Println(p.Name)
  }

  func main() {
    var mike Human = Person{"Mike"}
    mike.Say()
  }

  ```
- **インターフェースのメソッドで構造体の中身を書き換える場合**
  - Sayメソッドの処理に追加。p.Nameに「Mr.」を加える。
  - 構造体の中身を書き換えることになるので、Personの前に\*をつけてポインタレシーバーする必要がある。
  - メソッドの変更のみだとエラーとなる。
  - Sayメソッドはポインタレシーバーとなるので、main関数から呼び出すSayメソッドを呼びだすさいに、アドレスとして渡す必要がある。

    ```go
    package main

    import "fmt"

    type Human interface {
      Say()
    }

    type Person struct {
      Name string
    }

    func (p *Person) Say() {
      p.Name = "Mr." + p.Name
      fmt.Println(p.Name)
    }

    func main() {
      var mike Human = &Person{"Mike"}
      mike.Say()
    }

    ```

### 4-3-2 ダックタイピング
- Humanというインターフェースを、引数として受け付ける関数DriveCarを作っていく
- 引数humanのSayメソッドは、戻り値が「Mr.Mike」であれば「Run」そうでなければ「Get Out」を表示する。

  ```go
  package main

  import "fmt"

  type Human interface {
    Say() string
  }

  type Person struct {
    Name string
  }

  func (p *Person) Say() string {
    p.Name = "Mr." + p.Name
    fmt.Println(p.Name)
    return p.Name
  }

  func DriveCar(human Human) {
    if human.Say() == "Mr.Mike" {
      fmt.Println("Run")
    } else {
      fmt.Println("Get out")
    }
  }

  func main() {
    var mike Human = &Person{"Mike"}
    var x Human = &Person{"X"}
    DriveCar(mike)
    DriveCar(x)
  }

  ```

## 4-4 型アサーションとswitch typeを使う
- Goではメソッドを持たない空のインターフェースを作成できる。
- 空のインターフェースには、どのような型の値でも入れることができるので、引数にどんな型が入るのか分からないときに活用できる
- そして、空のインターフェースに入れた値を特定の値として使うときに必要なのが型アサーションです。
- ここでは、型アサーションの方法と、switch typeと呼ばれる活用方法について記載する

### 4-4-1 型アサーションについて学ぼう
- 最初に、iというインターフェースを引数にするdoという関数を作る
  - 引数は「interface{}」**空のインターフェース**。どのような型でも引数として受け付ける
- do関数で引数を2倍にして変数iiに代入、初期化し変数iiの値を表示してみよう
  ```go
  func do(i interface{}) {
    ii := i * 2 // invalid operation: i * 2 (mismatched types interface{} and untyped int)
    fmt.Println(ii)
  }

  func main() {
    do(10)
  }

  ```
- 空のインターフェースが持つ値をintやstringといった具体的な型として扱うため、**型アサーション(Type Assertion)**という仕組みを使う
- 「ii := i.(int)」とすることで、iの値をint型として扱い、変数iiに代入する。
  ```go
  func do(i interface{}) {
    ii := i.(int)
    ii *= 2
    fmt.Println(ii)
  }

  func main() {
    do(10)
  }

  ```
- **文字列への型アサーション**
  - 文字列への型アサーションも確認してみよう
  - main関数内で実行するdo関数の引数「"Mike"」とし、do関数内で引数をstring型に型アサーションして変数ssに代入します。文字列に「!」を足して、表示を確認してみよう

    ```go
    package main

    import "fmt"

    func do(i interface{}) {
      ss := i.(string)
      fmt.Println(ss + "!")
    }

    func main() {
      do("Mike")
    }

    ```

### 4-4-2 switch typeで型ごとに処理を実行しよう
- 型アサーションを用いるコードは、int型を処理したいか、文字列を処理したいかわかりにくくなる
- 異なる型に対応できるようにコードを書き換えてみよう
- switch文で、実行結果に「v := i.(type)」と書きiをtype型アサーションした結果をvに代入する
- なお、i.(type)という書き方は、switchと一緒でなければ使えない。そのため、i.(type)を単独で書くとエラーになる。switchとtypeはセットであると覚える。

  ```go
  package main

  import "fmt"

  func do(i interface{}) {
    switch v := i.(type) {
    case int:
      fmt.Println(v * 2)
    case string:
      fmt.Println(v + "!")
    default:
      fmt.Printf("I don't know %T\n", v)
    }
  }

  func main() {
    do(10)
    do("Mike")
    do(true)
  }

  ```

## 4-5 Stringerで表示内容を変更しよう
- Stringerは、fmtパッケージに含まれるインターフェースです。
- このインターフェースにあるStringメソッドを実装すると、fmt.Printlnなどによる表示が変更される。

### 4-5-1 Stringerインターフェースを利用しよう
- **Stringer**はfmtパッケージのprint.goに書かれているインターフェースです。
- Stringerインターフェースは、Stringメソッドを持っている。
- Stringerインターフェースを実装した構造体をfmt.Println関数の引数とするとStringメソッドの返り値が出力される。

- **Stringerで表示内容を変更する**
  - NameとAgeを持った構造体Personを作成する
  - main関数内で初期化して変数mikeに代入し「fmt.Println(mike)」と実行すると「Mike 22」が表示される。

    ```go
    package main

    import "fmt"

    type Person struct {
      Name string
      Age  int
    }

    func main() {
      mike := Person{"Mike", 22}
      fmt.Println(mike)
    }

    ```
  - Personに紐づいたStringメソッドを作り返り値として「return fmt.Sprintf("My name is %v.", p.name)」とする

    ```go
    package main

    import "fmt"

    type Person struct {
      Name string
      Age  int
    }

    func (p Person) String() string {
      return fmt.Sprintf("My name is %v.", p.Name)
    }

    func main() {
      mike := Person{"Mike", 22}
      fmt.Println(mike)
    }

  ```
  
## 4-6 カスタムエラーを作成しよう