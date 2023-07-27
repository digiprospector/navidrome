package migrations

import (
	"database/sql"

	"github.com/navidrome/navidrome/log"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up20230726102934, Down20230726102934)
}

func Up20230726102934(tx *sql.Tx) error {
	log.Info("Creating album resume")
	_, err := tx.Exec(`
create table if not exists album_resume
(
	ar_id varchar(255) not null
		primary key,
	user_id varchar(255) default '' not null,
	album_id varchar(255) default '' not null,
	song_index integer default 0 not null,
	start_at integer default 0 not null
);`)
	return err
}

func Down20230726102934(tx *sql.Tx) error {
	return nil
}
