package main

import (
	"fmt"
<<<<<<< HEAD

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

type Stream struct {
	// TABLE `streams`
	UID        int    //`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	Service_ID int    //`service_id` INTEGER NOT NULL,
	Url        string //`url` TEXT,
	Title      string //`title` TEXT,
	StreamType string //`stream_type` TEXT,
	Duration   int    //`duration` INTEGER,
	Uploader   string //`uploader` TEXT,
	ThumbUrl   string //`thumbnail_url` TEXT
}

func MarshallStream(stmt *sqlite3.Stmt) (Stream, error) {
	var uid int
	var sid int
	var url string
	var title string
	var stype string
	var duration int
	var uploader string
	var turl string

	var stream Stream
	err := stmt.Scan(&uid, &sid, &url, &title, &stype, &duration, &uploader, &turl)
	if err != nil {
		return stream, err
	}
	stream = Stream{uid, sid, url, title, stype, duration, uploader, turl}
	return stream, nil
}

func GetStreams(conn *sqlite3.Conn) ([]Stream, error) {
	stmt, err := conn.Prepare("SELECT * FROM streams")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	var streams []Stream
	for {
		row, err := stmt.Step()
		if err != nil {
			fmt.Println(err)
		}
		if !row {
			break
		}

		stream, err := MarshallStream(stmt)
		if err != nil {
			fmt.Println(err)
		}
		streams = append(streams, stream)

	}
	fmt.Println(streams)
	return streams, nil
}

type Playlist struct {
	//TABLE `playlists`
	//`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `name` TEXT, `thumbnail_url` TEXT)
	//TABLE `playlist_stream_join` (`playlist_id` INTEGER NOT NULL, `stream_id` INTEGER NOT NULL, `join_index` INTEGER NOT NULL, PRIMARY KEY(`playlist_id`, `join_index`), FOREIGN KEY(`playlist_id`) REFERENCES `playlists`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED, FOREIGN KEY(`stream_id`) REFERENCES `streams`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED)
}

func MarshallPlaylist(stmt *sqlite3.Stmt) (Playlist, error) {
}

func GetPlaylists(conn *sqlite3.Conn) ([]Playlist, error) {
}

func main() {
	conn, err := sqlite3.Open("/media/dbs/newpipe.db")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	streams, err := GetStreams(conn)
	fmt.Println(streams)
}
