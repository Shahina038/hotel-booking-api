CREATE TYPE booking_status AS ENUM ('pending', 'confirmed', 'cancelled');

CREATE TABLE IF NOT EXISTS bookings (
    id          SERIAL PRIMARY KEY,
    user_id     INT             NOT NULL REFERENCES users(id)  ON DELETE CASCADE,
    room_id     INT             NOT NULL REFERENCES rooms(id)  ON DELETE CASCADE,
    check_in    DATE            NOT NULL,
    check_out   DATE            NOT NULL,
    total_price NUMERIC(10, 2)  NOT NULL,
    status      booking_status  NOT NULL DEFAULT 'pending',
    created_at  TIMESTAMP       NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP       NOT NULL DEFAULT NOW(),

    CONSTRAINT check_dates CHECK (check_out > check_in)
);
