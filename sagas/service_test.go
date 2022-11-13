package sagas

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderSagaService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stock := &MockStock{}
		stock.
			EXPECT().
			Reserve(12345, 11).
			Return(999, nil).
			Once()

		service := NewOrderSagaService(stock)
		require.NotNil(t, service)

		order := NewOrder(100, 0, 12345, 11)
		sagaId, err := service.Run(order)
		require.NoError(t, err)
		require.Equal(t, order.id, sagaId)

		info := service.SagaInfo(100)
		require.NotNil(t, info)
		require.Equal(t, 100, info.ID())
		require.Equal(t, 999, info.ReserveID())
		require.Equal(t, Log{
			"Reserve Process",
			"Reserve Success",
		}, info.Log())

		stock.AssertExpectations(t)
	})
	t.Run("reserve item failed", func(t *testing.T) {
		stock := &MockStock{}
		stock.
			EXPECT().
			Reserve(12345, 11).
			Return(0, errors.New("shit happens")).
			Once()

		service := NewOrderSagaService(stock)
		require.NotNil(t, service)

		order := NewOrder(100, 0, 12345, 11)
		sagaId, err := service.Run(order)
		require.EqualError(t, err, "shit happens")
		require.Zero(t, sagaId)

		info := service.SagaInfo(100)
		require.NotNil(t, info)
		require.Equal(t, 100, info.ID())
		require.Equal(t, Log{
			"Reserve Process",
			"Reserve Fail: shit happens",
		}, info.Log())

		stock.AssertExpectations(t)
	})
	t.Run("running multiple sagas", func(t *testing.T) {
		stock := &MockStock{}
		stock.EXPECT().Reserve(111111, 11).Return(1, nil).Once()
		stock.EXPECT().Reserve(111222, 22).Return(2, nil).Once()
		stock.EXPECT().Reserve(111333, 33).Return(3, nil).Once()

		service := NewOrderSagaService(stock)
		require.NotNil(t, service)

		sagaId, err := service.Run(NewOrder(100, 0, 111111, 11))
		require.NoError(t, err)
		require.Equal(t, 100, sagaId)

		sagaId, err = service.Run(NewOrder(200, 0, 111222, 22))
		require.NoError(t, err)
		require.Equal(t, 200, sagaId)

		sagaId, err = service.Run(NewOrder(300, 0, 111333, 33))
		require.NoError(t, err)
		require.Equal(t, 300, sagaId)

		stock.AssertExpectations(t)

	})
	t.Run("order already processed", func(t *testing.T) {
		stock := &MockStock{}
		stock.
			EXPECT().
			Reserve(12345, 11).
			Return(999, nil).
			Once()

		service := NewOrderSagaService(stock)
		require.NotNil(t, service)

		order := NewOrder(100, 0, 12345, 11)

		_, err := service.Run(order)
		require.NoError(t, err)

		_, err = service.Run(order)
		require.ErrorIs(t, err, ErrOrderAlreadyProcessed)

		stock.AssertExpectations(t)

	})
}
