package controllers

import (
  "database/sql"

  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "github.com/ngurajeka/ngurajeka.com/app"
  "github.com/revel/revel"
)

type GormController struct {
  *revel.Controller
  Txn *gorm.DB
}

func (c *GormController) Begin() revel.Result {
  txn := app.DB.Begin()
  if txn.Error != nil {
    panic(txn.Error)
  }

  c.Txn = txn
  return nil
}

func (c *GormController) Commit() revel.Result {
  if c.Txn == nil {
    return nil
  }

  c.Txn.Commit()
  if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }

  c.Txn = nil
  return nil
}

func (c *GormController) Rollback() revel.Result {
  if c.Txn == nil {
    return nil
  }

  c.Txn.Rollback()
  if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }

  c.Txn = nil
  return nil
}
