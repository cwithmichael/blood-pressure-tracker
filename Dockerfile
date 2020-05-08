FROM openjdk:14-jdk-alpine
WORKDIR /code
COPY . .
ENTRYPOINT ["./mvnw", "clean", "spring-boot:run"]