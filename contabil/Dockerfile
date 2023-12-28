FROM gradle:jdk17 AS build

# Set the working directory
ENV APP_HOME=/usr/app
WORKDIR $APP_HOME
COPY build.gradle settings.gradle gradlew $APP_HOME
COPY gradle $APP_HOME/gradle


#Build
RUN ./gradlew build || return 0
COPY . .
RUN gradle build


# Container with builded app
FROM openjdk:17-alpine

ENV ARTIFACT_NAME=contabil-0.0.1-SNAPSHOT.jar
ENV APP_HOME=/usr/app

WORKDIR $APP_HOME

COPY --from=build $APP_HOME/build/libs/$ARTIFACT_NAME .
COPY --from=build $APP_HOME/.env .

EXPOSE 8080

ENTRYPOINT exec java -jar ${ARTIFACT_NAME}

CMD ["java", "-jar", "${ARTIFACT_NAME}"]