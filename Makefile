make unix-install:
	go build -o out/ . 
	cp out/muuidwizard ~/go/bin