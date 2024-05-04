CREATE TABLE IF NOT EXISTS gyms (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255),
    sub_lvl INT
);

CREATE TABLE IF NOT EXISTS schedules (
    id SERIAL PRIMARY KEY,
    gym_id INT,
    day_of_week VARCHAR(50),
    start_time TIME,
    end_time TIME,
    FOREIGN KEY (gym_id) REFERENCES gyms(id)
);
