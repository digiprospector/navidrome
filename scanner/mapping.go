package scanner

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/deluan/sanitize"
	"github.com/navidrome/navidrome/conf"
	"github.com/navidrome/navidrome/consts"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
	"github.com/navidrome/navidrome/scanner/metadata"
	"github.com/navidrome/navidrome/utils"
)

type mediaFileMapper struct {
	rootFolder string
	genres     model.GenreRepository
}

func newMediaFileMapper(rootFolder string, genres model.GenreRepository) *mediaFileMapper {
	return &mediaFileMapper{
		rootFolder: rootFolder,
		genres:     genres,
	}
}

func parseAlbumSetting(path string, aj *map[string]interface{}) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	} else {
		err = json.Unmarshal(content, aj)
		if err != nil {
			log.Error("could not unmarshal json: %s\n", err)
		} else {
			value, ok := (*aj)["name_from_folder"]
			if ok && value == true {
				(*aj)["name"] = filepath.Base(filepath.Dir(path))
			}
		}
	}
}

func findAlbumSetting(path string, aj *map[string]interface{}) {
	dir := filepath.Dir(path)

	for dir != "." && dir != "/" {
		album_json_file := filepath.Join(dir, "album.json")
		if _, err := os.Stat(album_json_file); !os.IsNotExist(err) {
			parseAlbumSetting(album_json_file, aj)
			log.Debug("findAlbumSetting", "found", filepath.Join(dir, "album.json"))
			break
		}
		dir = filepath.Dir(dir)
	}
	log.Debug("findAlbumSetting", "aj", aj)
}

// TODO Move most of these mapping functions to setters in the model.MediaFile
func (s mediaFileMapper) toMediaFile(md metadata.Tags) model.MediaFile {
	var aj map[string]interface{}

	findAlbumSetting(md.FilePath(), &aj)
	mf := &model.MediaFile{}
	mf.ID = s.trackID(md)
	mf.Year, mf.Date, mf.OriginalYear, mf.OriginalDate, mf.ReleaseYear, mf.ReleaseDate = s.mapDates(md)
	mf.Title = s.mapTrackTitle(md)
	mf.Album = md.Album()
	n, n_ok := aj["name"]
	a, a_ok := aj["artist"]
	if n_ok && a_ok {
		mf.AlbumID = fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(fmt.Sprintf("%s\\%s", a, n)))))
	} else {
		mf.AlbumID = s.albumID(md, mf.ReleaseDate)
	}
	if a_ok {
		mf.Album = fmt.Sprintf("%s", n)
	} else {
		mf.Album = s.mapAlbumName(md)
	}
	mf.ArtistID = s.artistID(md)
	mf.Artist = s.mapArtistName(md)
	if a_ok {
		mf.AlbumArtistID = fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(fmt.Sprintf("%s", a)))))
		mf.AlbumArtist = fmt.Sprintf("%s", a)
	} else {
		mf.AlbumArtistID = s.albumArtistID(md)
		mf.AlbumArtist = s.mapAlbumArtistName(md)
	}
	mf.Genre, mf.Genres = s.mapGenres(md.Genres())
	mf.Compilation = md.Compilation()
	mf.TrackNumber, _ = md.TrackNumber()
	mf.DiscNumber, _ = md.DiscNumber()
	mf.DiscSubtitle = md.DiscSubtitle()
	mf.Duration = md.Duration()
	mf.BitRate = md.BitRate()
	mf.Channels = md.Channels()
	mf.Path = md.FilePath()
	mf.Suffix = md.Suffix()
	mf.Size = md.Size()
	mf.HasCoverArt = md.HasPicture()
	mf.SortTitle = md.SortTitle()
	mf.SortAlbumName = md.SortAlbum()
	mf.SortArtistName = md.SortArtist()
	mf.SortAlbumArtistName = md.SortAlbumArtist()
	mf.OrderTitle = strings.TrimSpace(sanitize.Accents(mf.Title))
	mf.OrderAlbumName = sanitizeFieldForSorting(mf.Album)
	mf.OrderArtistName = sanitizeFieldForSorting(mf.Artist)
	mf.OrderAlbumArtistName = sanitizeFieldForSorting(mf.AlbumArtist)
	mf.CatalogNum = md.CatalogNum()
	mf.MbzRecordingID = md.MbzRecordingID()
	mf.MbzReleaseTrackID = md.MbzReleaseTrackID()
	mf.MbzAlbumID = md.MbzAlbumID()
	mf.MbzArtistID = md.MbzArtistID()
	mf.MbzAlbumArtistID = md.MbzAlbumArtistID()
	mf.MbzAlbumType = md.MbzAlbumType()
	mf.MbzAlbumComment = md.MbzAlbumComment()
	mf.RGAlbumGain = md.RGAlbumGain()
	mf.RGAlbumPeak = md.RGAlbumPeak()
	mf.RGTrackGain = md.RGTrackGain()
	mf.RGTrackPeak = md.RGTrackPeak()
	mf.Comment = utils.SanitizeText(md.Comment())
	mf.Lyrics = utils.SanitizeText(md.Lyrics())
	mf.Bpm = md.Bpm()
	mf.CreatedAt = time.Now()
	mf.UpdatedAt = md.ModificationTime()

	return *mf
}

func sanitizeFieldForSorting(originalValue string) string {
	v := strings.TrimSpace(sanitize.Accents(originalValue))
	return utils.NoArticle(v)
}

func (s mediaFileMapper) mapTrackTitle(md metadata.Tags) string {
	if md.Title() == "" {
		s := strings.TrimPrefix(md.FilePath(), s.rootFolder+string(os.PathSeparator))
		e := filepath.Ext(s)
		return strings.TrimSuffix(s, e)
	}
	return md.Title()
}

func (s mediaFileMapper) mapAlbumArtistName(md metadata.Tags) string {
	switch {
	case md.AlbumArtist() != "":
		return md.AlbumArtist()
	case md.Compilation():
		return consts.VariousArtists
	case md.Artist() != "":
		return md.Artist()
	default:
		return consts.UnknownArtist
	}
}

func (s mediaFileMapper) mapArtistName(md metadata.Tags) string {
	if md.Artist() != "" {
		return md.Artist()
	}
	return consts.UnknownArtist
}

func (s mediaFileMapper) mapAlbumName(md metadata.Tags) string {
	name := md.Album()
	if name == "" {
		return consts.UnknownAlbum
	}
	return name
}

func (s mediaFileMapper) trackID(md metadata.Tags) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(md.FilePath())))
}

func (s mediaFileMapper) albumID(md metadata.Tags, releaseDate string) string {
	albumPath := strings.ToLower(fmt.Sprintf("%s\\%s", s.mapAlbumArtistName(md), s.mapAlbumName(md)))
	if !conf.Server.Scanner.GroupAlbumReleases {
		if len(releaseDate) != 0 {
			albumPath = fmt.Sprintf("%s\\%s", albumPath, releaseDate)
		}
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(albumPath)))
}

func (s mediaFileMapper) artistID(md metadata.Tags) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(s.mapArtistName(md)))))
}

func (s mediaFileMapper) albumArtistID(md metadata.Tags) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(s.mapAlbumArtistName(md)))))
}

func (s mediaFileMapper) mapGenres(genres []string) (string, model.Genres) {
	var result model.Genres
	unique := map[string]struct{}{}
	var all []string
	for i := range genres {
		gs := strings.FieldsFunc(genres[i], func(r rune) bool {
			return strings.ContainsRune(conf.Server.Scanner.GenreSeparators, r)
		})
		for j := range gs {
			g := strings.TrimSpace(gs[j])
			key := strings.ToLower(g)
			if _, ok := unique[key]; ok {
				continue
			}
			all = append(all, g)
			unique[key] = struct{}{}
		}
	}
	for _, g := range all {
		genre := model.Genre{Name: g}
		_ = s.genres.Put(&genre)
		result = append(result, genre)
	}
	if len(result) == 0 {
		return "", nil
	}
	return result[0].Name, result
}

func (s mediaFileMapper) mapDates(md metadata.Tags) (int, string, int, string, int, string) {
	year, date := md.Date()
	originalYear, originalDate := md.OriginalDate()
	releaseYear, releaseDate := md.ReleaseDate()

	// MusicBrainz Picard writes the Release Date of an album to the Date tag, and leaves the Release Date tag empty
	taggedLikePicard := (originalYear != 0) &&
		(releaseYear == 0) &&
		(year >= originalYear)
	if taggedLikePicard {
		return originalYear, originalDate, originalYear, originalDate, year, date
	}
	return year, date, originalYear, originalDate, releaseYear, releaseDate
}
