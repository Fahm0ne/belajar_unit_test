package service

import (
	"belajar_unit_test/entity"
	"belajar_unit_test/repository"
	"errors"
)

type Categoryservice struct {
	Repository repository.Categoryrepository
}

func (service Categoryservice) Get (id string) (*entity.Category , error) {
category := service.Repository.FindById(id)
if category == nil {
return category, errors.New("category not found")
} else {
return category , nil
}
}

