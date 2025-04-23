package repository

import (
	"database/sql"

	"github.com/live-kit-video-conference/model"
)

type RoomRepository interface {
	Create(room *model.Room) (int, error)
	GetByID(id int) (*model.Room, error)
	List() ([]*model.Room, error)
	Update(room *model.Room) error
	Delete(id int) error
}

type roomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) Create(room *model.Room) (int, error) {
	query := "INSERT INTO room (name, description, owner_id) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, room.Name, room.Description, room.OwnerID)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int(id), err
}

func (r *roomRepository) GetByID(id int) (*model.Room, error) {
	query := "SELECT id, name, description, owner_id, created_at FROM room WHERE id = ?"
	row := r.db.QueryRow(query, id)
	room := &model.Room{}
	err := row.Scan(&room.ID, &room.Name, &room.Description, &room.OwnerID, &room.CreatedAt)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *roomRepository) List() ([]*model.Room, error) {
	query := "SELECT id, name, description, owner_id, created_at FROM room"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*model.Room
	for rows.Next() {
		room := &model.Room{}
		err := rows.Scan(&room.ID, &room.Name, &room.Description, &room.OwnerID, &room.CreatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (r *roomRepository) Update(room *model.Room) error {
	query := "UPDATE room SET name = ?, description = ?, owner_id = ? WHERE id = ?"
	_, err := r.db.Exec(query, room.Name, room.Description, room.OwnerID, room.ID)
	return err
}

func (r *roomRepository) Delete(id int) error {
	query := "DELETE FROM room WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
