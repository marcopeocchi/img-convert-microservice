# img-convert-microservice

This is a simple microservice for on-the-fly image or PDF conversion.  
Everything is processed in-memory and no file is written to disk.

```sh
docker pull marcobaobao/img-convert-microservice
```
## Installation

```sh
docker run -d -p 8080:8080 --restart=unless-stopped --name img-conv-microservice marcobaobao/img-convert-microservice
```

## How it works
User supplies an image or PDF as a blob (the services elaborates the image of the **first** page of the document).
Depening on the supplied params the output image can be either into:
- jpeg
- png
- webp
- avif
- gif

## OpenAPI 3.0

This microservice is based on the OpenAPI 3.0 specification.  
Launch `http://localhost:8080` to view the related **swagger UI**


## Prometheus metrics
Anonymous metrics are collected.
Go memory stats, goroutines, `time_per_conversion_ns` and `processed_counter` collected and avaible from `/metrics` endpoint.
