package com.github.kapmahc;

import com.fasterxml.jackson.annotation.JsonProperty;
import io.dropwizard.Configuration;
import org.hibernate.validator.constraints.NotEmpty;

class H2OConfiguration extends Configuration {
    @NotEmpty
    @JsonProperty
    private String theme;

    @NotEmpty
    @JsonProperty
    private String host;

    @JsonProperty
    private boolean https;

    public String getTheme() {
        return theme;
    }

    public void setTheme(String theme) {
        this.theme = theme;
    }

    public String getHost() {
        return host;
    }

    public void setHost(String host) {
        this.host = host;
    }

    public boolean isHttps() {
        return https;
    }

    public void setHttps(boolean https) {
        this.https = https;
    }
}
