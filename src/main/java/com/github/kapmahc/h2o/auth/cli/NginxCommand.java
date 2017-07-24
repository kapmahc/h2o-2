package com.github.kapmahc.h2o.auth.cli;

import io.dropwizard.cli.Command;
import io.dropwizard.setup.Bootstrap;
import net.sourceforge.argparse4j.inf.Namespace;
import net.sourceforge.argparse4j.inf.Subparser;

public class NginxCommand  extends Command {
    public NginxCommand() {
        super("nginx", "Generate a nginx.conf file");
    }

    @Override
    public void configure(Subparser subparser) {
        subparser.addArgument("-H", "--host")
                .dest("host")
                .type(String.class)
                .required(true)
                .help("Hostname");
        subparser.addArgument("-s", "--ssl")
                .dest("ssl")
                .type(Boolean.class)
                .setDefault(false)
                .help("Enable ssl?");

    }

    @Override
    public void run(Bootstrap<?> bootstrap, Namespace namespace) throws Exception {
    System.out.println("hostname "+namespace.getString("host"));
    }
}
