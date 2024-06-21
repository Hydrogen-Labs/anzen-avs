package chainio

//go:generate mockgen -destination=./mocks/avs_subscriber.go -package=mocks anzen-avs/core/chainio AvsSubscriberer
//go:generate mockgen -destination=./mocks/avs_writer.go -package=mocks anzen-avs/core/chainio AvsWriterer
//go:generate mockgen -destination=./mocks/avs_reader.go -package=mocks anzen-avs/core/chainio AvsReaderer
