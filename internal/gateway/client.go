package gateway

import "github.com/gsilvasouza/ms-waller/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
