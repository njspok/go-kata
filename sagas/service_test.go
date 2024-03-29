package sagas

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
)

func TestOrderSagaService(t *testing.T) {
	t.Run("running multiple sagas", func(t *testing.T) {
		stock := &MockStock{}
		stock.EXPECT().Reserve(mock.Anything, mock.Anything).Return(1, nil).Once()
		stock.EXPECT().Reserve(mock.Anything, mock.Anything).Return(2, nil).Once()
		stock.EXPECT().Reserve(mock.Anything, mock.Anything).Return(3, nil).Once()

		payment := &MockPayment{}
		payment.EXPECT().Pay(mock.Anything, mock.Anything).Return(1, nil).Once()
		payment.EXPECT().Pay(mock.Anything, mock.Anything).Return(2, nil).Once()
		payment.EXPECT().Pay(mock.Anything, mock.Anything).Return(3, nil).Once()

		service := NewOrderSagaService(stock, payment)
		require.NotNil(t, service)

		sagaId, err := service.Run(NewOrder(100, 0, 111111, 11, 0))
		require.NoError(t, err)
		require.Equal(t, 100, sagaId)

		sagaId, err = service.Run(NewOrder(200, 0, 111222, 22, 0))
		require.NoError(t, err)
		require.Equal(t, 200, sagaId)

		sagaId, err = service.Run(NewOrder(300, 0, 111333, 33, 0))
		require.NoError(t, err)
		require.Equal(t, 300, sagaId)

		stock.AssertExpectations(t)
		payment.AssertExpectations(t)
	})
	t.Run("order already processed", func(t *testing.T) {
		stock := &MockStock{}
		stock.
			EXPECT().
			Reserve(12345, 11).
			Return(999, nil).
			Once()

		payment := &MockPayment{}
		payment.
			EXPECT().
			Pay(1, 1010).
			Return(111, nil).
			Once()

		service := NewOrderSagaService(stock, payment)
		require.NotNil(t, service)

		order := NewOrder(100, 1, 12345, 11, 1010)

		_, err := service.Run(order)
		require.NoError(t, err)

		_, err = service.Run(order)
		require.ErrorIs(t, err, ErrOrderAlreadyProcessed)

		stock.AssertExpectations(t)
		payment.AssertExpectations(t)
	})
}
