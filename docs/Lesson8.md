# Lesson 8 MVCモデル
- MVCモデルを例にとって説明していく。

## 8-1 MVCモデル
- MVCモデルは、Webアプリケーションの開発によく使われるアーキテクチャです。
- コードを以下の3つに分けアプリケーションを作成する
  - Model(モデル)
  - View(ビュー)
  - Controller(コントローラ)

### 8-1-1 アーキテクチャとは
- **アーキテクチャ(Architecture)**は「構造」という意味があり、アプリケーション開発では、アプリケーション(コードやハードウェア)の構造や設計思想を指す言葉。
- アプリケーションの代表的なアーキテクチャには、**MVCモデル、オニオンアーキテクチャ、ヘキサゴナルアーキテクチャ、クリーンアーキテクチャ**などがある。
- 業務内容に直接影響する処理(アプリケーションのコア部分の処理)を**ビジネスロジック**という。

### 8-1-2 MVCモデルとは
- **MVCモデル**とは、処理(コード)を**Model、View、Controller**の3つに分ける考え方。
- **アプリケーション内部のデータを、ユーザーが直接参照・編集する情報から分離できる**ことが特徴
- MVCモデルのメリット
  - 各部分が独立しているため、作業の分担や再利用がしやすい
  - UIの変更や機能の追加が他の部分に影響を与えにくい
  - テストやデバックが用意
- MVCモデルのデメリット
  - 設計や実装が複雑になり、コード量が増える場合がある。
  - 小規模な開発には不向きな場合がある。
  - 処理速度が遅くなる場合がある。
- **アルバムを操作するアプリケーション**
  - **REST API**(httpメソッドを通してやり取りする仕組み)を利用して作成する

## 8-2 アプリケーション作成の準備をしよう
- MVCモデルでアプリケーションを実装していく前に、まずは準備をしていこう。
- 必要なコードを自動生成したり、ロギングや環境変数の設定をしたり、データベースを作成したりなど、準備の手順について1つずつ解説する

### 8-2-1 APIのコードを自動生成しよう
- Goのツールである**oapi-codegen**を使用し、APIの定義ファイルを基にしてコードを生成していく
- APIは**OpenAPI**というWebアプリケーションにおいて一般的なAPIの仕様に従って、HTTPリクエストやHTTPレスポンスの定義していく
- ここでは、「api」というフォルダに次の構成で設定ファイルなどを作る
  - api
    - api.gen.go : oapi-codegenで自動生成
    - config.yaml : oapi-codegenの設定ファイル
    - openapi.yaml : 自動生成するAPIの設定ファイル
- opanapi.yamlでは次の情報を記述する

| 項目名     | 内容                                                      |
| :--------- | :-------------------------------------------------------- |
| openapi    | OpenAPIのバージョン                                       |
| info       | APIの情報                                                 |
| servers    | APIが利用可能なサーバーのURL                              |
| paths      | APIの設定(エンドポイントやメソッドなど)                   |
| components | APIで使用するスキーマ(データの構造や形式などの情報)の定義 |

- **oapi-codegenの実行**
  - oapi-codegenのインストール
    - go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
    - oapi-codegen -version
  - oapi-codegenの実行
    - oapi-codegen --config=./api/config.yaml ./api/openapi.yaml

### 8-2-2 ロギングの設定を作成しよう
- ロギングの設定を作成する
- プロジェクトフォルダに「pkg」フォルダを作成し、その中にさらに「logger」フォルダを作成する。
- 「logger.go」というファイルを作成してコードを記述する
- 「zap」という高速で構造化されたロギングパッケージを利用する

  ```go
  // Package logger provides logging functionality using zap.
  package logger

  import (
    "os"

    "go.uber.org/zap"
  )

  var (
    ZapLogger        *zap.Logger
    zapSugaredLogger *zap.SugaredLogger
  )

  func init() {
    cfg := zap.NewProductionConfig()
    logFile := os.Getenv("APP_LOG_FILE")
    if logFile != "" {
      cfg.OutputPaths = []string{"stderr", "logFile"}
    }

    ZapLogger = zap.Must(cfg.Build())
    if os.Getenv("APP_ENV") == "development" {
      ZapLogger = zap.Must(zap.NewDevelopment())
    }
    zapSugaredLogger = ZapLogger.Sugar()
  }

  func Sync() {
    err := zapSugaredLogger.Sync()
    if err != nil {
      zap.Error(err)
    }
  }

  func Info(msg string, keysAndValues ...interface{}) {
    zapSugaredLogger.Infow(msg, keysAndValues...)
  }

  func Debug(msg string, keysAndValues ...interface{}) {
    zapSugaredLogger.Debugw(msg, keysAndValues...)
  }

  func Warn(msg string, keysAndValues ...interface{}) {
    zapSugaredLogger.Warnw(msg, keysAndValues...)
  }

  func Error(msg string, keysAndValues ...interface{}) {
    zapSugaredLogger.Errorw(msg, keysAndValues...)
  }

  func Fatal(msg string, keysAndValues ...interface{}) {
    zapSugaredLogger.Fatalw(msg, keysAndValues...)
  }

  func Panic(msg string, keysAndValues ...interface{}) {
    zapSugaredLogger.Panicw(msg, keysAndValues...)
  }

  ```

### 8-2-3 アプリケーションの設定を作成しよう
- Goの環境変数などの設定を読み取るためのコードを作成する。
- プロジェクトフォルダに「configs」フォルダを作り、その中に「config.go」というファイルを作成する。
- **os.LookupEnv関数**を利用して、環境変数の値と存在を確認し
  - 環境変数の値が存在する場合、その値を利用。
  - 環境変数の値が存在しない場合、デフォルト値を利用。

    ```go
    // Package configs provides utility functions for configuration management, such as retrieving environment variables with default values.
    package configs

    import (
      "go-api-arch-mvc-template/pkg/logger"
      "os"
      "strconv"

      "go.uber.org/zap"
    )

    func GetEnvDefault(key, defVal string) string {
      val, err := os.LookupEnv(key)
      if !err {
        return defVal
      }
      return val
    }

    type ConfigList struct {
      Env                 string
      DBHost              string
      DBPort              int
      DBDriver            string
      DBName              string
      DBUser              string
      DBPassword          string
      APICorsAllowOrigins []string
    }

    func (c *ConfigList) IsDevelopment() bool {
      return c.Env == "development"
    }

    var Config ConfigList

    func LoadEnv() error {
      DBPort, err := strconv.Atoi(GetEnvDefault("MYSQL_PORT", "3306"))
      if err != nil {
        return err
      }

      Config = ConfigList{
        Env:                 GetEnvDefault("APP_ENV", "development"),
        DBDriver:            GetEnvDefault("DB_DRIVER", "mysql"),
        DBHost:              GetEnvDefault("DB_HOST", "0.0.0.0"),
        DBPort:              DBPort,
        DBUser:              GetEnvDefault("DB_USER", "app"),
        DBPassword:          GetEnvDefault("DB_PASSWORD", "password"),
        DBName:              GetEnvDefault("DB_NAME", "api_database"),
        APICorsAllowOrigins: []string{"http://0.0.0.0:8001"},
      }
      return nil
    }

    func init() {
      if err := LoadEnv(); err != nil {
        logger.Error("Failed to load env: ", zap.Error(err))
        panic(err)
      }
    }

    ```

## 8-3 モデルを実装しよう


## 8-4 コントローラを実装しよう


## 8-5 ビューを実装しよう

