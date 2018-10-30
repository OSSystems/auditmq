package storage

type StorageData map[string]*Buffer

type Storage struct {
	services map[string]StorageData
}

var storage *Storage

func init() {
	new()
}

func new() {
	storage = &Storage{
		services: map[string]StorageData{},
	}
}

func GetStorageData(service string) StorageData {
	if _, ok := storage.services[service]; !ok {
		storage.services[service] = make(StorageData)
	}

	return storage.services[service]
}

func (s StorageData) GetBuffer(field string) *Buffer {
	if _, ok := s[field]; !ok {
		s[field] = NewBuffer()
	}

	return s[field]
}
