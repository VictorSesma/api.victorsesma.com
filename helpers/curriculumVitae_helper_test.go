package helpers_test

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/leviatan89/api.victorsesma.com/helpers"
	"github.com/leviatan89/api.victorsesma.com/types"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCurriculumVitaes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.NoError(t, err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"description", "end_date", "name", "show_order", "start_date", "summary"}).
		AddRow("Volunteer in AEGEE", "2019-08-08 08:08:08", "Aegee coordinator", "1", "2019-08-08 08:08:08", "summary cool").
		AddRow("Volunteer in AEGEE 2", "2019-08-08 08:08:08", "Aegee coordinator 2", "2", "2019-08-08 08:08:08", "summary cool 2")

	mock.ExpectQuery("^SELECT (.+) FROM curriculum_vitae (.+)$").WillReturnRows(rows)

	result, err := helpers.GetCurriculumVitae(db)
	assert.NoError(t, err)

	expected := types.LifeEvents{
		types.LifeEvent{
			Description: "Volunteer in AEGEE",
			EndDate:     "2019-08-08 08:08:08",
			Name:        "Aegee coordinator",
			ShownOrder:  "1",
			StartDate:   "2019-08-08 08:08:08",
			Summary:     "summary cool",
		},
		types.LifeEvent{
			Description: "Volunteer in AEGEE 2",
			EndDate:     "2019-08-08 08:08:08",
			Name:        "Aegee coordinator 2",
			ShownOrder:  "2",
			StartDate:   "2019-08-08 08:08:08",
			Summary:     "summary cool 2",
		},
	}
	assert.Equal(t, expected, *result)
}
