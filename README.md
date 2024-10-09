# Diskreader

## Overview
Diskreader is a Go application designed to read files from a specified directory in a Kubernetes environment. It walks through the directory structure, skipping socket files and non-existent files, while reporting the size of successfully read files.

## Features
- Recursively traverses a specified directory (`/mnt/data`).
- Skips socket files and non-existent files.
- Reports the number of bytes read from each valid file.

## Prerequisites
- Docker
- Kubernetes cluster
- Persistent Volume and Persistent Volume Claim configured for your application
- Kubernetes CLI (`kubectl`) installed and configured

## Creating the Dockerfile
To build the Docker image, you need a Dockerfile. Use the following [Dockerfile](Dockerfile)

## Building the Docker Image
To build the Docker image for the Diskreader application, use the following command:

```bash

sudo docker build -t abhishekkamat27/diskreader:v2 .

```


## Logging into Docker
Before pushing the Docker image to your repository, log in using the following command:

```bash


docker login -u abhishekkamat27

```

## Pushing the Docker Image to Docker Hub
After successfully building the Docker image, push it to your Docker registry with the following command:

```bash


sudo docker push abhishekkamat27/diskreader:v2

```



