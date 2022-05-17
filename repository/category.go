package repository

import (
	"belajar_unit_test/entity"
)

type Categoryrepository interface {
	FindById(id string) *entity.Category
}