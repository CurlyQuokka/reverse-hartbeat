CERT_DIR_DEF := bin/certs
CN_DEF := localhost

CERTS_DIR ?= $(CERT_DIR_DEF)
CN ?= $(CN_DEF)

.PHONY: clean
clean:
	rm -rf certs bin

.PHONY: build-server
build-server:
	go build -o bin/server cmd/server/main.go

.PHONY: build-client
build-client:
	go build -o bin/client cmd/client/main.go

.PHONY: build
build: build-server build-client

.PHONY: certs
certs:
	rm -rf $(CERTS_DIR)
	mkdir -p $(CERTS_DIR) 2> /dev/null
	# ca.key
	openssl genrsa -out $(CERTS_DIR)/ca.key 2048
	# ca.crt
	openssl req -new -key $(CERTS_DIR)/ca.key -x509 -days 3650 -out $(CERTS_DIR)/ca.crt -subj /C=PL/ST=Szczecin/O="Localhost"/CN="Localhost Root" -addext "subjectAltName = DNS:localhost"
	# server.key
	openssl genrsa -out $(CERTS_DIR)/server.key 2048
	# server.csr
	openssl req -new -nodes -key $(CERTS_DIR)/server.key -out $(CERTS_DIR)/server.csr -subj /C=PL/ST=Zachodniopomorskie/L=Szczecin/O="Localhost Server"/CN=localhost -addext "subjectAltName = DNS:localhost"
	#  server.crt
	openssl x509 -copy_extensions copy -req -in $(CERTS_DIR)/server.csr -CA $(CERTS_DIR)/ca.crt -CAkey $(CERTS_DIR)/ca.key -CAcreateserial -out $(CERTS_DIR)/server.crt
	#  client.key
	openssl genrsa -out $(CERTS_DIR)/client.key 2048
	# client.csr
	openssl req -new -nodes -key $(CERTS_DIR)/client.key -out $(CERTS_DIR)/client.csr -subj /C=PL/ST=Zachodniopomorskie/L=Szczecin/O="Localhost Client"/CN=localhost -addext "subjectAltName = DNS:localhost"
	# client.crt
	openssl x509 -copy_extensions copy -req -in $(CERTS_DIR)/client.csr -CA $(CERTS_DIR)/ca.crt -CAkey $(CERTS_DIR)/ca.key -CAcreateserial -out $(CERTS_DIR)/client.crt

.PHONY: all
all: certs build
