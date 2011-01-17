all:
	make -C pkg/buffers install
	make -C cmd/hello
	make -C cmd/buffer

clean:
	make -C pkg/buffers clean
	make -C cmd/hello clean
	make -C cmd/buffer clean

.PHONY: all clean

