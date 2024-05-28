ALTER TABLE articles
    ADD CHECK (length(body) <= 2000);

ALTER TABLE comments
    ADD CHECK (length(body) <= 2000);