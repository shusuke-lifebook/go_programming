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
- 変数の代入とif文の条件を1行に纏めて書く場合、**if文の後に変数宣言し、;(セミコロン)で繋げて条件を書く
  ```go
  package main

  import "fmt"

  func by2(num int) string {
    if num%2 == 0 {
      return "ok"
    } else {
      return "no"
    }
  }

  func main() {
    result := by2(10)
    if result == "ok" {
      fmt.Println("great")
    }

    if result2 := by2(10); result2 == "ok" {
      fmt.Println("great2")
    }
  }

  ```

## 2-2 for文で処理を繰り返し実行しよう
- Goでの繰り返し処理はfor文を使う。
  - 次の繰り返しに進むcontinue
  - 繰り返しを抜けるbreak
- 他のプログラム言語と同じ

### 2-2-1 for文による繰り返し処理を作ろう
- Goでの**for**文は、forの後に「繰り返しで使う変数の初期化」「繰り返しを続ける条件」「繰り返すたびに実行する処理」を;(セミコロン)で区切って書く
  ```go
  package main

  import "fmt"

  func main() {
    for i := 0; i < 10; i++ {
      fmt.Println(i)
    }
  }

  ```

### 2-2-2 continue文で次の繰り返し進む処理を作ろう
- **continue**文は、処理の途中で次の繰り返し処理に進むときに使う
  ```go
  package main

  import "fmt"

  func main() {
    for i := 0; i < 10; i++ {
      if i == 3 {
        fmt.Println("continue")
        continue
      }
      fmt.Println(i)
    }
  }

  ```

### 2-2-3 break文で繰り返しを途中で抜けよう
- **break**文は、繰り返しを中断してfor文から抜けたいときにしようする
  ```go
  package main

  import "fmt"

  func main() {
    for i := 0; i < 10; i++ {
      if i == 3 {
        fmt.Println("continue")
        continue
      }

      if i > 5 {
        fmt.Println("break")
        break
      }
      fmt.Println(i)
    }
  }

  ```

### 2-2-4 for文の省略記法を使おう
- for文の1つ目の「変数の初期化」と3つ目の「繰り返すたびに実行する処理」は省略可能
  ```go
  package main

  import "fmt"

  func main() {
    sum := 1
    for sum < 10 {
      sum += sum
      fmt.Println(sum)
    }
    fmt.Println(sum)
  }

  ```

### 2-2-5 rangeで繰り返し処理を簡単に書こう
- for文と一緒に使うと便利なのが、**range**です。
  - スライスの例
  ```go
  package main

  import "fmt"

  func main() {
    l := []string{"python", "go", "java"}
    for i, v := range l {
      fmt.Println(i, v)
    }
  }

  ```
- マップの例
  ```go
  package main

  import "fmt"

  func main() {
    m := map[string]int{"apple": 100, "banana": 200}

    for k, v := range m {
      fmt.Println(k, v)
    }
  }

  ```

## 2-3 switch文で条件に応じた処理を実行しよう
- 条件に応じて別の処理を実行するには、if文を用いるほかにswitch文を使う方法がある。

### 2-3-1 switch文で条件ごとの処理を作ろう
- **switch文**で変数osの値を判定していく例。
  ```go
  package main

  import "fmt"

  func main() {
    os := "mac"
    switch os {
    case "mac":
      fmt.Println("Mac!!")
    case "windows":
      fmt.Println("Windows!!")
    default:
      fmt.Println("Default!!")
    }
  }

  ```
    
### 2-3-2 変数宣言とswitch文を纏めて書こう
- if文と同様に、**変数への代入とswitch文は、1行でまとめて書ける**
  ```go
  package main

  import "fmt"

  func getOsName() string {
    return "mac"
  }

  func main() {
    switch os := getOsName(); os {
    case "mac":
      fmt.Println("Mac!!")
    case "windows":
      fmt.Println("Windows!!")
    default:
      fmt.Println("Default!!")
    }
  }

  ```

### 2-3-3 条件式を書かないswitch文を作ろう
- **switch文に条件式を書かず、caseに判定条件を書く**
  ```go
  package main

  import (
    "fmt"
    "time"
  )

  func main() {
    t := time.Now()
    fmt.Println(t.Hour())
    switch {
    case t.Hour() < 12:
      fmt.Println("Morning")
    case t.Hour() < 17:
      fmt.Println("Afternoon")
    }
  }

  ```

## 2-4 defer文で処理を遅らせて実行しよう

## 2-5 ログを出力しよう

## 2-6 エラーハンドリングをしよう

## 2-7 panicとrecover