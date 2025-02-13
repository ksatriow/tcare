# TCare

This project demonstrates a simple microservices architecture using Go, Docker, Docker Compose, MySQL, and Redis. The setup includes an API Gateway that routes requests to different services: UserService, BillService, BookService, and ConsultationService.

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

## The Three Tiers

- **Presentation Layer (Client Tier)** 

    The top-most level of the application is the user interface (UI). This layer directly interacts with users, presenting the system's features and data. This top-level tier can run on a web browser, as desktop application, or a graphical user interface (GUI), **(frontend)**. It's responsible for collecting user inputs and displaying data. Technologies in this layer include **HTML**.

- **Business Logic Layer (Application Tier)**

    This layer processes user inputs, applies business logic, and makes decisions. It acts as an intermediary between the presentation layer and the database layer, ensuring the data exchanged between these layers is processed according to the business rules **(backend)**. Technologies used here include backend programming languages in this case **Go**.

- **Data Access Layer (Database Tier)**

    The bottom layer is where data is stored and retrieved. This layer handles database management and communication, ensuring data integrity and security. It can involve relational databases like **MySQL**.

##  Application Architecture

![AppArch](/docs/apparch.png)

### The Component

**API Gateway**
The gateway provides a unified entry point for client applications. It handles routing, filtering, and load balancing.

**Service Registry**
The service registry contains the details of all the services. The gateway discovers the service using the registry. For example, Consul, Eureka, Zookeeper, etc.

**Service Layer**
Each microservices serves a specific business function and can run on multiple instances. These services can be built using frameworks like Golang Mux, Spring Boot, NestJS, etc.

**Authorization Server**
Used to secure the microservices and manage identity and access control. Tools like Keycloak, Azure AD, and Okta can help over here.

**Data Storage**
Databases like PostgreSQL and MySQL can store application data generated by the services.

**Distributed Caching**
Caching is a great approach for boosting the application performance. Options include caching solutions like Redis, Couchbase, Memcached, etc.

**Async Microservices Communication**
Use platforms such as Kafka and RabbitMQ to support async communication between microservices.

**Metrics Visualization**
Microservices can be configured to publish metrics to Prometheus and tools like Grafana can help visualize the metrics.

**Log Aggregation and Visualization**
Logs generated by the services are aggregated using Logstash, stored in Elasticsearch, and visualized with Kibana.

##  Environment Architecture

![EnvArch](/docs/envarch.png)

### Environment Infrastructure

- **Terraform**: Terraform is an open-source infrastructure as code (IaC) tool used for building, changing, and versioning infrastructure safely and efficiently. It enables users to define and provision infrastructure resources such as virtual machines, networks, storage, and more using a declarative configuration language.

- **EKS Cluster** : Amazon Elastic Kubernetes Service (EKS) is a managed Kubernetes service provided by AWS. It allows users to deploy, manage, and scale containerized applications using Kubernetes on AWS infrastructure. EKS clusters provide features such as automated Kubernetes upgrades, integrated AWS IAM authentication, and seamless integration with other AWS services.

- **AWS CLI** : AWS Command Line Interface (CLI) is a unified tool that provides commands for interacting with various AWS services from the command line. It enables users to manage AWS resources, configure AWS settings, and automate tasks using scripts or command-line interfaces.

- **Kubectl** : Kubectl is a command-line tool used for interacting with Kubernetes clusters. It allows users to deploy and manage applications, inspect cluster resources, and troubleshoot issues within Kubernetes environments. Kubectl provides a wide range of commands for managing Kubernetes clusters, nodes, pods, services, and more.

- **Prometheus** : Prometheus is an open-source monitoring and alerting toolkit designed for monitoring containerized and cloud-native applications. It collects metrics from various sources, stores them in a time-series database, and enables querying and visualization of metrics data. Prometheus is widely used for monitoring Kubernetes clusters and applications deployed on them.

- **Grafana** : Grafana is an open-source analytics and visualization platform used for creating dashboards and graphs to monitor and analyze time-series data. It integrates with various data sources, including Prometheus, to visualize metrics collected from applications, servers, and other sources. Grafana provides interactive and customizable dashboards for monitoring infrastructure and applications effectively.

- **Helm** : Helm is a package manager for Kubernetes that streamlines the process of installing, upgrading, and managing applications on Kubernetes clusters. It uses charts, which are packages containing pre-configured Kubernetes resources, to simplify the deployment and management of complex applications.

- **VPC** : A virtual network you create in AWS to isolate your resources.

- **Public Subnet** : A subnet within a VPC where instances can have public IP addresses and access the internet directly.

- **Private Subnet** : A subnet within a VPC where instances cannot have public IP addresses by default. They are typically used for resources that don't need direct internet access but can communicate with the internet through a NAT Gateway (optional).

- **NAT Gateway** : A service that allows instances in private subnets to access the internet without directly exposing them publicly. It translates private IP addresses to a public IP for outbound traffic.

- **Internet Gateway**: A gateway that allows resources in your VPC to access the internet.

- **Route Table**: A table that defines how traffic is routed within your VPC. You can have separate route tables for public and private subnets to control traffic flow.

## Instruction use Terraform

```bash
cd Terraform
terraform init
terraform validate
terraform plan
terraform apply
```

## Instruction use Microservice

### Manual 

```bash
go run ApiGateway/main.go
go run BillService/main.go
go run BookService/main.go
go run ConsultService/main.go
go run UserService/main.go
```

### Docker

```bash
docker build -f ApiGateway/Dockerfile -t apigateway .
docker run -dit -p 8080:8080 apigateway

docker build -f BillService/Dockerfile -t bill .
docker run -dit -p 8083:8083 bill

docker build -f BookService/Dockerfile -t book .
docker run -dit -p 8082:8082 book

docker build -f ConsultService/Dockerfile -t consult .
docker run -dit -p 8084:8084 consult

docker build -f UserService/Dockerfile -t user .
docker run -dit -p 8081:8081 user

docker run --name redis -p 6379:6379 -d redis
```

### Docker Compose

```bash
docker compose build
docker compose up -d
```

##  Instruction use Helm

```sh
cd ApiGateway/
helm upgrade --install apigateway-chart -f values.yaml

cd UserService/
helm upgrade --install user-chart -f values.yaml

cd BookService/
helm upgrade --install book-chart -f values.yaml

cd BillService/
helm upgrade --install bill-chart -f values.yaml

cd ConsultService/
helm upgrade --install consult-chart -f values.yaml
```

