target=h2o
dist=dist

build:
	cargo build --release && strip target/release/$(target)
	mkdir -pv $(dist)
	cp -r target/release/$(target) db locales templates $(dist)/
	cd desktop && npm run build
	cp -r desktop/build $(dist)/public
	tar jcvf $(dist).tar.bz2 $(dist)


clean:
	-rm -r target/release/$(target) desktop/build $(dist) $(dist).tar.bz2
