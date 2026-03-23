# Lesson2 ステートメント
条件分岐のif文や繰り返し処理のfor文など、Goの基本的なステートメントについて説明する。

## 2-1 if文で条件分岐の処理を実行しよう
- Goでの条件分岐を書く際には「if」、「else」、「else if」を使います。

### 2-1-1 if文で条件分岐のプログラムを作ろう
- Goでの**if文**は、「if 条件式 {}」と書いて{}に条件に当てはまった時の処理を書く
  ```go
  package main

  import "fmt"

  func main() {
    num := 4
    if num%2 == 0 {
      fmt.Println("by 2")
    }
  }

  ```
- else節はif文の条件式に当てはまらない場合の処理を記述する
  ```go
  package main

  import "fmt"

  func main() {
    num := 5
    if num%2 == 0 {
      fmt.Println("by 2")
    } else {
      fmt.Println("else")
    }
  }

  ```
- else if節、if文とelseの間に書き、条件式と処理を書く
  ```go
  package main

  import "fmt"

  func main() {
    num := 9
    if num%2 == 0 {
      fmt.Println("by 2")
    } else if num%3 == 0 {
      fmt.Print("by 3")
    } else {
      fmt.Print("else")
    }
  }

  ```
### 2-1-2 複数の条件がある場合
- 複数の条件がある場合は**論理演算子**を使う。複数の条件にすべて当てはまるかを判断したいかは条件式を&&でつなげる。
  ```go
  package main

  import "fmt"

  func main() {
    x, y := 10, 10
    if x == 10 && y == 10 {
      fmt.Println("&&")
    }
  }

  ```
- 複数ある条件のうちいずれかに当てはまるかを判別したい場合は、条件式を||でつなげる。
  ```go
  package main

  import "fmt"

  func main() {
    x, y := 10, 10
    if x == 10 && y == 10 {
      fmt.Println("&&")
    }

    if x == 10 || y == 10 {
      fmt.Println("||")
    }
  }

  ```

### 2-1-3 if文の条件式を変数で宣言しよう

## 2-2 for文で処理を繰り返し実行しよう

## 2-3 switch文で条件に応じた処理を実行しよう

## 2-4 defer文で処理を遅らせて実行しよう

## 2-5 ログを出力しよう

## 2-6 エラーハンドリングをしよう

## 2-7 panicとrecover