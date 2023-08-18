package uploadprovider

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
	"ocean-app-be/common"
)

type cdProvider struct {
	cloudName string
	apiKey    string
	secret    string
	uploader  *cloudinary.Cloudinary
}

func NewCDProvider(cloudName, apiKey, secret string) *cdProvider {
	provider := &cdProvider{
		cloudName: cloudName,
		apiKey:    apiKey,
		secret:    secret,
	}

	cld, err := cloudinary.NewFromParams(provider.cloudName, provider.apiKey, provider.secret)

	if err != nil {
		log.Fatalln(err)
	}
	provider.uploader = cld

	return provider
}

func (provider *cdProvider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	resp, err := provider.uploader.Upload.Upload(ctx, data, uploader.UploadParams{
		PublicID: dst,
	})
	if err != nil {
		return nil, err
	}

	imageURL := resp.URL

	image := &common.Image{
		Url: imageURL,
	}

	return image, nil
}
