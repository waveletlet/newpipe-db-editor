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


## Probably relevant tables to pull from
CREATE TABLE `streams` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `service_id` INTEGER NOT NULL, `url` TEXT, `title` TEXT, `stream_type` TEXT, `duration` INTEGER, `uploader` TEXT, `thumbnail_url` TEXT);
INSERT INTO streams VALUES(787,0,'https://www.youtube.com/watch?v=tI_4NXQeEM4','After Workout Stretch Routine. 12 Minutes of the Best Stretches To Do After Exercise.','VIDEO_STREAM',842,'Caroline Jordan','https://i.ytimg.com/vi/tI_4NXQeEM4/maxresdefault.jpg');
CREATE TABLE `stream_history` (`stream_id` INTEGER NOT NULL, `access_date` INTEGER NOT NULL, `repeat_count` INTEGER NOT NULL, PRIMARY KEY(`stream_id`, `access_date`), FOREIGN KEY(`stream_id`) REFERENCES `streams`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE );
INSERT INTO stream_history VALUES(1,1537913795111,4);
CREATE TABLE `stream_state` (`stream_id` INTEGER NOT NULL, `progress_time` INTEGER NOT NULL, PRIMARY KEY(`stream_id`), FOREIGN KEY(`stream_id`) REFERENCES `streams`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE );
INSERT INTO stream_state VALUES(3,3286);
CREATE TABLE `playlists` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `name` TEXT, `thumbnail_url` TEXT);
INSERT INTO playlists VALUES(1,'yy','https://i.ytimg.com/vi/5rfII8w7gwQ/hqdefault.jpg?sqp=-oaymwEiCMQBEG5IWvKriqkDFQgBFQAAAAAYASUAAMhCPQCAokN4AQ==&rs=AOn4CLCkUufy4_LFBznXtL-Cg7fqeKEUDA');
CREATE TABLE `playlist_stream_join` (`playlist_id` INTEGER NOT NULL, `stream_id` INTEGER NOT NULL, `join_index` INTEGER NOT NULL, PRIMARY KEY(`playlist_id`, `join_index`), FOREIGN KEY(`playlist_id`) REFERENCES `playlists`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED, FOREIGN KEY(`stream_id`) REFERENCES `streams`(`uid`) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED);
INSERT INTO playlist_stream_join VALUES(4,56,0);
CREATE TABLE `remote_playlists` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `service_id` INTEGER NOT NULL, `name` TEXT, `url` TEXT, `thumbnail_url` TEXT, `uploader` TEXT, `stream_count` INTEGER);
CREATE UNIQUE INDEX `index_subscriptions_service_id_url` ON `subscriptions` (`service_id`, `url`);
CREATE INDEX `index_search_history_search` ON `search_history` (`search`);
CREATE UNIQUE INDEX `index_streams_service_id_url` ON `streams` (`service_id`, `url`);
CREATE INDEX `index_stream_history_stream_id` ON `stream_history` (`stream_id`);
CREATE INDEX `index_playlists_name` ON `playlists` (`name`);
CREATE UNIQUE INDEX `index_playlist_stream_join_playlist_id_join_index` ON `playlist_stream_join` (`playlist_id`, `join_index`);
CREATE INDEX `index_playlist_stream_join_stream_id` ON `playlist_stream_join` (`stream_id`);
CREATE INDEX `index_remote_playlists_name` ON `remote_playlists` (`name`);
CREATE UNIQUE INDEX `index_remote_playlists_service_id_url` ON `remote_playlists` (`service_id`, `url`);
