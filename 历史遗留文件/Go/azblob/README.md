# azblob

## 简介
**安装** : go get -u github.com/Azure/azure-storage-blob-go/azblob

**github地址**: https://github.com/Azure/azure-storage-blob-go

**微软文档**: https://docs.microsoft.com/zh-cn/azure/storage/blobs/storage-quickstart-blobs-go?tabs=windows
> 内容比较旧,一直没更新

## version
_0.12.0 -> 0.13.0_
- Validate echoed client request ID from the service
- Added new TransferManager option for UploadStreamToBlockBlob to fine-tune the concurrency and memory usage

## 第一部分 概念
TODO

## 第二部分 进阶
### 访问已存在的ContainerURL
> 采用SASUrl访问
```golang
credential := azblob.NewAnonymousCredential()

p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

URL, _ := url.Parse(
    "url?SAS")
//连接blob**
containerURL := azblob.NewContainerURL(*URL, p)

ctx := context.Background() // This example uses a never-expiring context
```
> 注意容器名称不等同于Azure Storage上目录的名称，可以通过右键点击Block文件的属性查看容器名称

### 将blob上传到容器
```golang
// Create a file to test the upload and download.
fmt.Printf("Creating a dummy file to test the upload and download\n")
data := []byte("hello world this is a blob\n")
fileName := randomString()
err = ioutil.WriteFile(fileName, data, 0700)
handleErrors(err)

// Here's how to upload a blob.
blobURL := containerURL.NewBlockBlobURL(fileName)
file, err := os.Open(fileName)
handleErrors(err)

// You can use the low-level Upload (PutBlob) API to upload files. Low-level APIs are simple wrappers for the Azure Storage REST APIs.
// Note that Upload can upload up to 256MB data in one shot. Details: https://docs.microsoft.com/rest/api/storageservices/put-blob
// To upload more than 256MB, use StageBlock (PutBlock) and CommitBlockList (PutBlockList) functions. 
// Following is commented out intentionally because we will instead use UploadFileToBlockBlob API to upload the blob
// _, err = blobURL.Upload(ctx, file, azblob.BlobHTTPHeaders{ContentType: "text/plain"}, azblob.Metadata{}, azblob.BlobAccessConditions{})
// handleErrors(err)

// The high-level API UploadFileToBlockBlob function uploads blocks in parallel for optimal performance, and can handle large files as well.
// This function calls StageBlock/CommitBlockList for files larger 256 MBs, and calls Upload for any file smaller
fmt.Printf("Uploading the file with blob name: %s\n", fileName)
_, err = azblob.UploadFileToBlockBlob(ctx, file, blobURL, azblob.UploadToBlockBlobOptions{
    BlockSize:   4 * 1024 * 1024,
    Parallelism: 16})
handleErrors(err)
```

### 列出容器中的 Blob并下载
```golang
// List the container that we have created above
fmt.Println("Listing the blobs in the container:")
for marker := (azblob.Marker{}); marker.NotDone(); {
    // Get a result segment starting with the blob indicated by the current Marker.
    listBlob, err := containerURL.ListBlobsFlatSegment(ctx, marker, azblob.ListBlobsSegmentOptions{})
    handleErrors(err)

    // ListBlobs returns the start of the next segment; you MUST use this to get
    // the next segment (after processing the current result segment).
    marker = listBlob.NextMarker

    // Process the blobs returned in this result segment (if the segment is empty, the loop body won't execute)
    for _, blobInfo := range listBlob.Segment.BlobItems {
            fmt.Println(" Blob name: " + blobInfo.Name + "")
            _, fileName := filepath.Split(blobInfo.Name)

            blobURL := containerURL.NewBlockBlobURL(blobInfo.Name)
            downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
            if err != nil {
                fmt.Println(err)
            }

            bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

            downloadedData := bytes.Buffer{}
            _, err = downloadedData.ReadFrom(bodyStream)
            if err != nil {
                fmt.Println(err)
            }

            ioutil.WriteFile(fileName, downloadedData.Bytes(), 0666)
    }
}
```

### 下载Blob
```golang
// Here's how to download the blob
downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
if err != nil {
    logger.Fatal(err)
}
// NOTE: automatically retries are performed if the connection fails
bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

// read the body into a buffer
downloadedData := bytes.Buffer{}
_, err = downloadedData.ReadFrom(bodyStream)
if err != nil {
    logger.Fatal(err)
}
```