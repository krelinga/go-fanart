package fanart

import "github.com/krelinga/go-jsonflex"

type Image jsonflex.Object

func (i Image) ID() (string, error) {
	return jsonflex.GetField(i, "id", jsonflex.AsString())
}

func (i Image) URL() (string, error) {
	return jsonflex.GetField(i, "url", jsonflex.AsString())
}

func (i Image) Lang() (string, error) {
	return jsonflex.GetField(i, "lang", jsonflex.AsString())
}

func (i Image) Likes() (string, error) {
	return jsonflex.GetField(i, "likes", jsonflex.AsString())
}

func (i Image) Disc() (string, error) {
	return jsonflex.GetField(i, "disc", jsonflex.AsString())
}

func (i Image) DiscType() (string, error) {
	return jsonflex.GetField(i, "disc_type", jsonflex.AsString())
}