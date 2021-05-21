package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/VanBur/tcp-chat/internal/message"
	"github.com/VanBur/tcp-chat/internal/server"
	"github.com/VanBur/tcp-chat/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendReceiveMessageToSelectedUser(t *testing.T) {
	const (
		network      = "tcp"
		address      = ":8008"
		senderName   = "tim"
		receiverName = "joe"
		messageText  = "Hello!"
	)

	s, err := server.New(network, address)
	require.NoError(t, err)

	go s.Serve()
	defer s.Stop()

	send, err := user.New(network, address)
	require.NoError(t, err)

	recv, err := user.New(network, address)
	require.NoError(t, err)

	sendRegMsg := message.Message{CommandType: message.Connect, User: senderName}
	err = send.Send(sendRegMsg)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	recvRegMsg := message.Message{CommandType: message.Connect, User: receiverName}
	err = recv.Send(recvRegMsg)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	broadcastMsg := message.Message{
		CommandType: message.Broadcast,
		User:        senderName,
		Msg:         &message.ChatMessage{To: receiverName, Text: messageText},
	}

	err = send.Send(broadcastMsg)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	msg, err := recv.Read()
	require.NoError(t, err)
	assert.Equal(t, broadcastMsg, *msg)
}

func TestSendReceiveMessageToAllUsers(t *testing.T) {
	const (
		network             = "tcp"
		address             = ":8008"
		senderName          = "tim"
		receiversUsersCount = 10
		messageText         = "Hello!"
	)

	s, err := server.New(network, address)
	require.NoError(t, err)

	go s.Serve()
	defer s.Stop()

	sender, err := user.New(network, address)
	require.NoError(t, err)

	sendRegMsg := message.Message{CommandType: message.Connect, User: senderName}
	err = sender.Send(sendRegMsg)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	receivers := make([]*user.User, 0, receiversUsersCount)
	for i := 0; i < receiversUsersCount; i++ {
		recv, err := user.New(network, address)
		require.NoError(t, err)

		receivers = append(receivers, recv)

		recvName := rand.Int()

		recvRegMsg := message.Message{CommandType: message.Connect, User: strconv.Itoa(recvName)}
		err = recv.Send(recvRegMsg)
		require.NoError(t, err)

		time.Sleep(time.Millisecond * 100)
	}

	broadcastMsg := message.Message{
		CommandType: message.Broadcast,
		User:        senderName,
		Msg:         &message.ChatMessage{Text: messageText},
	}

	err = sender.Send(broadcastMsg)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	for _, currReceiver := range receivers {
		msg, err := currReceiver.Read()
		require.NoError(t, err)
		assert.Equal(t, broadcastMsg, *msg)
	}
}

func TestSendReceiveMessageToDisconnectedUser(t *testing.T) {
	const (
		network      = "tcp"
		address      = ":8008"
		senderName   = "tim"
		receiverName = "joe"
		messageText  = "Hello!"
	)

	s, err := server.New(network, address)
	require.NoError(t, err)

	go s.Serve()
	defer s.Stop()

	send, err := user.New(network, address)
	require.NoError(t, err)

	recv, err := user.New(network, address)
	require.NoError(t, err)

	sendRegMsg := message.Message{CommandType: message.Connect, User: senderName}
	err = send.Send(sendRegMsg)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	recvRegMsg := message.Message{CommandType: message.Connect, User: receiverName}
	err = recv.Send(recvRegMsg)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	recvLeaveMsg := message.Message{CommandType: message.Disconnect, User: receiverName}
	err = recv.Send(recvLeaveMsg)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)

	broadcastMsg := message.Message{
		CommandType: message.Broadcast,
		User:        senderName,
		Msg:         &message.ChatMessage{To: receiverName, Text: messageText},
	}

	err = send.Send(broadcastMsg)
	require.NoError(t, err)
}
