package repository

import "database/sql"

type AppRepository struct {
	Room RoomRepository
	User UserRepository
}

func InjectRepository(db *sql.DB) *AppRepository {
	roomRepo := NewRoomRepository(db)
	userRepo := NewUserRepository(db)

	return &AppRepository{
		Room: roomRepo,
		User: userRepo,
	}
}
