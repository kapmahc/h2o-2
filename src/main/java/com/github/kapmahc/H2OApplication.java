package com.github.kapmahc;

import com.github.kapmahc.h2o.auth.cli.NginxCommand;
import com.github.kapmahc.h2o.auth.health.DatabaseHealth;
import com.github.kapmahc.h2o.auth.health.RabbitMQHealth;
import com.github.kapmahc.h2o.auth.health.RedisHealth;
import com.github.kapmahc.h2o.auth.models.Locale;
import com.github.kapmahc.h2o.auth.resources.LocaleResource;
import com.github.kapmahc.h2o.auth.dao.LocaleDao;
import com.github.kapmahc.h2o.auth.dao.impl.LocaleDaoImpl;
import com.github.kapmahc.h2o.auth.tasks.BackupTask;
import io.dropwizard.Application;
import io.dropwizard.db.DataSourceFactory;
import io.dropwizard.db.PooledDataSourceFactory;
import io.dropwizard.forms.MultiPartBundle;
import io.dropwizard.hibernate.HibernateBundle;
import io.dropwizard.migrations.MigrationsBundle;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;
import org.glassfish.hk2.utilities.binding.AbstractBinder;

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
        bt.addBundle(new MultiPartBundle());
        bt.addBundle(hibernate);
    }

    @Override
    public void run(final H2OConfiguration cfg,
                    final Environment env) {
        // inject
        env.jersey().register(new AbstractBinder() {
            @Override
            protected void configure() {
                bind(cfg).to(H2OConfiguration.class);
                bind(new LocaleDaoImpl(hibernate.getSessionFactory())).to(LocaleDao.class);
            }
        });

        // resources
        env.jersey().register(LocaleResource.class);
        // services
        // health checks
        env.healthChecks().register("redis", new RedisHealth());
        env.healthChecks().register("database", new DatabaseHealth());
        env.healthChecks().register("rabbitmq", new RabbitMQHealth());
        // tasks
        env.admin().addTask(new BackupTask());
    }

    private final HibernateBundle<H2OConfiguration> hibernate = new HibernateBundle<H2OConfiguration>(Locale.class) {
        @Override
        public DataSourceFactory getDataSourceFactory(H2OConfiguration cfg) {
            return cfg.getDatabase();
        }
    };


}
