pub mod db;
pub mod generate;
pub mod server;
pub mod worker;

use docopt;
use super::env;

docopt!(Args derive Debug, "
H2O - A complete open source e-commerce solution by rust language.

Usage:
  h2o server [--worker]
  h2o worker [--name=<wn>] [--threads=<tn>]
  h2o generate (config|nginx)
  h2o generate (migration|locale) --name=<fn>
  h2o db (create|migrate|rollback|status|drop)
  h2o (-h | --help)
  h2o (-v | --version)

Options:
  -h --help         Show this screen.
  --version         Show version.
  --worker          Start with a background-job worker.
  --name=<wn>       Background-job worker's name, default is hostname.
  --name=<fn>       Filename.
  --threads=<tn>    Threads in worker [default: 2].
");

fn parse(r: env::errors::Result<bool>) {
    match r {
        Ok(_) => println!("Done."),
        Err(e) => println!("{}", e),
    }
}

pub fn run() {
    let args: Args = Args::docopt().deserialize().unwrap_or_else(|e| e.exit());
    if args.flag_v || args.flag_version {
        println!("{}", env::version());
        return;
    }

    if args.cmd_generate {
        if args.cmd_config {
            parse(generate::config());
            return;
        }
        if args.cmd_nginx {
            parse(generate::nginx());
            return;
        }
        if args.cmd_locale {
            parse(generate::locale(&args.flag_name));
            return;
        }
        if args.cmd_migration {
            parse(generate::migration(&args.flag_name));
            return;
        }
    }

    if args.cmd_db {
        if args.cmd_create {
            parse(db::create());
            return;
        }
        if args.cmd_drop {
            parse(db::drop());
            return;
        }
        if args.cmd_migrate {
            parse(db::migrate());
            return;
        }
        if args.cmd_rollback {
            parse(db::rollback());
            return;
        }
        if args.cmd_status {
            parse(db::status());
            return;
        }
    }

    if args.cmd_server {
        parse(server::run(args.flag_worker));
        return;
    }
    if args.cmd_worker {
        match args.flag_threads.parse::<u32>() {
            Ok(threads) => parse(worker::run(&args.flag_name, threads)),
            Err(e) => println!("{}", e),
        }
        return;
    }
    // println!("{:?}", args);
    println!("Unknown!");
}
