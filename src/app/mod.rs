pub mod db;
pub mod generate;
pub mod server;
pub mod worker;

use docopt;

docopt!(Args derive Debug, "
H2O - A complete open source e-commerce solution by rust language.

Usage:
  h2o server [--worker]
  h2o worker [--name=<wn>] [--threads=<tn>]
  h2o generate (config|nginx|migration|locale|ssl)
  h2o db (create|connect|migrate|rollback|drop|status)
  h2o (-h | --help)
  h2o (-v | --version)

Options:
  -h --help     Show this screen.
  --version     Show version.
  --worker      Start with a background-job worker.
  --name=<wn>   Background-job worker's name, default is hostname.
  --threads=<tn>  Threads in worker [default: 2].
");

pub fn run() {
    let args: Args = Args::docopt().deserialize().unwrap_or_else(|e| e.exit());
    if args.flag_v || args.flag_version {
        println!("{}", super::env::version());
        return;
    }
    if args.cmd_generate {
        if args.cmd_nginx {
            generate::nginx();
            return;
        }
        if args.cmd_locale {
            generate::locale();
            return;
        }
        if args.cmd_migration {
            generate::migration();
            return;
        }
        if args.cmd_ssl {
            generate::ssl();
            return;
        }
    }
    println!("{:?}", args);
    // println!("{}", args.to_string());
    // args.help(true)

}
