package persistence

import (
	"errors"

	. "github.com/Masterminds/squirrel"
	"github.com/beego/beego/v2/client/orm"
	"github.com/google/uuid"
	"github.com/navidrome/navidrome/log"
)

const albumResumeTable = "album_resume"

func (r sqlRepository) arId(itemID ...string) And {
	return And{
		Eq{albumResumeTable + ".user_id": userId(r.ctx)},
		Eq{albumResumeTable + ".album_id": itemID},
	}
}

type MediaFile struct {
	ID string `structs:"id" json:"id"            orm:"pk;column(id)"`
}

// 记录album resume
func (r sqlRepository) SetAlbumResume(AlbumID string, trackID string) error {
	sql := Select().From("media_file").Where(Eq{"album_id": AlbumID}).
		OrderBy("order_album_name").
		OrderBy("release_date").
		OrderBy("disc_number").
		OrderBy("track_number").
		OrderBy("order_artist_name").
		OrderBy("order_album_name").
		OrderBy("title").
		Columns("id")

	var mfs []MediaFile
	err := r.queryAll(sql, &mfs)
	if err != nil {
		log.Info(r.ctx, "SetAlbumResume error")
	} else {
		found := false
		i := 0
		for i = 0; i < len(mfs); i++ {
			if mfs[i].ID == trackID {
				found = true
				break
			}
		}

		if found {
			upd := Update(albumResumeTable).Where(r.arId(AlbumID)).
				Set("song_index", i).
				Set("start_at", 0)
			c, err := r.executeSQL(upd)

			if c == 0 || errors.Is(err, orm.ErrNoRows) {
				values := map[string]interface{}{}
				values["ar_id"] = uuid.NewString()
				values["user_id"] = userId(r.ctx)
				values["album_id"] = AlbumID
				values["song_index"] = 0
				values["start_at"] = 0
				ins := Insert(albumResumeTable).SetMap(values)
				_, err = r.executeSQL(ins)
				if err != nil {
					return err
				}
			}
		}
	}

	return err
}

func (r sqlRepository) cleanAlbumResume() error {
	del := Delete(albumResumeTable)
	c, err := r.executeSQL(del)
	if err != nil {
		return err
	}
	if c > 0 {
		log.Debug(r.ctx, "Clean-up annotations", "table", r.tableName, "totalDeleted", c)
	}
	return nil
}
