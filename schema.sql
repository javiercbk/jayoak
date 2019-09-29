-- CREATE DATABASE jayoak WITH OWNER 'jayoak' ENCODING 'UTF8';

CREATE TYPE note AS ENUM ('A', 'B', 'C', 'D', 'E', 'F', 'G');

CREATE TABLE organizations(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE INDEX organizations_name_idx ON organizations (name);

CREATE TABLE users(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    organization_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT fk_user_organization FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_first_name ON users (name);
CREATE UNIQUE INDEX idx_users_email_org ON users (email, organization_id);

CREATE TABLE instruments(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    organization_id BIGINT NOT NULL,
    creator_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT fk_instrument_creator FOREIGN KEY (creator_id) REFERENCES users (id),
    CONSTRAINT fk_instrument_organization FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

CREATE UNIQUE INDEX idx_instruments_user_name ON instruments (name, organization_id);

CREATE TABLE sounds(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT,
    audio_file_name TEXT NOT NULL,
    audio_uuid TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    md5_file uuid NOT NULL,
    instrument_id BIGINT,
    -- We could insert a variation of a note
    note note,
    max_frequency NUMERIC(6,0),
    min_frequency NUMERIC(6,0),
    max_power_freq NUMERIC(6,0),
    max_power_value NUMERIC(16,8),
    organization_id BIGINT NOT NULL,
    creator_id BIGINT NOT NULL,
    processed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT cnst_name_or_note CHECK (name IS NOT NULL OR (instrument_id IS NOT NULL AND note IS NOT NULL)),
    CONSTRAINT fk_sound_instrument FOREIGN KEY (instrument_id) REFERENCES instruments (id),
    CONSTRAINT fk_sound_creator FOREIGN KEY (creator_id) REFERENCES users (id),
    CONSTRAINT fk_sound_organization FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

CREATE INDEX idx_sounds_name ON sounds (name);
CREATE INDEX idx_sounds_audio_file_name ON sounds (audio_file_name);
CREATE INDEX idx_sounds_note ON sounds (note);

CREATE TABLE frequencies(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    sound_id BIGINT NOT NULL,
    frequency INT NOT NULL,
    spl NUMERIC(16, 8) NOT NULL,
    CONSTRAINT cnst_frequency CHECK (frequency >= 0),
    CONSTRAINT fk_frequency_sound FOREIGN KEY (sound_id) REFERENCES sounds (id)
);

CREATE INDEX idx_frequencies_frequency ON frequencies (frequency);
CREATE INDEX idx_frequencies_frequency_spl ON frequencies (frequency, spl);