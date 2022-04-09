package colrepository

import (
	"codex/internal/pkg/domain"
)

var dbCollections = []domain.CollType{
	{
		Title:       "Топ 256",
		Description: "Вот такая вот подборочка :)",
		MovieList: []domain.MovieType{
			{
				Id:          "1",
				ImgHref:     "greenMile.png",
				Title:       "Зелёная миля",
				Info:        "1999, США. Драма",
				Rating:      "9.1",
				Description: "Пол Эджкомб — начальник блока смертников в тюрьме «Холодная гора», каждый из узников которого однажды проходит «зеленую милю» по пути к месту казни. Пол повидал много заключённых и надзирателей за время работы. Однако гигант Джон Коффи, обвинённый в страшном преступлении, стал одним из самых необычных обитателей блока.",
			},
			{
				Id:          "2",
				ImgHref:     "showshenkRedemption.png",
				Title:       "Побег из Шоушенка",
				Info:        "1994, США. Драма",
				Rating:      "8.9",
				Description: "Бухгалтер Энди Дюфрейн обвинён в убийстве собственной жены и её любовника. Оказавшись в тюрьме под названием Шоушенк, он сталкивается с жестокостью и беззаконием, царящими по обе стороны решётки. Каждый, кто попадает в эти стены, становится их рабом до конца жизни. Но Энди, обладающий живым умом и доброй душой, находит подход как к заключённым, так и к охранникам, добиваясь их особого к себе расположения.",
			},
		},
	},
}

var dbFilms = domain.FilmSelection{
	Coll: []domain.FilmType{
		{
			Description: "Top 256",
			ImgSrc:      "top.png",
			Page:        "1",
			Number:      "1",
		},
		{
			Description: "Приключения",
			ImgSrc:      "adventures.png",
			Page:        "2",
			Number:      "2",
		},
		{
			Description: "Для детей",
			ImgSrc:      "childish.png",
			Page:        "3",
			Number:      "3",
		},
		{
			Description: "Фильмы по комиксам",
			ImgSrc:      "comics.png",
			Page:        "4",
			Number:      "4",
		},
		{
			Description: "Драмы",
			ImgSrc:      "drama.png",
			Page:        "5",
			Number:      "5",
		},
		{
			Description: "Для всей семьи",
			ImgSrc:      "family.png",
			Page:        "6",
			Number:      "6",
		},
		{
			Description: "Рекомендации редакции",
			ImgSrc:      "ourTop.png",
			Page:        "7",
			Number:      "7",
		},
		{
			Description: "Романтические",
			ImgSrc:      "romantic.png",
			Page:        "8",
			Number:      "8",
		},
		{
			Description: "Спасение мира",
			ImgSrc:      "saveTheWorld.png",
			Page:        "9",
			Number:      "9",
		},
		{
			Description: "Советское кино",
			ImgSrc:      "soviet.png",
			Page:        "10",
			Number:      "10",
		},
		{
			Description: "Про шпионов",
			ImgSrc:      "spy.png",
			Page:        "11",
			Number:      "11",
		},
		{
			Description: "Сказки",
			ImgSrc:      "tales.png",
			Page:        "12",
			Number:      "12",
		},
	},
}

type dbCollectionsRepository struct {
	Collections []domain.CollType
	Films       domain.FilmSelection
}

func InitColRep() domain.CollectionsRepository {
	return &dbCollectionsRepository{
		Collections: dbCollections,
		Films:       dbFilms,
	}
}

func (cr *dbCollectionsRepository) GetCollection(id uint64) (domain.CollType, error) {
	return dbCollections[id], nil
}

func (cr *dbCollectionsRepository) GetFeed() (domain.FilmSelection, error) {
	return dbFilms, nil
}
