ALTER TABLE articles
DROP CONSTRAINT IF EXISTS articles_author_id_fkey;

DROP TABLE IF EXISTS articles;
