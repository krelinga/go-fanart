package fanart_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/krelinga/go-fanart"
)

type result[T comparable] struct {
	value T
	err   error
}

func newResult[T comparable](value T, err error) result[T] {
	return result[T]{value: value, err: err}
}

func (r result[T]) assertValueIs(t *testing.T, expected T) bool {
	t.Helper()
	if !r.assertNoError(t) {
		return false
	}
	if r.value != expected {
		t.Errorf("Expected value: %v, got: %v", expected, r.value)
		return false
	}
	return true
}

func (r result[T]) assertErrorIs(t *testing.T, expected error) bool {
	t.Helper()
	if r.err == nil {
		t.Errorf("Expected error: %v, got nil", expected)
		return false
	}
	if !errors.Is(r.err, expected) {
		t.Errorf("Expected error: %v, got: %v", expected, r.err)
		return false
	}
	return true
}

func (r result[T]) assertNoError(t *testing.T) bool {
	t.Helper()
	if r.err != nil {
		t.Errorf("Expected no error, got: %v", r.err)
		return false
	}
	return true
}

func TestGetMovie(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()

	// Replace with a valid movie ID for testing
	const movieID = "550" // Example: Fight Club

	movie, err := fanart.GetMovie(ctx, client, movieID)
	if err != nil {
		t.Fatalf("Failed to get movie: %v", err)
	}

	if movie == nil {
		t.Fatal("Expected movie, got nil")
	}

	newResult(movie.Name()).assertValueIs(t, "Fight Club")
	newResult(movie.TMDBID()).assertValueIs(t, "550")
	newResult(movie.IMDBID()).assertValueIs(t, "tt0137523")
	if posters, err := movie.MoviePoster(); err != nil {
		t.Errorf("Failed to get movie posters: %v", err)
	} else if len(posters) == 0 {
		t.Error("Expected at least one movie poster, got none")
	} else if poster, err := findImageById(t, posters, "50065"); err != nil {
		t.Error(err)
	} else {
		newResult(poster.ID()).assertValueIs(t, "50065")
		newResult(poster.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/movieposter/fight-club-522a5477c7bd3.jpg")
		newResult(poster.Lang()).assertValueIs(t, "en")
		newResult(poster.Likes()).assertValueIs(t, "15")
	}
	if logos, err := movie.HdMovieLogo(); err != nil {
		t.Errorf("Failed to get HD movie logos: %v", err)
	} else if len(logos) == 0 {
		t.Error("Expected at least one HD movie logo, got none")
	} else if logo, err := findImageById(t, logos, "12657"); err != nil {
		t.Error(err)
	} else {
		newResult(logo.ID()).assertValueIs(t, "12657")
		newResult(logo.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/hdmovielogo/fight-club-504c0530d5f93.png")
		newResult(logo.Lang()).assertValueIs(t, "en")
		newResult(logo.Likes()).assertValueIs(t, "8")
	}
	if discs, err := movie.MovieDisc(); err != nil {
		t.Errorf("Failed to get movie discs: %v", err)
	} else if len(discs) == 0 {
		t.Error("Expected at least one movie disc, got none")
	} else if disc, err := findImageById(t, discs, "25893"); err != nil {
		t.Error(err)
	} else {
		newResult(disc.ID()).assertValueIs(t, "25893")
		newResult(disc.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/moviedisc/fight-club-512e9ac9bf96d.png")
		newResult(disc.Lang()).assertValueIs(t, "en")
		newResult(disc.Likes()).assertValueIs(t, "8")
		newResult(disc.Disc()).assertValueIs(t, "1")
		newResult(disc.DiscType()).assertValueIs(t, "bluray")
	}
	if clearArts, err := movie.HdMovieClearArt(); err != nil {
		t.Errorf("Failed to get HD movie clear arts: %v", err)
	} else if len(clearArts) == 0 {
		t.Error("Expected at least one HD movie clear art, got none")
	} else if clearArt, err := findImageById(t, clearArts, "150127"); err != nil {
		t.Error(err)
	} else {
		newResult(clearArt.ID()).assertValueIs(t, "150127")
		newResult(clearArt.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/hdmovieclearart/fight-club-5721041e01e28.png")
		newResult(clearArt.Lang()).assertValueIs(t, "es")
		newResult(clearArt.Likes()).assertValueIs(t, "7")
	}
	if logos, err := movie.MovieLogo(); err != nil {
		t.Errorf("Failed to get movie logos: %v", err)
	} else if len(logos) == 0 {
		t.Error("Expected at least one movie logo, got none")
	} else if logo, err := findImageById(t, logos, "430"); err != nil {
		t.Error(err)
	} else {
		newResult(logo.ID()).assertValueIs(t, "430")
		newResult(logo.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/movielogo/fight-club-4f15e30ca6244.png")
		newResult(logo.Lang()).assertValueIs(t, "en")
		newResult(logo.Likes()).assertValueIs(t, "6")
	}
	if backgrounds, err := movie.MovieBackground(); err != nil {
		t.Errorf("Failed to get movie backgrounds: %v", err)
	} else if len(backgrounds) == 0 {
		t.Error("Expected at least one movie background, got none")
	} else if bg, err := findImageById(t, backgrounds, "119633"); err != nil {
		t.Error(err)
	} else {
		newResult(bg.ID()).assertValueIs(t, "119633")
		newResult(bg.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/moviebackground/fight-club-55e2393686745.jpg")
		newResult(bg.Lang()).assertValueIs(t, "")
		newResult(bg.Likes()).assertValueIs(t, "5")
	}
	if thumbs, err := movie.MovieThumb(); err != nil {
		t.Errorf("Failed to get movie thumbs: %v", err)
	} else if len(thumbs) == 0 {
		t.Error("Expected at least one movie thumb, got none")
	} else if thumb, err := findImageById(t, thumbs, "37711"); err != nil {
		t.Error(err)
	} else {
		newResult(thumb.ID()).assertValueIs(t, "37711")
		newResult(thumb.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/moviethumb/fight-club-51b0f879f12e2.jpg")
		newResult(thumb.Lang()).assertValueIs(t, "en")
		newResult(thumb.Likes()).assertValueIs(t, "5")
	}
	if arts, err := movie.MovieArt(); err != nil {
		t.Errorf("Failed to get movie arts: %v", err)
	} else if len(arts) == 0 {
		t.Error("Expected at least one movie art, got none")
	} else if art, err := findImageById(t, arts, "4863"); err != nil {
		t.Error(err)
	} else {
		newResult(art.ID()).assertValueIs(t, "4863")
		newResult(art.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/movieart/fight-club-4fd32550b72bf.png")
		newResult(art.Lang()).assertValueIs(t, "en")
		newResult(art.Likes()).assertValueIs(t, "4")
	}
	if banners, err := movie.MovieBanner(); err != nil {
		t.Errorf("Failed to get movie banners: %v", err)
	} else if len(banners) == 0 {
		t.Error("Expected at least one movie banner, got none")
	} else if banner, err := findImageById(t, banners, "202300"); err != nil {
		t.Error(err)
	} else {
		newResult(banner.ID()).assertValueIs(t, "202300")
		newResult(banner.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/moviebanner/fight-club-59901d2e28aa1.jpg")
		newResult(banner.Lang()).assertValueIs(t, "en")
		newResult(banner.Likes()).assertValueIs(t, "4")
	}
	if squares, err := movie.MovieSquare(); err != nil {
		t.Errorf("Failed to get movie squares: %v", err)
	} else if len(squares) == 0 {
		t.Error("Expected at least one movie square, got none")
	} else if square, err := findImageById(t, squares, "452731"); err != nil {
		t.Error(err)
	} else {
		newResult(square.ID()).assertValueIs(t, "452731")
		newResult(square.URL()).assertValueIs(t, "http://assets.fanart.tv/fanart/movies/550/moviesquare/fight-club-683c539c1f2ea.jpg")
		newResult(square.Lang()).assertValueIs(t, "en")
		newResult(square.Likes()).assertValueIs(t, "2")
	}

	t.Run("404MovieId", func(t *testing.T) {
		_, err := fanart.GetMovie(ctx, client, "00000000")
		if statusCodeErr, ok := err.(fanart.HttpStatusCodeError); ok {
			if statusCodeErr.StatusCode != 404 {
				t.Errorf("Expected 404 status code, got: %d", statusCodeErr.StatusCode)
			}
		}
	})
}

func findImageById(t *testing.T, images []fanart.Image, id string) (fanart.Image, error) {
	t.Helper()
	for i, img := range images {
		idResult := newResult(img.ID())
		if !idResult.assertNoError(t) {
			return nil, fmt.Errorf("failed to get ID for image at index %d: %w", i, idResult.err)
		}
		if idResult.value == id {
			return img, nil
		}
	}

	return nil, fmt.Errorf("image with ID %q not found", id)
}
