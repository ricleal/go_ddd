#!/bin/sh
sqlite3 test.db <<EOF
DROP TABLE IF EXISTS person;
CREATE TABLE person (
    -- uuid
    id text PRIMARY KEY,
    name text NOT NULL,
    age integer NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- trigger (updated_at)
CREATE TRIGGER tg_person_updated_at
AFTER UPDATE
ON person FOR EACH ROW
BEGIN
  UPDATE person SET updated_at = current_timestamp
    WHERE id = old.id;
END;
-- random data
INSERT INTO person (id, name, age) VALUES (uuid(), 'Mr. John Doe', 20);
INSERT INTO person (id, name, age) VALUES (uuid(), 'Ms. Jane Doe', 40);
.headers ON
select * from person;
EOF
