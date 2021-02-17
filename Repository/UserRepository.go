package Repository

import (
	"GoApp/Config"
	"GoApp/Model"
)

type UserRepository struct {

}


func  (*UserRepository) GetUserInfo (document string) (Model.User, error) {
	db, err := Config.GetConnection()
	if err != nil {
		return createUserErrorResponse(Model.DB_UNREACHABLE, err)
	}
	userQuery := "SELECT id, document, name, last_name FROM user WHERE document = (?)"
	stmt, err := db.Prepare(userQuery)
	defer db.Close()
	defer stmt.Close()

	if err != nil {
		return createUserErrorResponse(Model.DB_ERROR, err)
	}

	row := stmt.QueryRow(document)
	if row.Err() != nil {
		return createUserErrorResponse(Model.DB_ERROR, err)
	}

	var doc, name, lastname string
	var userId int
	rowErr := row.Scan(&userId, &doc, &name, &lastname)
	if rowErr != nil {
		return createUserErrorResponse(Model.NOT_FOUND, err)
	}

	return *Model.NewUser(userId, doc, name, lastname), nil
}

func  (*UserRepository) GetUserPhones (userId int) ([]Model.Phone, error){
	db, err := Config.GetConnection()
	if err != nil {
		return createPhoneErrorResponse(Model.DB_UNREACHABLE, err)
	}
	phonesQuery := "SELECT phone_number FROM phone WHERE owner = (?)"

	statement, err := db.Prepare(phonesQuery)
	if err != nil {
		return createPhoneErrorResponse(Model.DB_ERROR, err)
	}

	rows, err := statement.Query(userId)
	phones := make([]Model.Phone, 0)
	for rows.Next() {
		var phone string
		err = rows.Scan(&phone)
		if err != nil {
			return createPhoneErrorResponse(Model.DB_ERROR, err)
		}
		phones = append(phones, *Model.NewPhone(phone))
	}
	return phones,nil
}

func createUserErrorResponse(displayError Model.CustomError, ex error) (Model.User, error) {
	displayError.SetOrigError(ex)
	return Model.User{}, displayError
}

func createPhoneErrorResponse(displayError Model.CustomError, ex error) ([]Model.Phone, error) {
	displayError.SetOrigError(ex)
	return make([]Model.Phone,0), displayError
}
