package httphandler

import (
	"github.com/TomasCruz/users/internal/entities"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// func (h HTTPHandler) validateCreateUser(req entities.CreateUserReq) error {
// 	if err := h.checkSingleName(req.FirstName, "CreateUser.FirstName"); err != nil {
// 		return err
// 	}

// 	if err := h.checkSingleName(req.LastName, "CreateUser.LastName"); err != nil {
// 		return err
// 	}

// 	if _, err := bcrypt.Cost([]byte(req.PswdHash)); err != nil {
// 		return errors.Wrap(entities.ErrBadHashedPswd, err.Error())
// 	}

// 	if err := h.analyseEmail(req.Email); err != nil {
// 		return errors.Wrap(entities.ErrBadEmail, err.Error())
// 	}

// 	return nil
// }

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

// 	return nil
// }

// func (h HTTPHandler) checkSingleName(name, msg string) error {
// 	if name == "" {
// 		return errors.Wrap(entities.ErrEmptyName, msg)
// 	}

// 	return nil
// }

// func (h HTTPHandler) analyseEmail(email string) error {
// 	_, err := mail.ParseAddress(email)
// 	return errors.Wrap(entities.ErrBadEmail, err.Error())
// }

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
