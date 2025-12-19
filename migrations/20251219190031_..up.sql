

--      Column     |           Type           | Collation | Nullable |      Default
-- ----------------+--------------------------+-----------+----------+--------------------
--  id             | uuid                     |           | not null | uuid_generate_v4()
--  date_posted    | timestamp with time zone |           | not null |
--  body           | character varying(8000)  |           | not null |
--  video_id       | uuid                     |           |          |
--  profile_id     | uuid                     |           |          |
--  parent_comment | uuid                     |           |          |
--  likes          | integer                  |           | not null | 0
--  dislikes       | integer                  |           | not null | 0
--  responses      | integer                  |           | not null | 0
-- Indexes:
--     "comment_pkey" PRIMARY KEY, btree (id)
-- Foreign-key constraints:
--     "comment_parent_comment_fkey" FOREIGN KEY (parent_comment) REFERENCES comment(id)
--     "comment_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     "comment_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)
-- Referenced by:
--     TABLE "comment" CONSTRAINT "comment_parent_comment_fkey" FOREIGN KEY (parent_comment) REFERENCES comment(id)
--     TABLE "profile_comment_like_dislike" CONSTRAINT "profile_comment_like_dislike_comment_id_fkey" FOREIGN KEY (comment_id) REFERENCES comment(id)
--     TABLE "profile_comment_mention" CONSTRAINT "profile_comment_mention_comment_id_fkey" FOREIGN KEY (comment_id) REFERENCES comment(id)



--                    Table "public.contentinformation"
--    Column    |          Type           | Collation | Nullable | Default
-- -------------+-------------------------+-----------+----------+---------
--  title       | character varying(300)  |           | not null |
--  description | character varying(6000) |           | not null |
--  channelid   | character varying(100)  |           | not null |
--  published   | character varying(100)  |           | not null |
--  video_id    | uuid                    |           |          |
-- Foreign-key constraints:
--     "contentinformation_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)

--                          Table "public.login"
--    Column   |          Type          | Collation | Nullable | Default
-- ------------+------------------------+-----------+----------+---------
--  profile_id | uuid                   |           |          |
--  password   | character varying(100) |           |          |
-- Foreign-key constraints:
--     "login_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)

--                              Table "public.profile"
--    Column    |         Type          | Collation | Nullable |      Default
-- -------------+-----------------------+-----------+----------+--------------------
--  id          | uuid                  |           | not null | uuid_generate_v4()
--  username    | character varying(60) |           | not null |
--  displayname | character varying(60) |           | not null |
--  ischannel   | boolean               |           | not null | false
--  subscribers | integer               |           |          |
-- Indexes:
--     "profile_pkey" PRIMARY KEY, btree (id)
--     "profile_displayname_key" UNIQUE CONSTRAINT, btree (displayname)
--     "profile_username_key" UNIQUE CONSTRAINT, btree (username)
-- Referenced by:
--     TABLE "comment" CONSTRAINT "comment_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     TABLE "login" CONSTRAINT "login_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     TABLE "profile_comment_like_dislike" CONSTRAINT "profile_comment_like_dislike_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     TABLE "profile_comment_mention" CONSTRAINT "profile_comment_mention_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     TABLE "profile_video_like_dislike" CONSTRAINT "profile_video_like_dislike_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     TABLE "profile_video" CONSTRAINT "profile_video_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     TABLE "profile_view" CONSTRAINT "profile_view_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     TABLE "subscriber_subscribee" CONSTRAINT "subscriber_subscribee_subscribee_id_fkey" FOREIGN KEY (subscribee_id) REFERENCES profile(id)
--     TABLE "subscriber_subscribee" CONSTRAINT "subscriber_subscribee_subscriber_id_fkey" FOREIGN KEY (subscriber_id) REFERENCES profile(id)
--     TABLE "video" CONSTRAINT "video_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)


--         Table "public.profile_comment_like_dislike"
--    Column   |     Type     | Collation | Nullable | Default
-- ------------+--------------+-----------+----------+---------
--  profile_id | uuid         |           | not null |
--  comment_id | uuid         |           | not null |
--  status     | like_dislike |           | not null |
-- Indexes:
--     "profile_comment_like_dislike_pkey" PRIMARY KEY, btree (profile_id, comment_id)
-- Foreign-key constraints:
--     "profile_comment_like_dislike_comment_id_fkey" FOREIGN KEY (comment_id) REFERENCES comment(id)
--     "profile_comment_like_dislike_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)


--        Table "public.profile_comment_mention"
--    Column   | Type | Collation | Nullable | Default
-- ------------+------+-----------+----------+---------
--  profile_id | uuid |           | not null |
--  comment_id | uuid |           | not null |
-- Indexes:
--     "profile_comment_mention_pkey" PRIMARY KEY, btree (profile_id, comment_id)
-- Foreign-key constraints:
--     "profile_comment_mention_comment_id_fkey" FOREIGN KEY (comment_id) REFERENCES comment(id)
--     "profile_comment_mention_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)

--                               Table "public.video"
--    Column   |          Type          | Collation | Nullable |      Default
-- ------------+------------------------+-----------+----------+--------------------
--  id         | uuid                   |           | not null | uuid_generate_v4()
--  url        | character varying(500) |           | not null |
--  categoryid | character varying(100) |           |          |
--  duration   | integer                |           | not null |
--  tags       | jsonb                  |           |          |
--  profile_id | uuid                   |           |          |
--  comments   | integer                |           | not null | 0
--  visible    | boolean                |           | not null | true
-- Indexes:
--     "video_pkey" PRIMARY KEY, btree (id)
--     "video_url_key" UNIQUE CONSTRAINT, btree (url)
-- Foreign-key constraints:
--     "video_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
-- Referenced by:
--     TABLE "comment" CONSTRAINT "comment_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)
--     TABLE "contentinformation" CONSTRAINT "contentinformation_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)
--     TABLE "profile_video_like_dislike" CONSTRAINT "profile_video_like_dislike_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)
--     TABLE "profile_video" CONSTRAINT "profile_video_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)
--     TABLE "profile_view" CONSTRAINT "profile_view_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)
--     TABLE "statistic" CONSTRAINT "statistic_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)
--     TABLE "status" CONSTRAINT "status_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)
--     TABLE "thumbnail" CONSTRAINT "thumbnail_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)

--                       Table "public.thumbnail"
--   Column  |          Type          | Collation | Nullable | Default
-- ----------+------------------------+-----------+----------+---------
--  url      | character varying(500) |           | not null |
--  video_id | uuid                   |           |          |
-- Foreign-key constraints:
--     "thumbnail_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)

-- uploadstatus
--  enum_value
-- ------------
--  processing
--  error
--  complete

-- privacystatus
--  enum_value
-- ------------
--  private
--  public
-- (2 rows)
--                      Table "public.status"
--     Column     |     Type      | Collation | Nullable | Default
-- ---------------+---------------+-----------+----------+---------
--  uploadstatus  | uploadstatus  |           | not null |
--  privacystatus | privacystatus |           | not null |
--  video_id      | uuid          |           |          |
-- Foreign-key constraints:
--     "status_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)


--          Table "public.subscriber_subscribee"
--     Column     | Type | Collation | Nullable | Default
-- ---------------+------+-----------+----------+---------
--  subscriber_id | uuid |           | not null |
--  subscribee_id | uuid |           | not null |
-- Indexes:
--     "subscriber_subscribee_pkey" PRIMARY KEY, btree (subscriber_id, subscribee_id)
-- Foreign-key constraints:
--     "subscriber_subscribee_subscribee_id_fkey" FOREIGN KEY (subscribee_id) REFERENCES profile(id)
--     "subscriber_subscribee_subscriber_id_fkey" FOREIGN KEY (subscriber_id) REFERENCES profile(id)


--             Table "public.profile_view"
--    Column   | Type | Collation | Nullable | Default
-- ------------+------+-----------+----------+---------
--  profile_id | uuid |           | not null |
--  video_id   | uuid |           | not null |
-- Indexes:
--     "profile_view_pkey" PRIMARY KEY, btree (profile_id, video_id)
-- Foreign-key constraints:
--     "profile_view_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     "profile_view_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)

--         Table "public.profile_video_like_dislike"
--    Column   |     Type     | Collation | Nullable | Default
-- ------------+--------------+-----------+----------+---------
--  profile_id | uuid         |           | not null |
--  video_id   | uuid         |           | not null |
--  status     | like_dislike |           | not null |
-- Indexes:
--     "profile_video_like_dislike_pkey" PRIMARY KEY, btree (profile_id, video_id)
-- Foreign-key constraints:
--     "profile_video_like_dislike_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     "profile_video_like_dislike_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)


--             Table "public.profile_video"
--    Column   | Type | Collation | Nullable | Default
-- ------------+------+-----------+----------+---------
--  profile_id | uuid |           | not null |
--  video_id   | uuid |           | not null |
-- Indexes:
--     "profile_video_pkey" PRIMARY KEY, btree (profile_id, video_id)
-- Foreign-key constraints:
--     "profile_video_profile_id_fkey" FOREIGN KEY (profile_id) REFERENCES profile(id)
--     "profile_video_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)

--                Table "public.statistic"
--   Column   |  Type   | Collation | Nullable | Default 
-- -----------+---------+-----------+----------+---------
--  likes     | integer |           | not null | 0
--  dislikes  | integer |           | not null | 0
--  views     | integer |           | not null | 0
--  favorites | integer |           |          | 0
--  comments  | integer |           | not null | 0
--  video_id  | uuid    |           |          |
-- Foreign-key constraints:
--     "statistic_video_id_fkey" FOREIGN KEY (video_id) REFERENCES video(id)

-- Migration File: 001_initial_schema.sql
-- Create required extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enums
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uploadstatus') THEN
        CREATE TYPE uploadstatus AS ENUM ('processing', 'error', 'complete');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'privacystatus') THEN
        CREATE TYPE privacystatus AS ENUM ('private', 'public');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'like_dislike') THEN
        CREATE TYPE like_dislike AS ENUM ('like', 'dislike');
    END IF;
END$$;

-- Table: profile
CREATE TABLE IF NOT EXISTS profile (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username varchar(60) NOT NULL UNIQUE,
    displayname varchar(60) NOT NULL UNIQUE,
    ischannel boolean NOT NULL DEFAULT false,
    subscribers integer
);

-- Table: video
CREATE TABLE IF NOT EXISTS video (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    url varchar(500) NOT NULL UNIQUE,
    categoryid varchar(100),
    duration integer NOT NULL,
    tags jsonb,
    profile_id uuid REFERENCES profile(id),
    comments integer NOT NULL DEFAULT 0,
    visible boolean NOT NULL DEFAULT true
);


-- Table: statistic
CREATE TABLE IF NOT EXISTS statistic (
    video_id uuid PRIMARY KEY REFERENCES video(id),
    likes integer NOT NULL DEFAULT 0,
    dislikes integer NOT NULL DEFAULT 0,
    views integer NOT NULL DEFAULT 0,
    favorites integer DEFAULT 0,
    comments integer NOT NULL DEFAULT 0
);

-- Table: contentinformation
CREATE TABLE IF NOT EXISTS contentinformation (
    title varchar(300) NOT NULL,
    description varchar(6000) NOT NULL,
    channelid varchar(100) NOT NULL,
    published varchar(100) NOT NULL,
    video_id uuid REFERENCES video(id)
);

-- Table: login
CREATE TABLE IF NOT EXISTS login (
    profile_id uuid REFERENCES profile(id),
    password varchar(100)
);

-- Table: comment
CREATE TABLE IF NOT EXISTS comment (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    date_posted timestamptz NOT NULL,
    body varchar(8000) NOT NULL,
    video_id uuid REFERENCES video(id),
    profile_id uuid REFERENCES profile(id),
    parent_comment uuid REFERENCES comment(id),
    likes integer NOT NULL DEFAULT 0,
    dislikes integer NOT NULL DEFAULT 0,
    responses integer NOT NULL DEFAULT 0
);

-- Table: profile_comment_like_dislike
CREATE TABLE IF NOT EXISTS profile_comment_like_dislike (
    profile_id uuid NOT NULL REFERENCES profile(id),
    comment_id uuid NOT NULL REFERENCES comment(id),
    status like_dislike NOT NULL,
    PRIMARY KEY(profile_id, comment_id)
);

-- Table: profile_comment_mention
CREATE TABLE IF NOT EXISTS profile_comment_mention (
    profile_id uuid NOT NULL REFERENCES profile(id),
    comment_id uuid NOT NULL REFERENCES comment(id),
    PRIMARY KEY(profile_id, comment_id)
);

-- Table: thumbnail
CREATE TABLE IF NOT EXISTS thumbnail (
    url varchar(500) NOT NULL,
    video_id uuid REFERENCES video(id)
);

-- Table: status
CREATE TABLE IF NOT EXISTS status (
    uploadstatus uploadstatus NOT NULL,
    privacystatus privacystatus NOT NULL,
    video_id uuid REFERENCES video(id)
);

-- Table: subscriber_subscribee
CREATE TABLE IF NOT EXISTS subscriber_subscribee (
    subscriber_id uuid NOT NULL REFERENCES profile(id),
    subscribee_id uuid NOT NULL REFERENCES profile(id),
    PRIMARY KEY(subscriber_id, subscribee_id)
);

-- Table: profile_view
CREATE TABLE IF NOT EXISTS profile_view (
    profile_id uuid NOT NULL REFERENCES profile(id),
    video_id uuid NOT NULL REFERENCES video(id),
    PRIMARY KEY(profile_id, video_id)
);

-- Table: profile_video_like_dislike
CREATE TABLE IF NOT EXISTS profile_video_like_dislike (
    profile_id uuid NOT NULL REFERENCES profile(id),
    video_id uuid NOT NULL REFERENCES video(id),
    status like_dislike NOT NULL,
    PRIMARY KEY(profile_id, video_id)
);

-- Table: profile_video
CREATE TABLE IF NOT EXISTS profile_video (
    profile_id uuid NOT NULL REFERENCES profile(id),
    video_id uuid NOT NULL REFERENCES video(id),
    PRIMARY KEY(profile_id, video_id)
);
