package sagas

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestOrderSaga(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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

		order := NewOrder(100, 1, 12345, 11, 1010)
		saga := NewOrderSaga(order, stock, payment)
		require.NotNil(t, saga)

		err := saga.Run()
		require.NoError(t, err)

		require.Equal(t, 100, saga.ID())
		require.Equal(t, 999, saga.ReserveID())
		require.Equal(t, 111, saga.PayID())
		require.Equal(t, Log{
			"Reserve Process",
			"Reserve Success",
			"Pay Process",
			"Pay Success",
		}, saga.Log())

		err = saga.Run()
		require.ErrorIs(t, err, ErrSagaFinished)

		stock.AssertExpectations(t)
		payment.AssertExpectations(t)
	})
	t.Run("reserve item failed", func(t *testing.T) {
		stock := &MockStock{}
		stock.
			EXPECT().
			Reserve(12345, 11).
			Return(0, errors.New("shit happens")).
			Once()

		payment := &MockPayment{}

		order := NewOrder(100, 0, 12345, 11, 1050)
		saga := NewOrderSaga(order, stock, payment)
		require.NotNil(t, saga)

		err := saga.Run()
		require.EqualError(t, err, "shit happens")
		require.Equal(t, 100, saga.ID())
		require.Equal(t, Log{
			"Reserve Process",
			"Reserve Fail: shit happens",
		}, saga.Log())

		stock.AssertExpectations(t)
		payment.AssertExpectations(t)
	})
	t.Run("reserve item failed and retry success", func(t *testing.T) {
		stock := &MockStock{}
		stock.
			EXPECT().
			Reserve(12345, 11).
			Return(0, errors.New("shit happens")).
			Once()
		stock.
			EXPECT().
			Reserve(12345, 11).
			Return(999, nil).
			Once()

		payment := &MockPayment{}
		payment.
			EXPECT().
			Pay(1, 1050).
			Return(111, nil).
			Once()

		order := NewOrder(100, 1, 12345, 11, 1050)
		saga := NewOrderSaga(order, stock, payment)
		require.NotNil(t, saga)

		err := saga.Run()
		require.EqualError(t, err, "shit happens")

		err = saga.Run()
		require.NoError(t, err)

		require.Equal(t, Log{
			"Reserve Process",
			"Reserve Fail: shit happens",
			"Reserve Process",
			"Reserve Success",
			"Pay Process",
			"Pay Success",
		}, saga.Log())

		stock.AssertExpectations(t)
		payment.AssertExpectations(t)
	})
	t.Run("pay failed", func(t *testing.T) {
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
			Return(0, errors.New("shit happens")).
			Once()

		order := NewOrder(100, 1, 12345, 11, 1010)
		saga := NewOrderSaga(order, stock, payment)
		require.NotNil(t, saga)

		err := saga.Run()
		require.EqualError(t, err, "shit happens")
		require.Equal(t, 100, saga.ID())
		require.Equal(t, 999, saga.ReserveID())
		require.Equal(t, 0, saga.PayID())
		require.Equal(t, Log{
			"Reserve Process",
			"Reserve Success",
			"Pay Process",
			"Pay Fail: shit happens",
		}, saga.Log())

		stock.AssertExpectations(t)
		payment.AssertExpectations(t)
	})
	t.Run("pay failed and retry success", func(t *testing.T) {
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
			Return(0, errors.New("shit happens")).
			Once()
		payment.
			EXPECT().
			Pay(1, 1010).
			Return(111, nil).
			Once()

		order := NewOrder(100, 1, 12345, 11, 1010)
		saga := NewOrderSaga(order, stock, payment)
		require.NotNil(t, saga)

		err := saga.Run()
		require.EqualError(t, err, "shit happens")

		err = saga.Run()
		require.NoError(t, err)
		require.Equal(t, 100, saga.ID())
		require.Equal(t, 999, saga.ReserveID())
		require.Equal(t, 111, saga.PayID())
		require.Equal(t, Log{
			"Reserve Process",
			"Reserve Success",
			"Pay Process",
			"Pay Fail: shit happens",
			"Pay Process",
			"Pay Success",
		}, saga.Log())

		stock.AssertExpectations(t)
		payment.AssertExpectations(t)
	})
	t.Run("pay failed and rollback", func(t *testing.T) {
		stock := &MockStock{}
		payment := &MockPayment{}

		stock.
			EXPECT().
			Reserve(12345, 11).
			Return(999, nil).
			Once()
		payment.
			EXPECT().
			Pay(1, 1010).
			Return(0, errors.New("shit happens")).
			Once()
		stock.
			EXPECT().
			CancelReserve(999).
			Return(nil).
			Once()

		order := NewOrder(100, 1, 12345, 11, 1010)
		saga := NewOrderSaga(order, stock, payment)
		require.NotNil(t, saga)

		err := saga.Run()
		require.EqualError(t, err, "shit happens")

		err = saga.Rollback()
		require.NoError(t, err)
		require.Equal(t, 100, saga.ID())
		require.Equal(t, 999, saga.ReserveID())
		require.Equal(t, 0, saga.PayID())
		require.Equal(t, Log{
			"Reserve Process",
			"Reserve Success",
			"Pay Process",
			"Pay Fail: shit happens",
			"Reserve Rollback Start",
			"Reserve Rollback Success",
		}, saga.Log())

		stock.AssertExpectations(t)
		payment.AssertExpectations(t)
	})
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
