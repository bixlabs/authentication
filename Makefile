all: deps

deps:
		sh ./scripts/install_dep.sh
		sh ./scripts/install_air.sh
		dep ensure

clean:
		rm -r -f ./tmp

run:
		~/.air -c .air.config
