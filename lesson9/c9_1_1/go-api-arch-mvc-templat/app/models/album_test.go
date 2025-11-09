package models_test

import (
	"go-api-arch-mvc-template/app/models"
	"go-api-arch-mvc-template/pkg/tester"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type AlbumTestSuite struct {
	tester.DBSqliteSuite
	originalDB *gorm.DB
}

func TestAlbumTestSuite(t *testing.T) {
	suite.Run(t, new(AlbumTestSuite))
}

func (suite *AlbumTestSuite) SetupSuite() {
	suite.DBSqliteSuite.SetupSuite()
	suite.originalDB = models.DB
}

func (suite *AlbumTestSuite) AfterTest(suiteName, testName string) {
	models.DB = suite.originalDB
}

func Str2time(t string) time.Time {
	parsedTime, _ := time.Parse("2006-01-02", t)
	return parsedTime
}

func (suite *AlbumTestSuite) TestAlbum() {
	createdAlbum, err := models.CreateAlbum("Test", time.Now(), "sports")
	suite.Assert().Nil(err)
	suite.Assert().Equal("Test", createdAlbum.Title)
	suite.Assert().NotNil(createdAlbum.ReleaseDate)
	suite.Assert().NotNil(createdAlbum.Category.ID)
	suite.Assert().Equal(createdAlbum.Category.Name, "sports")
	// TODO: TBD: More tests
}
