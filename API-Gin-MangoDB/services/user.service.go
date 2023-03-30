package services

type UserService interface{
	CreateUser(*models.User) error {
		GetUser(*string) (*modrls.User, error)
		GetAll() ([]*models.User,error)
		UpdateUser(*models.User) error
		DeleteUser(*string) error  
	}
}

func (u *UserService) CreateUser(user *models.User) error {

	_, err := u.sercollerrtion.InsertOne(u.ctx,user)
	if err!= nil {
        return err
    }
}

func (u *UserService) GetUser() (name *string) ([]*models.User,error) {
	var users []*models.User
    query := bson.D{bson.E{Key:"name", Value :name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)

return user, err

}

func (u *UserService) GetAll() ([]*models.User,error) {
	return nil }

func (u *UserService) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserService) DeleteUser(name *string) error {
    return nil
}