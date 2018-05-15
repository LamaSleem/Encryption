package main_test

import (
	"reflect"
	mock "room-one/mocks"
	pb "room-one/proto"
	"testing"

	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
)

func TestReverese(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEncryptionServiceClient := mock.NewMockEncryptionServiceClient(ctrl)
	req := &pb.InputString{InitalInput: "HELLO"}
	mockEncryptionServiceClient.EXPECT().Reverse(context.Background(), req).Return(&pb.ReveresedString{FinalOutput: "OLLEH"}, nil).Times(1)
	testReverse(t, mockEncryptionServiceClient)
}

func testReverse(t *testing.T, client pb.EncryptionServiceClient) {
	r, err := client.Reverse(context.Background(), &pb.InputString{InitalInput: "HELLO"})
	if err != nil || r.FinalOutput != "OLLEH" {
		t.Errorf("MTest Reverese Failed")
	}
	t.Log("Reply : ", r.FinalOutput)
}

func TestEncrypt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockencryptClient := mock.NewMockEncryptionServiceClient(ctrl)
	req := &pb.RequestText{PlainMessage: "HELLO"}
	mockencryptClient.EXPECT().Encrypt(context.Background(), req).Return(&pb.EncryptedText{EncryptedMessage: "gqJeiNxRHZqNcPfgVV0q6s6_4hlu9pU6p8A_vQe76bI"}, nil).Times(1)
	testEncrypt(t, mockencryptClient)
}

func testEncrypt(t *testing.T, client pb.EncryptionServiceClient) {
	r, err := client.Encrypt(context.Background(), &pb.RequestText{PlainMessage: "HELLO"})
	if err != nil || (!reflect.DeepEqual(r.EncryptedMessage, "gqJeiNxRHZqNcPfgVV0q6s6_4hlu9pU6p8A_vQe76bI")) {
		t.Errorf("Error in Encryption")
	}
	t.Log("Reply : ", r.EncryptedMessage)
}
