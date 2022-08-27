package service

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/johnHPX/blog-hard-backend/internal/domain/models"
	"github.com/johnHPX/blog-hard-backend/internal/infra/repository"
	"github.com/johnHPX/blog-hard-backend/internal/infra/utils/messages"
	"github.com/johnHPX/validator-hard/pkg/validator"
)

type userServiceInterface interface {
	Store(name, telephone, nick, email, secret string) error
	StoreADM(name, telephone, nick, email, secret, kind string) error
	List(offset, limit, page int) ([]models.User, error)
	Count() (int, error)
	ListName(name string, offset, limit, page int) ([]models.User, error)
	CountName(name string) (int, error)
	Find(id string) (*models.User, error)
	Update(id, name, telefone, nick, email, kind string) error
	Remove(id string) error
	Login(emailOrNick, secret string) (string, error)
}

type userServiceImpl struct {
	UserID string
	Kind   string
}

func (s *userServiceImpl) Store(name, telephone, nick, email, secret string) error {

	val := validator.NewValidator()
	Name, err := val.CheckAnyData("nome", 255, name, true)
	if err != nil {
		return err
	}
	Telephone, err := val.CheckAnyData("telefone", 13, telephone, true)
	if err != nil {
		return err
	}
	Nick, err := val.CheckAnyData("nick", 255, nick, true)
	if err != nil {
		return err
	}
	Email, err := val.CheckAnyData("email", 255, email, true)
	if err != nil {
		return err
	}
	Password, err := val.CheckPassword(255, secret, "", "create")
	if err != nil {
		return err
	}

	repUser := repository.NewUserRepository()
	repPerson := repository.NewPersonRepository()

	err = repUser.CheckEmail(Email.(string))
	if err != nil {
		return err
	}

	err = repUser.CheckNick(Nick.(string))
	if err != nil {
		return err
	}

	uid := uuid.New()
	pid := uuid.New()

	e := &models.User{
		UserID: uid.String(),
		Person: models.Person{
			PersonID:  pid.String(),
			Name:      Name.(string),
			Telephone: Telephone.(string),
		},
		Nick:   Nick.(string),
		Email:  Email.(string),
		Secret: Password,
		Kind:   "user",
	}

	err = repUser.Store(e)
	if err != nil {
		return err
	}

	err = repPerson.Store(&e.Person, e.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) StoreADM(name, telephone, nick, email, secret, kind string) error {

	if s.Kind != "adm" {
		return errors.New(messages.AdmMessage)
	}

	val := validator.NewValidator()
	Name, err := val.CheckAnyData("nome", 255, name, true)
	if err != nil {
		return err
	}
	Telephone, err := val.CheckAnyData("telefone", 13, telephone, true)
	if err != nil {
		return err
	}
	Nick, err := val.CheckAnyData("nick", 255, nick, true)
	if err != nil {
		return err
	}
	Email, err := val.CheckAnyData("email", 255, email, true)
	if err != nil {
		return err
	}
	Password, err := val.CheckPassword(255, secret, "", "create")
	if err != nil {
		return err
	}
	KindVal, err := val.CheckAnyData("kind", 10, kind, true)
	if err != nil {
		return err
	}

	repUser := repository.NewUserRepository()
	repPerson := repository.NewPersonRepository()

	err = repUser.CheckEmail(Email.(string))
	if err != nil {
		return err
	}

	err = repUser.CheckNick(Nick.(string))
	if err != nil {
		return err
	}

	uid := uuid.New()
	pid := uuid.New()

	e := &models.User{
		UserID: uid.String(),
		Person: models.Person{
			PersonID:  pid.String(),
			Name:      Name.(string),
			Telephone: Telephone.(string),
		},
		Nick:   Nick.(string),
		Email:  Email.(string),
		Secret: Password,
		Kind:   KindVal.(string),
	}

	err = repUser.Store(e)
	if err != nil {
		return err
	}

	err = repPerson.Store(&e.Person, e.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) List(offset, limit, page int) ([]models.User, error) {

	if s.Kind != "adm" {
		return nil, errors.New(messages.AdmMessage)
	}

	repUser := repository.NewUserRepository()
	entities, err := repUser.List(offset, limit, page)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (s *userServiceImpl) Count() (int, error) {

	if s.Kind != "adm" {
		return 0, errors.New(messages.AdmMessage)
	}

	repUser := repository.NewUserRepository()
	count, err := repUser.Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *userServiceImpl) ListName(name string, offset, limit, page int) ([]models.User, error) {

	if s.Kind != "adm" {
		return nil, errors.New(messages.AdmMessage)
	}

	repUser := repository.NewUserRepository()
	entities, err := repUser.ListName(name, offset, limit, page)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (s *userServiceImpl) CountName(name string) (int, error) {

	if s.Kind != "adm" {
		return 0, errors.New(messages.AdmMessage)
	}

	repUser := repository.NewUserRepository()
	count, err := repUser.CountListName(name)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *userServiceImpl) Find(id string) (*models.User, error) {

	if s.UserID != id && s.Kind != "adm" {
		return nil, errors.New(messages.AnotherUser)
	}

	repUser := repository.NewUserRepository()
	user, err := repUser.Find(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) Update(id, name, telefone, nick, email, kind string) error {

	if s.UserID != id && s.Kind != "adm" {
		return errors.New(messages.AnotherUser)
	}

	val := validator.NewValidator()
	NameVal, err := val.CheckAnyData("nome", 255, name, true)
	if err != nil {
		return err
	}
	TelefoneVal, err := val.CheckAnyData("telefone", 13, telefone, true)
	if err != nil {
		return err
	}
	NickVal, err := val.CheckAnyData("nick", 255, nick, true)
	if err != nil {
		return err
	}
	EmailVal, err := val.CheckAnyData("email", 255, email, true)
	if err != nil {
		return err
	}
	KindVal, err := val.CheckAnyData("kind", 10, kind, true)
	if err != nil {
		return err
	}

	if s.Kind != "adm" {
		KindVal = "user"
	}

	// repositorys
	repUser := repository.NewUserRepository()
	repPerson := repository.NewPersonRepository()

	// find id person
	user, err := repUser.Find(id)
	if err != nil {
		return err
	}

	// Update person
	person := new(models.Person)
	person.PersonID = user.PersonID
	person.Name = NameVal.(string)
	person.Telephone = TelefoneVal.(string)
	err = repPerson.Update(person)
	if err != nil {
		return err
	}

	// Update User
	userEntity := new(models.User)
	userEntity.UserID = user.UserID
	userEntity.Nick = NickVal.(string)
	userEntity.Email = EmailVal.(string)
	userEntity.Kind = KindVal.(string)
	err = repUser.Update(userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) Remove(id string) error {

	if s.UserID != id && s.Kind != "adm" {
		return errors.New(messages.AnotherUser)
	}

	// repositorys
	repUser := repository.NewUserRepository()
	repPerson := repository.NewPersonRepository()

	// find id person
	user, err := repUser.Find(id)
	if err != nil {
		return err
	}

	// remove person
	err = repPerson.Remove(user.PersonID)
	if err != nil {
		return err
	}

	// remove user
	err = repUser.Remove(user.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) Login(emailOrNick, secret string) (string, error) {
	val := validator.NewValidator()
	EmailOrNickVal, err := val.CheckAnyData("email ou nick", 255, emailOrNick, true)
	if err != nil {
		return "", err
	}

	// repository
	repUser := repository.NewUserRepository()

	// finding user by email or nick
	user, err := repUser.FindByEmailOrNick(EmailOrNickVal.(string))
	if err != nil {
		return "", err
	}

	// checking if the password is correct
	_, err = val.CheckPassword(255, secret, user.Secret, "compare")
	if err != nil {
		return "", err
	}

	// create a token for this user
	svcAccess := NewAccessService()
	atoken, err := svcAccess.CreateAToken(user.UserID, user.Kind)
	if err != nil {
		return "", err
	}

	// creta a rtoken for this user
	rtoken, err := svcAccess.CreateRToken()
	if err != nil {
		return "", err
	}

	// access repository
	repAccess := repository.NewAccessRepository()
	// verific if exists a rtoken that user
	accessEntity, err := repAccess.FindToken(user.UserID)
	if err == nil {

		// verific if was blocked
		if accessEntity.IsBlocked {
			return "", errors.New(messages.UserBlocked)
		}

		// remove a old rtoken
		err := repAccess.RemoveToken(user.UserID)
		if err != nil {
			return "", err
		}
	}

	// salve rtoken in database
	err = repAccess.Store(rtoken, user.UserID, (time.Now().Add(time.Hour * 24 * 15)))
	if err != nil {
		return "", err
	}
	// return atoken
	return atoken, nil
}

func NewUserService(userID, kind string) userServiceInterface {
	return &userServiceImpl{
		UserID: userID,
		Kind:   kind,
	}
}