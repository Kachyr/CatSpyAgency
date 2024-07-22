# Build and start the application for development
# App will automatically rebild and start on save
runDev:
	@CompileDaemon -build="go build -o sca.exe" -command="./cmd/sca.exe" -directory=cmd -directory=./ -exclude=Makefile -exclude=.exe -exclude=.exe~ -exclude=.git exclude-dir=".trunk"
# Build the application for production
build:
	go build -o sca.exe ./cmd

# Run the application for production
run:
	./sca.exe

# Clean up build artifacts
clean:
	rm -f ./sca.exe
	rm -f ./sca.exe~