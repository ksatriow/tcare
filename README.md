# TCare

This project demonstrates a simple microservices architecture using Go, Docker, Docker Compose, MySQL, and Redis. The setup includes an API Gateway that routes requests to different services: UserService, BillService, BookService, and ConsultationService.

##  Application Architecture

![AppArch](/docs/apparch.png)

##  Environment Architecture

![EnvArch](/docs/envarch.png)

## Project Structure

```sh
├── ApiGateway
│   ├── Dockerfile
│   ├── apigateway-chart
│   │   ├── Chart.yaml
│   │   ├── templates
│   │   │   ├── deployment.yaml
│   │   │   ├── hpa.yaml
│   │   │   ├── ingress.yaml
│   │   │   └── service.yaml
│   │   └── values.yaml
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── BillService
│   ├── Dockerfile
│   ├── bill-chart
│   │   ├── Chart.yaml
│   │   ├── templates
│   │   │   ├── deployment.yaml
│   │   │   ├── hpa.yaml
│   │   │   ├── ingress.yaml
│   │   │   └── service.yaml
│   │   └── values.yaml
│   └── main.go
├── BookService
│   ├── Dockerfile
│   ├── book-chart
│   │   ├── Chart.yaml
│   │   ├── templates
│   │   │   ├── deployment.yaml
│   │   │   ├── hpa.yaml
│   │   │   ├── ingress.yaml
│   │   │   └── service.yaml
│   │   └── values.yaml
│   └── main.go
├── ConsultService
│   ├── Dockerfile
│   ├── consult-chart
│   │   ├── Chart.yaml
│   │   ├── templates
│   │   │   ├── deployment.yaml
│   │   │   ├── hpa.yaml
│   │   │   ├── ingress.yaml
│   │   │   └── service.yaml
│   │   └── values.yaml
│   └── main.go
├── README.md
├── Terraform
│   ├── eks_cluster.tf
│   ├── eks_node_group.tf
│   ├── iam_role.tf
│   ├── internet_gateway.tf
│   ├── provider.tf
│   ├── route.tf
│   ├── security_group.tf
│   ├── subnet.tf
│   └── vpc.tf
├── UserService
│   ├── Dockerfile
│   ├── main.go
│   └── user-chart
│       ├── Chart.yaml
│       ├── templates
│       │   ├── deployment.yaml
│       │   ├── hpa.yaml
│       │   ├── ingress.yaml
│       │   └── service.yaml
│       └── values.yaml
├── docker-compose.yml
├── docs
│   ├── apparch.png
│   └── envarch.png
├── go.mod
└── go.sum
```
