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