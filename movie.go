package fanart

import (
	"context"
	"fmt"

	"github.com/krelinga/go-jsonflex"
)

type Movie jsonflex.Object

func (m Movie) Name() (string, error) {
	return jsonflex.GetField(m, "name", jsonflex.AsString())
}

func (m Movie) TMDBID() (string, error) {
	return jsonflex.GetField(m, "tmdb_id", jsonflex.AsString())
}

func (m Movie) IMDBID() (string, error) {
	return jsonflex.GetField(m, "imdb_id", jsonflex.AsString())
}

func (m Movie) HdMovieLogo() ([]Image, error) {
	return jsonflex.GetField(m, "hdmovielogo", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (m Movie) MovieDisc() ([]Image, error) {
	return jsonflex.GetField(m, "moviedisc", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (m Movie) MovieLogo() ([]Image, error) {
	return jsonflex.GetField(m, "movielogo", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (m Movie) MoviePoster() ([]Image, error) {
	return jsonflex.GetField(m, "movieposter", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (m Movie) HdMovieClearArt() ([]Image, error) {
	return jsonflex.GetField(m, "hdmovieclearart", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (m Movie) MovieArt() ([]Image, error) {
	return jsonflex.GetField(m, "movieart", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (m Movie) MovieBackground() ([]Image, error) {
	return jsonflex.GetField(m, "moviebackground", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (m Movie) MovieBanner() ([]Image, error) {
	return jsonflex.GetField(m, "moviebanner", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (m Movie) MovieThumb() ([]Image, error) {
	return jsonflex.GetField(m, "moviethumb", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func (m Movie) MovieSquare() ([]Image, error) {
	return jsonflex.GetField(m, "moviesquare", jsonflex.AsArray(jsonflex.AsObject[Image]()))
}

func GetMovie(ctx context.Context, client Client, id string, options ...RequestOption) (Movie, error) {
	path := fmt.Sprintf("/v3/movies/%s", id)
	return client.GetObject(ctx, path, options...)
}