package presentation

import (
	"bookManagement/db"
	"bookManagement/logic"
)

type ServerManager struct {
	Db         *db.Db
	JwtManager *logic.JwtManager
}
