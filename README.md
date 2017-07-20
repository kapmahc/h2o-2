# H2O

A complete open source e-commerce solution by rust language.

## Install rust

```bash
sudo pacman -S rustup
rustup default nightly
# upgrade
rustup update && cargo update
# for development
rustup run nightly cargo install rustfmt-nightly
cargo install racer
rustup component add rust-src
```

## Build

```bash
git clone https://github.com/kapmahc/h2o.git
cd h2o
make
```

## Deployment

````bash
```
- Create database
```sql
psql -U postgres
CREATE DATABASE db-name WITH ENCODING = 'UTF8';
CREATE USER user-name WITH PASSWORD 'change-me';
GRANT ALL PRIVILEGES ON DATABASE db-name TO user-name;
```

- 'Peer authentication failed for user', open file "/etc/postgresql/9.5/main/pg_hba.conf" change line:

  ```
  local   all             all                                     peer  
  TO:
  local   all             all                                     md5
  ```

- Generate openssl certs

  ```bash
  openssl genrsa -out www.change-me.com.key 2048
  openssl req -new -x509 -key www.change-me.com.key -out www.change-me.com.crt -days 3650 # Common Name:*.change-me.com
````

## Editors

## Atom

- language-rust

## Documents

- [rust book](https://doc.rust-lang.org/book/)
- [24 days of rust](siciarz.net)
- [packages for rust](https://crates.io/)
- [rocket.rs](https://rocket.rs/guide/)
- [serde](https://serde.rs/)
- [rustfmt](https://github.com/rust-lang-nursery/rustfmt/blob/master/atom.md)
- [racer](https://atom.io/packages/autocomplete-racer)
