# Lesson3 ポインタ
- 変数以外にも、ポインタと呼ばれるものを使って、値を管理することがある。
- ポインタを使った基本的なプログラムについて学習する

## 3-1 ポインタを操作しよう
- 変数に値を格納すると、コンピュータはメモリの特定の位置(アドレス)を指定して、そこに値を保持する
- C言語など、プログラミング言語によっては、値を格納している場所(メモリ上のアドレス)を指すポインタを扱うことができる

### 3-1-1 ポインタでメモリ上のアドレスを参照する
- Goにおける**ポインタ**は変数に **&** をつけることでメモリ上のアドレスを確認できる。
- ポインタ変数は「var p \*int」のように「\*int」のように表現する。
- アドレスが指すメモリアドレスの中身を確認したいときは、ポインタ変数に\*をつける

  ```go
  package main

  import "fmt"

  func main() {
    var n int = 100
    fmt.Println(n)
    fmt.Println(&n)

    var p *int = &n
    fmt.Println(p)
    fmt.Println(*p)
  }

  ```

### 3-1-2 関数でポインタを受け取る
- 関数の引数がポインタを渡せるように\*をつける。例) 「x \*int」
- 関数内でポインタの実体を更新する\*x = 1のようにする
  ```go
  package main

  import "fmt"

  func one(x *int) {
    *x = 1
  }

  func main() {
    var n int = 100
    one(&n)
    fmt.Println(n)
  }

  ```

### 3-1-3 変数のアドレスと中身を表示する
- &と\*を使って変数のアドレスを表示したり、アドレスが指す実体を表示させてみよう
  - 変数nのアドレスを表示させるためには、「&n」と書く
  - &nのアドレスの中身を確認するには、「\*&n」と書く
    ```go
    package main

    import "fmt"

    func one(x *int) {
      *x = 1
    }

    func main() {
      var n int = 100
      one(&n)
      fmt.Println(n)
      fmt.Println(&n)
    }

    ```

## 3-2 new 関数 と make 関数の違い
- new関数はメモリを確保するための関数でポインタなどを作成する際に必要になる

### 3-2-1 newを使ってポインタのアドレスを確保する
- **new**関数を利用すると値を何も入れない状態で、メモリにポインタが入る領域を確保することができる
  - 例) 「var p *int = new(int)」
    ```go
    package main

    import "fmt"

    func main() {
      var p *int = new(int)
      fmt.Println(p)
    }

    ```
  - new関数を使わずにポインタ宣言すると<nil>となる
    ```go
    package main

    import "fmt"

    func main() {
      var p *int = new(int)
      fmt.Println(p)

      var p2 *int
      fmt.Println(p2)
    }

    ```

### 3-2-2 new関数とmake関数の違い
- new関数と似たものに、mapとスライスを作成するときに使うmake関数がある。
- コードを見てnew関数とmake関数の違いを確認しよう
  - **ポインタを返す場合は、new関数**
  - **そうでない場合は、make関数**
  ```go
  package main

  import "fmt"

  func main() {
    s := make([]int, 0)
    fmt.Printf("%T\n", s)

    m := make(map[string]int)
    fmt.Printf("%T\n", m)

    var p *int = new(int)
    fmt.Printf("%T\n", p)
  }

  ```

## 3-3 構造体で複数の値をまとめて扱う
- struct(構造体)は、複数の値をまとめたものです。
- Goには他の言語におけるクラスにあたるものがないが、構造体とメソッドを組み合わせることでオブジェクトのように扱うことができる

### 3-3-1 struct(構造体)
- **struct**(構造体)。まず、main関数の外でVertexという名前の構造体を作成し、{}の中に2つのint型の値X,Yを書く
- main関数で構造体Vertexを変数vに代入する。このとき、「Vertex{X:1, Y:2}」のように書いて、構造体Vertexを初期化する。
  ```go
  package main

  import "fmt"

  type Vertex struct {
    X int
    Y int
  }

  func main() {
    v := Vertex{X: 1, Y: 2}
    fmt.Println(v)
  }

- 構造体の中身の値を確認するには、「v.X」のように書く。
  ```go
  func main() {
    v := Vertex{X: 1, Y: 2}
    fmt.Println(v)
    fmt.Println(v.X, v.Y)
  }

  ```
- 構造体の中身を書き換える場合、v.X = 100のように書く
  ```go
  func main() {
    v := Vertex{X: 1, Y: 2}
    fmt.Println(v)
    fmt.Println(v.X, v.Y)
    v.X = 100
    fmt.Println(v.X, v.Y)
  }

  ```
- 構造体の一部だけを初期化して宣言することもできる。Xだけを初期化すると、Yの値はint型の初期値である0になる。
  ```go
  func main() {
    v2 := Vertex{X: 1}
    fmt.Println(v2)
  }
  ```
- string型の場合、初期値は空文字になる。
  ```go
  type Vertex struct {
    X int
    Y int
    S string
  }

  func main() {
    v2 := Vertex{X: 1}
    fmt.Println(v2)
  }

  ```
- フィールドを指定せずに初期する場合、「Vertex{1, 2, "test"}」のように、構造体に書いた順番に書く
- 