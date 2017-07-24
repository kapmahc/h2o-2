# H2O

A complete open source e-commerce solution.

## Install jdk

```bash
curl -s "https://get.sdkman.io" | zsh
sdk install java
sdk install maven
```

## Usage

```bash
git clone https://github.com/kapmahc/h2o.git
cd h2o
mvn package
java -jar target/h2o-1.0-SNAPSHOT.jar server config.yml
```

## Development
```bash
# download sources
mvn dependency:sources
```



- [sdkman](http://sdkman.io/usage.html)
- [dropwizard](http://www.dropwizard.io)
- [hibernate](https://docs.jboss.org/hibernate/orm/current/userguide/html_single/Hibernate_User_Guide.html)
- [dropwizard](http://www.dropwizard.io/1.1.2/docs/manual/index.html)
