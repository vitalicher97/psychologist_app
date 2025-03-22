
INSERT INTO psychologists (first_name, last_name, email, bio)
VALUES ('John', 'Doe', 'john.doe@example.com', 'Experienced family therapist.');

INSERT INTO specializations (name)
VALUES ('Family Therapy'), ('Cognitive-Behavioral Therapy');

INSERT INTO psychologist_specializations (psychologist_id, specialization_id)
VALUES (1, 1), (1, 2);

INSERT INTO availability (psychologist_id, day_of_week, start_time, end_time)
VALUES (1, 1, '09:00', '12:00'), (1, 3, '14:00', '18:00');

INSERT INTO consultation_pricing (psychologist_id, price, currency)
VALUES (1, 1000.00, 'UAH');

INSERT INTO appointments (psychologist_id, customer_id, start_time, end_time)
VALUES (1, 1, '10:00', '11:00');

-- Insert sample customers
INSERT INTO customers (first_name, last_name, email, phone, created_by, updated_by, created_at, updated_at) VALUES
('John', 'Doe', 'john.doe@example.com', '123-456-7890', 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Jane', 'Smith', 'jane.smith@example.com', '987-654-3210', 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Michael', 'Johnson', 'michael.johnson@example.com', '555-123-4567', 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Emily', 'Williams', 'emily.williams@example.com', '555-987-6543', 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('David', 'Brown', 'david.brown@example.com', '555-111-2222', 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

