CC=clang
CFLAGS=-g `llvm-config --cflags`
CXXFLAGS=-g `llvm-config --cxxflags`
LD=clang++
LDFLAGS=`llvm-config --cxxflags --ldflags --libs core executionengine mcjit asmprinter asmparser native irreader --system-libs`

all: main

%.c.o : %.c
	$(CC) $(CFLAGS) -c $< -o $<.o

%.cpp.o: %.cpp
	$(CC) $(CXXFLAGS) -c $< -o $<.o

main: main.cpp.o
	$(LD) $< $(LDFLAGS) -o $@

engine_test: engine.cpp.o engine_test.cpp.o memory_manager.cpp.o
	$(LD) $^ $(LDFLAGS) -o $@

sum: sum.c.o
	$(LD) $^ $(LDFLAGS) -o $@

clean:
	-rm -f *.o main sum engine_test

ir:
	clang -S -emit-llvm sample.c
