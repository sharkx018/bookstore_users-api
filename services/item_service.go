package services

var (
	ItemService itemServiceInterface = &itemService{}
)

type itemService struct{}

type itemServiceInterface interface {
	GetItem()
	SaveItem()
}

func (s *itemService) GetItem() {

}

func (s *itemService) SaveItem() {

}
