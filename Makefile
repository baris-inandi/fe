build:
	paru -S --removemake --skipreview --needed paruz
	rm -rf /usr/bin/fe
	go build
	mv fe /usr/bin/fe
	rm -rf fe