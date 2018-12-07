

all: deps

deps:
		chmod +x ./scripts/install_dep.sh && ./scripts/install_dep.sh
		chmod +x ./scripts/install_air.sh && ./scripts/install_air.sh
		chmod +x ./scripts/install_go_dependencies.sh && ./scripts/install_go_dependencies.sh
