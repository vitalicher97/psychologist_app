
CREATE DATABASE psychologist_app;

CREATE TABLE psychologists (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    profile_picture TEXT,
    bio TEXT,
	created_by INT,
	updated_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE specializations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
	created_by INT,
	updated_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE psychologist_specializations (
	id SERIAL PRIMARY KEY,
    psychologist_id INT REFERENCES psychologists(id) ON DELETE CASCADE,
    specialization_id INT REFERENCES specializations(id) ON DELETE CASCADE,
	created_by INT,
	updated_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE availabilities (
    id SERIAL PRIMARY KEY,
    psychologist_id INT REFERENCES psychologists(id) ON DELETE CASCADE,
    day_of_week INT NOT NULL CHECK (day_of_week BETWEEN 0 AND 6),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
	created_by INT,
	updated_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
    UNIQUE (psychologist_id, day_of_week, start_time, end_time)
);

CREATE TABLE consultation_pricing (
    id SERIAL PRIMARY KEY,
    psychologist_id INT REFERENCES psychologists(id) ON DELETE CASCADE,
    price DECIMAL(10, 2),
    currency VARCHAR(3),
	created_by INT,
	updated_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    psychologist_id INT REFERENCES psychologists(id) ON DELETE CASCADE,
    customer_id INT NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL,
    status VARCHAR(50),
	created_by INT,
	updated_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    UNIQUE (psychologist_id, appointment_date, start_time, end_time)
);

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    created_by INT,
    updated_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE customer_psychologist_prices (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL REFERENCES customers(id),
    psychologist_id INTEGER NOT NULL REFERENCES psychologists(id),
    fixed_price NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (customer_id, psychologist_id)
);

ALTER TABLE customer_psychologist_prices
DROP CONSTRAINT customer_psychologist_price_psychologist_id_fkey,
ADD CONSTRAINT customer_psychologist_price_psychologist_id_fkey
FOREIGN KEY (psychologist_id) REFERENCES psychologists(id) ON DELETE CASCADE;



