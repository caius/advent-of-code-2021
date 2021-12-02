TODAY_NUM ?= $(shell date +%d | sed -r 's/^0//')
TODAY_DIR ?= "day$(TODAY_NUM)"

.PHONY: today
today:
	mkdir -p $(TODAY_DIR)
	touch $(TODAY_DIR)/{example,input}.txt $(TODAY_DIR)/main.go
	echo ".PHONY: run\nrun:\n	go run main.go < input.txt\n\n.PHONY: example\nexample:\n	go run main.go < example.txt\n" > $(TODAY_DIR)/Makefile
