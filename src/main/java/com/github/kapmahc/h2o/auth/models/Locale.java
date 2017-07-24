package com.github.kapmahc.h2o.auth.models;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.io.Serializable;

public class Locale implements Serializable {
    @JsonProperty
    private Long id;
    private String code;
    private String message;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getCode() {
        return code;
    }

    public void setCode(String code) {
        this.code = code;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }
}
