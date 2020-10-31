<div align="center">
  <h1>Postback Delivery</h1>
  <blockquote>Service at small scale to distribute data. It accepts incoming HTTP requests (PHP service), 
it communicates through a Job Queue (Redis) with another service (Golang service), 
and this last service is in charge of delivering it to a specific endpoint. The response of these last request is then stored in a log file.</blockquote>
</div>

## 🚀 Usage and examples

#### Usage

```
docker-compose up -d    // runs services on background
docker-compose down     // stops services and remove them
```

#### Examples

- POST http://192.168.99.100:8000
- BODY:

```
{
    "endpoint":{
        "method":"GET",
        "url":"http://sample_domain_endpoint.com/data?title={mascot}&image={location}&foo={bar}"
    },
    "data":[
        {
            "mascot":"Gopher",
            "location":"https://blog.golang.org/gopher/gopher.png"
        }
    ]
}
```


## ⚙️ Tech Stack

- Golang
- PHP
- Redis
- Docker

