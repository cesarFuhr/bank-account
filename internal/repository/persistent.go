package repository

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/cesarFuhr/bank-account/internal/account"
)

// ErrAccNaoEncontrada erro de conta nao encontrada.
var ErrAccNaoEncontrada = errors.New("a conta solicitada nao foi encontrada")

// PersistentRepository é capaz de ler e escrever contas
// de forma persistente.
type PersistentRepository struct {
	db     *bolt.DB
	bucket string
}

// NewPersistentRepository cria um novo repositorio persistente
// e retorna um ponteiro para ele.
func NewPersistentRepository() (*PersistentRepository, error) {
	db, err := bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	bucketName := "accounts"
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	pr := PersistentRepository{
		db:     db,
		bucket: bucketName,
	}

	return &pr, nil
}

// Escrever escreve um registro de conta no banco de dados.
func (r *PersistentRepository) Escrever(acc account.Account) error {
	var buff bytes.Buffer
	if err := gob.NewEncoder(&buff).Encode(&acc); err != nil {
		return err
	}

	rawAcc, err := io.ReadAll(&buff)
	if err != nil {
		return err
	}

	id := strconv.Itoa(int(acc.ID))

	err = r.db.Update(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(r.bucket))

		return b.Put([]byte(id), rawAcc)
	})
	if err != nil {
		return nil
	}

	return nil
}

// Ler le um registro de conta e retorna pelo ID.
func (r *PersistentRepository) Ler(id uint) (account.Account, error) {
	strID := strconv.Itoa(int(id))

	var rawAcc []byte
	r.db.View(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(r.bucket))
		rawAcc = b.Get([]byte(strID))

		return nil
	})
	if rawAcc == nil {
		return account.Account{}, ErrAccNaoEncontrada
	}

	buff := bytes.NewBuffer(rawAcc)

	var acc account.Account
	if err := gob.NewDecoder(buff).Decode(&acc); err != nil {
		return account.Account{}, err
	}

	return acc, nil
}
