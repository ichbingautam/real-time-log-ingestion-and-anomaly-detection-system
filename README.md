# Real-Time Log Ingestion & Anomaly Detection System

[![Go Report Card](https://goreportcard.com/badge/github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system)](https://goreportcard.com/report/github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker Pulls](https://img.shields.io/docker/pulls/ichbingautam/log-processor)](https://hub.docker.com/r/ichbingautam/log-processor)

![Kafka + Go Architecture](docs/architecture.png)

A high-performance distributed system for real-time log processing and anomaly detection using Apache Kafka and Go.

## Features

- 🚀 **Million+ Events/Sec** throughput capability
- 🔍 **Real-Time Anomaly Detection** with rule engine
- ☁️ **Multi-Cloud Storage** (Elasticsearch + S3)
- 🔒 **Enterprise Security** (TLS + SASL)
- 📊 **Monitoring** with Prometheus/Grafana
- 🛠 **Auto-Scaling** Kubernetes deployment
- 🔄 **Multi-DC Replication** with MirrorMaker 2.0

## Tech Stack

![Apache Kafka](https://img.shields.io/badge/Apache%20Kafka-000?style=for-the-badge&logo=apachekafka)
![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Elasticsearch](https://img.shields.io/badge/ElasticSearch-005571?style=for-the-badge&logo=elasticsearch)
![AWS S3](https://img.shields.io/badge/Amazon_S3-FF9900?style=for-the-badge&logo=amazonaws&logoColor=white)

## Getting Started

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- Apache Kafka 3.0+
- Elasticsearch 8.0+

### Installation
```bash
# Clone repository
git clone https://github.com/ichbingautam/real-time-log-ingestion-and-anomaly-detection-system.git
cd real-time-log-ingestion-and-anomaly-detection-system

# Build project
make build

# Start dependencies
docker-compose -f deployments/docker-compose.yml up -d