-- This file is subject to change until the first real deployment of the bot.
-- Do not rely on these schema migrations remaining consistent until this
-- message has been removed.

BEGIN;

CREATE TABLE channels (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,

    user_id bigint NOT NULL UNIQUE,
    name text NOT NULL,
    bot_name text NOT NULL,
    active boolean NOT NULL,
    prefix text NOT NULL,
    bullet text,
    should_moderate boolean NOT NULL
);

CREATE INDEX channels_user_id_idx on channels (user_id);

CREATE TYPE access_level AS ENUM (
    'everyone',
    'subscriber',
    'moderator',
    'broadcaster',
    'admin'
);

CREATE TABLE simple_commands (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamptz DEFAULT NOW() NOT NULL,
    updated_at timestamptz DEFAULT NOW() NOT NULL,

    channel_id bigint REFERENCES channels (id) NOT NULL,

    name text NOT NULL,
    message text NOT NULL,
    editor text NOT NULL,
    access_level access_level NOT NULL,
    count bigint NOT NULL,

    UNIQUE (channel_id, name)
);

CREATE INDEX simple_commands_channel_id_idx ON simple_commands (channel_id);
CREATE INDEX simple_commands_name_idx on simple_commands (name);

COMMIT;
