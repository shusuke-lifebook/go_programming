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
	cateory, err := GetOrCreateCategory(categoryName)
	if err != nil {
		return nil, err
	}

	album := &Album{
		ReleaseDate: releaseDate,
		Title:       title,
		Category:    cateory,
		CategoryID:  cateory.ID,
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
