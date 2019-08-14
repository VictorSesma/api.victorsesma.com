package services

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/leviatan89/api.victorsesma.com/errors"
	"github.com/leviatan89/api.victorsesma.com/helpers"
	"github.com/leviatan89/api.victorsesma.com/types"
)

// CurriculumVitae is a handler for all curriculum vitae services
type CurriculumVitae struct {
	DB *sql.DB
}

// GetAll will fetch all the Curriculum Vitae events
func (c *CurriculumVitae) GetAll(r *http.Request, args *string, reply *types.LifeEvents) error {

	response, err := helpers.GetCurriculumVitae(c.DB)
	if err != nil {
		log.Println(errors.ErrBlogDB)
		return errors.ErrBlogDB
	}
	*reply = *response
	return nil
}
