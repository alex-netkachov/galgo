# Run it in the problem's folder as `make -f ../makefile`

export GOPATH=$(CURDIR)

SRC = $(shell find . -name '*.go')
TESTS = $(shell find . -name '*.in')

.PHONY: all format new $(TESTS)

all: format $(TESTS)

format:
	go fmt $(SRC)

$(TESTS):
	go run $(SRC) < $@ | diff -w $(@:.in=.out) -

# Create folder for the solution `make -f ../makefile new NAME=solution`
new:
ifeq ($(NAME),)
	@echo "Invalid syntax. The valid command should look like"
	@echo "  make -f ../makefile new NAME=solution"
else
	mkdir ../$(NAME)
	touch ../$(NAME)/$(NAME).go
	touch ../$(NAME)/$(NAME).in
	touch ../$(NAME)/$(NAME).out
endif
