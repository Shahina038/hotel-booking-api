CREATE TYPE room_type AS ENUM ('single', 'double', 'suite', 'deluxe');

CREATE TABLE IF NOT EXISTS rooms (
    id              SERIAL PRIMARY KEY,
    room_number     VARCHAR(10)     UNIQUE NOT NULL,
    type            room_type       NOT NULL,
    description     TEXT,
    price_per_night NUMERIC(10, 2)  NOT NULL,
    capacity        INT             NOT NULL DEFAULT 1,
    is_available    BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMP       NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP       NOT NULL DEFAULT NOW()
);
