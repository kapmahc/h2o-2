package com.github.kapmahc;

import com.github.kapmahc.h2o.auth.resources.LocaleResource;
import com.github.kapmahc.h2o.auth.cli.NginxCommand;
import com.github.kapmahc.h2o.auth.health.RedisHealth;
import com.github.kapmahc.h2o.auth.tasks.BackupTask;
import io.dropwizard.Application;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;

public class H2OApplication extends Application<H2OConfiguration> {

    public static void main(final String[] args) throws Exception {
        new H2OApplication().run(args);
    }

    @Override
    public String getName() {
        return "h2o";
    }

    @Override
    public void initialize(final Bootstrap<H2OConfiguration> bootstrap) {
        bootstrap.addCommand(new NginxCommand());
    }

    @Override
    public void run(final H2OConfiguration configuration,
                    final Environment environment) {
        final LocaleResource localeResource = new LocaleResource(configuration.getTheme());
        final RedisHealth redisHealth = new RedisHealth(configuration.getHost());

        environment.healthChecks().register("redis", redisHealth);
        environment.jersey().register(localeResource);
        environment.admin().addTask(new BackupTask());
    }

}
