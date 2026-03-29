# Lesson 5 ゴルーチン
- 並行処理のプログラムを簡単に書けることもGoの特徴の1つです。
- Goの並行処理で重要なゴルーチンについて基本的な使い方や例を記載する。

## 5-1 並行処理を作ろう
- ゴルーチンは、Goにおいて並行処理を行うための軽量のスレッド
- 並行処理のプログラムは、他の言語ではマルチプロセスやマルチスレッド、イベント駆動などとも呼ばれる
- Goでは並行処理を暗黙的に実行してくれるので、他言語のように深い知識を必要せず、並行処理のプログラムを書けるという特徴がある。

### 5-1-1 goroutine(ゴルーチン)で並行処理を実行しよう
- **goroutine**で並行処理を行うコードを書きましょう。string型の引数sを持つnormal関数を作成する
- for文でiが0～5になるまでループし、0.1秒のSleepと引数sの表示を行う
  ```go
  package main

  import (
    "fmt"
    "time"
  )

  func normal(s string) {
    for i := 0; i < 5; i++ {
      time.Sleep(100 * time.Millisecond)
      fmt.Println(s)
    }
  }

  func main() {
    normal("Hello")
  }

  ```
- 上記は、通常の順次処理なので、これを並行処理にするための準備していく
- normal関数をコピーして、関数名をgoroutineにした同じ処理の関数を作る
  ```go
  func goroutine(s string) {
    for i := 0; i < 5; i++ {
      time.Sleep(100 * time.Millisecond)
      fmt.Println(s)
    }
  }

  func main() {
    goroutine("world")
    normal("Hello")
  }

  ```
- 上記コードは、まだ並行処理になっていない。呼び出す関数の前に**go**をつけると並行処理になる
  ```go
  func main() {
    go goroutine("world")
    normal("Hello")
  }

  ```
- **ゴルーチンの実行前にプログラムが終了する場合**
  - 「go goroutine("world")」で生成された並行処理が実行される前にnormal関数が終わってしまう。
  - **ゴルーチンの処理が終わらなくてもプログラムが終了してしまう。**

### 5-1-2 sync.WaitGroupで並行処理を待機させよう 
- プログラムが途中終了することを避けるには、**sync.WaitGroup**を使う
  ```go
  package main

  import (
    "fmt"
    "sync"
  )

  func normal(s string) {
    for i := 0; i < 5; i++ {
      // time.Sleep(100 * time.Millisecond)
      fmt.Println(s)
    }
  }

  func goroutine(s string, wg *sync.WaitGroup) {
    for i := 0; i < 5; i++ {
      // time.Sleep(100 * time.Millisecond)
      fmt.Println(s)
    }
    wg.Done()
  }

  func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go goroutine("world", &wg)
    normal("Hello")
    wg.Wait()
  }

  ```
- **WaitGroupの注意点**
  - 「wg.Done()」をコメントアウトして実行すると、途中で止まってしまいエラーになる
  - VSCodeでデバックが途中で止まってしまった場合はF5キーなどを押して処理を進める

## 5-2 チャネルでゴルーチンと値のやりとりをしよう
- 並行で実行したゴルーチンと値の受け渡しをしたい場合、関数やメソッドのようにreturnで値を返すという方法でやりとりすることはできない
- ゴルーチンの間では代わりにChannel(チャネル)を使ってデータを受け渡す。

### 5-2-1 チャネルを使って値をやりとりしよう
- ゴルーチン間のデータのやりとりに用いる**Channel**について記載する。
- **ゴルーチン間でデータを受け渡すためには、チャネルを使う必要がある**
- main関数もゴルーチンによって動ている。➡ main関数を動かすゴルーチンを**メインゴルーチン**という
- **チャネルを使ったデータのやりとり**
  - **make関数**でチャネルを作る ➡ 「c := make(chan int)」とするとint型のチャネル変数cに代入される。
  - チャネルでゴルーチンから値を受けとって変数xに格納するために **<-演算子**を使って「x := <-c」と書く
  - 「goroutine1」関数内で、最後に、<-演算子を利用して「c <- sum」と書くことで、チャネル送信する。

    ```go
    package main

    import "fmt"

    func goroutine1(s []int, c chan int) {
      sum := 0
      for _, v := range s {
        sum += v
      }
      c <- sum
    }

    func main() {
      s := []int{1, 2, 3, 4, 5}
      c := make(chan int)
      go goroutine1(s, c)
      x := <-c
      fmt.Println(x)
    }

    ```
  - **同じ関数のゴルーチンを複数回呼び出す際のチャネル**
    - 1つの処理を何回か実行したい場合に使う
    - 今回はint型のチャネルを1つ作ったが、1つはint型のチャネル。もう1つはstring型のチャネルといったように異なるチャネルを複数作る場合もある。

    ```go
    package main

    import "fmt"

    func goroutine1(s []int, c chan int) {
      sum := 0
      for _, v := range s {
        sum += v
      }
      c <- sum
    }

    func main() {
      s := []int{1, 2, 3, 4, 5}
      c := make(chan int)
      go goroutine1(s, c)
      go goroutine1(s, c)
      x := <-c
      fmt.Println(x)
      y := <-c
      fmt.Println(y)
    }

    ```

### 5-2-2 バッファありチャネルを使って値をやりとりしよう
- 上記までのチャネルは**unbuffered channel(バッファなしチャネル)**といい、**バッファ(チャネルに入る値の数)を指定せずに作ったチャネルです。
- バッファを指定したチャネルが**buffered channle(バッファありチャネル)**。
  - 例)「ch := make(chan int, 2)」と書いてバッファを2と指定している。

    ```go
    // バッファ2つに3つ目までをチャネルに入れようとするとエラーとなる
    package main

    import "fmt"

    func main() {
      ch := make(chan int, 2)
      ch <- 100
      fmt.Println(len(ch))
      ch <- 200
      fmt.Println(len(ch))
      ch <- 300
      fmt.Println(len(ch))
    }

    ```
  - エラーを回避するには、「x := <-ch」で値を受信して値を取り出す必要がある。

    ```go
    package main

    import "fmt"

    func main() {
      ch := make(chan int, 2)
      ch <- 100
      fmt.Println(len(ch))
      ch <- 200
      fmt.Println(len(ch))

      x := <-ch
      fmt.Println(x)

      ch <- 300
      fmt.Println(len(ch))
    }

    ```

### 5-2-3 rangeとcloseでチャネルから値を取り出そう
- goroutine1関数ですべてのスライスの中身を合計してからチャネルに渡している
- これを合計していく途中過程もチャネルに送信して表示してみよう。
  - rangeはチャネルから値を送られてくるのを待ち続けるので、これ以上送信しないことを伝えるために**close**をする必要がある。

    ```go
    package main

    import "fmt"

    func goroutine1(s []int, c chan int) {
      sum := 0
      for _, v := range s {
        sum += v
        c <- sum
      }
      close(c)

    }

    func main() {
      s := []int{1, 2, 3, 4, 5}
      c := make(chan int)
      go goroutine1(s, c)
      for i := range c {
        fmt.Println(i)
      }
    }

    ```
  - **バッファありチャネルからrangeで値を取り出す**

    ```go
    package main

    import "fmt"

    func main() {
      ch := make(chan int, 2)
      ch <- 100
      fmt.Println(len(ch))
      ch <- 200
      fmt.Println(len(ch))
      close(ch)

      for c := range ch {
        fmt.Println(c)
      }
    }

    ```

## 5-3 2つのゴルーチンで値を送受信しよう
- プログラムにおいて
  - 値を送信する関数などの処理をProducerと呼ぶ
  - 受信する処理をConsumerと呼ぶ
- Goでは複数のゴルーチンにProducerとConsumerの役割を持たせた場合の並行処理のプログラムについて記載する。

### 5-3-1 ProducerとConsumerのゴルーチンを作ろう
- **Producer**と**Consumer**という2つの役割持つゴルーチンを作っていく。
  - いろいろなサーバーからログ解析結果をProducer側で取得
  - Consumer側に渡してログの処理や保存するようなアプリケーションなどのイメージ
- **Producerの処理**
  - チャネルを通じて値をconsumer関数のゴルーチンの送るproducer関数を作る
- **Consumerの処理**
  - producer関数のゴルーチンからチャネルを受け取って値を処理するconsumer関数を作る
  - チャネルの値はrangeで取り出す処理をする
  - producer関数から渡された値の処理が終わったということを「wg.Done()」を実行する
- **main関数の処理**
  - main関数では、まず、sync.WaitGroupを宣言する。
  - チャネルを ch := make(chan int)と書いて作る
  - producer関数をfor文で10回繰り返し
    - wg.Add(1)
    - producer関数を呼び出す
  - 最後にcosumer関数を呼び出す。
    - wg.Wait()
    - close(ch)
      ```go
      package main

      import (
        "fmt"
        "sync"
      )

      func producer(ch chan int, i int) {
        ch <- i * 2
      }

      func consumer(ch chan int, wg *sync.WaitGroup) {
        for i := range ch {
          fmt.Println("process", i*1000)
          wg.Done()
        }
      }

      func main() {
        var wg sync.WaitGroup
        ch := make(chan int)

        // Producer
        for i := 0; i < 10; i++ {
          wg.Add(1)
          go producer(ch, i)
        }

        // Consumer
        go consumer(ch, &wg)
        wg.Wait()
        close(ch)
      }

      ```

## 5-4 pipelineによる並行処理

## 5-5 selectでチャネルに応じた処理をしよう

## 5-6 selectでdefaultとbreakを使おう

## 5-7 sync.Mutexを使ったゴルーチンの処理