-- Explicit missing-file prune.
--
-- This script is run only when an admin user invokes the "Prune Missing"
-- library action. It removes database rows for files and aggregates that no
-- longer exist in the library.

-- 1. Delete tracks that the scanner marked missing.
DELETE FROM media_file
WHERE missing = TRUE;

-- 2. Delete albums that no longer have any tracks.
DELETE FROM album
WHERE NOT EXISTS (
    SELECT 1
    FROM media_file
    WHERE media_file.album_id = album.id
);

-- 3. Delete track-artist links that no longer point to existing tracks.
DELETE FROM media_file_artists
WHERE NOT EXISTS (
    SELECT 1
    FROM media_file
    WHERE media_file.id = media_file_artists.media_file_id
);

-- 4. Delete album-artist links that no longer point to existing albums.
DELETE FROM album_artists
WHERE NOT EXISTS (
    SELECT 1
    FROM album
    WHERE album.id = album_artists.album_id
);

-- 5. Delete artists that are no longer referenced as album artists or primary
-- track artists.
DELETE FROM artist
WHERE NOT EXISTS (
    SELECT 1
    FROM album_artists
    WHERE album_artists.artist_id = artist.id
      AND album_artists.role = 'albumartist'
)
AND NOT EXISTS (
    SELECT 1
    FROM media_file_artists
    WHERE media_file_artists.artist_id = artist.id
      AND media_file_artists.role = 'artist'
);
