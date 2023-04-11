This is  a go Project using gin and gorilla framework
you need to install go in your system https://go.dev/doc/install
after that you need to install docker https://www.docker.com/products/docker-desktop/
then there is  make file in this repository for you to create docker image and database you can follow it to create databse and tables.
you also need to install scoop using  irm get.scoop.sh | iex this command line so that you can install migrate then.
you need to initialize go mod using go mod init so when the different packages like gin and gorilla gives errors just run go mod tidy to automatically install those packages.
and the simply run go run cmd/main.go to run the project.
