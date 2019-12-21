package main

import (
	"fmt"

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

func GetStreams(conn *sqlite3.Conn) (map[int]*Stream, error) {
	streams := make(map[int]*Stream)
	// revisit whether i want a map or maybe it should be a slice with the UID as
	// index
	stmt, err := conn.Prepare("SELECT * FROM streams")
	if err != nil {
		return streams, err
	}
	defer stmt.Close()

	for {
		row, err := stmt.Step()
		if err != nil {
			return streams, err
		}
		if !row {
			break
		}

		stream, err := MarshallStream(stmt)
		if err != nil {
			return streams, err
		}
		streams[stream.UID] = &stream

	}
	return streams, nil
}

type Playlist struct {
	//TABLE `playlists`
	UID      int    //`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL
	Name     string //`name` TEXT
	ThumbUrl string //`thumbnail_url` TEXT

	//TABLE `playlist_stream_join`
	StreamList []*Stream //`stream_id` INTEGER NOT NULL
	// possibly want to do this differently
}

func MarshallPlaylist(stmt *sqlite3.Stmt) (Playlist, error) {
	var uid int
	var name string
	var thurl string
	var playlist Playlist

	err := stmt.Scan(&uid, &name, &thurl)
	if err != nil {
		return playlist, err
	}
	playlist = Playlist{uid, name, thurl, []*Stream{}}
	return playlist, nil
}

func GetPlaylists(conn *sqlite3.Conn, streams map[int]*Stream) (map[int]Playlist, error) {
	playlists := make(map[int]Playlist)
	stmt, err := conn.Prepare("SELECT * FROM playlists")
	if err != nil {
		return playlists, err
	}
	defer stmt.Close()

	for {
		row, err := stmt.Step()
		if err != nil {
			return playlists, err
		}
		if !row {
			break
		}

		playlist, err := MarshallPlaylist(stmt)
		if err != nil {
			return playlists, err
		}

		st, err := conn.Prepare(fmt.Sprintf("SELECT * FROM playlist_stream_join WHERE playlist_id = %v", playlist.UID))
		if err != nil {
			return playlists, err
		}
		for {
			row, err := st.Step()
			if err != nil {
				return playlists, err
			}
			if !row {
				break
			}

			var pid int
			var sid int
			var idx int
			err = st.Scan(&pid, &sid, &idx)
			if err != nil {
				return playlists, err
			}
			if len(playlist.StreamList) == idx {
				playlist.StreamList = append(playlist.StreamList, streams[sid])
			} else {
				fmt.Println("!!!!!! SHIT OUT OF ORDER !!!!!!!!!!!!!!")
				// I don't really expect to see this because I *think* everything is in
				// index order in the table, but I'm not sure so it should at least warn
			}
		}

		playlists[playlist.UID] = playlist

	}
	return playlists, nil
}

func main() {
	conn, err := sqlite3.Open("/media/dbs/newpipe.db")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	streams, err := GetStreams(conn)
	if err != nil {
		fmt.Println(err)
	}
	playlists, err := GetPlaylists(conn, streams)
	if err != nil {
		fmt.Println(err)
	}

	for _, playlist := range playlists {
		fmt.Printf("Playlist: %s\n", playlist.Name)
		for i, stream := range playlist.StreamList {
			fmt.Printf("%v: %s\n", i, stream.Title)
		}
	}

}
