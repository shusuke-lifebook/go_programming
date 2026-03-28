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

## 5-2 チャネルでゴルーチンと値のやりとりをしよう


## 5-3 2つのゴルーチンで値を送受信しよう

## 5-4 pipelineによる並行処理

## 5-5 selectでチャネルに応じた処理をしよう

## 5-6 selectでdefaultとbreakを使おう

## 5-7 sync.Mutexを使ったゴルーチンの処理