package reviews

import (
	"fmt"
	"strconv"
	"time"
)

type Review struct {
	Id           int       `form:"id" json:"id"`
	Presentation string    `form:"presentation" json:"presentation" binding:"required"`
	Rating       int       `form:"rating" json:"rating" binding:"required"`
	TheGood      string    `form:"good" json:"good" binding:"required"`
	TheBad       string    `form:"bad" json:"bad" binding:"required"`
	Presenter    string    `form:"presenter" json:"presenter" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateReview struct {
	Presentation string `form:"presentation" json:"presentation" binding:"required"`
	Rating       string `form:"rating" json:"rating" binding:"required"`
	TheGood      string `form:"good" json:"good" binding:"required"`
	TheBad       string `form:"bad" json:"bad" binding:"required"`
	Presenter    string `form:"presenter" json:"presenter" binding:"required"`
}

var reviews []Review = []Review{
	{
		Id:           0,
		Presentation: "First Review",
		Rating:       4,
		TheGood:      "Talked about the good stuff.",
		TheBad:       "Didn't talk about tailwindcss.",
		Presenter:    "John Doe",
		CreatedAt:    time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		Id:           1,
		Presentation: "Second Review",
		Rating:       5,
		TheGood:      "Was very informative.",
		TheBad:       "Didn't talk about tailwindcss.",
		Presenter:    "Jane Doe",
		CreatedAt:    time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
	},
	{
		Id:           2,
		Presentation: "Third Review",
		Rating:       5,
		TheGood:      "Wow, just realised I can build web application without react.",
		TheBad:       "No word on tailwindcss.",
		Presenter:    "Foo Bar",
		CreatedAt:    time.Date(2024, 2, 3, 0, 0, 0, 0, time.UTC),
	},
	{
		Id:           4,
		Presentation: "Fourth Review", Rating:       2,
		TheGood:      "is this required?",
		TheBad:       "Really, tailwindcss and fp programming are not mentioned.",
		Presenter:    "Lorem Ipsum",
		CreatedAt:    time.Date(2024, 2, 3, 0, 0, 0, 0, time.UTC),
	},
}

func All() []Review {
	return reviews
}

func Add(n CreateReview) (Review, error) {
	intRating, err := strconv.Atoi(n.Rating)
	if err != nil {
		return Review{}, fmt.Errorf("cannot convert id %s ", n.Rating)
	}

	review := Review{
		Id:           len(reviews),
		Presentation: n.Presentation,
		Rating:       intRating,
		TheGood:      n.TheGood,
		TheBad:       n.TheBad,
		Presenter:    n.Presenter,
		CreatedAt:    time.Now(),
	}
	reviews = append(reviews, review)

	return review, nil
}

func Delete(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("cannot convert id %d ", intID)
	}

	for idx, review := range reviews {
		if review.Id == intID {
			reviews = append(reviews[:idx], reviews[idx+1:]...)
			return nil
		}

	}
	return fmt.Errorf("there is no review  %d ", intID)
}
