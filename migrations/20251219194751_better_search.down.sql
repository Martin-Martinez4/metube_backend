-- Drop indexes (reverse order of creation)
DROP INDEX IF EXISTS profile_displayname_gin_indx;
DROP INDEX IF EXISTS ci_tags_gin_idx;
DROP INDEX IF EXISTS contentinformation_title_trgm_idx;

-- Re-add tags column to video (since UP dropped it)
ALTER TABLE video
ADD COLUMN IF NOT EXISTS tags jsonb;

-- Remove tags column from contentinformation
ALTER TABLE contentinformation
DROP COLUMN IF EXISTS tags;

-- We intentionally do NOT drop pg_trgm extension
