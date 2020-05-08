FROM maven:3.6.3-openjdk-11-slim as maven
WORKDIR /build

COPY ./pom.xml .

RUN mvn dependency:go-offline

COPY ./src/ ./src
RUN mvn package

FROM openjdk:11-jdk-slim
WORKDIR /code
COPY --from=maven /build/target/blood_pressure_tracker*.jar ./app.jar

CMD ["java", "-jar", "/code/app.jar"]