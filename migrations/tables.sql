
-- Table: subscriber_subscribee
IF EXISTS DROP subscriber_subscribee;

-- Table: profile_view
IF EXISTS DROP profile_view;

-- Table: profile_video_like_dislike
IF EXISTS DROP profile_video_like_dislike;

-- Table: profile_comment_like_dislike
IF EXISTS DROP profile_comment_like_dislike ;

-- Table: profile_comment_mention
IF EXISTS DROP profile_comment_mention;

-- Table: profile_video
IF EXISTS DROP profile_video;

-- Table: profile
IF EXISTS DROP profile;

-- Table: video
IF EXISTS DROP video;

-- Table: statistic
IF EXISTS DROP statistic;

-- Table: contentinformation
IF EXISTS DROP contentinformation;

-- Table: login
IF EXISTS DROP login ;

-- Table: comment
IF EXISTS DROP comment ;


-- Table: thumbnail
IF EXISTS DROP thumbnail;

-- Table: status
IF EXISTS DROP status;


DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uploadstatus') THEN
        DROP TYPE uploadstatus;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'privacystatus') THEN
        DROP TYPE privacystatus;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'like_dislike') THEN
        DROP TYPE like_dislike;
    END IF;
END$$;


