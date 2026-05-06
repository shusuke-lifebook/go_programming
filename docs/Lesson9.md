# Lesson 9 アプリケーションのテストを実施しよう
- 実際のアプリケーションでは、データベースなど周辺の周辺のアプリケーションが関連することがほとんど
- テストの際にはテスト用のデータベースなどを用意する必要がある
- モックやDockerコンテナを利用したテストについて記載する

## 9-1 SQLiteを使ってテストを実行しよう
- テストコードの実装を進めていく
- 軽量なデータベースであるSQLiteを使ったテストの方法について記載する

### 9-1-1 SQLiteを使ってモデルのユニットテスト
- **SQLiteによるテストの準備**
  - pkg/tester/配下にsqlite_suite.goを作成しコードを書いていく
  - DBSQLiteSuite構造体を定義し、**suite.Suite**構造体を埋め込んでいく
  - suite.Suite構造体は、**github.com/stretchr/testify/suiteパッケージ**
  - **SetupSuiteメソッド**は、suiteパッケージの機能により、テスト前に自動で実行されるメソッドです。
    - configs.Config.DBNameに「unittest.sqlite」と設定したあと、models.SetDatabase関数でSQLiteを指定してデータベースを初期化
    - models.GetModels関数でモデルの一覧を取得。gormパッケージのAutoMigrateメソッドでモデルの構造に対応したテーブルを作成する
  - **Assertメソッド**が返す。\*assert.Assertions型のメソッドを使うことで値が想定したものかどうかを確かめることができる

    | メソッド       | 内容                         |
    | :------------- | :--------------------------- |
    | Nilメソッド    | nilであることを検証          |
    | NotNilメソッド | nilでないことを検証          |
    | Equalメソッド  | 指定した値が等しいことを検証 |
    | Trueメソッド   | Trueであることを検証         |
    | JSONEqメソッド | JSONの内容を検証             |

  - **TearDownSuiteメソッド**は、suiteパッケージの機能でテスト後に実行されるメソッドです。
    - os.Remove関数でSQLiteのデータベースファイルを削除する。Assertメソッドでエラーがないことを確認する

  ```go
  // Package tester
  package tester

  import (
    "go-api-arch-mvc-template/app/models"
    "go-api-arch-mvc-template/configs"
    "os"

    "github.com/stretchr/testify/suite"
  )

  type DBSQLiteSuite struct {
    suite.Suite
  }

  func (suite *DBSQLiteSuite) SetupSuite() {
    configs.Config.DBName = "unittest.sqlite"
    err := models.SetDatabase(models.InstanceSqlLite)
    suite.Assert().Nil(err)

    for _, model := range models.GetModels() {
      err := models.DB.AutoMigrate(model)
      suite.Assert().Nil(err)
    }
  }

  func (suite *DBSQLiteSuite) TearDownSuite() {
    err := os.Remove(configs.Config.DBName)
    suite.Assert().Nil(err)
  }

  ```

- **Albumモデルをテストしよう**
  - app/models/album.goをテストするために、同じフォルダ内にalbum_test.goを作成してコードを書いていく
    - パッケージ名：テスト対象のmodelsとは別のmodels_testパッケージとする
      - **テスト対象のパッケージの外部からアクセスできるAPIだけをテストするためです。**
    - AlbumTestSuite構造体を定義する
      - tester.DBSQLiteSuite構造体とテスト前のデータベースの状態を保存するためのoriginalDBというフィールドを持つ
    - TestAlbumTestSuite関数は、**suite.Run関数**に\*testing.T型の引数tと、作成したAlbumTestSuite構造体を渡す
    - AlbumTestSuite構造体に**SetupSuiteメソッド**を作成する
      - suiteパッケージの機能で、テストの前に一度だけ実行されるメソッドです。
      - **AfterTestメソッド**は、suiteパッケージの機能で各テストケースの後に実行されるメソッドで、テスト前のデータベースの状態に戻していく
      - TestAlbum、TestAlbumMarshalのメソッドを作成しテストコードを記述する

  ```go
  package models_test

  import (
    "fmt"
    "go-api-arch-mvc-template/app/models"
    "go-api-arch-mvc-template/pkg/tester"
    "strings"
    "testing"
    "time"

    "github.com/stretchr/testify/suite"
    "gorm.io/gorm"
  )

  type AlbumTestSuite struct {
    tester.DBSQLiteSuite
    originalDB *gorm.DB
  }

  func TestAlbumTestSuite(t *testing.T) {
    suite.Run(t, new(AlbumTestSuite))
  }

  func (suite *AlbumTestSuite) SetupSuite() {
    suite.DBSQLiteSuite.SetupSuite()
    suite.originalDB = models.DB
  }

  func (suite *AlbumTestSuite) AfterTest(suiteName, testName string) {
    models.DB = suite.originalDB
  }

  func Str2time(t string) time.Time {
    parseTime, _ := time.Parse("2006-01-02", t)
    return parseTime
  }

  func (suite *AlbumTestSuite) TestAlbum() {
    createdAlbum, err := models.CreateAlbum("Test", time.Now(), "sports")
    suite.Assert().Nil(err)
    suite.Assert().Equal("Test", createdAlbum.Title)
    suite.Assert().NotNil(createdAlbum.ReleaseDate)
    suite.Assert().NotNil(createdAlbum.Category.ID)
    suite.Assert().Equal(createdAlbum.Category.Name, "sports")

    getAlbum, err := models.GetAlbum(createdAlbum.ID)
    suite.Assert().Nil(err)
    suite.Assert().Equal("Test", getAlbum.Title)
    suite.Assert().NotNil(getAlbum.ReleaseDate)
    suite.Assert().NotNil(getAlbum.Category.ID)
    suite.Assert().Equal(getAlbum.Category.Name, "sports")

    getAlbum.Title = "updated"
    err = getAlbum.Save()
    suite.Assert().Nil(err)
    updatedAlbum, err := models.GetAlbum(createdAlbum.ID)
    suite.Assert().Nil(err)
    suite.Assert().Equal("updated", updatedAlbum.Title)
    suite.Assert().NotNil(updatedAlbum.ReleaseDate)
    suite.Assert().NotNil(updatedAlbum.Category.ID)
    suite.Assert().Equal(updatedAlbum.Category.Name, "sports")

    err = updatedAlbum.Delete()
    suite.Assert().Nil(err)
    deletedAlbum, err := models.GetAlbum(updatedAlbum.ID)
    suite.Assert().Nil(deletedAlbum)
    suite.Assert().True(strings.Contains("record not found", err.Error()))

  }

  func (suite *AlbumTestSuite) TestAlbumMarshal() {
    album := models.Album{
      Title:       "Test",
      ReleaseDate: Str2time("2023-01-01"),
      Category:    &models.Category{Name: "sports"},
    }
    aniversary := time.Now().Year() - 2023
    albumJSON, err := album.MarshalJSON()
    suite.Assert().Nil(err)
    suite.Assert().JSONEq(fmt.Sprintf(`{
    "anniversary":%d,
    "category":{
      "id":0,"name":"sports"
    },
    "id":0,
    "releaseDate":"2023-01-01",
    "title":"Test"
    }`, aniversary), string(albumJSON))
  }

  ```

- **カテゴリーのモデルをテストしよう**
  - album_test.goと同様に作成していく
  
  ```go
  package models_test

  import (
    "go-api-arch-mvc-template/app/models"
    "go-api-arch-mvc-template/pkg/tester"
    "testing"

    "github.com/stretchr/testify/suite"
  )

  type CategoryTestSuite struct {
    tester.DBSQLiteSuite
  }

  func TestCategoryTestSuite(t *testing.T) {
    suite.Run(t, new(CategoryTestSuite))
  }

  func (suite *CategoryTestSuite) TestCategory() {
    category, err := models.GetOrCreateCategory("test")
    suite.Assert().Nil(err)
    suite.Assert().NotNil(category.ID)
    suite.Assert().Equal("test", category.Name)

    category2, err := models.GetOrCreateCategory("test")
    suite.Assert().Nil(err)
    suite.Assert().Equal("test", category2.Name)
    suite.Assert().Equal(category.ID, category2.ID)
  }

  ```

- **テストの実行**
  - go mod tidy
  - go test ./app/modeles/...

## 9-2 モックを使ってテストを実行しよう
- 実際のデータベースやコードを利用すると、テストが難しい場合がある
  - 例えば、データベースでエラーが起きた場合や特定の時間を参照する必要がある場合のテストなどです。
- そんなケースで活用できる。モックを使用したテストについて記載する

### 9-2-1 モックを作成しよう
- **go-sqlmock**パッケージを利用すると、実際のデータベースを使わずに、模擬的なデータベースがあるように振る舞うことができる。
- データベースの操作をモック化するMockDB関数を作成する
- go-sqlmockパッケージの**sqlmock.New関数**でモックデータベースとモックオブジェクトを生成し、gorm.Open関数でモックデータベースを扱えるようにしている。

- Clockインタフェースのモックを作成する。
  - 関数やメソッドによっては、**テストを実行する時刻によって結果が変わってしまうのを防ぐため、モックを使って固定した時刻でテストする必要がある**
  - **mockClock構造体**は、time.Time型のフィールドtを持つ構造体で、現在時刻を返すNowメソッドを持つことでClockインターフェースを実装しているため現在時刻の処理の代わりを作成できる
- [go-sqlmock](https://github.com/data-dog/go-sqlmock)
- go get github.com/DATA-DOG/go-sqlmock

  ```go
  package tester

  import (
    "go-api-arch-mvc-template/pkg/logger"
    "time"

    "github.com/DATA-DOG/go-sqlmock"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
  )

  func MockDB() (mock sqlmock.Sqlmock, mockGormDB *gorm.DB) {
    mockDB, mock, err := sqlmock.New(
      sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
    if err != nil {
      logger.Fatal(err.Error())
    }

    mockGormDB, err = gorm.Open(mysql.New(mysql.Config{
      DSN:                       "mock_db",
      DriverName:                "mysql",
      Conn:                      mockDB,
      SkipInitializeWithVersion: true,
    }), &gorm.Config{})
    if err != nil {
      logger.Fatal(err.Error())
    }
    return mock, mockGormDB
  }

  type mockClock struct {
    t time.Time
  }

  func NewMockClock(t time.Time) mockClock {
    return mockClock{t}
  }

  func (m mockClock) Now() time.Time {
    return m.t
  }

  ```

### 9-2-2 モックを使ってアルバムをテストしよう
- モッククロックを使用して、AnniversaryメソッドをテストするTestAnniversaryメソッドを追加する
- ReleaseDateフィールドの日時を少しずつ変更して、業務mockedClockとの差の計算があっているかを確認しましょう

- 続いてモックデータベースを使用したテストを追加する
  - AlbumTestSuite構造体にMockDBメソッドを追加して、モックデータベースを利用できるようにする。
- エラーの確認
  - **ExpectQueryメソッド**と**WillReturnErrorメソッド**で特定のクエリに対して**特定のエラーを返す**ように設定する
  - **ExpectExecメソッド**と**ExpectBeginメソッド**、**ExpectRollbackメソッド**、**ExpectCommitメソッド**

    ```go
    package models_test

    import (
      "errors"
      "fmt"
      "go-api-arch-mvc-template/app/models"
      "go-api-arch-mvc-template/pkg/tester"
      "regexp"
      "strings"
      "testing"
      "time"

      "github.com/DATA-DOG/go-sqlmock"
      "github.com/stretchr/testify/suite"
      "gorm.io/gorm"
    )

    type AlbumTestSuite struct {
      tester.DBSQLiteSuite
      originalDB *gorm.DB
    }

    func TestAlbumTestSuite(t *testing.T) {
      suite.Run(t, new(AlbumTestSuite))
    }

    func (suite *AlbumTestSuite) SetupSuite() {
      suite.DBSQLiteSuite.SetupSuite()
      suite.originalDB = models.DB
    }

    func (suite *AlbumTestSuite) AfterTest(suiteName, testName string) {
      models.DB = suite.originalDB
    }

    func Str2time(t string) time.Time {
      parseTime, _ := time.Parse("2006-01-02", t)
      return parseTime
    }

    func (suite *AlbumTestSuite) TestAlbum() {
      createdAlbum, err := models.CreateAlbum("Test", time.Now(), "sports")
      suite.Assert().Nil(err)
      suite.Assert().Equal("Test", createdAlbum.Title)
      suite.Assert().NotNil(createdAlbum.ReleaseDate)
      suite.Assert().NotNil(createdAlbum.Category.ID)
      suite.Assert().Equal(createdAlbum.Category.Name, "sports")

      getAlbum, err := models.GetAlbum(createdAlbum.ID)
      suite.Assert().Nil(err)
      suite.Assert().Equal("Test", getAlbum.Title)
      suite.Assert().NotNil(getAlbum.ReleaseDate)
      suite.Assert().NotNil(getAlbum.Category.ID)
      suite.Assert().Equal(getAlbum.Category.Name, "sports")

      getAlbum.Title = "updated"
      err = getAlbum.Save()
      suite.Assert().Nil(err)
      updatedAlbum, err := models.GetAlbum(createdAlbum.ID)
      suite.Assert().Nil(err)
      suite.Assert().Equal("updated", updatedAlbum.Title)
      suite.Assert().NotNil(updatedAlbum.ReleaseDate)
      suite.Assert().NotNil(updatedAlbum.Category.ID)
      suite.Assert().Equal(updatedAlbum.Category.Name, "sports")

      err = updatedAlbum.Delete()
      suite.Assert().Nil(err)
      deletedAlbum, err := models.GetAlbum(updatedAlbum.ID)
      suite.Assert().Nil(deletedAlbum)
      suite.Assert().True(strings.Contains("record not found", err.Error()))

    }

    func (suite *AlbumTestSuite) TestAlbumMarshal() {
      album := models.Album{
        Title:       "Test",
        ReleaseDate: Str2time("2023-01-01"),
        Category:    &models.Category{Name: "sports"},
      }
      aniversary := time.Now().Year() - 2023
      albumJSON, err := album.MarshalJSON()
      suite.Assert().Nil(err)
      suite.Assert().JSONEq(fmt.Sprintf(`{
      "anniversary":%d,
      "category":{
        "id":0,"name":"sports"
      },
      "id":0,
      "releaseDate":"2023-01-01",
      "title":"Test"
      }`, aniversary), string(albumJSON))
    }

    func (suite *AlbumTestSuite) TestAnniversary() {
      mockedClock := tester.NewMockClock(Str2time("2022-04-01"))

      album := models.Album{ReleaseDate: Str2time("2022-04-01")}
      suite.Assert().Equal(0, album.Anniversary(mockedClock))
      album = models.Album{ReleaseDate: Str2time("2021-04-02")}
      suite.Assert().Equal(0, album.Anniversary(mockedClock))
      album = models.Album{ReleaseDate: Str2time("2021-04-01")}
      suite.Assert().Equal(1, album.Anniversary(mockedClock))
      album = models.Album{ReleaseDate: Str2time("2020-04-02")}
      suite.Assert().Equal(1, album.Anniversary(mockedClock))
      album = models.Album{ReleaseDate: Str2time("2020-04-01")}
      suite.Assert().Equal(2, album.Anniversary(mockedClock))
    }

    func (suite *AlbumTestSuite) MockDB() sqlmock.Sqlmock {
      mock, mockGormDB := tester.MockDB()
      models.DB = mockGormDB
      return mock
    }

    func (suite *AlbumTestSuite) TestAlbumCreateFailure() {
      mockDB := suite.MockDB()
      mockDB.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories` WHERE `categories`.`name` = ? ORDER BY `categories`.`id` LIMIT ?")).WithArgs("sports", 1).WillReturnError(errors.New("create error"))
      createdAlbum, err := models.CreateAlbum("Test", Str2time("2023-01-01"), "sports")
      suite.Assert().Nil(createdAlbum)
      suite.Assert().NotNil(err)
      suite.Assert().Equal("create error", err.Error())
    }

    func (suite *AlbumTestSuite) TestAlbumGetFailure() {
      mockDB := suite.MockDB()
      mockDB.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `albums` WHERE `albums`.`id` = ? ORDER BY `albums`.`id` LIMIT ?")).WithArgs(1, 1).WillReturnError(errors.New("get error"))

      album, err := models.GetAlbum(1)
      suite.Assert().Nil(album)
      suite.Assert().NotNil(err)
      suite.Assert().Equal("get error", err.Error())
    }

    func (suite *AlbumTestSuite) TestAlbumSaveFailure() {
      mockDB := suite.MockDB()
      mockDB.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories` WHERE `categories`.`name` = ? ORDER BY `categories`.`id` LIMIT ?")).WithArgs("sports", 1).WillReturnError(errors.New("save error"))

      album := models.Album{
        Title:       "updated",
        ReleaseDate: Str2time("2023-01-01"),
        Category:    &models.Category{Name: "sports"},
      }

      err := album.Save()
      suite.Assert().NotNil(err)
      suite.Assert().Equal("save error", err.Error())
    }

    func (suite *AlbumTestSuite) TestAlbumDeleteFailure() {
      mockDB := suite.MockDB()
      mockDB.ExpectBegin()
      mockDB.ExpectExec("DELETE FROM `albums` WHERE id = ?").WithArgs(0).WillReturnError(errors.New("delete error"))
      mockDB.ExpectRollback()
      mockDB.ExpectCommit()

      album := models.Album{
        Title:       "Test",
        ReleaseDate: Str2time("2023-01-01"),
        Category:    &models.Category{Name: "sports"},
      }
      err := album.Delete()
      suite.Assert().NotNil(err)
      suite.Assert().Equal("delete error", err.Error())
    }

    ```


## 9-3 MySQLを使ってテストを実行しよう
- 機能によってはSQLiteではなくMySQLを使用してテストをしたい場合があるため、テスト用にMySQLをDockerコンテナで立ち上がてテストを実行する方法について記載する

### 9-3-1 MySQLを使用したテストの実行
- SQLiteでも十分なテストは可能。SQLiteとは異なる機能を使っている場合には、実際に運用するMySQLでテストを行いたいことがある。その際にはMySQLでテストできるように、**testcontainers**パッケージを利用して行う
- **テスト用のMySQLをDockerコンテナで立ち上げる**
  - [testcontainers-go](https://github.com/testcontainers/testcontainers-go)
  - go get github.com/testcontainers/testcontainers-go
- **MySQLでテストを実行する**
  - go mod tidy
  - go test -v ./app/models/...
  ```go
  package tester

  import (
    "context"
    "fmt"
    "log"
    "net"
    "time"

    "github.com/stretchr/testify/suite"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/wait"

    "go-api-arch-mvc-template/app/models"
    "go-api-arch-mvc-template/configs"
  )

  func CheckPort(host string, port int) bool {
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
    if conn != nil {
      conn.Close()
      return false
    }
    if err != nil {
      return true
    }
    return false
  }

  func WaitForPort(host string, port int, timeout time.Duration) bool {
    deadline := time.Now().Add(timeout)
    for time.Now().Before(deadline) {
      if CheckPort(host, port) {
        return true
      }
      time.Sleep(1 * time.Second)
    }
    return false
  }

  type DBMySQLSuite struct {
    suite.Suite
    mySQLContainer testcontainers.Container
    ctx            context.Context
  }

  func (suite *DBMySQLSuite) SetupTestContainers() (err error) {
    WaitForPort(configs.Config.DBHost, configs.Config.DBPort, 10*time.Second)
    suite.ctx = context.Background()
    req := testcontainers.ContainerRequest{
      Image: "mysql:8",
      Env: map[string]string{
        "MYSQL_DATABASE":             configs.Config.DBName,
        "MYSQL_USER":                 configs.Config.DBUser,
        "MYSQL_PASSWORD":             configs.Config.DBPassword,
        "MYSQL_ALLOW_EMPTY_PASSWORD": "yes",
      },
      ExposedPorts: []string{"3306/tcp"},
      WaitingFor:   wait.ForLog("port: 3306  MySQL Community Server - GPL."),
    }

    suite.mySQLContainer, err = testcontainers.GenericContainer(suite.ctx, testcontainers.GenericContainerRequest{
      ContainerRequest: req,
      Started:          true,
    })

    if err != nil {
      log.Fatal(err.Error())
    }

    return nil
  }

  // func (suite *DBMySQLSuite) SetupSuite() {
  // 	err := suite.SetupTestContainers()
  // 	suite.Assert().Nil(err)

  // 	err = models.SetDatabase(models.InstanceMySQL)
  // 	suite.Assert().Nil(err)

  // 	for _, model := range models.GetModels() {
  // 		err := models.DB.AutoMigrate(model)
  // 		suite.Assert().Nil(err)
  // 	}
  // }

  func (suite *DBMySQLSuite) SetupSuite() {
    // 1. コンテナ起動
    err := suite.SetupTestContainers()
    suite.Require().Nil(err)

    // 2. コンテナの host / port を取得
    port, err := suite.mySQLContainer.MappedPort(suite.ctx, "3306/tcp")
    suite.Require().Nil(err)

    host, err := suite.mySQLContainer.Host(suite.ctx)
    suite.Require().Nil(err)

    // 3. DSN を生成
    dsn := fmt.Sprintf(
      "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
      configs.Config.DBUser,
      configs.Config.DBPassword,
      host,
      port.Port(),
      configs.Config.DBName,
    )

    // 4. DSN を直接渡して DB を初期化
    err = models.SetDatabaseDSN(dsn)
    suite.Require().Nil(err)

    // 5. マイグレーション
    for _, model := range models.GetModels() {
      err := models.DB.AutoMigrate(model)
      suite.Require().Nil(err)
    }
  }

  func (suite *DBMySQLSuite) TearDownSuite() {
    if suite.mySQLContainer == nil {
      return
    }
    err := suite.mySQLContainer.Terminate(suite.ctx)
    suite.Assert().Nil(err)
  }

  ```


## 9-4 コントローラのテストを実行しよう
- コントローラについてもテストを作成して行こう
- テスト用にHTTPリクエストを作成し、コントローラの処理が呼び出した結果が、想定したHTTPレスポンスであるかどうかテストするための方法について記載する

### 9-4-1 ヘルスチェックのコントローラのテストを実行しよう
- コントローラのテスト実行するには、**net/http/httptestパッケージ**を使用してHTTPのリクエストとレスポンスを確認する
- ヘルスチェックについてテストを実施する。app/controllers/health_test.goを作成し、TestHealthHandler関数を定義します。
  - **httptest.NewRecorder関数**でHTTPレスポンスを記録するためのオブジェクトを作成する
  - **gin.CreateTestContext関数**でリクエストやレスポンスを管理するgin.Context型の構造体を作成する

## 9-5 インテグレーションテストを実行しよう