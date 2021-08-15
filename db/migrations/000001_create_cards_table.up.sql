-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS decks
(
    deck_id     UUID           NOT NULL PRIMARY KEY,
    shuffled    BOOL           NOT NULL DEFAULT FALSE,
    remaining   INTEGER        NOT NULL,
    card_codes  TEXT[]         NOT NULL
);