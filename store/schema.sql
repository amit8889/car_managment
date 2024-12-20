-- Creating the engine table
CREATE TABLE IF NOT EXISTS engines (
    engine_id VARCHAR(36)  UNIQUE,  
    displacement FLOAT NOT NULL, 
    no_of_cylinder INT NOT NULL, 
    car_range FLOAT NOT NULL,
    CONSTRAINT chk_displacement CHECK (displacement > 0),
    CONSTRAINT chk_no_of_cylinder CHECK (no_of_cylinder >= 1),
    CONSTRAINT chk_car_range CHECK (car_range > 0)
);

-- Creating the car table
CREATE TABLE IF NOT EXISTS cars (
    id VARCHAR(36) UNIQUE, 
    name VARCHAR(255) NOT NULL,  
    year VARCHAR(4) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    fuel_type VARCHAR(255) NOT NULL,
    engine_id VARCHAR(36) NOT NULL,  -- Foreign key referencing engine_id
    price FLOAT NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT chk_fuel_type CHECK (fuel_type IN ('petrol', 'diesel', 'electric', 'hybrid')),
    CONSTRAINT chk_price CHECK (price > 0)
);


-- Inserting data into the engines table
INSERT INTO engines (engine_id, displacement, no_of_cylinder, car_range) VALUES
('00000000-4333-4394-4343-434343434343', 2.0, 4, 500.0),
('00000000-4333-4944-4343-434343434343', 3.5, 6, 650.0),
('00000000-4633-4344-4343-434343434343', 4.0, 8, 700.0);

-- Inserting data into the cars table
INSERT INTO cars (id, name, year, brand, fuel_type, engine_id, price) VALUES
('00000000-4333-4343-4343-434343434343', 'Model X', '2023', 'Tesla', 'electric', '00000000-4333-4394-4343-434343434343', 75000.0),
('00000000-4333-4343-4243-434343434343', 'Civic', '2022', 'Honda', 'petrol', '00000000-4333-4944-4343-434343434343', 25000.0),
('00000000-4333-4343-4543-434343434343', 'Mustang', '2021', 'Ford', 'petrol', '00000000-4633-4344-4343-434343434343', 55000.0);
