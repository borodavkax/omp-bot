package transaction

import (
	"errors"
	"github.com/ozonmp/omp-bot/internal/model/bank/transaction"
	"log"
	"sync"
)

type Service interface {
	Describe(transactionID uint64) (*transaction.Transaction, error)
	List(cursor uint64, limit uint64) ([]transaction.Transaction, error)
	Create(transaction.Transaction) (uint64, error)
	Update(transactionID uint64, transaction transaction.Transaction) error
	Remove(transactionID uint64) (bool, error)
}

var (
	ErrOutOfBoundError = errors.New("cursor out of bound")
	ErrNotExists       = errors.New("given transactionID not exists")
	ErrNotImplemented  = errors.New("not implemented yet")
)

type DummyService struct {
	mtx      sync.RWMutex
	newID    uint64
	entities []transaction.Transaction
}

func NewDummyTransactionService() *DummyService {
	return &DummyService{newID: 0, entities: transaction.AllTransactions}
}

func (s *DummyService) Describe(transactionID uint64) (*transaction.Transaction, error) {
	for i := 0; i < len(s.entities); i++ {
		if transactionID == s.entities[i].ID {
			return &s.entities[i], nil
		}
	}
	log.Printf("transaction.DummyService.Describe not exists ID: %v", transactionID)
	return nil, ErrNotExists
}

func (s *DummyService) List(cursor uint64, limit uint64) ([]transaction.Transaction, error) {
	if limit == 0 {
		return s.entities, nil
	}
	l := uint64(len(s.entities))
	if cursor >= l {
		log.Printf("transaction.DummyService.List out of bound: %v, %v", cursor, l)
		return nil, ErrOutOfBoundError
	}
	if cursor+limit >= l {
		return s.entities[cursor:], nil
	}
	return s.entities[cursor : limit+cursor], nil
}

func (s *DummyService) Update(transactionID uint64, transaction transaction.Transaction) error {
	log.Printf("transaction.DummyService.List not implemented: %v, %v", transactionID, transaction)
	return ErrNotImplemented
}

func (s *DummyService) Create(tx transaction.Transaction) (uint64, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	newID := s.getNextID()
	t := transaction.Transaction{
		ID:              newID,
		BankAccountFrom: tx.BankAccountFrom,
		BankAccountTo:   tx.BankAccountTo,
		CreatedAt:       tx.CreatedAt,
		Amount:          tx.Amount,
		Currency:        tx.Currency,
	}
	s.entities = append(s.entities, t)
	return newID, nil
}

func (s *DummyService) Remove(transactionID uint64) (bool, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	for i := 0; i < len(s.entities); i++ {
		if transactionID == s.entities[i].ID {
			s.entities = append(s.entities[:i], s.entities[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (s *DummyService) getNextID() uint64 {
	if s.newID == 0 {
		if len(s.entities) > 0 {
			s.newID = s.entities[len(s.entities)-1].ID + 1
		} else {
			s.newID = 1
		}
	} else {
		s.newID += 1
	}
	return s.newID
}
