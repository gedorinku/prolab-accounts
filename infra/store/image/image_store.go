package imagestore

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"image"
	_ "image/gif" // for image
	_ "image/jpeg"
	_ "image/png"

	"github.com/minio/minio-go"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/infra/store"
)

type imageStoreImpl struct {
	ctx        context.Context
	cli        *minio.Client
	bucketName string
}

// NewImageStore returns new image store
func NewImageStore(ctx context.Context, cli *minio.Client, bucket string) store.ImageStore {
	return &imageStoreImpl{
		ctx:        ctx,
		cli:        cli,
		bucketName: bucket,
	}
}

func (s *imageStoreImpl) CreateImage(img []byte) (filename string, err error) {
	r := bytes.NewReader(img)
	_, ext, err := image.DecodeConfig(r)
	if err != nil {
		return "", errors.WithStack(err)
	}
	name, err := generateFilename(ext)
	if err != nil {
		return "", errors.WithStack(err)
	}

	opt := minio.PutObjectOptions{
		ContentType: "image/" + ext,
	}
	_, err = s.cli.PutObjectWithContext(s.ctx, s.bucketName, name, r, r.Size(), opt)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return name, nil
}

func (s *imageStoreImpl) DeleteImage(filename string) error {
	err := s.cli.RemoveObject(s.bucketName, filename)
	return errors.WithStack(err)
}

func generateFilename(ext string) (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.WithStack(err)
	}

	res := base64.RawURLEncoding.EncodeToString(b)

	return string(res) + "." + ext, nil
}
