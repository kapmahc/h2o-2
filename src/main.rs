extern crate h2o;
// extern crate env_logger;
extern crate syslog;
extern crate log;

fn main() {
    // env_logger::init().unwrap();
    syslog::init(
        syslog::Facility::LOG_USER,
        log::LogLevelFilter::max(),
        None).unwrap();
    h2o::app::run()
}
