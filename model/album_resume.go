package model

type AlbumResume struct {
	SongIndex int `structs:"-" json:"songIndex"   `
	StartAt   int `structs:"-" json:"startAt"   `
}

type AlbumResumeRepository interface {
	SetAlbumResume(AlbumID string, trackID string) error
}
