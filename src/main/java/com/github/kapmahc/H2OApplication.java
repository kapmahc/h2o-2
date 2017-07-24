package com.github.kapmahc;

import com.github.kapmahc.h2o.auth.resources.LocaleResource;
import com.github.kapmahc.h2o.auth.cli.NginxCommand;
import com.github.kapmahc.h2o.auth.health.RedisHealth;
import com.github.kapmahc.h2o.auth.tasks.BackupTask;
import io.dropwizard.Application;
import io.dropwizard.db.PooledDataSourceFactory;
import io.dropwizard.migrations.MigrationsBundle;
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
    public void initialize(final Bootstrap<H2OConfiguration> bt) {
        bt.addCommand(new NginxCommand());
        bt.addBundle(new MigrationsBundle<H2OConfiguration>() {
            @Override
            public PooledDataSourceFactory getDataSourceFactory(H2OConfiguration cfg) {
                return cfg.getDatabase();
            }
        });
    }

    @Override
    public void run(final H2OConfiguration cfg,
                    final Environment env) {
        final LocaleResource localeResource = new LocaleResource();
        final RedisHealth redisHealth = new RedisHealth();

        env.healthChecks().register("redis", redisHealth);
        env.jersey().register(localeResource);
        env.admin().addTask(new BackupTask());
    }

}
