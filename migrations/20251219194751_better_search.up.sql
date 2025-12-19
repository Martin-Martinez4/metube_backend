CREATE EXTENSION IF NOT EXISTS pg_trgm;
SET pg_trgm.similarity_threshold = 0.2;

CREATE INDEX contentinformation_title_trgm_idx
ON contentinformation
USING GIN (title gin_trgm_ops);

