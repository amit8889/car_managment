-- Creating the engine table
CREATE TABLE IF NOT EXISTS engine (
    engine_id INT PRIMARY KEY,  
    displacement FLOAT NOT NULL, 
    no_of_cylinder INT NOT NULL, 
    car_range FLOAT NOT NULL,
    CONSTRAINT chk_displacement CHECK (displacement > 0),
    CONSTRAINT chk_no_of_cylinder CHECK (no_of_cylinder >= 1),
    CONSTRAINT chk_car_range CHECK (car_range > 0)
);

-- Creating the car table
CREATE TABLE IF NOT EXISTS car (
    id INT PRIMARY KEY, 
    name VARCHAR(255) NOT NULL,  
    year VARCHAR(4) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    fuel_type VARCHAR(255) NOT NULL,
    engine_id INT NOT NULL,  -- Foreign key referencing engine_id
    price FLOAT NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT chk_fuel_type CHECK (fuel_type IN ('petrol', 'diesel', 'electric', 'hybrid')),
    CONSTRAINT chk_price CHECK (price > 0)
);


INSERT INTO engine (engine_id, displacement, no_of_cylinder, car_range) VALUES
(1, 2.0, 4, 500.0),
(2, 3.5, 6, 650.0),
(3, 4.0, 8, 700.0);
-- Inserting data into car table
INSERT INTO car (id, name, year, brand, fuel_type, engine_id, price) VALUES
(1, 'Model X', '2023', 'Tesla', 'electric', 1, 75000.0),
(2, 'Civic', '2022', 'Honda', 'petrol', 2, 25000.0),
(3, 'Mustang', '2021', 'Ford', 'petrol', 3, 55000.0);