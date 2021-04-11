default: dev 

test: ensure
	bash scripts/test.sh

dev: ensure
	bash scripts/dev.sh


.PHONY: ensure build test 