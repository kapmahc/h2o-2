extern crate h2o;
// extern crate env_logger;
extern crate syslog;

fn main() {
    // env_logger::init().unwrap();
    syslog::init(
        syslog::Facility::LOG_USER,
        h2o::env::config::log_level(),
        None,
    ).unwrap();
    h2o::app::run()
}
