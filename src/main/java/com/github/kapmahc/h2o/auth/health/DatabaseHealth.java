package com.github.kapmahc.h2o.auth.health;

import com.codahale.metrics.health.HealthCheck;

public class DatabaseHealth extends HealthCheck {

    @Override
    protected Result check() throws Exception {
        return Result.healthy();
    }

    public DatabaseHealth() {
    }
}
