package repository

import "sync"

type InMemmoryDb struct{
	lock sync.RWMutex
	Data map[string]interface{}
}


func (idb *InMemmoryDb) SaveTransaction() error{

	return nil
}