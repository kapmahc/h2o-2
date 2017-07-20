target=h2o
dist=dist

build:
	cargo build --release && strip target/release/$(target)
	mkdir -pv $(dist)
	cp -r target/release/$(target) package.json db locales themes templates $(dist)/
	tar jcvf $(dist).tar.bz2 $(dist)


clean:
	-rm -rf $(dist) $(dist).tar.bz2
