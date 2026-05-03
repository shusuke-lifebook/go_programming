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

## 9-3 MySQLを使ってテストを実行しよう

## 9-4 コントローラのテストを実行しよう

## 9-5 インテグレーションテストを実行しよう