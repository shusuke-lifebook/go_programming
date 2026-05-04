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

## 9-4 コントローラのテストを実行しよう

## 9-5 インテグレーションテストを実行しよう