package com.github.kapmahc.h2o.auth.health;

import com.codahale.metrics.health.HealthCheck;

public class RedisHealth extends HealthCheck {
    @Override
    protected Result check() throws Exception {
        if (host == null) {
            return Result.unhealthy("host is null");
        }
        return Result.healthy();
    }

    public RedisHealth(String host) {
        this.host = host;
    }

    private final String host;

}
