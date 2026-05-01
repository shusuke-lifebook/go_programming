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

### 8-2-4 データベースの設定をしよう
- **MySQL**というデータベースを使用してデータを保存していく
- **init.sql**の作成
  - external-apps/db
    ```sql
    CREATE DATABASE IF NOT EXISTS api_database;

    USE api_database;

    CREATE TABLE albums (
        id INT PRIMARY KEY AUTO_INCREMENT,
        release_date DATE,
        category_id INT,
        title VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

    CREATE TABLE categories (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(10),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    ```
- **MySQLの立ち上げ**
  - **Docker**を使用してMySQLを立ち上げる
  - Dockerを利用することで、**コンテナ**と呼ばれる仮想環境を作成することができる
  - **Docker Compose**というツールで、複数のコンテナを一括で起動および管理も可能。
- docker-compose.yaml
  ```yaml
  services:
    mysql:
      image: mysql:8.0
      container_name: mysql
      ports:
        - 3306:3306
      environment:
        MYSQL_USER: app
        MYSQL_PASSWORD: password
        MYSQL_DATABASE: api_database
        MYSQL_ALLOW_EMPTY_PASSWORD: yes
      healthcheck:
        test:
          [
            "CMD",
            "mysqladmin",
            "ping",
            "-h",
            "localhost",
            "-u",
            "root",
            "-p$MYSQL_ROOT_PASSWORD",
          ]
        interval: 3s
        timeout: 5s
        retries: 5
      # restart: always
      restart: no
      volumes:
        - ./external-apps/db/:/docker-entrypoint-initdb.d
      networks:
        - api-network
    mysql-cli:
      image: mysql:8.0
      command: mysql -hmysql -uapp -ppassword api_database

      depends_on:
        mysql:
          condition: service_healthy
      networks:
        - api-network
  networks:
    api-network:
      driver: bridge

  ```
  - mysqlの起動
    - docker compose up -d
  - mysqlへの接続
    - docker compose run mysql-cli
      ```console
      docker compose run mysql-cli
      WARN[0000] The "MYSQL_ROOT_PASSWORD" variable is not set. Defaulting to a blank string. 
      [+]  1/1t 1/11
      ✔ Container mysql Running                                                                                                     0.0s
      Container mysql Waiting 
      Container mysql Healthy 
      Container go-api-arch-mvc-template-mysql-cli-run-efda2e2a4a35 Creating 
      Container go-api-arch-mvc-template-mysql-cli-run-efda2e2a4a35 Created 
      mysql: [Warning] Using a password on the command line interface can be insecure.
      Reading table information for completion of table and column names
      You can turn off this feature to get a quicker startup with -A

      Welcome to the MySQL monitor.  Commands end with ; or \g.
      Your MySQL connection id is 26
      Server version: 8.0.45 MySQL Community Server - GPL

      Copyright (c) 2000, 2026, Oracle and/or its affiliates.

      Oracle is a registered trademark of Oracle Corporation and/or its
      affiliates. Other names may be trademarks of their respective
      owners.

      Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

      mysql> 
      ```
  - MySQLから抜ける
    - docker compose down

## 8-3 モデルを実装しよう
- ビジネスロジックを担当するモデルのコードを書いていく。
- 今回はアルバム情報を作成するコードを作成します。
- モデルの処理結果をデータベースに接続するコードを合わせて書いていく。

### 8-3-1 データベースに接続しよう
- プロジェクトフォルダに「app」フォルダを作成し、その中に「models」フォルダを作ってモデルを書く。
- 「app」
  - 「models」
    - db.go: データベースに接続するコード
    - album.go: Albumモデルのコード
    - category.go: Categoryモデルのコード
- **MySQLとSQLiteの両方が使える**ようにコードを実装する
- **GORM**という**ORM(Object-Relational Mapping:オブジェクト関係マッピング)ライブラリ**を利用する
  - インストール
    - go get -u gorm.io/gorm
    - go get -u gorm.io/driver/sqlite
    - go get -u gorm.io/driver/mysql

      ```go
      // Package models
      package models

      import (
        "errors"
        "fmt"
        "go-api-arch-mvc-template/configs"

        "gorm.io/driver/mysql"
        "gorm.io/driver/sqlite"
        "gorm.io/gorm"
      )

      const (
        InstanceSqlLite int = iota
        InstanceMySQL
      )

      var (
        DB                            *gorm.DB
        errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
      )

      func GetModels() []interface{} {
        return []interface{}{&Album{}, &Category{}}
      }

      func NewDatabaseSQLFactory(instance int) (db *gorm.DB, err error) {
        switch instance {
        case InstanceMySQL:
          dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
            configs.Config.DBUser,
            configs.Config.DBPassword,
            configs.Config.DBHost,
            configs.Config.DBPort,
            configs.Config.DBName)
          db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        case InstanceSqlLite:
          db, err = gorm.Open(sqlite.Open(configs.Config.DBName), &gorm.Config{})
        default:
          return nil, errInvalidSQLDatabaseInstance
        }
        return db, err
      }

      func SetDatabase(instance int) (err error) {
        db, err := NewDatabaseSQLFactory(instance)
        if err != nil {
          return err
        }
        DB = db
        return nil
      }

      ```

### 8-3-2 Categoryモデルを作成しよう
- categoriesテーブルに対応したCategoryモデルを作成する
- 構造体Categoryはint型のIDとstring型のNameという二つのフィールドを持つ。
- 変数DBの**FirstOrCreateメソッド**を実行する
  ```go
  package models

  type Category struct {
    ID   int
    Name string
  }

  func GetOrCreateCategory(name string) (*Category, error) {
    var category Category
    tx := DB.FirstOrCreate(&category, Category{Name: name})
    if tx.Error != nil {
      return nil, tx.Error
    }
    return &category, nil
  }

  ```

### 8-3-3 時間に関する処理を作成する
- アルバムのデータには日時に関する情報が含まれるので、時間に関する処理を作る。
- 汎用的な処理なので、プロジェクトフォルダ直下の「pkg」フォルダに「times.go」を作成する。
- **うるう年を考慮してアルバムのリリース日を考える**
- **時刻を取得するClockインタフェースと構造体RealClockを定義する**
- **うるう年かどうかを判定するisLeap関数**
- **リリース日の年内の経過日数を調整するGetAdjustedReleaseDay関数**

  ```go
  // Package pkg
  package pkg

  import (
    "time"
  )

  type Clock interface {
    Now() time.Time
  }

  type RealClock struct{}

  func (RealClock) Now() time.Time {
    return time.Now()
  }

  func isLeap(date time.Time) bool {
    year := date.Year()
    if year%400 == 0 {
      return true
    } else if year%100 == 0 {
      return false
    } else if year%4 == 0 {
      return true
    }
    return false
  }

  func GetAdjustedReleaseDay(releaseDate time.Time, now time.Time) int {
    releaseDay := releaseDate.YearDay()
    currentDay := now.YearDay()
    if isLeap(releaseDate) && !isLeap(now) && releaseDay >= 60 {
      return releaseDay - 1
    }
    if isLeap(now) && !isLeap(releaseDate) && currentDay >= 60 {
      return releaseDay + 1
    }
    return releaseDay
  }

  ```

### 8-3-4 Albumモデルを作成しよう
- 「app/models/」に「album.go」を作成し、albumsテーブルに対応するコードを作成していく
- アルバムの情報を保持する構造体Albumは、次のような5つのフィールドを持つ

| フィールド  | 型          | 説明                   |
| :---------- | :---------- | :--------------------- |
| ID          | int型       | アルバムのID           |
| Title       | string型    | アルバムのタイトル     |
| ReleaseDate | time.Time型 | アルバムのリリース日   |
| CategoryID  | int型       | アルバムのカテゴリーID |
| Category    | *Category型 | アルバムのカテゴリー   |

- **経過年数を表すAnniversaryメソッド**
- **構造体をJSONに変換するMarshalJSONメソッド**
- **アルバムを作成するCreateAlbum関数**
  - \*gorm.DB型の**Createメソッド**でデータベースに保存する
- **アルバムの情報を取得するGetAlbum関数**
  - \*gorm.DB型の**Preloadメソッド**でカテゴリを取得しておき、Firstメソッドで引数に指定したIDで検索して最初のレコードを返す
- **アルバムを保持するSaveメソッド**
  - \*gorm.DB型の**Saveメソッド**でデータベースを更新する
- **アルバムを削除するDeleteメソッド**
  - \*gorm.DB型の**Whereメソッド**でレコードを検索して、**Deleteメソッド**で該当するレコードを削除する

  ```go
  package models

  import (
    "encoding/json"
    "go-api-arch-mvc-template/api"
    "go-api-arch-mvc-template/pkg"
    "time"
  )

  type Album struct {
    ID          int
    Title       string
    ReleaseDate time.Time
    CategoryID  int
    Category    *Category
  }

  func (a *Album) Anniversary(clock pkg.Clock) int {
    now := clock.Now()
    years := now.Year() - a.ReleaseDate.Year()
    releaseDay := pkg.GetAdjustedReleaseDay(a.ReleaseDate, now)
    if now.YearDay() < releaseDay {
      years -= 1
    }
    return years
  }

  func (a *Album) MarshalJSON() ([]byte, error) {
    return json.Marshal(&api.AlbumResponse{
      Id:          a.ID,
      Title:       a.Title,
      Anniversary: a.Anniversary(pkg.RealClock{}),
      ReleaseDate: api.ReleaseDate{Time: a.ReleaseDate},
      Category: api.Category{
        Id:   &a.Category.ID,
        Name: api.CategoryName(a.Category.Name),
      },
    })
  }

  func CreateAlbum(title string, releaseDate time.Time, categoryName string) (*Album, error) {
    category, err := GetOrCreateCategory(categoryName)
    if err != nil {
      return nil, err
    }

    album := &Album{
      ReleaseDate: releaseDate,
      Title:       title,
      Category:    category,
      CategoryID:  category.ID,
    }
    if err := DB.Create(album).Error; err != nil {
      return nil, err
    }
    return album, nil
  }

  func GetAlbum(ID int) (*Album, error) {
    var album = Album{}
    if err := DB.Preload("Category").First(&album, ID).Error; err != nil {
      return nil, err
    }
    return &album, nil
  }

  func (a *Album) Save() error {
    category, err := GetOrCreateCategory(a.Category.Name)
    if err != nil {
      return err
    }
    a.Category = category
    a.CategoryID = category.ID

    if err := DB.Save(&a).Error; err != nil {
      return err
    }
    return nil
  }

  func (a *Album) Delete() error {
    if err := DB.Where("id = ?", &a.ID).Delete(&a).Error; err != nil {
      return err
    }
    return nil
  }

  ```

## 8-4 コントローラを実装しよう
- コントローラは、ビューから送られてきたデータをモデルに渡し、モデルから返ってきた処理をビューに渡す役割です。
- 今回は、ビューにあたるAPIからデータを受け取ってモデルへと渡し、結果をビューへと返す。
- 実際にコードを作成していく

### 8-4-1 Albumコントローラを作成しよう
- プロジェクトフォルダに「controllers」というフォルダを作成し、album.goにコードを書いていく
- 空の構造体としてAlbumHandler型を定義する
  - **ハンドラー**： HTTPリクエストを受け取ってレスポンスを返す処理を担当する
- **CreateAlbumメソッド**
  - アルバムを作成するリクエストがあったときに呼び出されるメソッド
  - リクエストやレスポンスなどの情報を保持する **\*gin.Context型**の引数を受け取る
- **GetAlbumByIdメソッド**
  - GetAlbumIdメソッドは、指定したIDのアルバム情報を取得するリクエストがあったときに呼び出されるメソッドです。
- **UpdateAlbumByIdメソッド**
  - UpdateAlbumByIdメソッドは、指定したIDのアルバムを更新するリクエストがあったときに呼び出されるメソッドです。
- **DeleteAlbumByIdメソッド**
  - DeleteAlbumByIdメソッドは、指定したIDでアルバムを削除するリクエストがあったときに呼び出されるメソッドです。

  ```go
  // Package controllers
  package controllers

  import (
    "go-api-arch-mvc-template/api"
    "go-api-arch-mvc-template/app/models"
    "go-api-arch-mvc-template/pkg/logger"
    "net/http"

    "github.com/gin-gonic/gin"
  )

  type AlbumHandler struct{}

  func (a *AlbumHandler) CreateAlbum(c *gin.Context) {
    var requestBody api.CreateAlbumJSONRequestBody
    if err := c.ShouldBindJSON(requestBody); err != nil {
      logger.Warn(err.Error())
      c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
      return
    }
    createdAlbum, err := models.CreateAlbum(
      requestBody.Title,
      requestBody.ReleaseDate.Time,
      string(requestBody.Category.Name))
    if err != nil {
      logger.Error(err.Error())
      c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
      return
    }
    c.JSON(http.StatusCreated, createdAlbum)
  }

  func (a *AlbumHandler) GetAlbumById(c *gin.Context, ID int) {
    album, err := models.GetAlbum(ID)
    if err != nil {
      logger.Error(err.Error())
      c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
    }
    c.JSON(http.StatusOK, album)
  }

  func (a *AlbumHandler) UpdateAlbumById(c *gin.Context, ID int) {
    var requestBody api.UpdateAlbumByIdJSONRequestBody
    if err := c.ShouldBindJSON(&requestBody); err != nil {
      logger.Warn(err.Error())
      c.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
      return
    }

    album, err := models.GetAlbum(ID)
    if err != nil {
      logger.Error(err.Error())
      c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
      return
    }

    if requestBody.Category != nil {
      album.Category.Name = string(requestBody.Category.Name)
    }

    if requestBody.Title != nil {
      album.Title = *requestBody.Title
    }

    if err := album.Save(); err != nil {
      logger.Error(err.Error())
      c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
      return
    }
    c.JSON(http.StatusOK, album)
  }

  func (a *AlbumHandler) DeleteAlbumById(c *gin.Context, ID int) {
    album := models.Album{ID: ID}

    if err := album.Delete(); err != nil {
      logger.Error(err.Error())
      c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
      return
    }
    c.JSON(http.StatusNoContent, nil)
  }

  ```

## 8-5 ビューを実装しよう
- ユーザーからの入力を受け取り、結果をユーザーに返す役割であるビューを実装する
- Webアプリケーションでは、処理の結果をHTMLなどに反映して画面に表示する処理を担う
- APIのサーバー起動
- Swaggerというツールの画面とターミナルでのCurlコマンドでユーザーとやり取りする処理を作る

### 8-5-1 APIサーバーを起動しよう
- プロジェクトフォルダにmain.goを作成し、サーバーの起動処理と合わせてビューを書いていく
- [swag](https://github.com/swaggo/swag)
  - go install github.com/swaggo/swag/cmd/swag@latest
- サーバー起動時の処理としてデータベースに接続する
- **github.com/gin-gonic/gin**パッケージ
  - **Webフレームワーク**と呼ばれるWebアプリケーションを作成するためのパッケージです。
  - **gin.Default**関数でHTTPリクエスト振り分けるための**ルーターを作成する**
- **Swaggerの準備**
  - api.GetSwagger関数で**Swagger**の仕様を取得する
  - Swaggerとは、REST APIのドキュメントを生成するツールで、WebページからAPIの確認や実行ができる
- **APIのルーティング(振り分け)**
  - ルーターの**Groupメソッド**で「/api」ではじめるURLをグループ化し、さらにその中で「/api/v1」ではじまるAPIのサブグループを作成する。
  - gin-middlewareパッケージの**OapiRequestValidator関数**によって、変数swaggerのAPI仕様に基づく**バリデーションを行う**
- **サーバーの起動処理**
  - 無名関数をゴルーチンとして実行する。**ListenAndServeメソッド**でサーバーを起動し、リクエストを待ち受ける
- **サーバー終了処理**
  - os.Signal型(オペレーティングシステムのシグナル)を受けるチャネルを作成する
  - **os/signalパッケージのsignal.Notify関数で、SIGINTとSIGTERMのシグナルがあれば、quitチャネルに送信する**
  - **context.WithTimeout関数**で2秒のタイムアウトを持つ**コンテキスト**(複数のゴルーチン間でキャンセルタイムアウトなどのシグナルを伝えるためのもの)
- **APIサーバーの起動**
  - APIの動作確認
    - docker compose up -d
    - go run main.go
    - http://localhost:8080/swagger/index.html