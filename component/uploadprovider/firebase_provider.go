package uploadprovider

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"

	"google.golang.org/api/option"
	"log"
	"net/http"
	"ocean-app-be/common"
)

type firebaseProvider struct {
	bucketName    string
	domain        string
	storageClient *storage.Client
}

func NewFirebaseProvider(bucketName, APIKey string) *firebaseProvider {
	provider := &firebaseProvider{
		bucketName: bucketName,
	}
	opt := option.WithCredentialsJSON([]byte(APIKey))

	client, err := storage.NewClient(context.Background(), opt)
	if err != nil {
		log.Fatalln(err)
	}

	provider.storageClient = client

	return provider
}

func (provider *firebaseProvider) SaveFileUploaded(ctx context.Context, data []byte, fileName, url string) (*common.Image, error) {

	ctx = context.Background()
	//fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	storageRef := provider.storageClient.Bucket(provider.bucketName).Object(fileName)

	writer := storageRef.NewWriter(ctx)

	writer.ContentType = fileType

	if _, err := writer.Write(data); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	img := &common.Image{
		Url: fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", provider.bucketName, url),
	}

	return img, nil
}
