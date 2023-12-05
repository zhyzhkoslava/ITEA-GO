package testing

import (
	"errors"
	"testing"
	"time"
)

func TestOrderSuccessProcess(t *testing.T) {
	customer := Customer{Email: "example@gmail.com"}
	order := NewOrder(customer)

	err := order.Process()
	if err != nil {
		t.Errorf("Unexpected error processing order: %s", err)
	}

	if order.Status != Processing {
		t.Errorf("Expected status %s after processing, got %s", Processing, order.Status)
	}
}

func TestOrderAlreadyProcess(t *testing.T) {
	customer := Customer{Email: "example@gmail.com"}
	order := NewOrder(customer)

	err := order.Process()
	if err != nil {
		t.Errorf("Unexpected error processing order: %s", err)
	}

	actualErr := order.Process()
	t.Log("Actual Error:", actualErr)

	// expectedErr := "Cannot process order. Invalid status."
	// if actualErr.Error() != expectedErr {
	// 	t.Errorf("Expected error processing an already processing order, got %v", actualErr)
	// }
	expectedErr := errors.New("Cannot process order. Invalid status.")
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error processing an already processing order, got %v", err)
	}
}

func TestOrderMarkAsSuccess(t *testing.T) {
	customer := Customer{Email: "example@gmail.com"}
	order := NewOrder(customer)

	err := order.Process()
	if err != nil {
		t.Errorf("Error processing order: %s", err)
	}

	err = order.MarkAsSuccess()
	if err != nil {
		t.Errorf("Error marking order as success: %s", err)
	}

	if order.Status != Success {
		t.Errorf("Expected status %s after marking as success, got %s", Success, order.Status)
	}
}

func TestOrderMarkAsSuccessFailure(t *testing.T) {
	customer := Customer{Email: "example@gmail.com"}
	order := NewOrder(customer)

	err := order.MarkAsSuccess()

	expectedErrMsg := "Cannot mark order as success. Invalid status."
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Expected error marking as success without processing, got %v", err)
	}
}

func TestOrderMarkAsFailFailure(t *testing.T) {
	customer := Customer{Email: "example@gmail.com"}
	order := NewOrder(customer)

	err := order.MarkAsFail()

	expectedErrMsg := "Cannot mark order as fail. Invalid status."
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Expected error marking as fail without processing, got %v", err)
	}
}

func TestOrderMarkAsFailSuccessful(t *testing.T) {
	customer := Customer{Email: "example@gmail.com"}
	order := NewOrder(customer)

	err := order.Process()
	if err != nil {
		t.Errorf("Error processing order: %s", err)
	}

	err = order.MarkAsFail()
	if err != nil {
		t.Errorf("Error marking order as fail: %s", err)
	}

	if order.Status != Fail {
		t.Errorf("Expected status %s after marking as fail, got %s", Fail, order.Status)
	}
}

func TestOrderString(t *testing.T) {
	tests := []struct {
		status Status
		result string
	}{
		{Initiated, "initiated"},
		{Processing, "processing"},
		{Success, "success"},
		{Fail, "fail"},
	}

	for _, test := range tests {
		if result := test.status.String(); result != test.result {
			t.Errorf("Expected %s, got %s", test.result, result)
		}
	}
}

func TestNewOrder(t *testing.T) {
	customer := Customer{Email: "example@gmail.com"}
	order := NewOrder(customer)

	if order.Customer != customer {
		t.Errorf("Expected customer %v, got %v", customer, order.Customer)
	}

	if order.Status != Initiated {
		t.Errorf("Expected status %s, got %s", Initiated, order.Status)
	}

	if !order.CreatedAt.Before(time.Now()) {
		t.Errorf("Expected CreatedAt to be before the current time")
	}

	if !order.CreatedAt.Equal(order.UpdatedAt) {
		t.Errorf("Expected CreatedAt and UpdatedAt to be identical")
	}
}
