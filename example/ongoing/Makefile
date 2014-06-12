
# available tags: console
TAGS = console

BUILD_DIR = build
SRC_DIR = src
TARGET = main.exe

all:
	@mkdir -p $(BUILD_DIR)
	@echo -n "Building..."; go build -o $(BUILD_DIR)/$(TARGET) -tags '$(TAGS)' ./$(SRC_DIR)
	@echo "done."

run:
	@$(BUILD_DIR)/$(TARGET)

clean:
	@rm -rf $(BUILD_DIR)

.PHONY: all run clean
