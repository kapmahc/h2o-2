
use super::super::env;

pub fn run(cfg: &env::config::Config, name: &str, threads: usize) -> env::errors::Result<bool> {
    println!("Start worker {}[{}]-{}", name, threads, cfg.env());
    Ok(true)
}
