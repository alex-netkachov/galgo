.PHONY: all clean

bin/primsmstsub: primsmstsub.cpp
	g++ primsmstsub.cpp -o bin/primsmstsub -std=c++11

all: bin/primsmstsub primsmstsub1.in primsmstsub2.in
	./bin/primsmstsub < primsmstsub1.in | diff -w primsmstsub1.out -
	./bin/primsmstsub < primsmstsub2.in | diff -w primsmstsub2.out -

test1: bin/primsmstsub primsmstsub1.in
	./bin/primsmstsub < primsmstsub1.in

test2: bin/primsmstsub primsmstsub2.in
	./bin/primsmstsub < primsmstsub2.in

clean:
	rm bin/primsmstsub
