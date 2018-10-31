package storage

import (
	"errors"
)

type Services map[string]Service

type Storage interface {
	AddService(srv Service) Service
	GetService(srv string) (Service, error)
	Services() Services
}

type storage struct {
	services Services
}

var storageInstance Storage

func GetStorage() (Storage, error) {
	if storageInstance == nil {
		return nil, errors.New("Storage not initialized")
	}

	return storageInstance, nil
}

func (s *storage) AddService(srv Service) Service {
	if val, ok := s.services[srv.Name()]; ok {
		return val
	}

	s.services[srv.Name()] = srv
	return srv
}

func (s *storage) GetService(name string) (Service, error) {
	val, ok := s.services[name]
	if !ok {
		return nil, errors.New("Service not found")
	}

	return val, nil
}

func (s *storage) Services() Services {
	return s.services
}

func newStorage() Storage {
	return &storage{
		services: make(Services),
	}
}
