package templates

var DockerComposeTemplate = `services:
  app:
    build: ./infrastructure/docker/Dockerfile
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=production
      - MONGO_URI=mongodb://${MONGO_USER}:${MONGO_PASSWORD}@mongo:27017
      - MONGO_DB_NAME=${MONGO_DB_NAME}
    volumes:
      - ./config:/app/config
    depends_on:
      mongo:
        condition: service_healthy
    networks:
      - app-network
    restart: unless-stopped

  mongo:
    image: mongo:latest
    environment:
      - MONGO_INITDB_ROOT_USERNAME=${MONGO_USER}
      - MONGO_INITDB_ROOT_PASSWORD=${MONGO_PASSWORD}
    volumes:
      - mongo-data:/data/db
    networks:
      - app-network
    healthcheck:
      test: echo 'db.runCommand("ping").ok' | mongosh localhost:27017/test --quiet
      interval: 10s
      timeout: 10s
      retries: 5
      start_period: 40s
    restart: unless-stopped

volumes:
  mongo-data:

networks:
  app-network:
    driver: bridge`

func NewDockerComposeTemplate() string {
	return DockerComposeTemplate
}
