ALTER TABLE authors
DROP CONSTRAINT IF EXISTS articles_author_id_fkey;

DROP TABLE IF EXISTS authors;