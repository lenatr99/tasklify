package auth

import (
	"errors"
	"log"
	"tasklify/internal/database"
	"time"

	"github.com/alexedwards/argon2id"
)

func AuthenticateUser(username, password string) (bool, error) {
	loginTime := time.Now()

	user, err := database.GetDatabase().GetUser(username)
	if err != nil {
		return false, err
	}

	match, err := argon2id.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		log.Println(err)
		return false, errors.New("no matching username and password")
	}

	var userLastLogin = &database.User{}
	userLastLogin.ID = user.ID
	userLastLogin.LastLogin = &loginTime
	err = database.GetDatabase().UpdateUser(userLastLogin)
	if err != nil {
		return false, err
	}

	return match, nil
}

func CreateUser(username, password, firstName, lastName, email string, systemRoleName string) error {
	passwordHash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return err
	}

	systemRole := database.SystemRoles.Parse(systemRoleName)
	if systemRole == (&database.SystemRole{}) {
		return errors.New("system role not found")
	}

	var user = &database.User{
		Username:   username,
		Password:   passwordHash,
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		SystemRole: *systemRole,
	}

	return database.GetDatabase().CreateUser(user)
}

func UpdateUser(issuerUsername, issuerPassword string, id uint, username, password, firstName, lastName, email, systemRole *string) error {
	ok, err := AuthenticateUser(issuerUsername, issuerPassword)
	if err != nil {
		return err
	}

	issuerUser, err := database.GetDatabase().GetUser(issuerUsername)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("you are not authenticated")
	}

	var user = &database.User{}
	user.ID = id

	if username != nil {
		user.Username = *username
	}

	if password != nil {
		passwordHash, err := argon2id.CreateHash(*password, argon2id.DefaultParams)
		if err != nil {
			return err
		}

		user.Username = passwordHash
	}

	if firstName != nil {
		user.FirstName = *firstName
	}

	if lastName != nil {
		user.LastName = *lastName
	}

	if email != nil {
		user.Email = *email
	}

	if systemRole != nil {
		err = GetAuthorization().HasPermission(database.SystemRoles.WrappedValue(issuerUser.SystemRole), "/system/user/system-role", "u")
		if err != nil {
			return err
		}

		systemRoleObj := database.SystemRoles.Parse(*systemRole)
		if systemRoleObj == (&database.SystemRole{}) {
			return errors.New("system role not found")
		}

		user.SystemRole = *systemRoleObj
	}

	return database.GetDatabase().UpdateUser(user)
}