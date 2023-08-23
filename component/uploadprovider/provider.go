package uploadprovider

import (
	"context"
	"ocean-app-be/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, fileName string) (*common.Image, error)
}
