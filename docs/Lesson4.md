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

## 4-3 インターフェースを使ったプログラムをつくろ

## 4-4 型アサーションとswitch typeを使う

## 4-5 Stringerで表示内容を変更しよう

## 4-6 カスタムエラーを作成しよう