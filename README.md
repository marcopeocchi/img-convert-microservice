# pdf-covert-microservice

This is a simple microservice for on-the-fly PDF conversion to images.  
Everything is processed in-memory and no file is written to disk.

```sh
docker pull marcobaobao/pdf-convert-microservice
```
## Installation

```sh
docker run -d -p 8083:8083 --restart=unless-stopped --name pdf-microservice marcobaobao/pdf-convert-microservice
```

## How it works
User supplies a PDF as a blob, the services elaborates the image of the **first** page of the document.
Depening on the supplied params the output image can be either into:
- jpeg
- png
- webp

## OpenAPI 3.0

This microservice is based on the OpenAPI 3.0 specification.  
Launch `http://localhost:8083` to view the related **swagger UI**
