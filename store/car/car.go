package store

import (
	"context"
	"database/sql"

	"github.com/amit8889/car_managemnt/models"
	"github.com/google/uuid"
)

type CarStore struct {
	DB *sql.DB
}

func New(db *sql.DB) CarStore {
	return CarStore{DB: db}
}
func (s CarStore) GetCarById(ctx context.Context, id string) (models.Car, error) {
	var car models.Car
	var engine models.Engine
	err := s.DB.QueryRowContext(ctx, `SELECT * FROM cars WHERE id = '00000000-4333-4343-4343-434343434343'`).Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuleType, &car.Engine.EngineID, &car.Price, &car.CreatedAt, &car.UpdatedAt)
	if err != nil {
		return car, err
	}
	err = s.DB.QueryRowContext(ctx, `SELECT * FROM engines WHERE engine_id = ?`, car.Engine.EngineID).Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCylinder, &engine.CarRang)
	if err != nil {
		return car, err
	}
	return car, nil

}
func (s CarStore) GetAllCars(ctx context.Context, page int32) ([]models.Car, int32, error) {
	var cars []models.Car
	var count int32
	err := s.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM cars`).Scan(&count)
	if err != nil {
		return cars, 0, err
	}
	rows, err := s.DB.QueryContext(ctx, `
		SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.Price, c.created_at, c.updated_at, 
		       e.engine_id, e.displacement, e.no_of_cylinder, e.car_range
		FROM cars c
		JOIN engines e ON c.engine_id = e.engine_id
		LIMIT ?, ?`, page*10, 10)

	if err != nil {
		return cars, count, err
	}
	defer rows.Close()

	// Process the results
	for rows.Next() {
		var car models.Car
		var engine models.Engine
		err = rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuleType, &car.Price,
			&car.CreatedAt, &car.UpdatedAt,
			&engine.EngineID, &engine.Displacement, &engine.NoOfCylinder, &engine.CarRang)
		if err != nil {
			return cars, count, err
		}
		car.Engine = engine
		cars = append(cars, car)
	}
	if err := rows.Err(); err != nil {
		return cars, count, err
	}

	return cars, count, nil
}

func (s CarStore) CreateCar(ctx context.Context, car models.Car) error {
	var engine models.Engine = car.Engine

	// Create transaction
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Insert into engines table
	_, err = tx.ExecContext(ctx, `
		INSERT INTO engines (engine_id, displacement, no_of_cylinder, car_range)
		VALUES (?, ?, ?, ?)`, engine.EngineID, engine.Displacement, engine.NoOfCylinder, engine.CarRang)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert into cars table
	_, err = tx.ExecContext(ctx, `
		INSERT INTO cars (name, year, brand, fuel_type, price, engine_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, car.Name, car.Year, car.Brand, car.FuleType, car.Price, car.Engine.EngineID, car.CreatedAt, car.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (c CarStore) DeleteCarById(ctx context.Context, carId uuid.UUID) error {

	tx, err := c.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.ExecContext(ctx, `
		DELETE FROM engines
		WHERE engine_id = (SELECT engine_id FROM cars WHERE id = ?);
	`, carId)
	if err != nil {
		return err
	}

	// Then, delete from the cars table
	_, err = tx.ExecContext(ctx, `
		DELETE FROM cars
		WHERE id = ?;
	`, carId)
	if err != nil {
		return err
	}
	return nil
}
