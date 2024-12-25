CREATE TABLE IF NOT EXSISTS bookings (
    id SERIAL PRIMARY KEY,
    client VARCHAR(100),
    hotel VARCHAR(100),
    checkin DATE,
    checkout DATE
);
