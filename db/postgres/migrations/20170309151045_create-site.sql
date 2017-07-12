-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE leave_words (
  id         BIGSERIAL PRIMARY KEY,
  body       TEXT                        NOT NULL,
  type       VARCHAR(8)                  NOT NULL DEFAULT 'markdown',
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);


CREATE TABLE links (
  id BIGSERIAL PRIMARY KEY,
  href VARCHAR(255) NOT NULL,
  label VARCHAR(255) NOT NULL,
  loc VARCHAR(16) NOT NULL,
  sort_order INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_links_loc ON links (loc);

CREATE TABLE cards (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  summary VARCHAR(2048) NOT NULL,
  action VARCHAR(32) NOT NULL,
  href VARCHAR(255) NOT NULL,
  logo VARCHAR(255) NOT NULL,
  loc VARCHAR(16) NOT NULL,
  sort_order INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_cards_loc ON cards (loc);

CREATE TABLE votes (
  id BIGSERIAL PRIMARY KEY,
  resource_type VARCHAR(255) NOT NULL,
  resource_id BIGINT NOT NULL,
  points INT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX idx_votes_resource_type ON votes (resource_type);
CREATE UNIQUE INDEX idx_votes_resource ON votes (resource_type, resource_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE votes;
DROP TABLE cards;
DROP TABLE links;
DROP TABLE leave_words;
