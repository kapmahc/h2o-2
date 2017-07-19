target=h2o

build:
	cargo build --release && strip target/release/$(target)



