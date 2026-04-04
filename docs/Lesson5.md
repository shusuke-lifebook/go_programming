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
- 並行処理のプログラムの作り方はさまざまですが、その中の一つにpipeline(パイプライン)というものがある。
- Goでは、複数のチャネルを用意し、順番に値を渡していくような並行処理をpiplineという。

### 5-4-1 pipelineを使って並行処理をしよう
- 並行処理の方法の一つである**pipeline**パターン。
- main関数からゴルーチンを1つ立ち上げ、最初に立ち上げたゴルーチンが処理する。処理した結果をチャネルに渡して次のゴルーチンで処理する。これを繰り返して最終結果をmain関数に渡す。
  ```go
  package main

  import "fmt"

  func producer(first chan int) {
    defer close(first)
    for i := 0; i < 10; i++ {
      first <- i
    }
  }

  func multi2(first chan int, second chan int) {
    defer close(second)
    for i := range first {
      second <- i * 2
    }
  }

  func multi4(second chan int, third chan int) {
    defer close(third)
    for i := range second {
      third <- i * 4
    }
  }

  func main() {
    first := make(chan int)
    second := make(chan int)
    third := make(chan int)

    go producer(first)
    go multi2(first, second)
    go multi4(second, third)
    for result := range third {
      fmt.Println(result)
    }
  }

  ```

## 5-5 selectでチャネルに応じた処理をしよう
- 複数のチャネルを使って複数のゴルーチンとやりとりするとき受信したチャネルによって処理を分岐させる場合がある。
- その場合、selectを使うことでチャネルごとの処理を書くことができる
- ここでは、selectを用いたプログラムがどのようなものかについて記述する

### 5-5-1 selectの使い方を学ぼう
- 複数のゴルーチンがありますが、それぞれ別チャネルでデータを受信する。
- 例えば、複数のゴルーチンからネットワークのパケットを受信するようなイメージです。
- このとき、それぞれの処理をブロッキングしないように、**select**を使う

  ```go
  package main

  import (
    "fmt"
    "time"
  )

  func goroutine1(ch chan<- string) {
    for {
      ch <- "packet from 1"
      time.Sleep(1 * time.Second)
    }
  }

  func goroutine2(ch chan<- string) {
    for {
      ch <- "packet from 2"
      time.Sleep(1 * time.Second)
    }
  }

  func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go goroutine1(c1)
    go goroutine2(c2)

    for {
      select {
      case msg1 := <-c1:
        fmt.Println(msg1)
      case msg2 := <-c2:
        fmt.Println(msg2)
      }
    }
  }

  ```

## 5-6 selectでdefaultとbreakを使おう
- selectでは、チャネルに応じてそれぞれの処理を作る。
- どのチャネルでもないときの処理はdefaultを使って書くことができる
- selectを使った処理におけるdefaultの使用例を記述する
- また、for文とselect文を使うときにラベルを使って途中で処理を抜ける方法について記載する

### 5-6-1 selectとdefaultでどのチャネルでもない処理を書こう
- selectで**default**を使うと、どのチャネルでもないときに実行したい処理を書ける
- ここでは、「A Tour of Go」というGoの公式チュートリアルで公開されているコードを例に記載する。
  - **time.Tick**と**time.After**は、それぞれ設定した時間に値が送信されるチャネルを返す。
    - time.Tickは設定した時間ごとに周期的にチャネルを送信する
    - time.Afterは設定した時間が経過したタイミングで値を送信する

  ```go
  package main

  import (
    "fmt"
    "time"
  )

  func main() {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    for {
      select {
      case <-tick:
        fmt.Println("tick.")
      case <-boom:
        fmt.Println("BOOM!")
        return
      default:
        fmt.Println("    .")
        time.Sleep(50 * time.Millisecond)
      }
    }
  }

  ```

### 5-6-2 for文とselect文を途中で抜けよう
- 例えば、for文の外に、「fmt.Println("##########")」と書いて実行しても先ほどの表示結果は変わらない
- これは、**returnの時点で処理が終了**してしまい、そのあとの「fmt.Println("##########")」は実行されない

  ```go
  package main

  import (
    "fmt"
    "time"
  )

  func main() {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    for {
      select {
      case <-tick:
        fmt.Println("tick.")
      case <-boom:
        fmt.Println("BOOM!")
        return
      default:
        fmt.Println("    .")
        time.Sleep(50 * time.Millisecond)
      }
    }
    fmt.Println("##########")
  }

  ```
- **returnの代わりにbreakを用いる**
  - selectの中でreturnの代わりにbreakを書いた場合、for文から抜け出せず無限ループになってしまうので注意が必要 

- **ラベルを使用してforループを抜ける**
  ```go
  package main

  import (
    "fmt"
    "time"
  )

  func main() {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
  OuterLoop:
    for {
      select {
      case <-tick:
        fmt.Println("tick.")
      case <-boom:
        fmt.Println("BOOM!")
        break OuterLoop
      default:
        fmt.Println("    .")
        time.Sleep(50 * time.Millisecond)
      }
    }
    fmt.Println("##########")
  }

  ```

## 5-7 sync.Mutexを使ったゴルーチンの処理
- 複数のゴルーチン間でのデータのやりとりにはチャネルを使ってきました。
- マップなどの値を異なるゴルーチンから読み込んだり、書き込んだりする場合、場合によっては同じタイミングで読み書きをしてしまうためエラーが起こりやすくなる。
- そのため、同時に読み書きが起こらないようにプログラムを調整する必要がある
- ここでは、sync.Mutexを使ってことなるゴルーチンから値を読み書きするための方法について記載する。

### 5-7-1 異なるゴルーチンから同じ構造体を書き換えよう

  ```go
  package main

  import (
    "fmt"
    "time"
  )

  func main() {
    c := make(map[string]int)
    go func() {
      for i := 0; i < 10; i++ {
        c["key"] += 1
      }
    }()
    go func() {
      for i := 0; i < 10; i++ {
        c["key"] += 1
      }
    }()

    time.Sleep(1 * time.Second)
    fmt.Println(c, c["key"])
  }

  ```
- **sync.Mutex**
  - 2つのゴルーチンから1つのマップを読み込んだり書き換えたりしようとすると、問題が起きるため、**sync.Mutex**を使う必要がある

    ```go
    package main

    import (
      "fmt"
      "sync"
      "time"
    )

    type Counter struct {
      v   map[string]int
      mux sync.Mutex
    }

    func (c *Counter) Inc(key string) {
      c.mux.Lock()
      defer c.mux.Unlock()
      c.v[key]++
    }

    func (c *Counter) Value(key string) int {
      c.mux.Lock()
      defer c.mux.Unlock()
      return c.v[key]
    }

    func main() {
      c := Counter{v: make(map[string]int)}
      go func() {
        for i := 0; i < 10; i++ {
          c.Inc("Key")
        }
      }()
      go func() {
        for i := 0; i < 10; i++ {
          c.Inc("key")
        }
      }()
      time.Sleep(1 * time.Second)
      fmt.Println(c.v, c.Value("Key"))
    }

    ```