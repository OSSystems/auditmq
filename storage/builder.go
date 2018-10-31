package storage

import (
	"github.com/rodrigoapereira/auditmq/config"
)

type StorageBuilder struct {
	storage    Storage
	baseConfig config.DataFields
}

func (b *StorageBuilder) buildServices() {
	for _, data := range b.baseConfig {
		b.storage.AddService(NewService(data.Owner))

		for service := range data.Replicas {
			b.storage.AddService(NewService(service))
		}
	}
}

func (b *StorageBuilder) buildFields() error {
	for fieldName, fieldOpts := range b.baseConfig {
		ownerService, err := b.storage.GetService(fieldOpts.Owner)
		if err != nil {
			return err
		}

		err = ownerService.AddField(fieldName, fieldOpts, nil)
		if err != nil {
			return err
		}

		for srvName, dataOptions := range fieldOpts.Replicas {
			srv, err := b.storage.GetService(srvName)
			if err != nil {
				return err
			}

			err = srv.AddFieldWithOptions(fieldName, fieldOpts, ownerService, dataOptions)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *StorageBuilder) Build() error {
	b.buildServices()
	err := b.buildFields()
	if err != nil {
		return err
	}

	return nil
}

func InitializeStorage(cfg config.DataFields) Storage {
	storageInstance = newStorage()
	builder := &StorageBuilder{
		storage:    storageInstance,
		baseConfig: cfg,
	}

	builder.Build()
	return storageInstance
}
