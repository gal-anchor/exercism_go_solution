package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a random private key in the range [2, p-2].
func PrivateKey(p *big.Int) *big.Int {
	one := big.NewInt(1)
	max := new(big.Int).Sub(p, one) // p - 1
	min := big.NewInt(2)

	for {
		priv, err := rand.Int(rand.Reader, max) // [0, p-2]
		if err != nil {
			panic(err)
		}
		priv.Add(priv, min) // [2, p-1]
		if priv.Cmp(max) < 0 {
			return priv
		}
	}
}

// PublicKey calculates the public key: g^private mod p
func PublicKey(private, p *big.Int, g int64) *big.Int {
	base := big.NewInt(g)
	pub := new(big.Int).Exp(base, private, p)
	return pub
}

// NewPair generates a random private key and corresponding public key.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priv := PrivateKey(p)
	pub := PublicKey(priv, p, g)
	return priv, pub
}

// SecretKey computes the shared secret: public2^private1 mod p
func SecretKey(private1, public2, p *big.Int) *big.Int {
	secret := new(big.Int).Exp(public2, private1, p)
	return secret
}
