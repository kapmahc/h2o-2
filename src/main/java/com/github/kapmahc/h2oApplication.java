package com.github.kapmahc;

import io.dropwizard.Application;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;

public class h2oApplication extends Application<h2oConfiguration> {

    public static void main(final String[] args) throws Exception {
        new h2oApplication().run(args);
    }

    @Override
    public String getName() {
        return "h2o";
    }

    @Override
    public void initialize(final Bootstrap<h2oConfiguration> bootstrap) {
        // TODO: application initialization
    }

    @Override
    public void run(final h2oConfiguration configuration,
                    final Environment environment) {
        // TODO: implement application
    }

}
