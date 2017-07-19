pub mod db;
pub mod generate;
pub mod server;
pub mod worker;

extern crate serde;
extern crate docopt;

use self::docopt::Docopt;

const USAGE: &'static str = "
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
";

#[derive(Debug, Deserialize)]
struct Args {
    cmd_generate: bool,
    cmd_nginx: bool,
    cmd_config: bool,
    cmd_migration: bool,
    cmd_locale: bool,
    cmd_ssl: bool,

    cmd_db: bool,
    cmd_create: bool,
    cmd_connect: bool,
    cmd_migrate: bool,
    cmd_rollback: bool,
    cmd_drop: bool,
    cmd_status: bool,

    cmd_server: bool,
    flag_worker: bool,

    cmd_worker: bool,
    flag_name: Vec<String>,
    flag_threads: isize,

    flag_version: bool,
    flag_v: bool,
}

pub fn run() {
    let args: Args = Docopt::new(USAGE)
        .and_then(|d| d.deserialize())
        .unwrap_or_else(|e| e.exit());
    println!("{:?}", args);
}
