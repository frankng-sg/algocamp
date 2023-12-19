TARGET := main

.PHONY: sample input all clean re

$(TARGET):
	go build -o $(TARGET)

all: $(TARGET)

sample: re
	./$(TARGET) sample.txt

input: re
	./$(TARGET) input.txt

clean:
	rm -f $(TARGET)

re: clean all
