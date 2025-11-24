package handler

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}
type statusResponse struct {
	Status string `json:"status"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var (
	ErrMovieNotFound = errors.New("movie not found")
	ErrInvalidID     = errors.New("invalid movie ID")
)
var movies = []Movie{
	{
		ID:    "1",
		Isbn:  "978-3-16-148410-0",
		Title: "Inception",
		Director: &Director{
			ID:        "101",
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:    "2",
		Isbn:  "978-1-4028-9462-6",
		Title: "The Shawshank Redemption",
		Director: &Director{
			ID:        "102",
			Firstname: "Frank",
			Lastname:  "Darabont",
		},
	},
}

func (h *Handler) getAllMovies(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(movies)
}
func (h *Handler) getByID(c *gin.Context) {
	id := c.Param("id")
	c.Writer.Header().Set("Content-Type", "application/json")
	for _, val := range movies {
		if val.ID == id {
			json.NewEncoder(c.Writer).Encode(val)
			break
		}
	}
}

func (h *Handler) createMovie(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(c.Request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(c.Writer).Encode(movie)
}

func (h *Handler) deletMovie(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	id := c.Param("id")
	for i, val := range movies {
		if val.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			json.NewEncoder(c.Writer).Encode(id)
			break
		}
	}
}

func (h *Handler) updateMovie(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(c.Request.Body).Decode(&movie)
	for _, val := range movies {
		if val.ID == movie.ID {
			val.Isbn = movie.Isbn
			val.Title = movie.Title
			val.Director.Firstname = movie.Director.Firstname
			val.Director.Lastname = movie.Director.Lastname

			break
		}
	}
	json.NewEncoder(c.Writer).Encode(movie)
}

// func newErrorResponse(c *gin.Context, statusCode int, message string) {
// 	c.AbortWithStatusJSON(statusCode, errorResponse{message})
// }
