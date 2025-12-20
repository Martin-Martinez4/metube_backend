CREATE EXTENSION IF NOT EXISTS pg_trgm;
SET pg_trgm.similarity_threshold = 0.2;

CREATE INDEX contentinformation_title_trgm_idx
ON contentinformation
USING GIN (title gin_trgm_ops);

ALTER TABLE contentinformation
ADD COLUMN tags text[] NOT NULL DEFAULT '{}';

CREATE INDEX ci_tags_gin_idx
ON contentinformation
USING GIN (tags);

ALTER TABLE video DROP COLUMN tags;

CREATE INDEX profile_displayname_gin_indx
ON profile
USING GIN(displayname gin_trgm_ops);
