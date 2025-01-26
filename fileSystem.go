package aspera

import (
	"context"
	"net/http"
)

type FileSystemService struct {
	Client *Client
}

func NewFileSystemService(client *Client) *FileSystemService {
	return &FileSystemService{
		Client: client,
	}
}

func (f *FileSystemService) StartTransfer(ctx context.Context, transferSpec *TransferSpec) error {
	endpoint, ok := endpoints["startTransfer"]
	if !ok {
		return ErrEndpointNotFound
	}

	req, err := f.Client.NewRequest(http.MethodPost, endpoint.URL(), transferSpec)
	if err != nil {
		return err
	}

	return f.Client.Do(ctx, req, nil)
}
