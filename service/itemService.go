package service

import "gorm.io/gorm"

type ItemService struct {
	db    *gorm.DB
	table string
}

func NewItemService() *ItemService {
	return &ItemService{
		//db:    data.GetDatabase(),
		//table: "users",
	}
}

/*
func (service *UserService) FindOneByUid(uid string) *model.User {
	var user model.User

	res := service.db.First(&user, "uid = ?", uid)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}
*/
