package httphandler

import (
	"github.com/TomasCruz/users/internal/entities"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

// func (p Presenter) validateCreateUser(req entities.CreateUserReq) error {
// 	if err := p.checkSingleName(req.FirstName, "CreateUser.FirstName"); err != nil {
// 		return err
// 	}

// 	if err := p.checkSingleName(req.LastName, "CreateUser.LastName"); err != nil {
// 		return err
// 	}

// 	if err := p.checkSingleName(req.NickName, "CreateUser.NickName"); err != nil {
// 		return err
// 	}

// 	if _, err := bcrypt.Cost([]byte(req.PswdHash)); err != nil {
// 		log.Error(err)
// 		return errors.Wrap(entities.ErrBadHashedPswd, err.Error())
// 	}

// 	if err := p.analyseEmail(req.Email); err != nil {
// 		log.Error(err)
// 		return errors.Wrap(entities.ErrBadEmail, err.Error())
// 	}

// 	countryLength := len(req.Country)
// 	if countryLength != 3 && countryLength != 2 {
// 		log.Error(entities.ErrCountryLength)
// 		return errors.Wrap(entities.ErrCountryLength, strconv.Itoa(countryLength))
// 	}

// 	return nil
// }

// func (p Presenter) validateUpdateUser(req entities.UpdateUserReq) error {
// 	if req.FirstName != nil {
// 		if err := p.checkSingleName(*req.FirstName, "UpdateUser.FirstName"); err != nil {
// 			return err
// 		}
// 	}

// 	if req.LastName != nil {
// 		if err := p.checkSingleName(*req.LastName, "UpdateUser.LastName"); err != nil {
// 			return err
// 		}
// 	}

// 	if req.NickName != nil {
// 		if err := p.checkSingleName(*req.NickName, "UpdateUser.NickName"); err != nil {
// 			return err
// 		}
// 	}

// 	if req.PswdHash != nil {
// 		if _, err := bcrypt.Cost([]byte(*req.PswdHash)); err != nil {
// 			log.Error(err)
// 			return errors.Wrap(entities.ErrBadHashedPswd, err.Error())
// 		}
// 	}

// 	if req.Email != nil {
// 		if err := p.analyseEmail(*req.Email); err != nil {
// 			log.Error(err)
// 			return errors.Wrap(entities.ErrBadEmail, err.Error())
// 		}
// 	}

// 	if req.Country != nil {
// 		countryLength := len(*req.Country)
// 		if countryLength != 3 && countryLength != 2 {
// 			log.Error(entities.ErrCountryLength)
// 			return errors.Wrap(entities.ErrCountryLength, strconv.Itoa(countryLength))
// 		}
// 	}

// 	return nil
// }

// func (p Presenter) checkSingleName(name, msg string) error {
// 	if name == "" {
// 		err := errors.Wrap(entities.ErrEmptyName, msg)
// 		log.Error(err)
// 		return err
// 	}

// 	return nil
// }

// func (p Presenter) analyseEmail(email string) error {
// 	_, err := mail.ParseAddress(email)
// 	log.Error(err)
// 	return err
// }

func (h HTTPHandler) validateUUID(uuidString string) (uuid.UUID, error) {
	if err := uuid.Validate(uuidString); err != nil {
		log.Error(err)
		return uuid.UUID{}, errors.Wrap(entities.ErrBadUUID, err.Error())
	}

	// uuidString is valid
	u, err := uuid.Parse(uuidString)
	if err != nil {
		log.Error(err)
		return uuid.UUID{}, errors.Wrap(entities.ErrBadUUID, err.Error())
	}

	return u, nil
}
