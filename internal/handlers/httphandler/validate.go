package httphandler

import (
	"fmt"
	"net/mail"

	"github.com/TomasCruz/users/internal/entities"
	"github.com/TomasCruz/users/utils/errlog"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func (h HTTPHandler) validateCreateUser(req entities.CreateUserReq) error {
	if err := h.checkSingleName(req.FirstName, "CreateUser.FirstName"); err != nil {
		return err
	}

	if err := h.checkSingleName(req.LastName, "CreateUser.LastName"); err != nil {
		return err
	}

	if _, err := bcrypt.Cost([]byte(req.PswdHash)); err != nil {
		return errors.Wrap(entities.ErrBadHashedPswd, err.Error())
	}

	if err := h.analyseEmail(req.Email); err != nil {
		return errors.Wrap(entities.ErrBadEmail, err.Error())
	}

	countryLength := len(req.Country)
	if countryLength != 3 && countryLength != 2 {
		err := errors.WithStack(entities.ErrCountryLength)
		errlog.Error(err, fmt.Sprintf("wrong country length %d", countryLength))
		return err
	}

	return nil
}

// func (h HTTPHandler) validateUpdateUser(req entities.UpdateUserReq) error {
// 	if req.FirstName != nil {
// 		if err := h.checkSingleName(*req.FirstName, "UpdateUser.FirstName"); err != nil {
// 			return err
// 		}
// 	}

// 	if req.LastName != nil {
// 		if err := h.checkSingleName(*req.LastName, "UpdateUser.LastName"); err != nil {
// 			return err
// 		}
// 	}

// 	if req.PswdHash != nil {
// 		if _, err := bcrypt.Cost([]byte(*req.PswdHash)); err != nil {
// 			return errors.Wrap(entities.ErrBadHashedPswd, err.Error())
// 		}
// 	}

// 	if req.Email != nil {
// 		if err := h.analyseEmail(*req.Email); err != nil {
// 			return errors.Wrap(entities.ErrBadEmail, err.Error())
// 		}
// 	}

// 	countryLength := len(req.Country)
// 	if countryLength != 3 && countryLength != 2 {
// 		err := errors.WithStack(entities.ErrCountryLength)
// 		errlog.Error(err, fmt.Sprintf("wrong country length %d", strconv.Itoa(countryLength)))
// 		return err
// 	}

// 	return nil
// }

func (h HTTPHandler) checkSingleName(name, msg string) error {
	if name == "" {
		return errors.Wrap(entities.ErrEmptyName, msg)
	}

	return nil
}

func (h HTTPHandler) analyseEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.Wrap(entities.ErrBadEmail, err.Error())
	}

	return nil
}

func (h HTTPHandler) validateUUID(uuidString string) (uuid.UUID, error) {
	if err := uuid.Validate(uuidString); err != nil {
		return uuid.UUID{}, errors.Wrap(entities.ErrBadUUID, err.Error())
	}

	// uuidString is valid
	u, err := uuid.Parse(uuidString)
	if err != nil {
		return uuid.UUID{}, errors.Wrap(entities.ErrBadUUID, err.Error())
	}

	return u, nil
}
