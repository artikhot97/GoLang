
package main

import (
    "fmt"
    "net/http"
    "bytes"
      "github.com/aws/aws-sdk-go/aws"
      "github.com/aws/aws-sdk-go/aws/session"
      "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/private/protocol/rest"
    "encoding/json"
)

type UrlResponse struct {
    Url string `json:Url`
}

const (
    AWS_S3_REGION = "your-s3-regoin-name"
    AWS_S3_BUCKET = "your_bucket_name"
)

var s3session *s3.S3

func UploadFile(rw http.ResponseWriter , req* http.Request){
    fmt.Println("In received file func ....")

    s3session := s3.New(session.Must(session.NewSession(&aws.Config{
      Region: aws.String(AWS_S3_REGION),
    })))

    file, parseMultipart,err := req.FormFile("input_file")
    check(err)
    defer file.Close()
    size := parseMultipart.Size
    buffer := make([]byte, size)
    file.Read(buffer)
    

    fmt.Println(parseMultipart.Filename)
    result, _ := s3session.PutObject(&s3.PutObjectInput{
            Body: bytes.NewReader(buffer),
            Bucket: aws.String(AWS_S3_BUCKET),
            Key: aws.String(parseMultipart.Filename),
//             ACL: aws.String("public-read"), // could be private if you want it to be access by only authorized users
          })

      fmt.Println(result)
    res, _ := s3session.GetObjectRequest(&s3.GetObjectInput{
        Bucket: aws.String(AWS_S3_BUCKET),
        Key:    aws.String(parseMultipart.Filename),
    })
        rest.Build(res)
        uploadedResourceLocation := res.HTTPRequest.URL.String()
        fmt.Println(uploadedResourceLocation)
    // Format response url in json
    metaData := UrlResponse{Url:uploadedResourceLocation}
    jsonString, err := json.MarshalIndent(metaData, "", " ")
    fmt.Println(jsonString)
    rw.Write([]byte(string(jsonString)))
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
    fmt.Println("Running...")
    http.HandleFunc("/uploadFile" , UploadFile)
    http.ListenAndServe(":7000" , nil)
}
