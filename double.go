package perc

import ( /* TODO */ )

// DoubleTransform represents an instance of the SRTP double-encryption
// transform used by PERC.  It provides most SRTP guarantees end-to-end, but
// allows a middlebox with partial key material to perform some modifications.
//
// For now, we only implement the endpoint logic.  Middlebox logic will be
// slightly more complicated, since it will have to deal with changing
// parameters and adding the OHB if necessary.
type DoubleTransform struct {
	// TODO keys, ciphers, etc.
}

// NewDoubleTransform constructs a new transform from cryptographic parameters
// (e.g., keys, salts).
func NewDoubleTransform( /* TODO */ ) *DoubleTransform {
	// TODO
	return nil
}

// Encrypt performs inner and outer encryption transforms on the payload
// XXX(rlb): I think we can do this in place; else return ([]byte, error)
// XXX(rlb): We will probably want to factor this into inner and outer encrypt
//           methods later, to facilitate the middlebox function
func (dt DoubleTransform) Encrypt(payload []byte) error {
	// TODO inner encrypt
	// TODO outer enrypt
	return nil
}

// Encrypt performs inner and outer encryption transforms on the payload
// XXX(rlb): I think we can do this in place; else return ([]byte, error)
// XXX(rlb): We will probably want to factor this into inner and outer decrypt
//           methods later, to facilitate the middlebox function
func (dt DoubleTransform) Decrypt(payload []byte) error {
	// TODO inner encrypt
	// TODO outer enrypt
	return nil
}
