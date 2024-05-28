ALTER TABLE articles_votes
    ADD CONSTRAINT unique_article_author UNIQUE (article_id, author_id);

ALTER TABLE comments_votes
    ADD CONSTRAINT unique_comment_author UNIQUE (comment_id, author_id);