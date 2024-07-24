PROTO_DIR=./protos
PROTO_FILE=$(PROTO_DIR)/currency.proto
OUT_DIR=$(PROTO_DIR)/currency
PROTO_GEN_FILES=$(OUT_DIR)/currency/*.pb.go

all: protos

protos: $(PROTO_GEN_FILES)

$(PROTO_GEN_FILES): $(PROTO_FILE)
	protoc -I $(PROTO_DIR) $(PROTO_FILE) --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR)

clean:
	rm -f $(PROTO_GEN_FILES)

.PHONY: all protos clean

