package items

import (
	"api/app/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestItem(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	expectedItem := models.Item{ID: "100", Name: "algo1", Description: "algo"}
	rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow("100", "algo1", "algo")
	mock.ExpectQuery("SELECT id, name, description FROM items WHERE id = ?").WithArgs("100").WillReturnRows(rows)
	i, err := Is.Item("100")
	assert.Nil(t, err)
	assert.Equal(t, &expectedItem, i)
	assert.Nil(t, mock.ExpectationsWereMet())

}
func TestItemFail(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	mock.ExpectQuery("SELECT id, name, description FROM items WHERE id = ?").WithArgs("100").WillReturnError(errors.New("error"))
	_, err := Is.Item("100")
	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func TestItems(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	expectedItems := []*models.Item{&models.Item{ID: "100", Name: "algo1", Description: "algo"}}

	rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow("100", "algo1", "algo")
	mock.ExpectQuery("SELECT \\* FROM items").WillReturnRows(rows)
	its, err := Is.Items()
	assert.Nil(t, err)
	assert.Equal(t, expectedItems, its)
	assert.Nil(t, mock.ExpectationsWereMet())

}
func TestItemsFail(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	mock.ExpectQuery("SELECT \\* FROM items").WillReturnError(errors.New("error"))
	_, err := Is.Items()
	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestItemsFailRows(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow(nil, nil, nil)

	mock.ExpectQuery("SELECT \\* FROM items").WillReturnRows(rows)
	_, err := Is.Items()
	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreateItem(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	item := models.Item{ID: "100", Name: "algo1", Description: "algo"}

	mock.ExpectPrepare(`INSERT INTO items`)
	mock.ExpectExec(`INSERT INTO items`).WithArgs("algo1", "algo").WillReturnResult(sqlmock.NewResult(1, 1))
	err := Is.CreateItem(&item)
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func TestCreateItemFailPrepare(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	item := models.Item{ID: "100", Name: "algo1", Description: "algo"}

	mock.ExpectPrepare(`INSERT INTO items`).WillReturnError(errors.New("error")) //no llega al Exec
	err := Is.CreateItem(&item)
	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func TestCreateItemFailExec(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	item := models.Item{ID: "100", Name: "algo1", Description: "algo"}

	mock.ExpectPrepare(`INSERT INTO items`)
	mock.ExpectExec(`INSERT INTO items`).WithArgs("algo1", "algo").WillReturnError(errors.New("error"))

	err := Is.CreateItem(&item)
	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func TestDeleteItemService(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	//item := models.Item{ID: "100", Name: "algo1", Description: "algo"}
	mock.ExpectExec("DELETE FROM items WHERE id = ?").WithArgs("100").WillReturnResult(sqlmock.NewResult(1, 1))
	err := Is.DeleteItem("100")
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteItemServiceFail(t *testing.T) {
	db, mock, _ := sqlmock.New()
	Is = &ItemService{DB: db} //mockeo la db
	defer db.Close()
	//item := models.Item{ID: "100", Name: "algo1", Description: "algo"}
	mock.ExpectExec("DELETE FROM items WHERE id = ?").WithArgs("100").WillReturnError(errors.New("Error"))
	err := Is.DeleteItem("100")
	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}
