package uuid

import (
	"github.com/google/uuid"
)

type GeneratorInterface interface {
	GenerateUUID() (string, error)
	GenerateOrderID() (string, error)
}

// type Generator struct {
// 	db *sql.DB
// }

// func NewGeneratorUUID(db *sql.DB) *Generator {
// 	return &Generator{
// 		db: db,
// 	}
// }

func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// exists, err := g.checkIDExists(id.String())
	// if err != nil {
	// 	return "", err
	// }

	// if exists {
	// 	return g.GenerateUUID()
	// }

	return id.String(), nil
}

// func (g *Generator) checkIDExists(ID string) (bool, error) {
// 	var count int64
// 	err := g.db.QueryRow("SELECT COUNT(*) FROM staffs WHERE id = $1", ID).Scan(&count)
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }
