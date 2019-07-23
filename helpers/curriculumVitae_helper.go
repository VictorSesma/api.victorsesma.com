package helpers

import (
	"database/sql"
	"log"

	"github.com/leviatan89/api.victorsesma.com/types"
)

// GetCurriculumVitae will get a list of all the live events in the DB
func GetCurriculumVitae(db *sql.DB) (*types.LifeEvents, error) {
	query := `
		SELECT description, end_date, name, show_order, start_date, summary
		FROM curriculum_vitae
		WHERE section_type = 'work_experience'
		ORDER BY show_order;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	result := types.LifeEvents{}

	for rows.Next() {
		le := types.LifeEvent{}
		err := rows.Scan(&le.Description, &le.EndDate, &le.Name, &le.ShownOrder, &le.StartDate, &le.Summary)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, le)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &result, nil
}
