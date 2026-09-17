BIN_DIR := bin
MODULES := $(patsubst %/go.mod,%,$(wildcard */go.mod))

.PHONY: all build clean $(MODULES)

all: build

build: $(MODULES)

$(MODULES):
	@mkdir -p $(BIN_DIR)
	cd $@ && go build -o ../$(BIN_DIR)/$@ .
	cp $@/policy.rego $(BIN_DIR)/$@.rego

clean:
	rm -rf $(BIN_DIR)
