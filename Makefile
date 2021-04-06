default: dev 

test: ensure
	bash scripts/test.sh


.PHONY: ensure build test 