# Software Architecture for Big Data - Exercise 6

This exercise we are going to add an object storage to our setup. 
We are going to use [Minio](https://hub.docker.com/r/minio/minio) as an S3 compatible local object storage provider.
To simplify things, we are going to call this object store S3.

For every posted drink order we want to create a "receipt" markdown file, in addition to storing the order in the database.
The receipt file should conform to the following format:

```markdown
# Order: 11

| Created At      | Drink ID | Amount |
|-----------------|----------|--------|
| Nov 12 17:12:39 | 1        | 3      |

Thanks for drinking with us!
```

## Todo
- [ ] Add a Minio instance to the existing `docker-compose.yml`
  - Expose port 8500
  - Persist storage using a volume
  - Set the correct environment variables
- [ ] Create a markdown compatible receipt string in `model/order.go` 
  - Make use of [fmt.Sprintf](https://gobyexample.com/string-formatting) to fill the string with data
- [ ] Add Put and Get receipts functionality to the `rest/api.go` file
  - Add `PutObject` to `PostOrder` route
  - Add `GetObject` to `GetReceiptFile` route and serve the requested receipt markdown file
- [ ] Verify that files are stored properly using an S3 compatible viewer

## Tipps and Tricks
Use the `debug.env` for development, **NOT** in your `docker-compose.yml`.

To rebuild the Orderservice Dockerfile when using Docker Compose, use the following command:
```bash
docker compose build orderservice
docker compose up -d --force-recreate orderservice 
```

Beware that the dashboard will no be served when developing locally!
Only the OpenAPI definition is reachable at http://localhost:3000/openapi/index.html.

### Object Storage (Minio / S3)
To view files in your local S3 / object storage in your IDE, install the 
[Remote File Systems](https://plugins.jetbrains.com/plugin/21706-remote-file-systems) plugin.
Use the username and password supplied in `debug.env` or your `docker-compose.yml`.

![s3_plugin_goland.png](solution/res/s3_plugin_goland.png)

A [similar plugin](https://marketplace.visualstudio.com/items?itemName=seriousbenentertainment.minio) also exists for Visual Studio Code.
