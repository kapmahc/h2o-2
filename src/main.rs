extern crate h2o;
extern crate env_logger;

fn main() {
    env_logger::init().unwrap();
    h2o::app::run();
}
