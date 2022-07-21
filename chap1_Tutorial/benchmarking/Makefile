bench: build
	go test -bench=. -benchmem 


build:
	go build -o efficient_echo ../efficient_echo.go
	go build -o inefficient_echo1 ../inefficient_echo1.go
	go build -o inefficient_echo2 ../inefficient_echo2.go


clean: 
	rm efficient_echo inefficient_echo*
