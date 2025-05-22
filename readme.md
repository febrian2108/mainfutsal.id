# Sistem Booking mainfutsal.id

## Goal
Sistem ini dirancang untuk memberikan kemudahan kepada para pecinta futsal dalam membooking lapangan futsal secara online dengan mudah. Dengan menggunakan sistem ini, semua proses mulai dari pemesanan hingga pembayaran dapat dilakukan melalui sistem ini.

## Scope
Scope untuk sistem ini yaitu dimulai dari mengelola lapangan dan jadwalnya kemudian pemesanan sampai pembayaran.

## Tech Stack

| Category            | Tech                          |
|---------------------|-------------------------------|
| Database Engine     | PostgreSQL                    |
| Programming Language| Golang                       |
| Framework           | Gin & Gorm                   |
| API Protocol        | Rest API                     |
| Message Broker      | Apache Kafka                 |
| Containerization    | Docker                      |
| CI/CD               | Jenkins                     |
| Cloud Computing     | Google Cloud Provider (GCP) |
| Payment Gateway     | Midtrans                    |
| Storage             | Google Cloud Storage (GCS)  |
| Store Data Key Value| Consul KV                   |
| Frontend            | Next.Js                     |

## Flow

1. Admin → Login  
2. Admin → Manage Field  
3. Admin → Manage Field Schedule  
4. Customer → Register  
5. Customer → Login  
6. Customer → Create Order  
7. Customer → Pay Order  
8. Done  

## API Documentation

[Link API Documentation di Postman](https://winter-capsule-185405.postman.co/workspace/77705732-d788-4091-8113-c8a93983adbf/documentation/34660606-c10a7dc1-81ec-44cd-a9c2-59a4ef4d33a7)

## Entity Relationship Diagram (ERD)

[Link ERD Diagram](https://viewer.diagrams.net/?border=0&tags=%7B%7D&lightbox=1&highlight=0000ff&edit=_blank&layers=1&nav=1&title=erd%20booking%20field%20schedule.drawio#Uhttps%3A%2F%2Fdrive.google.com%2Fuc%3Fid%3D1bqD-fS_KVA_a1a6nl-NuxLstZ-zxffST%26export%3Ddownload)

## Backend Service Flow Diagram

[Link Backend Flow Diagram](https://viewer.diagrams.net/?border=0&tags=%7B%7D&lightbox=1&highlight=0000ff&edit=_blank&layers=1&nav=1&title=flow.drawio#Uhttps%3A%2F%2Fdrive.google.com%2Fuc%3Fid%3D1bIPRjL7gIAJlEQEtlI6nl-bigP29X1nf%26export%3Ddownload)

---

# Backend Service

### 1. User Service

User service adalah service untuk mengelola data user baik admin ataupun customer. Dimana endpointnya sebagai berikut:

| Method | Endpoint          | Role           |
|--------|-------------------|----------------|
| GET    | api/v1/auth/user  | Admin & Customer |
| GET    | api/v1/auth/:uuid | Admin & Customer |
| POST   | api/v1/auth/login | Admin & Customer |
| POST   | api/v1/auth/register | Customer       |
| PUT    | api/v1/auth/:uuid | Admin & Customer |

**Library**

| Name      | Library                            |
|-----------|----------------------------------|
| Gorm      | gorm.io/gorm                     |
| Postgres  | gorm.io/driver/postgres          |
| Env       | github.com/joho/godotenv         |
| UUID      | github.com/google/uuid           |
| JWT       | github.com/golang-jwt/jwt/v5     |
| Logger    | github.com/sirupsen/logrus       |
| Cobra     | github.com/spf13/cobra           |
| Viper     | github.com/spf13/viper           |
| Limitter  | github.com/didip/tollbooth       |
| Gin       | github.com/gin-gonic/gin          |
| Hashi Corp| github.com/hashicorp/consul/api  |

**Command Install**

```bash
go get gorm.io/gorm gorm.io/gorm \
gorm.io/driver/postgres \
github.com/joho/godotenv \
github.com/google/uuid \
github.com/golang-jwt/jwt/v5 \
github.com/sirupsen/logrus \
github.com/spf13/cobra \
github.com/spf13/viper \
github.com/didip/tollbooth \
github.com/gin-gonic/gin \
github.com/hashicorp/consul/api \
github.com/hashicorp/consul/sdk
```

### 2. Field Service

Field service adalah service untuk mengelola data lapangan dan juga jadwal lapangan yang tersedia. Endpointnya sebagai berikut:

| Method | Endpoint                      | Role             |
|--------|-------------------------------|------------------|
| GET    | api/v1/field/pagination       | Admin            |
| GET    | api/v1/field                  | Admin            |
| GET    | api/v1/field/:id              | Admin            |
| POST   | api/v1/field                  | Admin            |
| PUT    | api/v1/field/:id              | Admin            |
| DELETE | api/v1/field/:id              | Admin            |
| GET    | api/v1/field/schedule         | Admin & Customer |
| GET    | api/v1/field/schedule/:uuid   | Admin & Customer |
| GET    | api/v1/field/schedule/lists/:uuid | Admin & Customer |
| POST   | /api/v1/field/schedule/generate-one-month | Admin  |
| POST   | api/v1/field/schedule         | Admin            |
| PUT    | api/v1/field/schedule/:uuid   | Admin            |
| PATCH  | api/v1/field/schedule/status/:uuid | Customer    |
| DELETE | api/v1/field/schedule/:uuid   | Admin            |
| GET    | api/v1/time                   | Admin            |
| GET    | api/v1/time/:uuid             | Admin            |
| POST   | api/v1/time                   | Admin            |

**Library**

| Name          | Library                          |
|---------------|---------------------------------|
| Gorm          | gorm.io/gorm                    |
| Postgres      | gorm.io/driver/postgres         |
| Env           | github.com/joho/godotenv        |
| UUID          | github.com/google/uuid          |
| Go Request    | github.com/parnurzeal/gorequest@v0.2.16 |
| Logger        | github.com/sirupsen/logrus      |
| Cobra         | github.com/spf13/cobra          |
| Viper         | github.com/spf13/viper          |
| Limitter      | github.com/didip/tollbooth      |
| Gin           | github.com/gin-gonic/gin         |
| Pq            | github.com/lib/pq               |
| Config Currency| github.com/dustin/go-humanize  |
| Hashi Corp    | github.com/hashicorp/consul/api |
| Hashi Corp SDK| github.com/hashicorp/consul/sdk |

**Command Install**

```bash
go get gorm.io/gorm \
gorm.io/driver/postgres \
github.com/joho/godotenv \
github.com/google/uuid \
github.com/parnurzeal/gorequest@v0.2.16 \
github.com/sirupsen/logrus \
github.com/spf13/cobra \
github.com/spf13/viper \
github.com/didip/tollbooth \
github.com/gin-gonic/gin \
github.com/lib/pq \
github.com/dustin/go-humanize \
github.com/hashicorp/consul/api \
github.com/hashicorp/consul/sdk \
cloud.google.com/go/storage \
google.golang.org/api/option
```

### 3. Order Service

Order service adalah service untuk mengelola data order atau booking yang akan/telah diorder oleh customer. Endpointnya sebagai berikut:

| Method | Endpoint           | Role             |
|--------|--------------------|------------------|
| GET    | api/v1/order       | Admin & Customer |
| GET    | api/v1/order/:uuid | Admin & Customer |
| GET    | api/v1/order/user  | Customer         |
| POST   | api/v1/order       | Customer         |

**Library**

| Name      | Library                           |
|-----------|----------------------------------|
| Gorm      | gorm.io/gorm                    |
| Postgres  | gorm.io/driver/postgres         |
| Env       | github.com/joho/godotenv        |
| UUID      | github.com/google/uuid          |
| Go Request| github.com/parnurzeal/gorequest@v0.2.16 |
| Logger    | github.com/sirupsen/logrus      |
| Cobra     | github.com/spf13/cobra          |

**Command Install**

```bash
go get gorm.io/gorm \
gorm.io/driver/postgres \
github.com/joho/godotenv \
github.com/google/uuid \
github.com/parnurzeal/gorequest@v0.2.16 \
github.com/sirupsen/logrus \
github.com/spf13/cobra
```

### 4. Payment Service

Payment service adalah service untuk mengelola data pembayaran dari customer yang telah membooking lapangan. Dimana endpointnya sebagai berikut:

| Method | Endpoint             | Role             |
|--------|----------------------|------------------|
| GET    | api/v1/payment       | Admin            |
| GET    | api/v1/payment/:uuid | Admin            |
| POST   | api/v1/payment       | Customer         |
| POST   | api/v1/payment/webhook | System          |

**Library**

| Name           | Library                          |
|----------------|---------------------------------|
| Gorm           | gorm.io/gorm                    |
| Postgres       | gorm.io/driver/postgres         |
| Env            | github.com/joho/godotenv        |
| UUID           | github.com/google/uuid          |
| Go Request     | github.com/parnurzeal/gorequest@v0.2.16 |
| Logger         | github.com/sirupsen/logrus      |
| Cobra          | github.com/spf13/cobra          |
| Viper          | github.com/spf13/viper          |
| Limitter       | github.com/didip/tollbooth      |
| Midtrans       | github.com/midtrans/midtrans-go |
| Apache Kafka   | github.com/IBM/sarama            |
| Google Cloud Storage | cloud.google.com/go/storage  |
| Golang Option  | google.golang.org/api/option    |
| Wkhtmltopdf    | github.com/SebastiaanKlippert/go-wkhtmltopdf |
| Gin            | github.com/gin-gonic/gin         |
| Config Currency| github.com/dustin/go-humanize   |
| Hashi Corp     | github.com/hashicorp/consul/api |
| Hashi Corp SDK | github.com/hashicorp/consul/sdk |

**Command Install**

```bash
go get gorm.io/gorm \
gorm.io/driver/postgres \
github.com/joho/godotenv \
github.com/google/uuid \
github.com/parnurzeal/gorequest@v0.2.16 \
github.com/sirupsen/logrus \
github.com/spf13/cobra \
github.com/spf13/viper \
github.com/didip/tollbooth \
github.com/midtrans/midtrans-go \
github.com/IBM/sarama \
cloud.google.com/go/storage \
google.golang.org/api/option \
github.com/SebastiaanKlippert/go-wkhtmltopdf \
github.com/gin-gonic/gin \
github.com/dustin/go-humanize \
github.com/hashicorp/consul/api \
github.com/hashicorp/consul/sdk
