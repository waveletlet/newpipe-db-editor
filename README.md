# newpipe-db-editor

Manipulate the database output by NewPipe (https://github.com/TeamNewPipe/NewPipe)

I add videos to playlists pretty haphazardly and have some nebulous idea of
being able to better organize my playlists from a computer than a phone.

## Design sketch
- Command line flags for dumping db info
- Plaintext playlist export
- Interactive TUI
  - Sort bookmarked videos/playlists by various criteria
- Playlist name (default)
  - Video length
  - Video name/url
  - For showing which playlists a given video is in
  - Generate playlists from sorted output
  - Example: From an unsorted playlist of workout videos, make a "short workout"
  playlist of all videos between 10-15 minutes


## What does the database look like?

type 	|	name	|	tbl_name 	|	rootpage 	|	sql
table	|	android_metadata	|	android_metadata	|	3	|	CREATE TABLE android_metadata (locale TEXT)
table	|	room_master_table	|	room_master_table	|	4	|	CREATE TABLE room_master_table (id INTEGER PRIMARY KEY,identity_hash TEXT)
*table	|	subscriptions	|	subscriptions	|	5	|	CREATE TABLE `subscriptions` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `service_id` INTEGER NOT NULL, `url` TEXT, `name` TEXT, `avatar_url` TEXT, `subscriber_count` INTEGER, `description` TEXT)*
table	|	sqlite_sequence	|	sqlite_sequence	|	6	|	CREATE TABLE sqlite_sequence(name,seq)
index	|	index_subscriptions_service_id_url	|	subscriptions	|	7	|	CREATE UNIQUE INDEX `index_subscriptions_service_id_url` ON `subscriptions` (`service_id`, `url`)
table	|	search_history	|	search_history	|	8	|	CREATE TABLE `search_history` (`id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `creation_date` INTEGER, `service_id` INTEGER NOT NULL, `search` TEXT)
index	|	index_search_history_search	|	search_history	|	9	|	CREATE INDEX `index_search_history_search` ON `search_history` (`search`)
*table	|	streams	|	streams	|	10	|	CREATE TABLE `streams` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `service_id` INTEGER NOT NULL, `url` TEXT, `title` TEXT, `stream_type` TEXT, `duration` INTEGER, `uploader` TEXT, `thumbnail_url` TEXT)*
index	|	index_streams_service_id_url	|	streams	|	11	|	CREATE UNIQUE INDEX `index_streams_service_id_url` ON `streams` (`service_id`, `url`)
table	|	stream_history	|	stream_history	|	12	|	CREATE TABLE `stream_history` (`stream_id` INTEGER NOT NULL, `access_date` INTEGER NOT NULL, `repeat_count` INTEGER NOT NULL, PRIMARY KEY(`stream_id`, `access_date`), FOREIGN KEY(`stream_id`) REFERENCES `streams`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE )
index	|	sqlite_autoindex_stream_history_1	|	stream_history	|	13	|	
index	|	index_stream_history_stream_id	|	stream_history	|	14	|	CREATE INDEX `index_stream_history_stream_id` ON `stream_history` (`stream_id`)
table	|	stream_state	|	stream_state	|	15	|	CREATE TABLE `stream_state` (`stream_id` INTEGER NOT NULL, `progress_time` INTEGER NOT NULL, PRIMARY KEY(`stream_id`), FOREIGN KEY(`stream_id`) REFERENCES `streams`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE )
*table	|	playlists	|	playlists	|	16	|	CREATE TABLE `playlists` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `name` TEXT, `thumbnail_url` TEXT)*
index	|	index_playlists_name	|	playlists	|	17	|	CREATE INDEX `index_playlists_name` ON `playlists` (`name`)
*table	|	playlist_stream_join	|	playlist_stream_join	|	18	|	CREATE TABLE `playlist_stream_join` (`playlist_id` INTEGER NOT NULL, `stream_id` INTEGER NOT NULL, `join_index` INTEGER NOT NULL, PRIMARY KEY(`playlist_id`, `join_index`), FOREIGN KEY(`playlist_id`) REFERENCES `playlists`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED, FOREIGN KEY(`stream_id`) REFERENCES `streams`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED)*
index	|	sqlite_autoindex_playlist_stream_join_1	|	playlist_stream_join	|	19	|	
index	|	index_playlist_stream_join_playlist_id_join_index	|	playlist_stream_join	|	20	|	CREATE UNIQUE INDEX `index_playlist_stream_join_playlist_id_join_index` ON `playlist_stream_join` (`playlist_id`, `join_index`)
index	|	index_playlist_stream_join_stream_id	|	playlist_stream_join	|	21	|	CREATE INDEX `index_playlist_stream_join_stream_id` ON `playlist_stream_join` (`stream_id`)
table	|	remote_playlists	|	remote_playlists	|	22	|	CREATE TABLE `remote_playlists` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `service_id` INTEGER NOT NULL, `name` TEXT, `url` TEXT, `thumbnail_url` TEXT, `uploader` TEXT, `stream_count` INTEGER)
index	|	index_remote_playlists_name	|	remote_playlists	|	23	|	CREATE INDEX `index_remote_playlists_name` ON `remote_playlists` (`name`)
index	|	index_remote_playlists_service_id_url	|	remote_playlists	|	24	|	CREATE UNIQUE INDEX `index_remote_playlists_service_id_url` ON `remote_playlists` (`service_id`, `url`)

