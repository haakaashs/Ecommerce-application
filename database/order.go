package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/haakaashs/antino-labs/config"
	"github.com/haakaashs/antino-labs/constants"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
	"gorm.io/gorm"
)

type OrderDb interface {
	CreateOrder(models.Order) (uint64, error)
	GetOrderById(uint64) (models.Order, error)
	UpdateOrderStatus(resources.OrderStatusUpdate) error
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

	if orderUpdate.OrderStatus == constants.CANCELLED {
		order, err := d.GetOrderById(orderUpdate.OrderId)
		if err != nil {
			log.Println("error in DB query: ", err.Error())
			return err
		} else if order.OrderStatus == constants.COMPLETED {
			log.Println("error in DB query: order already completed")
			return errors.New("order already completed")
		}
	}

	switch orderUpdate.OrderStatus {

	case constants.DISPATCHED:
		{
			result := d.db.Debug().Where("id=?", orderUpdate.OrderId).Updates(models.Order{OrderStatus: orderUpdate.OrderStatus, IsActive: &[]bool{true}[0]})
			if err := result.Error; err != nil {
				log.Println("error in DB query: ", err.Error())
				return err
			}
		}
	default:
		{
			fmt.Println("jill: ", orderUpdate.OrderStatus)
			result := d.db.Debug().Where("id=?", orderUpdate.OrderId).Updates(models.Order{OrderStatus: orderUpdate.OrderStatus, IsActive: &[]bool{false}[0]})
			if err := result.Error; err != nil {
				log.Println("error in DB query: ", err.Error())
				return err
			}
		}
	}

	log.Println("exit " + funcdesc)
	return nil
}
