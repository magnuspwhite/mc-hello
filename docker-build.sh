#! /bin/bash

TAG=$(<version)
echo "tag: ${TAG}"
IMAGE="localhost:5000/mc-hello"
IMAGE_WITH_TAG="${IMAGE}:${TAG}"
echo "image with tag: ${IMAGE_WITH_TAG}"

docker build . -t "${IMAGE_WITH_TAG}"
docker push "${IMAGE_WITH_TAG}"
