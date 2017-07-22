# H2O

A complete open source e-commerce solution.

## Install jdk

```bash
curl -s "https://get.sdkman.io" | zsh
sdk install java
sdk install gradle
```

## Build

```bash
git clone https://github.com/kapmahc/h2o.git
gradle build
```

## Deployment

```bash
vi application-product.properties
vi logback.xml
java -jar -Dspring.profiles.active=production h2o-*.jar
```

- [sdkman](http://sdkman.io/usage.html)
- [spring boot](https://spring.io/guides)
