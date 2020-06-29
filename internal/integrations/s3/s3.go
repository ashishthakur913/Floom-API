package s3

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 interface {
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

type s3 struct {
	session *session.Session
}

func Initialize() (S3, error) {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsKey, awsSecret, ""),
	})

	if err != nil {
		return nil, err
	}
	s3Uploader := &s3{
		session: session,
	}
	return s3Uploader, nil
}

/*
	Uploads the image to S3 and returns the imageName
*/
func (s *s3) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	uploader := s3manager.NewUploader(s.session)

	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	extension := filepath.Ext(fileHeader.Filename)

	fileNameToBeStored := s.getEncodedFileName(fileHeader.Filename)
	if len(extension) != 0 {
		fileNameToBeStored = fileNameToBeStored + extension
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(fmt.Sprintf("%s/%s", folder, fileNameToBeStored)),
		Body:        bytes.NewReader(buffer),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(http.DetectContentType(buffer)),
	})
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
		return "", err
	}
	fmt.Printf("file uploaded to, %v\n", result.Location)

	return fileNameToBeStored, nil
}

func (s *s3) getEncodedFileName(fileName string) string {
	fileName = fmt.Sprintf("%s_%d_%d_%f", fileName, time.Now().UnixNano(), rand.Float64())
	hash := md5.Sum([]byte(fileName))
	return hex.EncodeToString(hash[:])
}

func GetImageURL(imageName string) string {
	return host + "/" + folder + "/" + imageName
}

func GetFileName(imageName string) string {
	fileName := strings.Split(imageName, host+"/"+folder+"/")
	if len(fileName) > 1 {
		return fileName[1]
	}
	return imageName
}
