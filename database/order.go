package database

import (
	"errors"
	"log"

	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
	"gorm.io/gorm"
)

type OrderDb interface {
	CreateOrder(order models.Order) (uint64, error)
	GetOrderById(orderId uint64) (order models.Order, err error)
	UpdateOrderStatus(orderUpdate resources.OrderStatusUpdate) error
}

type orderDb struct {
	db *gorm.DB
}

func NewOrderDb() *orderDb {
	return &orderDb{
		db: config.Conn,
	}
}

func (d *orderDb) CreateOrder(order models.Order) (uint64, error) {
	funcdesc := "CreateOrder"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Create(&order)
	if err := result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return order.ID, err
	}

	log.Println("exit " + funcdesc)
	return order.ID, nil
}

func (d *orderDb) GetOrderById(orderId uint64) (order models.Order, err error) {
	funcdesc := "GetOrderById"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("id=?", orderId).Find(&order)
	if err = result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return order, err
	} else if order.ID == 0 {
		return order, errors.New("user id doesn't exist")
	}

	log.Println("exit " + funcdesc)
	return order, nil
}

func (d *orderDb) UpdateOrderStatus(orderUpdate resources.OrderStatusUpdate) error {
	funcdesc := "UpdateOrderStatus"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Where("id=?", orderUpdate.OrderId).Update("order_status=?", orderUpdate.OrderStatus)
	if err := result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return err
	}

	log.Println("exit " + funcdesc)
	return nil
}
